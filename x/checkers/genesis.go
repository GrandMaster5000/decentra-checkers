package checkers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tolik22869/checkers/x/checkers/keeper"
	"github.com/tolik22869/checkers/x/checkers/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextGame != nil {
		k.SetNextGame(ctx, *genState.NextGame)
	}
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, elem)
	}
	// Set all the playerInfo
	for _, elem := range genState.PlayerInfoList {
		k.SetPlayerInfo(ctx, elem)
	}
	// Set if defined
	if genState.Leaderboard != nil {
		k.SetLeaderboard(ctx, *genState.Leaderboard)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextGame
	nextGame, found := k.GetNextGame(ctx)
	if found {
		genesis.NextGame = &nextGame
	}
	genesis.StoredGameList = k.GetAllStoredGame(ctx)
	genesis.PlayerInfoList = k.GetAllPlayerInfo(ctx)
	// Get all leaderboard
	leaderboard, found := k.GetLeaderboard(ctx)
	if found {
		genesis.Leaderboard = &leaderboard
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
