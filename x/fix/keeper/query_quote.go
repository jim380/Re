package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/fix/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QuoteAll(goCtx context.Context, req *types.QueryAllQuoteRequest) (*types.QueryAllQuoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var quotes []types.Quote
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	quoteStore := prefix.NewStore(store, types.GetQuoteKey())

	pageRes, err := query.Paginate(quoteStore, req.Pagination, func(key []byte, value []byte) error {
		var quote types.Quote
		if err := k.cdc.Unmarshal(value, &quote); err != nil {
			return err
		}

		quotes = append(quotes, quote)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllQuoteResponse{Quote: quotes, Pagination: pageRes}, nil
}

func (k Keeper) Quote(goCtx context.Context, req *types.QueryGetQuoteRequest) (*types.QueryGetQuoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	quote, found := k.GetQuote(ctx, req.QuoteReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetQuoteResponse{Quote: quote}, nil
}

// Query quotes by sessionID
func (k Keeper) QuotesBySessionID(goCtx context.Context, req *types.QuerySessionByIDQuoteRequest) (*types.QuerySessionByIDQuoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// get all quotes and return for each sessionID
	var quotesBySessionID []types.Quote
	allQuote := k.GetAllQuote(ctx)
	for _, quotes := range allQuote {
		// check for if the requested sessionID matches with any sessionID returned from GetAllQuote()
		if quotes.SessionID == req.SessionID {
			quotesBySessionID = append(quotesBySessionID, quotes)
		}
	}

	return &types.QuerySessionByIDQuoteResponse{Quote: quotesBySessionID}, nil
}
