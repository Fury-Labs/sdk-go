package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fury-Labs/sdk-go/client/common"
	exchangeclient "github.com/Fury-Labs/sdk-go/client/exchange"
	spotExchangePB "github.com/Fury-Labs/sdk-go/exchange/spot_exchange_rpc/pb"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
	skip := uint64(0)
	limit := int32(10)

	req := spotExchangePB.OrdersRequest{
		MarketId: marketId,
		Skip:     skip,
		Limit:    limit,
	}

	res, err := exchangeClient.GetSpotOrders(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
