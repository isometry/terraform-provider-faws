// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53resolver_test

import (
	"context"
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfroute53resolver "github.com/isometry/terraform-provider-faws/internal/service/route53resolver"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccRoute53ResolverFirewallConfig_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.FirewallConfig
	resourceName := "aws_route53_resolver_firewall_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53ResolverServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckFirewallConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallConfigConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirewallConfigExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "firewall_fail_open", "ENABLED"),
					acctest.CheckResourceAttrAccountID(ctx, resourceName, names.AttrOwnerID),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRoute53ResolverFirewallConfig_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.FirewallConfig
	resourceName := "aws_route53_resolver_firewall_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53ResolverServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckFirewallConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallConfigConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirewallConfigExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfroute53resolver.ResourceFirewallConfig(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckFirewallConfigDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53ResolverClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route53_resolver_firewall_config" {
				continue
			}

			config, err := tfroute53resolver.FindFirewallConfigByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			if config.FirewallFailOpen == awstypes.FirewallFailOpenStatusDisabled {
				return nil
			}

			return fmt.Errorf("Route53 Resolver Firewall Config still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckFirewallConfigExists(ctx context.Context, n string, v *awstypes.FirewallConfig) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Route53 Resolver Firewall Config ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53ResolverClient(ctx)

		output, err := tfroute53resolver.FindFirewallConfigByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccFirewallConfigConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 0), `
resource "aws_route53_resolver_firewall_config" "test" {
  resource_id        = aws_vpc.test.id
  firewall_fail_open = "ENABLED"
}
`)
}
