// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/macie2/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/errs"
	tfmacie2 "github.com/isometry/terraform-provider-faws/internal/service/macie2"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccClassificationJob_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_nameGenerated(bucketName, string(awstypes.JobTypeOneTime)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					acctest.CheckResourceAttrNameGenerated(resourceName, names.AttrName),
					resource.TestCheckResourceAttr(resourceName, names.AttrNamePrefix, "terraform-"),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeOneTime)),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_definitions.0.buckets.0", bucketName),
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

func testAccClassificationJob_Name_Generated(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_nameGenerated(bucketName, string(awstypes.JobTypeOneTime)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					acctest.CheckResourceAttrNameGenerated(resourceName, names.AttrName),
					resource.TestCheckResourceAttr(resourceName, names.AttrNamePrefix, "terraform-"),
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

func testAccClassificationJob_NamePrefix(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	namePrefix := "tf-acc-test-prefix-"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_namePrefix(bucketName, namePrefix, string(awstypes.JobTypeOneTime)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					acctest.CheckResourceAttrNameFromPrefix(resourceName, names.AttrName, namePrefix),
					resource.TestCheckResourceAttr(resourceName, names.AttrNamePrefix, namePrefix),
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

func testAccClassificationJob_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_nameGenerated(bucketName, string(awstypes.JobTypeOneTime)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfmacie2.ResourceClassificationJob(), resourceName),
				),
			},
		},
	})
}

func testAccClassificationJob_Status(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output, macie2Output2 macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_status(bucketName, string(awstypes.JobStatusRunning)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_definitions.0.buckets.0", bucketName),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
				),
			},
			{
				Config: testAccClassificationJobConfig_status(bucketName, string(awstypes.JobStatusUserPaused)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output2),
					testAccCheckClassificationJobNotRecreated(&macie2Output, &macie2Output2),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusUserPaused)),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_definitions.0.buckets.0", bucketName),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
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

func testAccClassificationJob_complete(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	description := "Description of a test"
	descriptionUpdated := "Updated Description of a test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_complete(bucketName, string(awstypes.JobStatusRunning), description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_definitions.0.buckets.0", bucketName),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, "sampling_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "initial_run", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.key", "OBJECT_EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.values.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.key", "TAG"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.1.value", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.key", "OBJECT_EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.values.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.key", "TAG"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.1.value", "test2"),
				),
			},
			{
				Config: testAccClassificationJobConfig_complete(bucketName, string(awstypes.JobStatusRunning), descriptionUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, descriptionUpdated),
					resource.TestCheckResourceAttr(resourceName, "sampling_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "initial_run", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.key", "OBJECT_EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.0.simple_scope_term.0.values.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.key", "TAG"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.excludes.0.and.1.tag_scope_term.0.tag_values.1.value", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.key", "OBJECT_EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.0.simple_scope_term.0.values.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.key", "TAG"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.scoping.0.includes.0.and.1.tag_scope_term.0.tag_values.1.value", "test2"),
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

func testAccClassificationJob_WithTags(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	bucketName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_completeTags(bucketName, string(awstypes.JobStatusRunning)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key", names.AttrValue),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsAllPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.Key", names.AttrValue),
				),
			},
			{
				Config: testAccClassificationJobConfig_completeTags(bucketName, string(awstypes.JobStatusUserPaused)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusUserPaused)),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key", names.AttrValue),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsAllPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.Key", names.AttrValue),
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

