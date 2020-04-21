package reputation

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdkTypes.Querier {
	return func(ctx sdkTypes.Context, path []string, req abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {

		default:
			return nil, sdkTypes.ErrUnknownRequest("unknown bank query endpoint")
		}
	}
}