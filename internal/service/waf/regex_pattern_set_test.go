// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waf_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	awstypes "github.com/aws/aws-sdk-go-v2/service/waf/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfwaf "github.com/isometry/terraform-provider-faws/internal/service/waf"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

// Serialized acceptance tests due to WAF account limits
// https://docs.aws.amazon.com/waf/latest/developerguide/limits.html
func TestAccWAFRegexPatternSet_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]func(t *testing.T){
		acctest.CtBasic:      testAccRegexPatternSet_basic,
		"changePatterns":     testAccRegexPatternSet_changePatterns,
		"noPatterns":         testAccRegexPatternSet_noPatterns,
		acctest.CtDisappears: testAccRegexPatternSet_disappears,
	}

	acctest.RunSerialTests1Level(t, testCases, 0)
}

func testAccRegexPatternSet_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.RegexPatternSet
	patternSetName := fmt.Sprintf("tfacc-%s", sdkacctest.RandString(5))
	resourceName := "aws_waf_regex_pattern_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WAFServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRegexPatternSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRegexPatternSetConfig_basic(patternSetName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegexPatternSetExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "waf", regexache.MustCompile(`regexpatternset/.+`)),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, patternSetName),
					resource.TestCheckResourceAttr(resourceName, "regex_pattern_strings.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "one"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "two"),
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

func testAccRegexPatternSet_changePatterns(t *testing.T) {
	ctx := acctest.Context(t)
	var before, after awstypes.RegexPatternSet
	patternSetName := fmt.Sprintf("tfacc-%s", sdkacctest.RandString(5))
	resourceName := "aws_waf_regex_pattern_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WAFServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRegexPatternSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRegexPatternSetConfig_basic(patternSetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckRegexPatternSetExists(ctx, resourceName, &before),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, patternSetName),
					resource.TestCheckResourceAttr(resourceName, "regex_pattern_strings.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "one"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "two"),
				),
			},
			{
				Config: testAccRegexPatternSetConfig_changes(patternSetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckRegexPatternSetExists(ctx, resourceName, &after),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, patternSetName),
					resource.TestCheckResourceAttr(resourceName, "regex_pattern_strings.#", "3"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "two"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "three"),
					resource.TestCheckTypeSetElemAttr(resourceName, "regex_pattern_strings.*", "four"),
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

func testAccRegexPatternSet_noPatterns(t *testing.T) {
	ctx := acctest.Context(t)
	var patternSet awstypes.RegexPatternSet
	patternSetName := fmt.Sprintf("tfacc-%s", sdkacctest.RandString(5))
	resourceName := "aws_waf_regex_pattern_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WAFServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRegexPatternSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRegexPatternSetConfig_nos(patternSetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckRegexPatternSetExists(ctx, resourceName, &patternSet),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, patternSetName),
					resource.TestCheckResourceAttr(resourceName, "regex_pattern_strings.#", "0"),
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

func testAccRegexPatternSet_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.RegexPatternSet
	patternSetName := fmt.Sprintf("tfacc-%s", sdkacctest.RandString(5))
	resourceName := "aws_waf_regex_pattern_set.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.WAFServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRegexPatternSetDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRegexPatternSetConfig_basic(patternSetName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegexPatternSetExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfwaf.ResourceRegexPatternSet(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckRegexPatternSetExists(ctx context.Context, n string, v *awstypes.RegexPatternSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).WAFClient(ctx)

		output, err := tfwaf.FindRegexPatternSetByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckRegexPatternSetDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_waf_regex_pattern_set" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).WAFClient(ctx)

			_, err := tfwaf.FindRegexPatternSetByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("WAF Regex Pattern Set %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccRegexPatternSetConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "aws_waf_regex_pattern_set" "test" {
  name                  = %[1]q
  regex_pattern_strings = ["one", "two"]
}
`, name)
}

func testAccRegexPatternSetConfig_changes(name string) string {
	return fmt.Sprintf(`
resource "aws_waf_regex_pattern_set" "test" {
  name                  = %[1]q
  regex_pattern_strings = ["two", "three", "four"]
}
`, name)
}

func testAccRegexPatternSetConfig_nos(name string) string {
	return fmt.Sprintf(`
resource "aws_waf_regex_pattern_set" "test" {
  name = %[1]q
}
`, name)
}
