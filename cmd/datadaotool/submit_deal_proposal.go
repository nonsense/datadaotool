package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/davecgh/go-spew/spew"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/market"
	"github.com/ipfs/go-cid"
	"github.com/nonsense/fevmtest/dc"
	"github.com/urfave/cli/v2"
)

var (
	// chain id -- ideally fetch this from chain, but seems like rpc is not supported just yet
	// testnet
	chainId = big.NewInt(31415926)

	// hyperspace
	//chainId = big.NewInt(3141)

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
		&cli.StringFlag{
			Name:  "version",
			Usage: "",
			Value: "1",
		},
		&cli.StringFlag{
			Name:  "location_ref",
			Usage: "",
			Value: "http://webserver.com/carfile.car",
		},
	},
	Action: func(cctx *cli.Context) error {
		endpoint := cctx.String("endpoint")

		cl, err := rpc.Dial(endpoint)
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
			return fmt.Errorf("broken pk: %v", err)
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
			GasLimit:  uint64(10000000000),
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

		//head := tipset.Height()

		head := abi.ChainEpoch(0)

		startEpoch := head + abi.ChainEpoch(5760)
		endEpoch := startEpoch + 521280 // startEpoch + 181 days
		l, err := market.NewLabelFromBytes(rootCid.Bytes())
		if err != nil {
			return fmt.Errorf("new label err: %w", err)
		}
		lbytes, err := l.ToBytes()
		if err != nil {
			return fmt.Errorf("label to string err: %w", err)
		}

		clientAddr, err := address.NewFromString(cctx.String("client")) // i guess the f4 address to the contract that should verify the deal?
		if err != nil {
			return err
		}
		providerCollateral := uint64(0)

		// params marshalling
		paramsRecord := struct {
			LocationRef      string
			CarSize          *big.Int
			SkipIpniAnnounce bool
		}{
			cctx.String("location_ref"),
			big.NewInt(int64(carFileSize)),
			cctx.Bool("skip-ipni-announce"),
		}

		paramsVersion1, _ := ethabi.NewType("tuple", "paramsVersion1", []ethabi.ArgumentMarshaling{
			{Name: "location_ref", Type: "string"},
			{Name: "car_size", Type: "uint256"},
			{Name: "skip_ipni_announce", Type: "bool"},
		})

		params, err := ethabi.Arguments{
			{Type: paramsVersion1, Name: "paramsVersion1"},
		}.Pack(&paramsRecord)
		if err != nil {
			return err
		}

		spew.Dump(params)

		dealProposal := dc.DealClientDealProposal{
			PieceCid:        pieceCid.Bytes(),
			PaddedPieceSize: pieceSize,
			VerifiedDeal:    false,
			Client:          clientAddr.Bytes(),
			// note that there is no Provider,
			Label:                lbytes,
			StartEpoch:           uint64(startEpoch),
			EndEpoch:             uint64(endEpoch),
			StoragePricePerEpoch: cctx.Uint64("storage-price"),
			ProviderCollateral:   providerCollateral,
			ClientCollateral:     0,
			Version:              cctx.String("version"),
			Params:               params,
		}

		tx, err := dealclient.MakeDealProposal(opts, dealProposal)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println(tx.Hash())

		time.Sleep(15 * time.Second)

		hash := tx.Hash()
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			panic(err)
		}

		spew.Dump(receipt)

		return nil
	},
}
