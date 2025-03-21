// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tags

import (
	"context"

	"github.com/isometry/terraform-provider-faws/internal/types/option"
)

// InContext represents the tagging information kept in Context.
type InContext struct {
	DefaultConfig *DefaultConfig
	IgnoreConfig  *IgnoreConfig
	// TagsIn holds tags specified in configuration. Typically this field includes any default tags and excludes system tags.
	TagsIn option.Option[KeyValueTags]
	// TagsOut holds tags returned from AWS, including any ignored or system tags.
	TagsOut option.Option[KeyValueTags]
}

// NewContext returns a Context enhanced with tagging information.
func NewContext(ctx context.Context, defaultConfig *DefaultConfig, ignoreConfig *IgnoreConfig) context.Context {
	v := InContext{
		DefaultConfig: defaultConfig,
		IgnoreConfig:  ignoreConfig,
		TagsIn:        option.None[KeyValueTags](),
		TagsOut:       option.None[KeyValueTags](),
	}

	return context.WithValue(ctx, tagKey, &v)
}

func FromContext(ctx context.Context) (*InContext, bool) {
	v, ok := ctx.Value(tagKey).(*InContext)
	return v, ok
}

type keyType int

var tagKey keyType
