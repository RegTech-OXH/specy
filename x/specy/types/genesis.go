package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ExecutorList: []Executor{},
		TaskList:     []Task{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in executor
	executorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ExecutorList {
		index := string(ExecutorKey(elem.Address))
		if _, ok := executorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for executor")
		}
		executorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in task
	taskIndexMap := make(map[string]struct{})

	for _, elem := range gs.TaskList {
		index := string(TaskKey(elem.TaskHash))
		if _, ok := taskIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for task")
		}
		taskIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
