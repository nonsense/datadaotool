package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/sha3"

	"encoding/json"

	cliutil "github.com/filecoin-project/lotus/cli/util"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/gateway"

	lcli "github.com/filecoin-project/lotus/cli"
)

var (
	topicHash = paddedEthHash(ethTopicHash("DealProposalCreate(bytes32)"))
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Starts an event monitor",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		ctx := context.Background()

		subCh := gateway.NewEthSubHandler()

		api, closer, err := lcli.GetFullNodeAPIV1(cctx, cliutil.FullNodeWithEthSubscribtionHandler(subCh))
		if err != nil {
			return err
		}
		defer closer()

		var topicSpec ethtypes.EthTopicSpec
		topicSpec = append(topicSpec, ethtypes.EthHashList{topicHash})

		subParam, err := json.Marshal(ethtypes.EthSubscribeParams{
			EventType: "logs",
			Params:    &ethtypes.EthSubscriptionParams{Topics: topicSpec},
		})

		fmt.Println("topicHash: ", topicHash)
		fmt.Println("calling eth subscribe with subParam: ", string(subParam))
		subID, err := api.EthSubscribe(ctx, subParam)
		if err != nil {
			return err
		}

		responseCh := make(chan ethtypes.EthSubscriptionResponse, 1)

		err = subCh.AddSub(ctx, subID, func(ctx context.Context, resp *ethtypes.EthSubscriptionResponse) error {
			responseCh <- *resp
			return nil
		})

		for resp := range responseCh {
			fmt.Println("event sub response triggered")
			spew.Dump(resp)
		}

		return nil

	},
}

func paddedEthHash(orig []byte) ethtypes.EthHash {
	if len(orig) > 32 {
		panic("exceeds EthHash length")
	}
	var ret ethtypes.EthHash
	needed := 32 - len(orig)
	copy(ret[needed:], orig)
	return ret
}

func ethTopicHash(sig string) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(sig))
	return hasher.Sum(nil)
}
