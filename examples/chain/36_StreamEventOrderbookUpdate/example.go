package main

import (
	"fmt"
	exchangetypes "github.com/Fury-Labs/sdk-go/chain/exchange/types"
	chainclient "github.com/Fury-Labs/sdk-go/client/chain"
	"github.com/Fury-Labs/sdk-go/client/common"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet", "")
	tmRPC, err := rpchttp.New("http://139.178.68.147:26657", "/websocket")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000kai"),
	)
	if err != nil {
		panic(err)
	}

	//0x74b17b0d6855feba39f1f7ab1e8bad0363bd510ee1dcc74e40c2adfe1502f781
	//0x74ee114ad750f8429a97e07b5e73e145724e9b21670a7666625ddacc03d6758d
	//0x26413a70c9b78a495023e5ab8003c9cf963ef963f6755f8b57255feb5744bf31
	marketIds := []string{
		"0x74b17b0d6855feba39f1f7ab1e8bad0363bd510ee1dcc74e40c2adfe1502f781",
	}

	orderbookCh := make(chan exchangetypes.Orderbook, 10000)
	go chainClient.StreamOrderbookUpdateEvents(chainclient.SpotOrderbook, marketIds, orderbookCh)
	for {
		ob := <-orderbookCh
		fmt.Println(ob)
	}
}
