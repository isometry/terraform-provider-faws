// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package licensemanager_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/envvar"
	tflicensemanager "github.com/isometry/terraform-provider-faws/internal/service/licensemanager"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccGrantAccepter_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	licenseARN := envvar.SkipIfEmpty(t, licenseARNKey, envVarLicenseARNKeyError)
	principal := envvar.SkipIfEmpty(t, principalKey, envVarPrincipalKeyError)
	homeRegion := envvar.SkipIfEmpty(t, homeRegionKey, envVarHomeRegionError)
	resourceName := "aws_licensemanager_grant_accepter.test"
	resourceGrantName := "aws_licensemanager_grant.test"

	providers := make(map[string]*schema.Provider)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LicenseManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesNamedAlternate(ctx, t, providers),
		CheckDestroy:             acctest.CheckWithNamedProviders(testAccCheckGrantAccepterDestroyWithProvider(ctx), providers),
		Steps: []resource.TestStep{
			{
				Config: testAccGrantAccepterConfig_basic(licenseARN, rName, principal, homeRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckGrantAccepterExists(ctx, resourceName, acctest.NamedProviderFunc(acctest.ProviderName, providers)),
					resource.TestCheckResourceAttrPair(resourceName, "grant_arn", resourceGrantName, names.AttrARN),
					resource.TestCheckResourceAttrSet(resourceName, "allowed_operations.0"),
					resource.TestCheckResourceAttrPair(resourceName, "home_region", resourceGrantName, "home_region"),
					resource.TestCheckResourceAttr(resourceName, "license_arn", licenseARN),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrName, resourceGrantName, names.AttrName),
					resource.TestCheckResourceAttrPair(resourceName, "parent_arn", resourceGrantName, "parent_arn"),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrPrincipal, resourceGrantName, names.AttrPrincipal),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrStatus),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrVersion),
				),
			},
			{
				Config:            testAccGrantAccepterConfig_basic(licenseARN, rName, principal, homeRegion),
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGrantAccepter_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	licenseARN := envvar.SkipIfEmpty(t, licenseARNKey, envVarLicenseARNKeyError)
	principal := envvar.SkipIfEmpty(t, principalKey, envVarPrincipalKeyError)
	homeRegion := envvar.SkipIfEmpty(t, homeRegionKey, envVarHomeRegionError)
	resourceName := "aws_licensemanager_grant_accepter.test"

	providers := make(map[string]*schema.Provider)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LicenseManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesNamedAlternate(ctx, t, providers),
		CheckDestroy:             acctest.CheckWithNamedProviders(testAccCheckGrantAccepterDestroyWithProvider(ctx), providers),
		Steps: []resource.TestStep{
			{
				Config: testAccGrantAccepterConfig_basic(licenseARN, rName, principal, homeRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckGrantAccepterExists(ctx, resourceName, acctest.NamedProviderFunc(acctest.ProviderName, providers)),
					acctest.CheckResourceDisappears(ctx, acctest.NamedProvider(acctest.ProviderName, providers), tflicensemanager.ResourceGrantAccepter(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckGrantAccepterExists(ctx context.Context, n string, providerF func() *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := providerF().Meta().(*conns.AWSClient).LicenseManagerClient(ctx)

		_, err := tflicensemanager.FindReceivedGrantByARN(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccCheckGrantAccepterDestroyWithProvider(ctx context.Context) acctest.TestCheckWithProviderFunc {
	return func(s *terraform.State, provider *schema.Provider) error {
		conn := provider.Meta().(*conns.AWSClient).LicenseManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_licensemanager_grant_accepter" {
				continue
			}

			_, err := tflicensemanager.FindReceivedGrantByARN(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("License Manager Grant Accepter %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccGrantAccepterConfig_basic(licenseARN, rName, principal, homeRegion string) string {
	principalArn, _ := arn.Parse(principal)
	roleARN := arn.ARN{
		Partition: principalArn.Partition,
		Service:   "iam",
		AccountID: principalArn.AccountID,
		Resource:  "role/OrganizationAccountAccessRole",
	}
	return acctest.ConfigCompose(
		acctest.ConfigNamedRegionalProvider(acctest.ProviderNameAlternate, homeRegion),
		fmt.Sprintf(`
provider %[1]q {
	assume_role {
		role_arn = %[2]q
	}
}`, acctest.ProviderName, roleARN),
		fmt.Sprintf(`
resource "aws_licensemanager_grant_accepter" "test" {
  grant_arn = aws_licensemanager_grant.test.arn
}

data "aws_licensemanager_received_license" "test" {
  provider    = awsalternate
  license_arn = %[1]q
}

locals {
  allowed_operations = [for i in data.aws_licensemanager_received_license.test.received_metadata[0].allowed_operations : i if i != "CreateGrant"]
}

resource "aws_licensemanager_grant" "test" {
  provider = awsalternate

  name               = %[2]q
  allowed_operations = local.allowed_operations
  license_arn        = data.aws_licensemanager_received_license.test.license_arn
  principal          = %[3]q
}
`, licenseARN, rName, principal),
	)
}
