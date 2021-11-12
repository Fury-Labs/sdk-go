all:

copy-exchange-client:
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/derivative_exchange_rpc

	cp -r ../injective-exchange/api/gen/grpc/injective_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb

.PHONY: copy-exchange-client

copy-chain-types:
	cp ../injective-core/injective-chain/modules/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../injective-core/injective-chain/modules/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../injective-core/injective-chain/modules/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../injective-core/injective-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../injective-core/injective-chain/modules/oracle/types/*.go chain/oracle/types
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../injective-core/injective-chain/modules/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	echo "👉 Replace injective-core/injective-chain/modules with sdk-go/chain"
	echo "👉 Replace injective-core/injective-chain/types with sdk-go/chain/types"
