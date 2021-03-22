package types

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

const (
	// module name
	ModuleName = "exchange"

	// StoreKey to be used when creating the KVStore
	StoreKey  = ModuleName
	TStoreKey = "transient_exchange"
)
const SpotPriceDecimalPlaces = 18

var (
	// Keys for store prefixes
	DepositKey                = []byte{0x01} // prefix for each key to a Deposit
	TECSeenHashStoreKeyPrefix = []byte{0x02} // prefix for the seen state of a TEC transaction hash

	SpotMarketKey                       = []byte{0x11} // prefix for each key to a spot market by (isEnabled, marketID)
	SpotLimitOrdersKey                  = []byte{0x12} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotMarketOrdersKey                 = []byte{0x13} // prefix for each key to a spot order, by (marketID, direction, price level, order hash)
	SpotLimitOrdersBySubaccountIndexKey = []byte{0x14} // prefix for each key to a spot order index, by (marketID, direction, subaccountID, order hash)
	SpotMarketOrderIndicatorKey         = []byte{0x15} // prefix for each key to a spot market order indicator, by marketID and direction
	SpotOrderHashKey                    = []byte{0x16} // prefix for each key for an archived spot order hash

	DerivativeMarketKey                       = []byte{0x21} // prefix for each key to a derivative market by (exchange address, isEnabled, marketID)
	DerivativeLimitOrdersKey                  = []byte{0x22} // prefix for each key to an derivative limit order, by (marketID, direction, price level, order hash)
	DerivativeMarketOrdersKey                 = []byte{0x23} // prefix for each key to a derivative order, by (marketID, direction, price level, order hash)
	DerivativeLimitOrdersBySubaccountIndexKey = []byte{0x24} // prefix for each key to a derivative order index, by (marketID, direction, subaccountID, order hash)
	DerivativeMarketOrderIndicatorKey         = []byte{0x25} // prefix for each key to a derivative market order indicator, by marketID and direction
	DerivativePositionKey                     = []byte{0x26} // prefix for each key to a Position

	DerivativeOrderHashKey = []byte{0x26} // prefix for each key for a derivative order hash

)

// GetDepositKey provides the key to obtain a given subaccount's deposits for a given denom
func GetDepositKey(subaccountID common.Hash, denom string) []byte {
	return append(subaccountID.Bytes(), []byte(denom)...)
}

func GetSpotMarketKey(isEnabled bool) []byte {
	return append(SpotMarketKey, enabledPrefix(isEnabled)...)
}

func GetSpotMarketTransientMarketsKeyPrefix(marketId common.Hash, isBuy bool) []byte {
	return append(SpotMarketKey, MarketDirectionPrefix(marketId, isBuy)...)
}

// TODO: properly compute this without using decimal.Decimal
func getPaddedPrice(price sdk.Dec) string {
	priceString := price.String()
	temp, _ := decimal.NewFromString(priceString)
	priceString = temp.StringFixed(SpotPriceDecimalPlaces)
	leftSide := fmt.Sprintf("%032s", temp.Floor().String())
	priceComponents := strings.Split(priceString, ".")
	return leftSide + "." + priceComponents[1]
}

func GetSpotLimitOrderHashFromKey(spotLimitOrderKey []byte) common.Hash {
	return common.BytesToHash(spotLimitOrderKey[len(spotLimitOrderKey)-common.HashLength:])
}

func GetLimitOrderByPriceKeyPrefix(marketId common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return GetOrderByPriceKeyPrefix(marketId, isBuy, price, orderHash)
}

func GetLimitOrderBySubaccountKey(marketId common.Hash, isBuy bool, subaccountId, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketId, isBuy), subaccountId.Bytes()...), orderHash.Bytes()...)
}

