package main

import (
	"encoding/hex"
	"fmt"

	mbig "math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nonsense/datadaotool/dc"
	"github.com/urfave/cli/v2"
)

var getDealProposalCmd = &cli.Command{
	Name:  "get-deal-proposal",
	Usage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "contract",
			Usage: "",
			Value: "0x690b9a9e9aa1c9db991c7721a92d351db4fac990",
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

		id := common.HexToHash(cctx.String("id"))

		var id32 [32]byte
		copy(id32[:], id.Bytes()[:32])

		res, err := dealclient.GetDealProposal(nil, id32)
		if err != nil {
			return err
		}

		spew.Dump(res)

		fmt.Println()

		fmt.Println(hex.EncodeToString(res))

		return nil
	},
}
