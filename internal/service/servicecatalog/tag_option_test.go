// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalog_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	awstypes "github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/errs"
	tfservicecatalog "github.com/isometry/terraform-provider-faws/internal/service/servicecatalog"
	"github.com/isometry/terraform-provider-faws/names"
)

// add sweeper to delete known test servicecat tag options

func TestAccServiceCatalogTagOption_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_servicecatalog_tag_option.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagOptionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccTagOptionConfig_basic(rName, "värde", "active = true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagOptionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde"),
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

func TestAccServiceCatalogTagOption_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_servicecatalog_tag_option.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagOptionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccTagOptionConfig_basic(rName, "värde", "active = true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagOptionExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfservicecatalog.ResourceTagOption(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccServiceCatalogTagOption_update(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_servicecatalog_tag_option.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	// UpdateTagOption() is very particular about what it receives. Only fields that change should
	// be included or it will throw servicecatalog.ErrCodeDuplicateResourceException, "already exists"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagOptionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccTagOptionConfig_basic(rName, "värde ett", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde ett"),
				),
			},
			{
				Config: testAccTagOptionConfig_basic(rName, "värde två", "active = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde två"),
				),
			},
			{
				Config: testAccTagOptionConfig_basic(rName, "värde två", "active = false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtFalse),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName), // cannot be updated in place
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde två"),
				),
			},
			{
				Config: testAccTagOptionConfig_basic(rName, "värde två", "active = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName), // cannot be updated in place
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde två"),
				),
			},
			{
				Config: testAccTagOptionConfig_basic(rName2, "värde ett", "active = true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName2),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde ett"),
				),
			},
		},
	})
}

func TestAccServiceCatalogTagOption_notActive(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_servicecatalog_tag_option.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagOptionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccTagOptionConfig_basic(rName, "värde ett", "active = false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "active", acctest.CtFalse),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, rName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrOwner),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "värde ett"),
				),
			},
		},
	})
}

func testAccCheckTagOptionDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).ServiceCatalogClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_servicecatalog_tag_option" {
				continue
			}

			input := &servicecatalog.DescribeTagOptionInput{
				Id: aws.String(rs.Primary.ID),
			}

			output, err := conn.DescribeTagOption(ctx, input)

			if errs.IsA[*awstypes.ResourceNotFoundException](err) {
				continue
			}

			if err != nil {
				return fmt.Errorf("error getting Service Catalog Tag Option (%s): %w", rs.Primary.ID, err)
			}

			if output != nil {
				return fmt.Errorf("Service Catalog Tag Option (%s) still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccCheckTagOptionExists(ctx context.Context, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ServiceCatalogClient(ctx)

		input := &servicecatalog.DescribeTagOptionInput{
			Id: aws.String(rs.Primary.ID),
		}

		_, err := conn.DescribeTagOption(ctx, input)

		if err != nil {
			return fmt.Errorf("error describing Service Catalog Tag Option (%s): %w", rs.Primary.ID, err)
		}

		return nil
	}
}

func testAccTagOptionConfig_basic(key, value, active string) string {
	return fmt.Sprintf(`
resource "aws_servicecatalog_tag_option" "test" {
  key   = %[1]q
  value = %[2]q
  %[3]s
}
`, key, value, active)
}
