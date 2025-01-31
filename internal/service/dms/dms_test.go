// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dms_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func init() {
	acctest.RegisterServiceErrorCheckFunc(names.DMSServiceID, testAccErrorCheckSkip)
}

// testAccErrorCheckSkip skips DMS tests that have error messages indicating unsupported features
func testAccErrorCheckSkip(t *testing.T) resource.ErrorCheckFunc {
	return acctest.ErrorCheckSkipMessagesContaining(t,
		// Serverless DMS in GovCloud
		"SERVERLESS feature is not available",
	)
}
