package specy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"specy/x/specy/keeper"
	"specy/x/specy/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the executor
	for _, elem := range genState.ExecutorList {
		k.SetExecutor(ctx, elem)
	}
	// Set all the task
	for _, elem := range genState.TaskList {
		k.SetTask(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ExecutorList = k.GetAllExecutor(ctx)
	genesis.TaskList = k.GetAllTask(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
