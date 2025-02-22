// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	awstypes "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/enum"
	"github.com/isometry/terraform-provider-faws/internal/errs"
	"github.com/isometry/terraform-provider-faws/internal/sweep"
	"github.com/isometry/terraform-provider-faws/internal/sweep/awsv2"
	"github.com/isometry/terraform-provider-faws/internal/sweep/sdk"
	"github.com/isometry/terraform-provider-faws/names"
)

const (
	loggingKeyServiceNamespace = "service_namespace"
)

func RegisterSweepers() {
	awsv2.Register("aws_appautoscaling_policy", sweepPolicy)

	awsv2.Register("aws_appautoscaling_target", sweepTarget,
		"aws_appautoscaling_policy",
	)
}

func sweepPolicy(ctx context.Context, client *conns.AWSClient) ([]sweep.Sweepable, error) {
	conn := client.AppAutoScalingClient(ctx)

	var sweepResources []sweep.Sweepable
	r := resourcePolicy()

	for _, serviceNamespace := range enum.EnumValues[awstypes.ServiceNamespace]() {
		ctx = tflog.SetField(ctx, loggingKeyServiceNamespace, serviceNamespace)
		tflog.Debug(ctx, "listing by service namespace")
		input := applicationautoscaling.DescribeScalingPoliciesInput{
			ServiceNamespace: serviceNamespace,
		}
		pages := applicationautoscaling.NewDescribeScalingPoliciesPaginator(conn, &input)
		for pages.HasMorePages() {
			page, err := pages.NextPage(ctx)
			if errs.IsAErrorMessageContains[*awstypes.ValidationException](err, "at 'serviceNamespace' failed to satisfy constraint") {
				tflog.Info(ctx, "Skipping service namespace", map[string]any{
					"error": err.Error(),
				})
				break
			}
			if err != nil {
				return nil, err
			}

			for _, policies := range page.ScalingPolicies {
				d := r.Data(nil)
				d.SetId(aws.ToString(policies.PolicyName)) // unused
				d.Set(names.AttrName, policies.PolicyName)
				d.Set(names.AttrResourceID, policies.ResourceId)
				d.Set("scalable_dimension", policies.ScalableDimension)
				d.Set("service_namespace", policies.ServiceNamespace)

				sweepResources = append(sweepResources, sdk.NewSweepResource(r, d, client))
			}
		}
	}

	return sweepResources, nil
}

func sweepTarget(ctx context.Context, client *conns.AWSClient) ([]sweep.Sweepable, error) {
	conn := client.AppAutoScalingClient(ctx)

	var sweepResources []sweep.Sweepable
	r := resourceTarget()

	for _, serviceNamespace := range enum.EnumValues[awstypes.ServiceNamespace]() {
		ctx = tflog.SetField(ctx, loggingKeyServiceNamespace, serviceNamespace)
		tflog.Debug(ctx, "listing by service namespace")
		input := applicationautoscaling.DescribeScalableTargetsInput{
			ServiceNamespace: serviceNamespace,
		}
		pages := applicationautoscaling.NewDescribeScalableTargetsPaginator(conn, &input)
		for pages.HasMorePages() {
			page, err := pages.NextPage(ctx)
			if errs.IsAErrorMessageContains[*awstypes.ValidationException](err, "at 'serviceNamespace' failed to satisfy constraint") {
				tflog.Info(ctx, "Skipping service namespace", map[string]any{
					"error": err.Error(),
				})
				break
			}
			if err != nil {
				return nil, err
			}

			for _, target := range page.ScalableTargets {
				d := r.Data(nil)
				d.SetId(aws.ToString(target.ResourceId))
				d.Set("scalable_dimension", target.ScalableDimension)
				d.Set("service_namespace", target.ServiceNamespace)

				sweepResources = append(sweepResources, sdk.NewSweepResource(r, d, client))
			}
		}
	}

	return sweepResources, nil
}
