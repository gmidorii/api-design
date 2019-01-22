//+build wireinject

package provider

import (
	"github.com/google/wire"

	"github.com/gmidorii/api-design/wire/app"
	"github.com/gmidorii/api-design/wire/repository/db"
)

func InitHelloApp() (app.Hello, error) {
	wire.Build(db.NewUser, app.NewHello)
	return app.Hello{}, nil
}