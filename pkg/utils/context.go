package utils

import (
	"context"
)

type ContextKey string

const (
	PreloadersKey ContextKey = "preloads"
)

func LoadersFromContext(ctx context.Context) []string {
	if loaders, ok := ctx.Value(PreloadersKey).([]string); ok {
		return loaders
	}

	return []string{}
}
