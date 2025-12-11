package identity

import (
	idApi "go-compose-dev/pkg/compose-identifier/api"
	idModels "go-compose-dev/pkg/compose-identifier/models"
)

type apiIdentityManager = *idModels.IdentityManager

// compose-identifier.api.Identifier
type Identifier = idApi.Identifier // Public API of the composer

var getOrCreateIdentityManager = idModels.GetOrCreateIdentityManager