func testAccClassificationJob_BucketCriteria(t *testing.T) {
	ctx := acctest.Context(t)
	var macie2Output macie2.DescribeClassificationJobOutput
	resourceName := "aws_macie2_classification_job.test"
	currentAccount := "data.aws_caller_identity.current"
	description := "Description of a test"
	descriptionUpdated := "Updated Description of a test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClassificationJobDestroy(ctx),
		ErrorCheck:               acctest.ErrorCheck(t, names.Macie2ServiceID),
		Steps: []resource.TestStep{
			{
				Config: testAccClassificationJobConfig_bucketCriteria(string(awstypes.JobStatusRunning), description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, "sampling_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "initial_run", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.comparator", "NE"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.key", "S3_BUCKET_SHARED_ACCESS"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.0", "EXTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.1", "INTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.2", "NOT_SHARED"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.3", "UNKNOWN"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.comparator", "NE"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.1.value", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.key", "ACCOUNT_ID"),
					resource.TestCheckResourceAttrPair(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.values.0", currentAccount, names.AttrAccountID),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.1.value", "test2"),
				),
			},
			{
				Config: testAccClassificationJobConfig_bucketCriteria(string(awstypes.JobStatusRunning), descriptionUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClassificationJobExists(ctx, resourceName, &macie2Output),
					resource.TestCheckResourceAttr(resourceName, "job_type", string(awstypes.JobTypeScheduled)),
					resource.TestCheckResourceAttr(resourceName, "job_status", string(awstypes.JobStatusRunning)),
					acctest.CheckResourceAttrRFC3339(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, descriptionUpdated),
					resource.TestCheckResourceAttr(resourceName, "sampling_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "initial_run", acctest.CtTrue),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.comparator", "NE"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.key", "S3_BUCKET_SHARED_ACCESS"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.0", "EXTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.1", "INTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.2", "NOT_SHARED"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.0.simple_criterion.0.values.3", "UNKNOWN"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.comparator", "NE"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.excludes.0.and.1.tag_criterion.0.tag_values.1.value", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.key", "ACCOUNT_ID"),
					resource.TestCheckResourceAttrPair(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.0.simple_criterion.0.values.0", currentAccount, names.AttrAccountID),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.comparator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.0.key", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.1.key", "test2"),
					resource.TestCheckResourceAttr(resourceName, "s3_job_definition.0.bucket_criteria.0.includes.0.and.1.tag_criterion.0.tag_values.1.value", "test2"),
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

func testAccCheckClassificationJobExists(ctx context.Context, resourceName string, macie2Session *macie2.DescribeClassificationJobOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Macie2Client(ctx)
		input := &macie2.DescribeClassificationJobInput{JobId: aws.String(rs.Primary.ID)}

		resp, err := conn.DescribeClassificationJob(ctx, input)

		if err != nil {
			return err
		}

		if resp == nil {
			return fmt.Errorf("macie ClassificationJob %q does not exist", rs.Primary.ID)
		}

		*macie2Session = *resp

		return nil
	}
}

func testAccCheckClassificationJobDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Macie2Client(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_macie2_classification_job" {
				continue
			}
			input := &macie2.DescribeClassificationJobInput{JobId: aws.String(rs.Primary.ID)}

			resp, err := conn.DescribeClassificationJob(ctx, input)

			if errs.IsAErrorMessageContains[*awstypes.AccessDeniedException](err, "Macie is not enabled") ||
				errs.IsAErrorMessageContains[*awstypes.ValidationException](err, "cannot update cancelled job for job") {
				continue
			}

			if err != nil {
				return err
			}

			if resp != nil && resp.JobStatus != awstypes.JobStatusCancelled {
				return fmt.Errorf("macie ClassificationJob %q still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccCheckClassificationJobNotRecreated(i, j *macie2.DescribeClassificationJobOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if !aws.ToTime(i.CreatedAt).Equal(aws.ToTime(j.CreatedAt)) {
			return fmt.Errorf("Macie Classification Job recreated")
		}

		return nil
	}
}

func testAccClassificationJobConfig_nameGenerated(bucketName, jobType string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_macie2_classification_job" "test" {
  depends_on = [aws_macie2_account.test]
  job_type   = %[2]q
  s3_job_definition {
    bucket_definitions {
      account_id = data.aws_caller_identity.current.account_id
      buckets    = [aws_s3_bucket.test.bucket]
    }
  }
}
`, bucketName, jobType)
}

func testAccClassificationJobConfig_namePrefix(nameBucket, namePrefix, jobType string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_macie2_classification_job" "test" {
  name_prefix = %[2]q
  job_type    = %[3]q
  s3_job_definition {
    bucket_definitions {
      account_id = data.aws_caller_identity.current.account_id
      buckets    = [aws_s3_bucket.test.bucket]
    }
  }
  depends_on = [aws_macie2_account.test]
}
`, nameBucket, namePrefix, jobType)
}

