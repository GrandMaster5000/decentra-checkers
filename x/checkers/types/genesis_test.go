package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tolik22869/checkers/x/checkers/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				NextGame: &types.NextGame{
					IdValue: 25,
				},
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				PlayerInfoList: []types.PlayerInfo{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				Leaderboard: &types.Leaderboard{
					Winners: []*types.WinningPlayer{},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedGame",
			genState: &types.GenesisState{
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated playerInfo",
			genState: &types.GenesisState{
				PlayerInfoList: []types.PlayerInfo{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredGameList: []types.StoredGame{},
			NextGame:       &types.NextGame{"", uint64(1), "-1", "-1"},
		},
		types.DefaultGenesis())
}
