package client

import (
	"fmt"
	"google.golang.org/grpc/credentials"
)

type network struct {
	LcdEndpoint          string
	TmEndpoint           string
	ChainGrpcEndpoint    string
	ChainTlsCert         credentials.TransportCredentials
	ExchangeGrpcEndpoint string
	ExchangeTlsCert      credentials.TransportCredentials
	ChainId              string
	Fee_denom            string
	Name                 string
}

func LoadNetwork(name string, node string) network {
	if name == "devnet" {
		return network{
			LcdEndpoint:          "https://devnet.lcd.injective.dev",
			TmEndpoint:           "https://devnet.tm.injective.dev:443",
			ChainGrpcEndpoint:    "tcp://devnet.injective.dev:9900",
			ExchangeGrpcEndpoint: "tcp://devnet.injective.dev:9910",
			ChainId:              "injective-777",
			Fee_denom:            "inj",
			Name:                 "devnet",
		}
	} else if name == "testnet" {
		validNodes := []string{"sentry0", "sentry1", "k8s"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}

		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, exchangeGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert credentials.TransportCredentials
		if node == "k8s" {
			lcdEndpoint = "https://k8s.testnet.lcd.injective.network"
			tmEndpoint = "https://k8s.testnet.tm.injective.network:443"
			chainGrpcEndpoint = "tcp://k8s.testnet.chain.grpc.injective.network:443"
			chainTlsCert = LoadTlsCert("client/cert/testnet.crt", chainGrpcEndpoint)
			exchangeGrpcEndpoint = "tcp://k8s.testnet.exchange.grpc.injective.network:443"
			exchangeTlsCert = LoadTlsCert("client/cert/testnet.crt", exchangeGrpcEndpoint)
		} else {
			lcdEndpoint = fmt.Sprintf("http://%s.injective.dev:10337", node)
			tmEndpoint = fmt.Sprintf("http://%s.injective.dev:26657", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.dev:9900", node)
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.dev:9910", node)
		}

		return network{
			LcdEndpoint:          lcdEndpoint,
			TmEndpoint:           tmEndpoint,
			ChainGrpcEndpoint:    chainGrpcEndpoint,
			ChainTlsCert:         chainTlsCert,
			ExchangeGrpcEndpoint: exchangeGrpcEndpoint,
			ExchangeTlsCert:      exchangeTlsCert,
			ChainId:              "injective-888",
			Fee_denom:            "inj",
			Name:                 "testnet",
		}
	} else if name == "mainnet" {
		validNodes := []string{"lb", "sentry0", "sentry1", "sentry2", "sentry3"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}

		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, exchangeGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert credentials.TransportCredentials
		if node == "lb" {
			lcdEndpoint = "https://lb.mainnet.lcd.injective.network"
			lcdEndpoint = fmt.Sprintf("https://lb.mainnet.tm.injective.network:443", node)
			chainGrpcEndpoint = "tcp://lb.mainnet.chain.grpc.injective.network:443"
			chainTlsCert = LoadTlsCert("client/cert/mainnet.crt", chainGrpcEndpoint)
			exchangeGrpcEndpoint = "tcp://lb.mainnet.exchange.grpc.injective.network:443"
			exchangeTlsCert = LoadTlsCert("client/cert/mainnet.crt", exchangeGrpcEndpoint)
		} else {
			lcdEndpoint = fmt.Sprintf("http://%s.injective.network:10337", node)
			tmEndpoint = fmt.Sprintf("http://%s.injective.network:26657", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9900", node)
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9910", node)
		}

		return network{
			LcdEndpoint:          lcdEndpoint,
			TmEndpoint:           tmEndpoint,
			ChainGrpcEndpoint:    chainGrpcEndpoint,
			ChainTlsCert:         chainTlsCert,
			ExchangeGrpcEndpoint: exchangeGrpcEndpoint,
			ExchangeTlsCert:      exchangeTlsCert,
			ChainId:              "injective-1",
			Fee_denom:            "inj",
			Name:                 "mainnet",
		}
	}

	return network{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
