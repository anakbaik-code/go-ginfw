//go:build wireinject
// +build wireinject

package bootstrap

import (
	"database/sql"
	"go-fwgin/internal/config"
	"go-fwgin/internal/database"
	"go-fwgin/internal/modules/category"
	"go-fwgin/internal/modules/event"
	"go-fwgin/internal/modules/user"

	"github.com/google/wire"
)

func InitializeApp() (*App, func(), error) {
	wire.Build(
		// DB provides
		config.LoadConfig,
		config.NewMySQL,
		database.New,
		wire.Bind(new(database.DBTX), new(*sql.DB)),

		// Route Group
		user.UserSet,
		category.CategorySet,
		event.EventSet,

		wire.Struct(new(App), "*"),
	)
	return nil, nil, nil
}
