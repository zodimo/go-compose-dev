package api

import (
	"go-compose-dev/pkg/compose-identifier/models"
)

func NewIdentifier() Identifier {
	return models.NewIdentifier()
}

func NewScopedIdentifier(scope string) Identifier {
	return models.NewScopedIdentifier(scope)
}
