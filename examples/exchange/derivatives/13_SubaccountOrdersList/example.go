package main

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	subaccountId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"

	req := derivativeExchangePB.SubaccountOrdersListRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}

	res, err := exchangeClient.GetSubaccountDerivativeOrdersList(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
