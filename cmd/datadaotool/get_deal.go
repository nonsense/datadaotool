package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nonsense/fevmtest/dc"
	"github.com/urfave/cli/v2"
)

var getDealCmd = &cli.Command{
	Name:  "get-deal",
	Usage: "Fetch deal state",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "contract",
			Usage: "",
			Value: "0x690b9a9e9aa1c9db991c7721a92d351db4fac990",
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

		id := common.HexToHash(cctx.String("id"))

		var id32 [32]byte
		copy(id32[:], id.Bytes()[:32])

		proposal, err := dealclient.GetDealProposal(nil, id32)
		if err != nil {
			panic(err)
		}

		//params, err := dealclient.GetDealParams(nil, id32)
		//if err != nil {
		//panic(err)
		//}

		fmt.Println()
		spew.Dump(proposal)
		//spew.Dump(params)

		return nil
	},
}
