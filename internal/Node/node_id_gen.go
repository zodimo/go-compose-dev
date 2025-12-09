package node

import idApi "go-compose-dev/pkg/compose-identifier/api"

func NewNodeID() NodeID {
	return idApi.NewIdentifier()
}
