package main

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/pkg/api"
)

type Composable = api.Composable
type Composer = api.Composer
type Modifier = ui.Modifier

func UI() Composable {
	return func(c Composer) Composer {
		// return textfield.BasicTextField()
		return nil
	}

}