func GetOrderByPriceKeyPrefix(marketId common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketId, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

// SpotMarketDirectionPriceHashPrefix turns a marketID + direction + price + order hash to prefix used to get a spot order from the store.
func SpotMarketDirectionPriceHashPrefix(marketID common.Hash, isBuy bool, price sdk.Dec, orderHash common.Hash) []byte {
	return append(append(MarketDirectionPrefix(marketID, isBuy), []byte(getPaddedPrice(price))...), orderHash.Bytes()...)
}

func GetDerivativeMarketKey(isEnabled bool) []byte {
	return append(DerivativeMarketKey, enabledPrefix(isEnabled)...)
}

// TODO @albert rename obviously
func enabledPrefix(isEnabled bool) []byte {
	isEnabledByte := byte(0)
	if isEnabled {
		isEnabledByte = byte(1)
	}
	return []byte{isEnabledByte}
}

// OrdersByMarketDirectionPriceOrderHashPrefix turns a marketID + direction + price + order hash to prefix used to get an order from the store.
func OrdersByMarketDirectionPriceOrderHashPrefix(marketID common.Hash, orderHash common.Hash, price *big.Int, isLong bool) []byte {
	return append(ordersByMarketDirectionPricePrefix(marketID, price, isLong), orderHash.Bytes()...)
}

// orderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of exchange against a particular marketID, direction and price
func ordersByMarketDirectionPricePrefix(marketID common.Hash, price *big.Int, isLong bool) []byte {
	return append(MarketDirectionPrefix(marketID, isLong), common.LeftPadBytes(price.Bytes(), 32)...)
}

// OrderIndexByMarketDirectionSubaccountOrderHashPrefix turns a marketID + direction + subaccountID + order hash to prefix used to get an order from the store.
func OrderIndexByMarketDirectionSubaccountOrderHashPrefix(marketID common.Hash, isLong bool, subaccountID common.Hash, orderHash common.Hash) []byte {
	return append(OrderIndexByMarketDirectionSubaccountPrefix(marketID, subaccountID, isLong), orderHash.Bytes()...)
}

// OrderIndexByMarketDirectionSubaccountPrefix allows to obtain prefix of exchange against a particular marketID, subaccountID and direction
func OrderIndexByMarketDirectionSubaccountPrefix(marketID common.Hash, subaccountID common.Hash, isLong bool) []byte {
	return append(MarketDirectionPrefix(marketID, isLong), subaccountID.Bytes()...)
}

// MarketDirectionHashPrefix allows to obtain prefix against a particular marketID, direction, order hash
func MarketDirectionHashPrefix(marketID common.Hash, isLong bool, hash common.Hash) []byte {
	direction := byte(0)
	if isLong {
		direction = byte(1)
	}
	return append(append(marketID.Bytes(), direction), hash.Bytes()...)
}

// MarketDirectionPrefix allows to obtain prefix against a particular marketID, direction
func MarketDirectionPrefix(marketID common.Hash, isLong bool) []byte {
	direction := byte(0)
	if isLong {
		direction = byte(1)
	}
	return append(marketID.Bytes(), direction)
}

func GetMarketIdDirectionFromKey(key []byte) (marketId common.Hash, isBuy bool) {
	marketID := common.BytesToHash(key[len(SpotMarketKey) : common.HashLength+len(SpotMarketKey)])
	isBuyBytes := key[common.HashLength+len(SpotMarketKey):]
	return marketID, bytes.Equal(isBuyBytes, []byte{byte(0)})
}

// MarginInfoByExchangeSubaccountBaseCurrencyPrefix provides the prefix key to obtain a given subaccount's base currency
// margin info in a given exchange
func MarginInfoByExchangeSubaccountBaseCurrencyPrefix(subaccountID common.Hash, baseCurrencyAddress common.Address) []byte {
	return append(MarginInfoByExchangeSubaccountPrefix(subaccountID), baseCurrencyAddress.Bytes()...)
}

// MarginInfoByExchangeSubaccountPrefix provides the prefix key to obtain a given subaccount's margin info in a given exchange
func MarginInfoByExchangeSubaccountPrefix(subaccountID common.Hash) []byte {
	return subaccountID.Bytes()
}

// PositionBySubaccountMarketIDPrefix provides the prefix key to obtain a position for a given subaccount in a given market
func PositionBySubaccountMarketIDPrefix(subaccountID common.Hash, marketID common.Hash) []byte {
	return append(subaccountID.Bytes(), marketID.Bytes()...)
}

func GetSubaccountAndMarketIDFromPositionKey(key []byte) (subaccountID, marketID common.Hash) {
	keyOffsetLen := len(DerivativePositionKey)
	subaccountOffsetLen := keyOffsetLen + common.HashLength
	subaccountID = common.BytesToHash(key[keyOffsetLen:subaccountOffsetLen])
	marketID = common.BytesToHash(key[subaccountOffsetLen : subaccountOffsetLen+common.HashLength])
	return subaccountID, marketID
}

// SpotMarketsStoreKey turns a pair hash to key used to get it from the store.
func SpotMarketsStoreKey(hash common.Hash) []byte {
	return append(SpotMarketKey, hash.Bytes()...)
}
