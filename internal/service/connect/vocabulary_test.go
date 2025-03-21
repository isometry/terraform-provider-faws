// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect_test

import (
	"context"
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/connect/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfconnect "github.com/isometry/terraform-provider-faws/internal/service/connect"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccVocabulary_basic(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	var v awstypes.Vocabulary
	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")

	content := "Phrase\tIPA\tSoundsLike\tDisplayAs\nLos-Angeles\t\t\tLos Angeles\nF.B.I.\tɛ f b i aɪ\t\tFBI\nEtienne\t\teh-tee-en\t"
	languageCode := "en-US"

	resourceName := "aws_connect_vocabulary.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ConnectServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVocabularyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyConfig_basic(rName, rName2, content, languageCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARNFormat(ctx, resourceName, names.AttrARN, "connect", "instance/{instance_id}/vocabulary/{vocabulary_id}"),
					resource.TestCheckResourceAttr(resourceName, names.AttrContent, content),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrInstanceID, "aws_connect_instance.test", names.AttrID),
					resource.TestCheckResourceAttr(resourceName, names.AttrLanguageCode, languageCode),
					resource.TestCheckResourceAttrSet(resourceName, "last_modified_time"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName2),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrState),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "Value1"),
					resource.TestCheckResourceAttrSet(resourceName, "vocabulary_id"),
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

func testAccVocabulary_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	var v awstypes.Vocabulary
	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")

	content := "Phrase\tIPA\tSoundsLike\tDisplayAs\nLos-Angeles\t\t\tLos Angeles\nF.B.I.\tɛ f b i aɪ\t\tFBI\nEtienne\t\teh-tee-en\t"
	languageCode := "en-US"

	resourceName := "aws_connect_vocabulary.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ConnectServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVocabularyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyConfig_basic(rName, rName2, content, languageCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfconnect.ResourceVocabulary(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccVocabulary_updateTags(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	var v awstypes.Vocabulary
	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")

	content := "Phrase\tIPA\tSoundsLike\tDisplayAs\nLos-Angeles\t\t\tLos Angeles\nF.B.I.\tɛ f b i aɪ\t\tFBI\nEtienne\t\teh-tee-en\t"
	languageCode := "en-US"

	resourceName := "aws_connect_vocabulary.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ConnectServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVocabularyDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyConfig_basic(rName, rName2, content, languageCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "Value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVocabularyConfig_tags(rName, rName2, content, languageCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "Value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "Value2a"),
				),
			},
			{
				Config: testAccVocabularyConfig_tagsUpdate(rName, rName2, content, languageCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "Value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "Value2b"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key3", "Value3"),
				),
			},
		},
	})
}

func testAccCheckVocabularyExists(ctx context.Context, n string, v *awstypes.Vocabulary) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ConnectClient(ctx)

		output, err := tfconnect.FindVocabularyByTwoPartKey(ctx, conn, rs.Primary.Attributes[names.AttrInstanceID], rs.Primary.Attributes["vocabulary_id"])

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckVocabularyDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_connect_vocabulary" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).ConnectClient(ctx)

			_, err := tfconnect.FindVocabularyByTwoPartKey(ctx, conn, rs.Primary.Attributes[names.AttrInstanceID], rs.Primary.Attributes["vocabulary_id"])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Connect Vocabulary %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccVocabularyConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_connect_instance" "test" {
  identity_management_type = "CONNECT_MANAGED"
  inbound_calls_enabled    = true
  instance_alias           = %[1]q
  outbound_calls_enabled   = true
}
`, rName)
}

func testAccVocabularyConfig_basic(rName, rName2, content, languageCode string) string {
	return acctest.ConfigCompose(
		testAccVocabularyConfig_base(rName),
		fmt.Sprintf(`
resource "aws_connect_vocabulary" "test" {
  instance_id   = aws_connect_instance.test.id
  name          = %[1]q
  content       = %[2]q
  language_code = %[3]q

  tags = {
    "Key1" = "Value1"
  }
}
`, rName2, content, languageCode))
}

func testAccVocabularyConfig_tags(rName, rName2, content, languageCode string) string {
	return acctest.ConfigCompose(
		testAccVocabularyConfig_base(rName),
		fmt.Sprintf(`
resource "aws_connect_vocabulary" "test" {
  instance_id   = aws_connect_instance.test.id
  name          = %[1]q
  content       = %[2]q
  language_code = %[3]q

  tags = {
    "Key1" = "Value1"
    "Key2" = "Value2a"
  }
}
`, rName2, content, languageCode))
}

func testAccVocabularyConfig_tagsUpdate(rName, rName2, content, languageCode string) string {
	return acctest.ConfigCompose(
		testAccVocabularyConfig_base(rName),
		fmt.Sprintf(`
resource "aws_connect_vocabulary" "test" {
  instance_id   = aws_connect_instance.test.id
  name          = %[1]q
  content       = %[2]q
  language_code = %[3]q

  tags = {
    "Key1" = "Value1"
    "Key2" = "Value2b"
    "Key3" = "Value3"
  }
}
`, rName2, content, languageCode))
}
