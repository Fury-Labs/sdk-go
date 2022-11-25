all:

copy-exchange-client:
	rm -rf exchange/*
	mkdir -p exchange/meta_rpc
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/auction_rpc
	mkdir -p exchange/oracle_rpc
	mkdir -p exchange/insurance_rpc
	mkdir -p exchange/explorer_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/derivative_exchange_rpc

	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_meta_rpc/pb exchange/meta_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_auction_rpc/pb exchange/auction_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_oracle_rpc/pb exchange/oracle_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_insurance_rpc/pb exchange/insurance_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_explorer_rpc/pb exchange/explorer_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../kaiju-indexer/api/gen/grpc/kaiju_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb

.PHONY: copy-exchange-client

copy-chain-types:
	cp ../kaiju-core/kaiju-chain/modules/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/oracle/types/*.go chain/oracle/types
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/wasmx/types/*.go chain/wasmx/types
	rm -rf chain/wasmx/types/*test.go  rm -rf chain/wasmx/types/*gw.go
	cp ../kaiju-core/kaiju-chain/modules/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go

	echo "👉 Replace kaiju-core/kaiju-chain/modules with sdk-go/chain"
	echo "👉 Replace kaiju-core/kaiju-chain/types with sdk-go/chain/types"
