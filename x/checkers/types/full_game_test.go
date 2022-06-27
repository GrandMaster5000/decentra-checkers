package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	rules "github.com/tolik22869/checkers/x/checkers/rules"
	"github.com/tolik22869/checkers/x/checkers/types"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *types.StoredGame {
	return &types.StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "b",
	}
}

func TestCanGetAddressCreator(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	creator, err2 := GetStoredGame1().GetCreatorAddress()
	require.Equal(t, aliceAddress, creator)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongCreator(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Creator = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	creator, err := storedGame.GetCreatorAddress()
	require.Nil(t, creator)
	require.EqualError(t,
		err,
		"creator address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: checksum failed. Expected 3xn9d3, got 3xn9d4.")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressBlack(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, bobAddress, black)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h"
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: checksum failed. Expected xqhc8g, got xqhc8h.")
	require.EqualError(t, storedGame.Validate(), err.Error())
}
