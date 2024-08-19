package keeper

type Querier struct {
	Keeper
}

func NewQuerierImpl(keeper Keeper) *Querier {
	return &Querier{
		keeper,
	}
}
