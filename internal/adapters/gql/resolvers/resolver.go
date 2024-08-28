package graph

import (
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SystemFieldService *ports.ISystemFieldService
	TransactorRepo     *database.IDatabaseTransactor
}
