// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batch_test

import (
	"context"
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/batch/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfbatch "github.com/isometry/terraform-provider-faws/internal/service/batch"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccBatchSchedulingPolicy_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var schedulingPolicy1 awstypes.SchedulingPolicyDetail
	resourceName := "aws_batch_scheduling_policy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BatchServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSchedulingPolicyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSchedulingPolicyConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSchedulingPolicyExists(ctx, resourceName, &schedulingPolicy1),
					acctest.CheckResourceAttrRegionalARNFormat(ctx, resourceName, names.AttrARN, "batch", "scheduling-policy/{name}"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.compute_reservation", "1"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.share_decay_seconds", "3600"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.share_distribution.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// add one more share_distribution block
				Config: testAccSchedulingPolicyConfig_basic2(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSchedulingPolicyExists(ctx, resourceName, &schedulingPolicy1),
					acctest.CheckResourceAttrRegionalARNFormat(ctx, resourceName, names.AttrARN, "batch", "scheduling-policy/{name}"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.compute_reservation", "1"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.share_decay_seconds", "3600"),
					resource.TestCheckResourceAttr(resourceName, "fair_share_policy.0.share_distribution.#", "2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
				),
			},
		},
	})
}

func TestAccBatchSchedulingPolicy_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var schedulingPolicy1 awstypes.SchedulingPolicyDetail
	resourceName := "aws_batch_scheduling_policy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BatchServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSchedulingPolicyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSchedulingPolicyConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSchedulingPolicyExists(ctx, resourceName, &schedulingPolicy1),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfbatch.ResourceSchedulingPolicy(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckSchedulingPolicyExists(ctx context.Context, n string, v *awstypes.SchedulingPolicyDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).BatchClient(ctx)

		output, err := tfbatch.FindSchedulingPolicyByARN(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckSchedulingPolicyDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_batch_scheduling_policy" {
				continue
			}
			conn := acctest.Provider.Meta().(*conns.AWSClient).BatchClient(ctx)

			_, err := tfbatch.FindSchedulingPolicyByARN(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Batch Scheduling Policy %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccSchedulingPolicyConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_batch_scheduling_policy" "test" {
  name = %[1]q

  fair_share_policy {
    compute_reservation = 1
    share_decay_seconds = 3600

    share_distribution {
      share_identifier = "A1*"
      weight_factor    = 0.1
    }
  }

  tags = {
    "Name" = "Test Batch Scheduling Policy"
  }
}
`, rName)
}

func testAccSchedulingPolicyConfig_basic2(rName string) string {
	return fmt.Sprintf(`
resource "aws_batch_scheduling_policy" "test" {
  name = %[1]q

  fair_share_policy {
    compute_reservation = 1
    share_decay_seconds = 3600

    share_distribution {
      share_identifier = "A1*"
      weight_factor    = 0.1
    }

    share_distribution {
      share_identifier = "A2"
      weight_factor    = 0.2
    }
  }

  tags = {
    "Name" = "Test Batch Scheduling Policy"
  }
}
`, rName)
}
