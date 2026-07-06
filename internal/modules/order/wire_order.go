package order

import "github.com/google/wire"

var OrderSet = wire.NewSet(
	NewRepositoryOrder,
)
