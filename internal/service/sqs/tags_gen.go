// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/logging"
	tftags "github.com/isometry/terraform-provider-faws/internal/tags"
	"github.com/isometry/terraform-provider-faws/internal/types/option"
	"github.com/isometry/terraform-provider-faws/names"
)

// listTags lists sqs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *sqs.Client, identifier string, optFns ...func(*sqs.Options)) (tftags.KeyValueTags, error) {
	input := sqs.ListQueueTagsInput{
		QueueUrl: aws.String(identifier),
	}

	output, err := conn.ListQueueTags(ctx, &input, optFns...)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists sqs service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).SQSClient(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// map[string]string handling

// Tags returns sqs service tags.
func Tags(tags tftags.KeyValueTags) map[string]string {
	return tags.Map()
}

// KeyValueTags creates tftags.KeyValueTags from sqs service tags.
func KeyValueTags(ctx context.Context, tags map[string]string) tftags.KeyValueTags {
	return tftags.New(ctx, tags)
}

// getTagsIn returns sqs service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) map[string]string {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets sqs service tags in Context.
func setTagsOut(ctx context.Context, tags map[string]string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// createTags creates sqs service tags for new resources.
func createTags(ctx context.Context, conn *sqs.Client, identifier string, tags map[string]string, optFns ...func(*sqs.Options)) error {
	if len(tags) == 0 {
		return nil
	}

	return updateTags(ctx, conn, identifier, nil, tags, optFns...)
}

// updateTags updates sqs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *sqs.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*sqs.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SQS)
	if len(removedTags) > 0 {
		input := sqs.UntagQueueInput{
			QueueUrl: aws.String(identifier),
			TagKeys:  removedTags.Keys(),
		}

		_, err := conn.UntagQueue(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SQS)
	if len(updatedTags) > 0 {
		input := sqs.TagQueueInput{
			QueueUrl: aws.String(identifier),
			Tags:     Tags(updatedTags),
		}

		_, err := conn.TagQueue(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates sqs service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).SQSClient(ctx), identifier, oldTags, newTags)
}
