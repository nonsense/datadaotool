package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	mbig "math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	"github.com/nonsense/datadaotool/dc"
	"github.com/urfave/cli/v2"
)

var (
	chainId *mbig.Int

	// hyperspace
	// 3141

	// testnet
	// 31415926
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
		&cli.Int64Flag{
			Name:     "start-epoch",
			Usage:    "start epoch by when the deal should be proved by provider on-chain",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:  "duration",
			Usage: "duration of the deal in epochs",
			Value: 518400, // default is 2880 * 180 == 180 days
		},
		&cli.Int64Flag{
			Name:  "storage-price",
			Usage: "storage price in attoFIL per epoch per GiB",
			Value: 0,
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "whether the deal funds should come from verified client data-cap",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "remove-unsealed-copy",
			Usage: "indicates that an unsealed copy of the sector in not required for fast retrieval",
			Value: false,
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
		&cli.Int64Flag{
			Name:  "chain-id",
			Usage: "network id",
			Value: 3141,
		},
	},
	Action: func(cctx *cli.Context) error {
		endpoint := cctx.String("endpoint")

		chainId = mbig.NewInt(cctx.Int64("chain-id"))

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
			GasFeeCap: mbig.NewInt(100000000),
			GasTipCap: mbig.NewInt(100000000),
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

		ps := cctx.Uint64("piece-size")
		if ps == 0 {
			return fmt.Errorf("must provide piece-size parameter for CAR url")
		}

		payloadCidStr := cctx.String("payload-cid")
		_, err = cid.Parse(payloadCidStr)
		if err != nil {
			return fmt.Errorf("parsing payload cid %s: %w", payloadCidStr, err)
		}

		carFileSize := cctx.Uint64("car-size")
		if carFileSize == 0 {
			return fmt.Errorf("size of car file cannot be 0")
		}

		duration := cctx.Uint64("duration")
		if duration == 0 {
			return fmt.Errorf("duration of deal cant be 0")
		}

		startEpoch := abi.ChainEpoch(cctx.Int64("start-epoch"))
		endEpoch := startEpoch + abi.ChainEpoch(duration) // startEpoch + 181 days

		dr := dc.DealRequest{
			PieceCid:             pieceCid.Bytes(),
			PieceSize:            ps,
			VerifiedDeal:         cctx.Bool("verified"),
			Label:                payloadCidStr,
			StartEpoch:           int64(startEpoch),
			EndEpoch:             int64(endEpoch),
			StoragePricePerEpoch: mbig.NewInt(cctx.Int64("storage-price")),
			ProviderCollateral:   mbig.NewInt(cctx.Int64("provider-collateral")),
			ClientCollateral:     mbig.NewInt(0),
			ExtraParamsVersion:   1,
			ExtraParams: dc.ExtraParamsV1{
				LocationRef:        cctx.String("location_ref"),
				CarSize:            carFileSize,
				SkipIpniAnnounce:   cctx.Bool("skip-ipni-announce"),
				RemoveUnsealedCopy: cctx.Bool("remove-unsealed-copy"),
			},
		}

		tx, err := dealclient.MakeDealProposal(opts, dr)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println(tx.Hash())

		time.Sleep(60 * time.Second)

		hash := tx.Hash()
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			return err
		}

		spew.Dump(receipt)

		return nil
	},
}
