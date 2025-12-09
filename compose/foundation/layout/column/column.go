package column

import "go-compose-dev/pkg/api"

func Column(content api.Composable, options ...ColumnOption) api.Composable {
	return func(c api.Composer) api.Composer {
		return c
	}
}
