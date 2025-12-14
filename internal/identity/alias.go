package identity

import (
	idApi "github.com/zodimo/go-compose/pkg/compose-identifier/api"
	idModels "github.com/zodimo/go-compose/pkg/compose-identifier/models"
)

type apiIdentityManager = *idModels.IdentityManager

// compose-identifier.api.Identifier
type Identifier = idApi.Identifier // Public API of the composer

var getOrCreateIdentityManager = idModels.GetOrCreateIdentityManager
