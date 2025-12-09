package zipper

import (
	idModels "go-compose-dev/pkg/compose-identifier/models"
)

var IDManager IdentityManager

func init() {
	IDManager = idModels.GeOrCreateIdentityManager("zipper")
}
