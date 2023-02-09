package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/market"
	lcli "github.com/filecoin-project/lotus/cli"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/filecoin-project/lotus/gateway"
	"github.com/ipfs/go-cid"
	"github.com/nonsense/fevmtest/dc"
	"github.com/urfave/cli/v2"
)

var (
	// chain id -- ideally fetch this from chain, but seems like rpc is not supported just yet
	chainId = big.NewInt(31415926)

	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//log.Fatal(err)
	//}
)

var submitDealProposalCmd = &cli.Command{
	Name:  "submit-deal",
	Usage: "Submit a deal on chain",
	Flags: []cli.Flag{
		//&cli.StringFlag{
		//Name:     "provider",
		//Usage:    "storage provider on-chain address",
		//Required: true,
		//},
		&cli.StringFlag{
			Name:     "commp",
			Usage:    "commp of the CAR file",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "piece-size",
			Usage:    "size of the CAR file as a padded piece",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "car-size",
			Usage:    "size of the CAR file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "payload-cid",
			Usage:    "root CID of the CAR file",
			Required: true,
		},
		&cli.IntFlag{
			Name:        "start-epoch",
			Usage:       "start epoch by when the deal should be proved by provider on-chain",
			DefaultText: "current chain head + 2 days",
		},
		&cli.IntFlag{
			Name:  "duration",
			Usage: "duration of the deal in epochs",
			Value: 518400, // default is 2880 * 180 == 180 days
		},
		&cli.IntFlag{
			Name:  "provider-collateral",
			Usage: "deal collateral that storage miner must put in escrow; if empty, the min collateral for the given piece size will be used",
		},
		&cli.Int64Flag{
			Name:  "storage-price",
			Usage: "storage price in attoFIL per epoch per GiB",
			Value: 1,
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "whether the deal funds should come from verified client data-cap",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "remove-unsealed-copy",
			Usage: "indicates that an unsealed copy of the sector in not required for fast retrieval",
			Value: false,
		},
		&cli.StringFlag{
			Name:  "client",
			Usage: "client address to be used to initiate the deal - should point to contract?",
		},
		&cli.BoolFlag{
			Name:  "skip-ipni-announce",
			Usage: "indicates that deal index should not be announced to the IPNI(Network Indexer)",
			Value: false,
		},
		&cli.Int64Flag{
			Name:  "provider-collateral",
			Usage: "provider collateral (chain dependent)",
			Value: 0, // or 147276332 on mainnet?
		},
		&cli.StringFlag{
			Name:  "contract",
			Usage: "",
			Value: "0x690b9a9e9aa1c9db991c7721a92d351db4fac990",
		},
	},
	Action: func(cctx *cli.Context) error {

		//header := make(http.Header)
		//header.Add("Authorization", "Bearer "+token)

		ctx := context.Background()

		subCh := gateway.NewEthSubHandler()

		api, closer, err := lcli.GetFullNodeAPIV1(cctx, cliutil.FullNodeWithEthSubscribtionHandler(subCh))
		if err != nil {
			return err
		}
		defer closer()

		endpoint := cctx.String("endpoint")

		//ctx := context.Background()
		cl, err := rpc.Dial(endpoint)
		//cl, err := rpc.DialWebsocketWithHeader(ctx, endpoint, "", header) // fork
		if err != nil {
			return err
		}

		client := ethclient.NewClient(cl)

		DealClientContractAddr := common.HexToAddress(cctx.String("contract"))
		dealclient, err := dc.NewDealClient(DealClientContractAddr, client)
		if err != nil {
			return err
		}

		privateKey := cctx.String("private-key")
		pk, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return err
		}

		publicKey := pk.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		me := crypto.PubkeyToAddress(*publicKeyECDSA)

		signFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewLondonSigner(chainId), pk)
		}

		opts := &bind.TransactOpts{
			From:      me,
			GasFeeCap: big.NewInt(100000000),
			GasTipCap: big.NewInt(100000000),
			Signer:    signFn,
			Nonce:     nil,
			Value:     nil,
			GasLimit:  uint64(1000000000),
		}

		tipset, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}

		commp := cctx.String("commp")
		pieceCid, err := cid.Parse(commp)
		if err != nil {
			return fmt.Errorf("parsing commp '%s': %w", commp, err)
		}

		pieceSize := cctx.Uint64("piece-size")
		if pieceSize == 0 {
			return fmt.Errorf("must provide piece-size parameter for CAR url")
		}

		payloadCidStr := cctx.String("payload-cid")
		rootCid, err := cid.Parse(payloadCidStr)
		if err != nil {
			return fmt.Errorf("parsing payload cid %s: %w", payloadCidStr, err)
		}

		carFileSize := cctx.Uint64("car-size")
		if carFileSize == 0 {
			return fmt.Errorf("size of car file cannot be 0")
		}

		head := tipset.Height()

		startEpoch := head + abi.ChainEpoch(5760)
		endEpoch := startEpoch + 521280 // startEpoch + 181 days
		l, err := market.NewLabelFromString(rootCid.String())
		if err != nil {
			return err
		}
		lbytes, err := l.ToBytes()
		if err != nil {
			return err
		}

		clientAddr, err := address.NewFromString(cctx.String("client")) // i guess the f4 address to the contract that should verify the deal?
		if err != nil {
			return err
		}

		//ProviderCollateral:   policy.MaxProviderCollateral, // this value is chain dependent and is 0 on devnet, and currently a multiple of 147276332 on mainnet
		providerCollateral := uint64(0)

		dealParams := dc.DealClientDealParams{
			LocationRef:        "http://mywebserver/myfile.car",
			RemoveUnsealedCopy: false,
			SkipIPNIAnnounce:   false,
		}

		dealProposal := dc.DealClientDealProposal{
			PieceCid:        pieceCid.Bytes(),
			PaddedPieceSize: pieceSize,
			VerifiedDeal:    false,
			Client:          clientAddr.Bytes(),
			// note that there is no Provider,
			Label:                lbytes,
			StartEpoch:           uint64(startEpoch),
			EndEpoch:             uint64(endEpoch),
			StoragePricePerEpoch: 1,
			ProviderCollateral:   providerCollateral,
			ClientCollateral:     0,
		}

		tx, err := dealclient.MakeDealProposalWithParams(opts, dealProposal, dealParams)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println(tx.Hash())

		return nil
	},
}
