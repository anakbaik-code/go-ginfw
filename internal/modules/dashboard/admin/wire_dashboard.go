package admin

import (
	"github.com/google/wire"
)

var DashboardAdminSet = wire.NewSet(
	NewRepositoryDashboardAdmin,
	NewServiceDashboardAdmin,
	NewDashboardHandler,
)
