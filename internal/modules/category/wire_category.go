package category

import "github.com/google/wire"

var CategorySet = wire.NewSet(
	NewRepositoryCategory,
	NewServiceCategory,
	NewHandlerCategory,
)
