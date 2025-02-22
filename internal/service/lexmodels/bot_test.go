// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexmodels_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tflexmodels "github.com/isometry/terraform-provider-faws/internal/service/lexmodels"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func init() {
	acctest.RegisterServiceErrorCheckFunc(names.LexModelsServiceID, testAccErrorCheckSkip)
}

func testAccErrorCheckSkip(t *testing.T) resource.ErrorCheckFunc {
	return acctest.ErrorCheckSkipMessagesContaining(t,
		"You can't set the enableModelImprovements field to false",
	)
}

func TestAccLexModelsBot_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					testAccCheckBotNotExists(ctx, testBotID, "1"),

					resource.TestCheckResourceAttr(rName, "abort_statement.#", "1"),
					resource.TestCheckResourceAttrSet(rName, names.AttrARN),
					resource.TestCheckResourceAttrSet(rName, "checksum"),
					resource.TestCheckResourceAttr(rName, "child_directed", acctest.CtFalse),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.#", "0"),
					resource.TestCheckResourceAttr(rName, "create_version", acctest.CtFalse),
					acctest.CheckResourceAttrRFC3339(rName, names.AttrCreatedDate),
					resource.TestCheckResourceAttr(rName, names.AttrDescription, "Bot to order flowers on the behalf of a user"),
					resource.TestCheckResourceAttr(rName, "detect_sentiment", acctest.CtFalse),
					resource.TestCheckResourceAttr(rName, "enable_model_improvements", acctest.CtFalse),
					resource.TestCheckResourceAttr(rName, "failure_reason", ""),
					resource.TestCheckResourceAttr(rName, "idle_session_ttl_in_seconds", "300"),
					resource.TestCheckResourceAttr(rName, "intent.#", "1"),
					acctest.CheckResourceAttrRFC3339(rName, names.AttrLastUpdatedDate),
					resource.TestCheckResourceAttr(rName, "locale", "en-US"),
					resource.TestCheckResourceAttr(rName, names.AttrName, testBotID),
					resource.TestCheckResourceAttr(rName, "nlu_intent_confidence_threshold", "0"),
					resource.TestCheckResourceAttr(rName, "process_behavior", "SAVE"),
					resource.TestCheckResourceAttr(rName, names.AttrStatus, "NOT_BUILT"),
					resource.TestCheckResourceAttr(rName, names.AttrVersion, tflexmodels.BotVersionLatest),
					resource.TestCheckResourceAttr(rName, "voice_id", ""),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_Version_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]func(t *testing.T){
		"LexBot_createVersion":         testAccBot_createVersion,
		"LexBotAlias_botVersion":       testAccBotAlias_botVersion,
		"DataSourceLexBot_withVersion": testAccBotDataSource_withVersion,
		"DataSourceLexBotAlias_basic":  testAccBotAliasDataSource_basic,
	}

	acctest.RunSerialTests1Level(t, testCases, 0)
}

func testAccBot_createVersion(t *testing.T) {
	ctx := acctest.Context(t)
	var v1, v2 lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	// If this test runs in parallel with other Lex Bot tests, it loses its description
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v1),
					testAccCheckBotNotExists(ctx, testBotID, "1"),
					resource.TestCheckResourceAttr(rName, names.AttrVersion, tflexmodels.BotVersionLatest),
					resource.TestCheckResourceAttr(rName, names.AttrDescription, "Bot to order flowers on the behalf of a user"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_createVersion(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExistsWithVersion(ctx, rName, "1", &v2),
					resource.TestCheckResourceAttr(rName, names.AttrVersion, "1"),
					resource.TestCheckResourceAttr(rName, names.AttrDescription, "Bot to order flowers on the behalf of a user"),
				),
			},
		},
	})
}

func TestAccLexModelsBot_abortStatement(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_abortStatement(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "abort_statement.#", "1"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.#", "1"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.0.content", "Sorry, I'm not able to assist at this time"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.0.content_type", "PlainText"),
					resource.TestCheckNoResourceAttr(rName, "abort_statement.0.message.0.group_number"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.response_card", ""),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_abortStatementUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.#", "2"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.0.content", "Sorry, I'm not able to assist at this time"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.0.content_type", "PlainText"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.0.group_number", "1"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.1.content", "Sorry, I'm not able to assist at this time. Good bye."),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.1.content_type", "PlainText"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.message.1.group_number", "1"),
					resource.TestCheckResourceAttr(rName, "abort_statement.0.response_card", "Sorry, I'm not able to assist at this time"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_clarificationPrompt(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_clarificationPrompt(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.#", "1"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.max_attempts", "2"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.message.#", "1"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.message.#", "1"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.message.0.content", "I didn't understand you, what would you like to do?"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.message.0.content_type", "PlainText"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.response_card", ""),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"abort_statement.0.message.0.group_number",
					"clarification_prompt.0.message.0.group_number",
				},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_clarificationPromptUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.max_attempts", "3"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.message.#", "2"),
					resource.TestCheckResourceAttr(rName, "clarification_prompt.0.response_card", "I didn't understand you, what would you like to do?"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"abort_statement.0.message.0.group_number",
					"clarification_prompt.0.message.0.group_number",
				},
			},
		},
	})
}

