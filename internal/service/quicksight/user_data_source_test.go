// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksight_test

import (
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	tfquicksight "github.com/isometry/terraform-provider-faws/internal/service/quicksight"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccQuickSightUserDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_quicksight_user.test"
	dataSourceName := "data.aws_quicksight_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.QuickSightServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrUserName, resourceName, names.AttrUserName),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(dataSourceName, names.AttrEmail, acctest.DefaultEmailAddress),
					resource.TestCheckResourceAttr(dataSourceName, names.AttrNamespace, tfquicksight.DefaultUserNamespace),
					resource.TestCheckResourceAttr(dataSourceName, "identity_type", string(awstypes.IdentityTypeQuicksight)),
					resource.TestCheckResourceAttrSet(dataSourceName, "principal_id"),
					resource.TestCheckResourceAttr(dataSourceName, "user_role", string(awstypes.UserRoleReader)),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_quicksight_user" "test" {
  user_name     = %[1]q
  email         = %[2]q
  identity_type = "QUICKSIGHT"
  user_role     = "READER"
}

data "aws_quicksight_user" "test" {
  user_name = aws_quicksight_user.test.user_name
}
`, rName, acctest.DefaultEmailAddress)
}
