package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/urfave/cli/v2"

	cliutil "github.com/filecoin-project/lotus/cli/util"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/gateway"

	lcli "github.com/filecoin-project/lotus/cli"
)

var getDealCmd = &cli.Command{
	Name:  "get-deal",
	Usage: "",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		ctx := context.Background()

		subCh := gateway.NewEthSubHandler()

		api, closer, err := lcli.GetFullNodeAPIV1(cctx, cliutil.FullNodeWithEthSubscribtionHandler(subCh))
		if err != nil {
			return err
		}
		defer closer()

		_from := "0xD66E69D6FeD4202DdB21266C5ECBe3B2Acd738a4"
		_to := "0xf54d13e2afb212da562e3ab4bfb4372c63913eb6" // contract
		_params := ""

		// GetDealParams is a free data retrieval call binding the contract method 0x62798bb9.
		// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.

		id := "a9b075aa51a2013c5b2e936d1f0a7418b195a3e49f39cfc923b4dc2bfcf1204e"
		_params = "0x62798bb9" + id
		_params = "0xf4b2e4d8a9b075aa51a2013c5b2e936d1f0a7418b195a3e49f39cfc923b4dc2bfcf1204e"

		fromEthAddr, err := ethtypes.ParseEthAddress(_from)

		toEthAddr, err := ethtypes.ParseEthAddress(_to)
		if err != nil {
			return err
		}

		params, err := ethtypes.DecodeHexString(_params)
		if err != nil {
			return err
		}

		res, err := api.EthCall(ctx, ethtypes.EthCall{
			From: &fromEthAddr,
			To:   &toEthAddr,
			Data: params,
		}, "latest")
		if err != nil {
			fmt.Println("Eth call fails, return val: ", res)
			return err
		}

		fmt.Println("Result: ", hex.EncodeToString(res))

		return nil

	},
}
