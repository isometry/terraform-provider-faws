// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflow

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/sweep"
	"github.com/isometry/terraform-provider-faws/internal/sweep/awsv2"
	"github.com/isometry/terraform-provider-faws/internal/sweep/sdk"
	"github.com/isometry/terraform-provider-faws/names"
)

func RegisterSweepers() {
	awsv2.Register("aws_appflow_flow", sweepInstanceProfile)
}

func sweepInstanceProfile(ctx context.Context, client *conns.AWSClient) ([]sweep.Sweepable, error) {
	conn := client.AppFlowClient(ctx)

	var sweepResources []sweep.Sweepable
	r := resourceFlow()

	pages := appflow.NewListFlowsPaginator(conn, &appflow.ListFlowsInput{})
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, flow := range page.Flows {
			d := r.Data(nil)
			d.SetId(aws.ToString(flow.FlowArn))
			d.Set(names.AttrName, flow.FlowName)

			sweepResources = append(sweepResources, sdk.NewSweepResource(r, d, client))
		}
	}

	return sweepResources, nil
}
