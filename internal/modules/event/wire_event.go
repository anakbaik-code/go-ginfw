package event

import "github.com/google/wire"

var EventSet = wire.NewSet(
	NewRepositoryEvent,
	NewServiceEvent,
	NewHandlerEvent,
)
