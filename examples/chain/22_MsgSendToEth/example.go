package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Fury-Labs/sdk-go/client/common"

	peggytypes "github.com/Fury-Labs/sdk-go/chain/peggy/types"
	chainclient "github.com/Fury-Labs/sdk-go/client/chain"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
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

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	ethDest := "0xaf79152ac5df276d9a8e1e2e22822f9713474902"

	amount := sdktypes.Coin{
		Denom: "kai", Amount: sdktypes.NewInt(5000000000000000000), // 5 KAI
	}
	bridgeFee := sdktypes.Coin{
		Denom: "kai", Amount: sdktypes.NewInt(2000000000000000000), // 2 KAI
	}

	msg := &peggytypes.MsgSendToEth{
		Sender:    senderAddress.String(),
		Amount:    amount,
		EthDest:   ethDest,
		BridgeFee: bridgeFee,
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
