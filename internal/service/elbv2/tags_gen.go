// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package elbv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/logging"
	tftags "github.com/isometry/terraform-provider-faws/internal/tags"
	"github.com/isometry/terraform-provider-faws/internal/types/option"
	"github.com/isometry/terraform-provider-faws/names"
)

// listTags lists elbv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *elasticloadbalancingv2.Client, identifier string, optFns ...func(*elasticloadbalancingv2.Options)) (tftags.KeyValueTags, error) {
	input := elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{identifier},
	}

	output, err := conn.DescribeTags(ctx, &input, optFns...)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagDescriptions[0].Tags), nil
}

// ListTags lists elbv2 service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).ELBV2Client(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns elbv2 service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	result := make([]awstypes.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from elasticloadbalancingv2 service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns elbv2 service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets elbv2 service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// createTags creates elbv2 service tags for new resources.
func createTags(ctx context.Context, conn *elasticloadbalancingv2.Client, identifier string, tags []awstypes.Tag, optFns ...func(*elasticloadbalancingv2.Options)) error {
	if len(tags) == 0 {
		return nil
	}

	return updateTags(ctx, conn, identifier, nil, KeyValueTags(ctx, tags), optFns...)
}

// updateTags updates elbv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *elasticloadbalancingv2.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*elasticloadbalancingv2.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.ELBV2)
	if len(removedTags) > 0 {
		input := elasticloadbalancingv2.RemoveTagsInput{
			ResourceArns: []string{identifier},
			TagKeys:      removedTags.Keys(),
		}

		_, err := conn.RemoveTags(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.ELBV2)
	if len(updatedTags) > 0 {
		input := elasticloadbalancingv2.AddTagsInput{
			ResourceArns: []string{identifier},
			Tags:         Tags(updatedTags),
		}

		_, err := conn.AddTags(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates elbv2 service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).ELBV2Client(ctx), identifier, oldTags, newTags)
}
