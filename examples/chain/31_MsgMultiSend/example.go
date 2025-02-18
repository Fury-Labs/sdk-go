package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Fury-Labs/sdk-go/client/common"

	chainclient "github.com/Fury-Labs/sdk-go/client/chain"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.kaijud",
		"kaijud",
		"file",
		"kai-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	// initialize grpc client

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	// prepare tx msg

	msg := &banktypes.MsgMultiSend{
		Inputs: []banktypes.Input{
			{
				Address: senderAddress.String(),
				Coins: []sdktypes.Coin{{
					Denom: "kai", Amount: sdktypes.NewInt(1000000000000000000)}, // 1 KAI
				},
			},
			{
				Address: senderAddress.String(),
				Coins: []sdktypes.Coin{{
					Denom: "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5", Amount: sdktypes.NewInt(1000000)}, // 1 USDT
				},
			},
		},
		Outputs: []banktypes.Output{
			{
				Address: "kai1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
				Coins: []sdktypes.Coin{{
					Denom: "kai", Amount: sdktypes.NewInt(1000000000000000000)}, // 1 KAI
				},
			},
			{
				Address: "kai1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
				Coins: []sdktypes.Coin{{
					Denom: "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5", Amount: sdktypes.NewInt(1000000)}, // 1 USDT
				},
			},
		},
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000kai"),
	)

	if err != nil {
		fmt.Println(err)
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee, "KAI")
}
