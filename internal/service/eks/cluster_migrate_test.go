// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eks_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	tfeks "github.com/isometry/terraform-provider-faws/internal/service/eks"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestClusterStateUpgradeV0(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	testCases := []struct {
		testName string
		rawState map[string]interface{}
		want     map[string]interface{}
	}{
		{
			testName: "empty state",
			rawState: map[string]interface{}{},
			want: map[string]interface{}{
				"bootstrap_self_managed_addons": acctest.CtTrue,
			},
		},
		{
			testName: "non-empty state",
			rawState: map[string]interface{}{
				names.AttrName:    "testing",
				names.AttrVersion: "1.1.0",
			},
			want: map[string]interface{}{
				"bootstrap_self_managed_addons": acctest.CtTrue,
				names.AttrName:                  "testing",
				names.AttrVersion:               "1.1.0",
			},
		},
		{
			testName: "bootstrap_self_managed_addons set",
			rawState: map[string]interface{}{
				"bootstrap_self_managed_addons": acctest.CtFalse,
				names.AttrName:                  "testing",
				names.AttrVersion:               "1.1.0",
			},
			want: map[string]interface{}{
				"bootstrap_self_managed_addons": acctest.CtFalse,
				names.AttrName:                  "testing",
				names.AttrVersion:               "1.1.0",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()

			got, err := tfeks.ClusterStateUpgradeV0(ctx, testCase.rawState, nil)

			if err != nil {
				t.Errorf("err = %q", err)
			} else if diff := cmp.Diff(got, testCase.want); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
