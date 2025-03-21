// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoidp_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccCognitoIDPUserGroupsDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_cognito_user_groups.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckIdentityProvider(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.CognitoIDPServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccUserGroupsDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "groups.#", "2"),
				),
			},
		},
	})
}

func testAccUserGroupsDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_cognito_user_pool" "test" {
  name = %q
}

resource "aws_cognito_user_group" "test_1" {
  name         = "%s-1"
  user_pool_id = aws_cognito_user_pool.test.id
  description  = "test 1"
}
resource "aws_cognito_user_group" "test_2" {
  name         = "%s-2"
  user_pool_id = aws_cognito_user_pool.test.id
  description  = "test 2"
}

data "aws_cognito_user_groups" "test" {
  user_pool_id = aws_cognito_user_group.test_1.user_pool_id
}
`, rName, rName, rName)
}
