package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fury-Labs/sdk-go/client/common"
	exchangeclient "github.com/Fury-Labs/sdk-go/client/exchange"
	accountPB "github.com/Fury-Labs/sdk-go/exchange/accounts_rpc/pb"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	accountAddress := "kai1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e"
	epoch := int64(1)

	req := accountPB.RewardsRequest{
		Epoch:          epoch,
		AccountAddress: accountAddress,
	}

	res, err := exchangeClient.GetRewards(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