func TestAccLexModelsBot_childDirected(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_childDirectedUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "child_directed", acctest.CtTrue),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_description(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_descriptionUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, names.AttrDescription, "Bot to order flowers"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_detectSentiment(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_detectSentimentUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "detect_sentiment", acctest.CtTrue),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_enableModelImprovements(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_enableModelImprovementsUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "enable_model_improvements", acctest.CtTrue),
					resource.TestCheckResourceAttr(rName, "nlu_intent_confidence_threshold", "0.5"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_idleSessionTTLInSeconds(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_idleSessionTTLInSecondsUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "idle_session_ttl_in_seconds", "600"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_intents(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intentMultiple(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intentMultiple(testBotID),
					testAccBotConfig_intentsUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "intent.#", "2"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_computeVersion(t *testing.T) {
	ctx := acctest.Context(t)
	var v1 lexmodelbuildingservice.GetBotOutput
	var v2 lexmodelbuildingservice.GetBotAliasOutput

	botResourceName := "aws_lex_bot.test"
	botAliasResourceName := "aws_lex_bot_alias.test"
	intentResourceName := "aws_lex_intent.test"
	intentResourceName2 := "aws_lex_intent.test_2"

	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_createVersion(testBotID),
					testAccBotConfig_intentMultiple(testBotID),
					testAccBotAliasConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExistsWithVersion(ctx, botResourceName, "1", &v1),
					resource.TestCheckResourceAttr(botResourceName, names.AttrVersion, "1"),
					resource.TestCheckResourceAttr(botResourceName, "intent.#", "1"),
					resource.TestCheckResourceAttr(botResourceName, "intent.0.intent_version", "1"),
					testAccCheckBotAliasExists(ctx, botAliasResourceName, &v2),
					resource.TestCheckResourceAttr(botAliasResourceName, "bot_version", "1"),
					resource.TestCheckResourceAttr(intentResourceName, names.AttrVersion, "1"),
				),
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intentMultipleSecondUpdated(testBotID),
					testAccBotConfig_multipleIntentsWithVersion(testBotID),
					testAccBotAliasConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExistsWithVersion(ctx, botResourceName, "2", &v1),
					resource.TestCheckResourceAttr(botResourceName, names.AttrVersion, "2"),
					resource.TestCheckResourceAttr(botResourceName, "intent.#", "2"),
					resource.TestCheckResourceAttr(botResourceName, "intent.0.intent_version", "1"),
					resource.TestCheckResourceAttr(botResourceName, "intent.1.intent_version", "2"),
					resource.TestCheckResourceAttr(botAliasResourceName, "bot_version", "2"),
					resource.TestCheckResourceAttr(intentResourceName, names.AttrVersion, "1"),
					resource.TestCheckResourceAttr(intentResourceName2, names.AttrVersion, "2"),
				),
			},
		},
	})
}

func TestAccLexModelsBot_locale(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_localeUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "locale", "en-GB"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_voiceID(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_voiceIdUpdate(testBotID),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					resource.TestCheckResourceAttr(rName, "voice_id", "Justin"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"abort_statement.0.message.0.group_number"},
			},
		},
	})
}

