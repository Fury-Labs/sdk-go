package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fury-Labs/sdk-go/client/common"
	exchangeclient "github.com/Fury-Labs/sdk-go/client/exchange"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	accountAddress := "kai14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	res, err := exchangeClient.GetPortfolio(ctx, accountAddress)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
