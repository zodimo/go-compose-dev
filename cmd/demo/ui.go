package main

import (
	"go-compose-dev/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {
	c.StartBlock("Main")

	return c.Build()

}
