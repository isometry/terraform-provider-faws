// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worklink_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfworklink "github.com/isometry/terraform-provider-faws/internal/service/worklink"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccWorkLinkWebsiteCertificateAuthorityAssociation_basic(t *testing.T) {
	acctest.Skip(t, "skipping test; Amazon WorkLink has been deprecated and will be removed in the next major release")

	ctx := acctest.Context(t)
	suffix := sdkacctest.RandStringFromCharSet(20, sdkacctest.CharSetAlpha)
	resourceName := "aws_worklink_website_certificate_authority_association.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WorkLinkServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWebsiteCertificateAuthorityAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWebsiteCertificateAuthorityAssociationConfig_basic(suffix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebsiteCertificateAuthorityAssociationExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(
						resourceName, "fleet_arn",
						"aws_worklink_fleet.test", names.AttrARN),
					resource.TestMatchResourceAttr(resourceName, names.AttrCertificate, regexache.MustCompile("^-----BEGIN CERTIFICATE-----")),
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

func TestAccWorkLinkWebsiteCertificateAuthorityAssociation_displayName(t *testing.T) {
	acctest.Skip(t, "skipping test; Amazon WorkLink has been deprecated and will be removed in the next major release")

	ctx := acctest.Context(t)
	suffix := sdkacctest.RandStringFromCharSet(20, sdkacctest.CharSetAlpha)
	resourceName := "aws_worklink_website_certificate_authority_association.test"
	displayName1 := fmt.Sprintf("tf-website-certificate-%s", sdkacctest.RandStringFromCharSet(5, sdkacctest.CharSetAlpha))
	displayName2 := fmt.Sprintf("tf-website-certificate-%s", sdkacctest.RandStringFromCharSet(5, sdkacctest.CharSetAlpha))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WorkLinkServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWebsiteCertificateAuthorityAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWebsiteCertificateAuthorityAssociationConfig_displayName(suffix, displayName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebsiteCertificateAuthorityAssociationExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDisplayName, displayName1),
				),
			},
			{
				Config: testAccWebsiteCertificateAuthorityAssociationConfig_displayName(suffix, displayName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebsiteCertificateAuthorityAssociationExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDisplayName, displayName2),
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

func TestAccWorkLinkWebsiteCertificateAuthorityAssociation_disappears(t *testing.T) {
	acctest.Skip(t, "skipping test; Amazon WorkLink has been deprecated and will be removed in the next major release")

	ctx := acctest.Context(t)
	suffix := sdkacctest.RandStringFromCharSet(20, sdkacctest.CharSetAlpha)
	resourceName := "aws_worklink_website_certificate_authority_association.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WorkLinkServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckWebsiteCertificateAuthorityAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccWebsiteCertificateAuthorityAssociationConfig_basic(suffix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWebsiteCertificateAuthorityAssociationExists(ctx, resourceName),
					testAccCheckWebsiteCertificateAuthorityAssociationDisappears(ctx, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckWebsiteCertificateAuthorityAssociationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).WorkLinkClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_worklink_website_certificate_authority_association" {
				continue
			}

			_, err := tfworklink.FindWebsiteCertificateAuthorityByARNAndID(ctx, conn, rs.Primary.Attributes["fleet_arn"], rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Worklink Website Certificate Authority Association(%s) still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckWebsiteCertificateAuthorityAssociationDisappears(ctx context.Context, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).WorkLinkClient(ctx)

		fleetArn, websiteCaID, err := tfworklink.DecodeWebsiteCertificateAuthorityAssociationResourceID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = tfworklink.FindWebsiteCertificateAuthorityByARNAndID(ctx, conn, fleetArn, websiteCaID)

		if err != nil {
			return err
		}

		stateConf := &retry.StateChangeConf{
			Pending:    []string{"DELETING"},
			Target:     []string{"DELETED"},
			Refresh:    tfworklink.WebsiteCertificateAuthorityAssociationStateRefresh(ctx, conn, websiteCaID, fleetArn),
			Timeout:    15 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 3 * time.Second,
		}

		_, err = stateConf.WaitForStateContext(ctx)

		return err
	}
}

func testAccCheckWebsiteCertificateAuthorityAssociationExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worklink Website Certificate Authority Association ID is set")
		}

		if _, ok := rs.Primary.Attributes["fleet_arn"]; !ok {
			return fmt.Errorf("WorkLink Fleet ARN is missing, should be set.")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).WorkLinkClient(ctx)
		fleetArn, websiteCaID, err := tfworklink.DecodeWebsiteCertificateAuthorityAssociationResourceID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = tfworklink.FindWebsiteCertificateAuthorityByARNAndID(ctx, conn, fleetArn, websiteCaID)

		return err
	}
}

func testAccWebsiteCertificateAuthorityAssociationConfig_basic(r string) string {
	return acctest.ConfigCompose(
		testAccFleetConfig_basic(r), `
resource "aws_worklink_website_certificate_authority_association" "test" {
  fleet_arn   = aws_worklink_fleet.test.arn
  certificate = file("test-fixtures/worklink-website-certificate-authority-association.pem")
}
`)
}

func testAccWebsiteCertificateAuthorityAssociationConfig_displayName(r, displayName string) string {
	return acctest.ConfigCompose(
		testAccFleetConfig_basic(r),
		fmt.Sprintf(`
resource "aws_worklink_website_certificate_authority_association" "test" {
  fleet_arn    = aws_worklink_fleet.test.arn
  certificate  = file("test-fixtures/worklink-website-certificate-authority-association.pem")
  display_name = "%s"
}
`, displayName))
}
