package v1tov2

import "github.com/tolik22869/checkers/x/checkers/types"

type GenesisStateV1 struct {
	Params         types.Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	NextGame       *types.NextGame    `protobuf:"bytes,2,opt,name=nextGame,proto3" json:"nextGame,omitempty"`
	StoredGameList []types.StoredGame `protobuf:"bytes,3,rep,name=storedGameList,proto3" json:"storedGameList"`
}
