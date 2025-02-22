// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fms_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/create"
	"github.com/isometry/terraform-provider-faws/internal/errs"
	tffms "github.com/isometry/terraform-provider-faws/internal/service/fms"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccFMSResourceSet_basic(t *testing.T) {
	ctx := acctest.Context(t)

	var resourceset fms.GetResourceSetOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_fms_resource_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckRegion(t, endpoints.UsEast1RegionID)
			acctest.PreCheckOrganizationsEnabled(ctx, t)
			acctest.PreCheckOrganizationManagementAccount(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.FMSServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckResourceSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSetConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceSetExists(ctx, resourceName, &resourceset),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrID),
					resource.TestCheckResourceAttrSet(resourceName, "resource_set.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "resource_set.0.resource_set_status"),
				),
			},
		},
	})
}

func testAccFMSResourceSet_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var resourceset fms.GetResourceSetOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_fms_resource_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckRegion(t, endpoints.UsEast1RegionID)
			acctest.PreCheckOrganizationsEnabled(ctx, t)
			acctest.PreCheckOrganizationManagementAccount(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.FMSServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckResourceSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSetConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceSetExists(ctx, resourceName, &resourceset),
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tffms.ResourceSet, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckResourceSetDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).FMSClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_fms_resource_set" {
				continue
			}

			_, err := tffms.FindResourceSetByID(ctx, conn, rs.Primary.ID)
			if errs.IsA[*types.ResourceNotFoundException](err) {
				return nil
			}
			if err != nil {
				return err
			}

			return create.Error(names.BCMDataExports, create.ErrActionCheckingDestroyed, tffms.ResNameResourceSet, rs.Primary.ID, errors.New("not destroyed"))
		}

		return nil
	}
}

func testAccCheckResourceSetExists(ctx context.Context, name string, resourceset *fms.GetResourceSetOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.FMS, create.ErrActionCheckingExistence, tffms.ResNameResourceSet, name, errors.New("not found"))
		}

		if rs.Primary.ID == "" {
			return create.Error(names.FMS, create.ErrActionCheckingExistence, tffms.ResNameResourceSet, name, errors.New("not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).FMSClient(ctx)
		resp, err := tffms.FindResourceSetByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return create.Error(names.FMS, create.ErrActionCheckingExistence, tffms.ResNameResourceSet, rs.Primary.ID, err)
		}

		*resourceset = *resp

		return nil
	}
}

func testAccResourceSetConfig_basic(rName string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_fms_admin_account" "test" {
  account_id = data.aws_caller_identity.current.account_id
}

resource "aws_fms_resource_set" "test" {
  depends_on = [aws_fms_admin_account.test]
  resource_set {
    name               = %[1]q
    resource_type_list = ["AWS::NetworkFirewall::Firewall"]
  }
}
`, rName)
}
