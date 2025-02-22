// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadiness_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfroute53recoveryreadiness "github.com/isometry/terraform-provider-faws/internal/service/route53recoveryreadiness"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccRoute53RecoveryReadinessRecoveryGroup_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoveryreadiness_recovery_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryReadinessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecoveryGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecoveryGroupConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "route53-recovery-readiness", regexache.MustCompile(`recovery-group/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cells.#", "0"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
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

func TestAccRoute53RecoveryReadinessRecoveryGroup_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoveryreadiness_recovery_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryReadinessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecoveryGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecoveryGroupConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfroute53recoveryreadiness.ResourceRecoveryGroup(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRoute53RecoveryReadinessRecoveryGroup_nestedCell(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNameCell := sdkacctest.RandomWithPrefix("tf-acc-test-cell")
	resourceName := "aws_route53recoveryreadiness_recovery_group.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryReadinessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecoveryGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecoveryGroupConfig_andCell(rName, rNameCell),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "route53-recovery-readiness", regexache.MustCompile(`recovery-group/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cells.#", "1"),
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

func TestAccRoute53RecoveryReadinessRecoveryGroup_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoveryreadiness_recovery_group.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryReadinessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecoveryGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecoveryGroupConfig_tags1(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRecoveryGroupConfig_tags2(rName, acctest.CtKey1, acctest.CtValue1Updated, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1Updated),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
			{
				Config: testAccRecoveryGroupConfig_tags1(rName, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
		},
	})
}

func TestAccRoute53RecoveryReadinessRecoveryGroup_timeout(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_route53recoveryreadiness_recovery_group.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53RecoveryReadinessServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRecoveryGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRecoveryGroupConfig_timeout(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecoveryGroupExists(ctx, resourceName),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "route53-recovery-readiness", regexache.MustCompile(`recovery-group/.+`)),
					resource.TestCheckResourceAttr(resourceName, "cells.#", "0"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
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

func testAccCheckRecoveryGroupDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53RecoveryReadinessClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route53recoveryreadiness_recovery_group" {
				continue
			}

			_, err := tfroute53recoveryreadiness.FindRecoveryGroupByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Route53 Recovery Readiness Recovery Group %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckRecoveryGroupExists(ctx context.Context, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53RecoveryReadinessClient(ctx)

		_, err := tfroute53recoveryreadiness.FindRecoveryGroupByName(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccRecoveryGroupConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53recoveryreadiness_recovery_group" "test" {
  recovery_group_name = %q
}
`, rName)
}

func testAccRecoveryGroupConfig_andCell(rName, rNameCell string) string {
	return fmt.Sprintf(`
resource "aws_route53recoveryreadiness_cell" "test" {
  cell_name = %[2]q
}

resource "aws_route53recoveryreadiness_recovery_group" "test" {
  recovery_group_name = %[1]q
  cells               = [aws_route53recoveryreadiness_cell.test.arn]
}
`, rName, rNameCell)
}

func testAccRecoveryGroupConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_route53recoveryreadiness_recovery_group" "test" {
  recovery_group_name = %[1]q
  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccRecoveryGroupConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_route53recoveryreadiness_recovery_group" "test" {
  recovery_group_name = %[1]q
  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccRecoveryGroupConfig_timeout(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53recoveryreadiness_recovery_group" "test" {
  recovery_group_name = %q

  timeouts {
    delete = "10m"
  }
}
`, rName)
}
