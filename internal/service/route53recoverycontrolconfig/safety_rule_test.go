// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolconfig_test

import (
	"context"
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfroute53recoverycontrolconfig "github.com/isometry/terraform-provider-faws/internal/service/route53recoverycontrolconfig"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccSafetyRule_assertionRule(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoverycontrolconfig_safety_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.Route53RecoveryControlConfigEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryControlConfigServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSafetyRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSafetyRuleConfig_routingControlAssertion(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSafetyRuleExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, names.AttrStatus, "DEPLOYED"),
					resource.TestCheckResourceAttr(resourceName, "wait_period_ms", "5000"),
					resource.TestCheckResourceAttr(resourceName, "asserted_controls.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "control_panel_arn", "aws_route53recoverycontrolconfig_control_panel.test", names.AttrARN),
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

func testAccSafetyRule_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoverycontrolconfig_safety_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.Route53RecoveryControlConfigEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryControlConfigServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSafetyRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSafetyRuleConfig_routingControlAssertion(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSafetyRuleExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfroute53recoverycontrolconfig.ResourceSafetyRule(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccSafetyRule_gatingRule(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoverycontrolconfig_safety_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.Route53RecoveryControlConfigEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryControlConfigServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSafetyRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSafetyRuleConfig_routingControlGating(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSafetyRuleExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, names.AttrStatus, "DEPLOYED"),
					resource.TestCheckResourceAttr(resourceName, "wait_period_ms", "5000"),
					resource.TestCheckResourceAttr(resourceName, "target_controls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gating_controls.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "control_panel_arn", "aws_route53recoverycontrolconfig_control_panel.test", names.AttrARN),
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

func testAccCheckSafetyRuleDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53RecoveryControlConfigClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route53recoverycontrolconfig_safety_rule" {
				continue
			}

			_, err := tfroute53recoverycontrolconfig.FindSafetyRuleByARN(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Route53RecoveryControlConfig Safety Rule (%s) not deleted", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckSafetyRuleExists(ctx context.Context, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53RecoveryControlConfigClient(ctx)

		_, err := tfroute53recoverycontrolconfig.FindSafetyRuleByARN(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccSafetyRuleConfig_routingControlAssertion(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53recoverycontrolconfig_cluster" "test" {
  name = %[1]q
}

resource "aws_route53recoverycontrolconfig_control_panel" "test" {
  name        = %[1]q
  cluster_arn = aws_route53recoverycontrolconfig_cluster.test.arn
}

resource "aws_route53recoverycontrolconfig_routing_control" "test" {
  name              = %[1]q
  cluster_arn       = aws_route53recoverycontrolconfig_cluster.test.arn
  control_panel_arn = aws_route53recoverycontrolconfig_control_panel.test.arn
}

resource "aws_route53recoverycontrolconfig_safety_rule" "test" {
  name              = %[1]q
  control_panel_arn = aws_route53recoverycontrolconfig_control_panel.test.arn
  wait_period_ms    = 5000
  asserted_controls = [aws_route53recoverycontrolconfig_routing_control.test.arn]

  rule_config {
    inverted  = false
    threshold = 0
    type      = "AND"
  }
}
`, rName)
}

func testAccSafetyRuleConfig_routingControlGating(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53recoverycontrolconfig_cluster" "test" {
  name = %[1]q
}

resource "aws_route53recoverycontrolconfig_control_panel" "test" {
  name        = %[1]q
  cluster_arn = aws_route53recoverycontrolconfig_cluster.test.arn
}

resource "aws_route53recoverycontrolconfig_routing_control" "test" {
  name              = %[1]q
  cluster_arn       = aws_route53recoverycontrolconfig_cluster.test.arn
  control_panel_arn = aws_route53recoverycontrolconfig_control_panel.test.arn
}

resource "aws_route53recoverycontrolconfig_safety_rule" "test" {
  name              = %[1]q
  control_panel_arn = aws_route53recoverycontrolconfig_control_panel.test.arn
  wait_period_ms    = 5000
  gating_controls   = [aws_route53recoverycontrolconfig_routing_control.test.arn]
  target_controls   = [aws_route53recoverycontrolconfig_routing_control.test.arn]

  rule_config {
    inverted  = false
    threshold = 0
    type      = "AND"
  }
}
`, rName)
}