func testAccClassificationJobConfig_complete(nameBucket, jobStatus, description string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_macie2_classification_job" "test" {
  job_type = "SCHEDULED"
  s3_job_definition {
    bucket_definitions {
      account_id = data.aws_caller_identity.current.account_id
      buckets    = [aws_s3_bucket.test.bucket]
    }
    scoping {
      excludes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
        and {
          tag_scope_term {
            comparator = "EQ"
            key        = "TAG"
            tag_values {
              key   = "test"
              value = "test"
            }
            tag_values {
              key   = "test2"
              value = "test2"
            }
            target = "S3_OBJECT"
          }
        }
      }
      includes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
        and {
          tag_scope_term {
            comparator = "EQ"
            key        = "TAG"
            tag_values {
              key   = "test"
              value = "test"
            }
            tag_values {
              key   = "test2"
              value = "test2"
            }
            target = "S3_OBJECT"
          }
        }
      }
    }
  }
  schedule_frequency {
    daily_schedule = true
  }
  sampling_percentage = 100
  initial_run         = true
  job_status          = %[2]q
  description         = %[3]q

  depends_on = [aws_macie2_account.test]
}
`, nameBucket, jobStatus, description)
}

func testAccClassificationJobConfig_status(nameBucket, jobStatus string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_macie2_classification_job" "test" {
  job_type = "SCHEDULED"
  s3_job_definition {
    bucket_definitions {
      account_id = data.aws_caller_identity.current.account_id
      buckets    = [aws_s3_bucket.test.bucket]
    }
    scoping {
      excludes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
      }
      includes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
      }
    }
  }
  schedule_frequency {
    daily_schedule = true
  }
  sampling_percentage = 100
  initial_run         = true
  job_status          = %[2]q

  depends_on = [aws_macie2_account.test]
}
`, nameBucket, jobStatus)
}

func testAccClassificationJobConfig_completeTags(nameBucket, jobStatus string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_macie2_classification_job" "test" {
  job_type = "SCHEDULED"
  s3_job_definition {
    bucket_definitions {
      account_id = data.aws_caller_identity.current.account_id
      buckets    = [aws_s3_bucket.test.bucket]
    }
    scoping {
      excludes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
      }
      includes {
        and {
          simple_scope_term {
            comparator = "EQ"
            key        = "OBJECT_EXTENSION"
            values     = ["test"]
          }
        }
      }
    }
  }
  schedule_frequency {
    daily_schedule = true
  }
  sampling_percentage = 100
  description         = "test"
  initial_run         = true
  job_status          = %[2]q
  tags = {
    Key = "value"
  }

  depends_on = [aws_macie2_account.test]
}
`, nameBucket, jobStatus)
}

func testAccClassificationJobConfig_bucketCriteria(jobStatus, description string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_macie2_account" "test" {}

resource "aws_macie2_classification_job" "test" {
  job_type = "SCHEDULED"
  s3_job_definition {
    bucket_criteria {
      excludes {
        and {
          simple_criterion {
            comparator = "NE"
            key        = "S3_BUCKET_SHARED_ACCESS"
            values = [
              "EXTERNAL",
              "INTERNAL",
              "NOT_SHARED",
              "UNKNOWN"
            ]
          }
        }
        and {
          tag_criterion {
            comparator = "NE"
            tag_values {
              key   = "test"
              value = "test"
            }
            tag_values {
              key   = "test2"
              value = "test2"
            }
          }
        }
      }

      includes {
        and {
          simple_criterion {
            comparator = "EQ"
            key        = "ACCOUNT_ID"
            values = [
              data.aws_caller_identity.current.account_id,
            ]
          }
        }
        and {
          tag_criterion {
            comparator = "EQ"
            tag_values {
              key   = "test"
              value = "test"
            }
            tag_values {
              key   = "test2"
              value = "test2"
            }
          }
        }
      }
    }
  }
  schedule_frequency {
    daily_schedule = true
  }
  sampling_percentage = 100
  initial_run         = true
  job_status          = %[1]q
  description         = %[2]q

  depends_on = [aws_macie2_account.test]
}
`, jobStatus, description)
}
