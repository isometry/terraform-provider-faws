// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"github.com/isometry/terraform-provider-faws/internal/conns"
)

var (
	_ WithMeta = (*withMeta)(nil)
)

type WithMeta interface {
	Meta() *conns.AWSClient
}

type withMeta struct {
	meta *conns.AWSClient
}

// Meta returns the provider Meta (instance data).
func (w *withMeta) Meta() *conns.AWSClient {
	return w.meta
}
