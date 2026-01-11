package dialog

import (
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/pkg/api"
)

func TextContent(content string) api.Composable {
	return text.BodyMedium(content)
}
