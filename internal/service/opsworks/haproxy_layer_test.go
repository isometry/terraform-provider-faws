// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworks_test

import (
	"context"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccOpsWorksHAProxyLayer_basic(t *testing.T) {
	acctest.Skip(t, "skipping test; Amazon OpsWorks has been deprecated and will be removed in the next major release")

	ctx := acctest.Context(t)
	var v awstypes.Layer
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_opsworks_haproxy_layer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.OpsWorks) },
		ErrorCheck:               acctest.ErrorCheck(t, names.OpsWorksServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckHAProxyLayerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHAProxyLayerConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLayerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "healthcheck_method", "OPTIONS"),
					resource.TestCheckResourceAttr(resourceName, "healthcheck_url", "/"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, "HAProxy"),
					resource.TestCheckResourceAttr(resourceName, "stats_enabled", acctest.CtTrue),
					resource.TestCheckResourceAttrSet(resourceName, "stats_password"),
					resource.TestCheckResourceAttr(resourceName, "stats_url", "/haproxy?stats"),
					resource.TestCheckResourceAttr(resourceName, "stats_user", "opsworks"),
				),
			},
		},
	})
}

// _disappears and _tags for OpsWorks Layers are tested via aws_opsworks_rails_app_layer.

func testAccCheckHAProxyLayerDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error { return testAccCheckLayerDestroy(ctx, "aws_opsworks_haproxy_layer", s) }
}

func testAccHAProxyLayerConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccLayerConfig_base(rName), `
resource "aws_opsworks_haproxy_layer" "test" {
  stack_id       = aws_opsworks_stack.test.id
  stats_password = "avoid-plaintext-passwords"

  custom_security_group_ids = aws_security_group.test[*].id
}
`)
}
