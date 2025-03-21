// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billing_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	tfmeta "github.com/isometry/terraform-provider-faws/internal/service/meta"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccBillingServiceAccountDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_billing_service_account.test"
	billingAccountID := "386209384616"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, tfmeta.PseudoServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceAccountDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, names.AttrID, billingAccountID),
					acctest.CheckResourceAttrGlobalARNAccountID(dataSourceName, names.AttrARN, billingAccountID, "iam", "root"),
				),
			},
		},
	})
}

const testAccServiceAccountDataSourceConfig_basic = `
data "aws_billing_service_account" "test" {}
`