func TestAccLexModelsBot_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v lexmodelbuildingservice.GetBotOutput
	rName := "aws_lex_bot.test"
	testBotID := "test_bot_" + sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.LexModelBuildingServiceEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.LexModelsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBotDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					testAccBotConfig_intent(testBotID),
					testAccBotConfig_basic(testBotID),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotExists(ctx, rName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tflexmodels.ResourceBot(), rName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckBotExistsWithVersion(ctx context.Context, rName, botVersion string, v *lexmodelbuildingservice.GetBotOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rName]
		if !ok {
			return fmt.Errorf("Not found: %s", rName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Lex Bot ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).LexModelsClient(ctx)

		output, err := tflexmodels.FindBotVersionByName(ctx, conn, rs.Primary.ID, botVersion)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckBotExists(ctx context.Context, rName string, output *lexmodelbuildingservice.GetBotOutput) resource.TestCheckFunc {
	return testAccCheckBotExistsWithVersion(ctx, rName, tflexmodels.BotVersionLatest, output)
}

func testAccCheckBotNotExists(ctx context.Context, botName, botVersion string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).LexModelsClient(ctx)

		_, err := tflexmodels.FindBotVersionByName(ctx, conn, botName, botVersion)

		if tfresource.NotFound(err) {
			return nil
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("Lex Box %s/%s still exists", botName, botVersion)
	}
}

func testAccCheckBotDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).LexModelsClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_lex_bot" {
				continue
			}

			output, err := conn.GetBotVersions(ctx, &lexmodelbuildingservice.GetBotVersionsInput{
				Name: aws.String(rs.Primary.ID),
			})

			if err != nil {
				return err
			}

			if output == nil || len(output.Bots) == 0 {
				return nil
			}

			return fmt.Errorf("Lex Bot %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccBotConfig_intent(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_intent" "test" {
  create_version = true
  name           = "%s"
  fulfillment_activity {
    type = "ReturnIntent"
  }
  sample_utterances = [
    "I would like to pick up flowers",
  ]
}
`, rName)
}

func testAccBotConfig_intentMultiple(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_intent" "test" {
  create_version = true
  name           = "%[1]s"
  fulfillment_activity {
    type = "ReturnIntent"
  }
  sample_utterances = [
    "I would like to pick up flowers",
  ]
}

resource "aws_lex_intent" "test_2" {
  create_version = true
  name           = "%[1]stwo"
  fulfillment_activity {
    type = "ReturnIntent"
  }
  sample_utterances = [
    "I would like to pick up flowers",
  ]
}
`, rName)
}

func testAccBotConfig_intentMultipleSecondUpdated(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_intent" "test" {
  create_version = true
  name           = "%[1]s"
  fulfillment_activity {
    type = "ReturnIntent"
  }
  sample_utterances = [
    "I would like to pick up flowers",
  ]
}

resource "aws_lex_intent" "test_2" {
  create_version = true
  name           = "%[1]stwo"
  fulfillment_activity {
    type = "ReturnIntent"
  }
  sample_utterances = [
    "I would like to return these flowers",
  ]
}
`, rName)
}

func testAccBotConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_createVersion(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed   = false
  create_version   = true
  description      = "Bot to order flowers on the behalf of a user"
  name             = "%s"
  process_behavior = "BUILD"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_abortStatement(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_abortStatementUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
      group_number = 1
    }
    message {
      content      = "Sorry, I'm not able to assist at this time. Good bye."
      content_type = "PlainText"
      group_number = 1
    }
    response_card = "Sorry, I'm not able to assist at this time"
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_clarificationPrompt(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  clarification_prompt {
    max_attempts = 2
    message {
      content      = "I didn't understand you, what would you like to do?"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_clarificationPromptUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  clarification_prompt {
    max_attempts = 3
    message {
      content      = "I didn't understand you, what would you like to do?"
      content_type = "PlainText"
      group_number = 1
    }
    message {
      content      = "I didn't understand you, can you re-phrase your request, please?"
      content_type = "PlainText"
      group_number = 1
    }
    response_card = "I didn't understand you, what would you like to do?"
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_childDirectedUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = true
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_descriptionUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_detectSentimentUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed   = false
  description      = "Bot to order flowers on the behalf of a user"
  detect_sentiment = true
  name             = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_enableModelImprovementsUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed                  = false
  description                     = "Bot to order flowers on the behalf of a user"
  enable_model_improvements       = true
  name                            = "%s"
  nlu_intent_confidence_threshold = 0.5
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_idleSessionTTLInSecondsUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed              = false
  description                 = "Bot to order flowers on the behalf of a user"
  idle_session_ttl_in_seconds = 600
  name                        = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_intentsUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
  intent {
    intent_name    = aws_lex_intent.test_2.name
    intent_version = aws_lex_intent.test_2.version
  }
}
`, rName)
}

func testAccBotConfig_localeUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed            = false
  description               = "Bot to order flowers on the behalf of a user"
  enable_model_improvements = true
  locale                    = "en-GB"
  name                      = "%s"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_voiceIdUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed = false
  description    = "Bot to order flowers on the behalf of a user"
  name           = "%s"
  voice_id       = "Justin"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
}
`, rName)
}

func testAccBotConfig_multipleIntentsWithVersion(rName string) string {
	return fmt.Sprintf(`
resource "aws_lex_bot" "test" {
  child_directed   = false
  create_version   = true
  description      = "Bot to order flowers on the behalf of a user"
  name             = "%s"
  process_behavior = "BUILD"
  abort_statement {
    message {
      content      = "Sorry, I'm not able to assist at this time"
      content_type = "PlainText"
    }
  }
  intent {
    intent_name    = aws_lex_intent.test.name
    intent_version = aws_lex_intent.test.version
  }
  intent {
    intent_name    = aws_lex_intent.test_2.name
    intent_version = aws_lex_intent.test_2.version
  }
}
`, rName)
}
