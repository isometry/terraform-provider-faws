// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rum

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rum"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/sweep"
	"github.com/isometry/terraform-provider-faws/internal/sweep/awsv2"
)

func RegisterSweepers() {
	resource.AddTestSweepers("aws_rum_app_monitor", &resource.Sweeper{
		Name: "aws_rum_app_monitor",
		F:    sweepAppMonitors,
	})
}

func sweepAppMonitors(region string) error {
	ctx := sweep.Context(region)
	client, err := sweep.SharedRegionalSweepClient(ctx, region)
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}
	conn := client.RUMClient(ctx)
	input := &rum.ListAppMonitorsInput{}
	sweepResources := make([]sweep.Sweepable, 0)

	pages := rum.NewListAppMonitorsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if awsv2.SkipSweepError(err) {
			log.Printf("[WARN] Skipping CloudWatch RUM App Monitor sweep for %s: %s", region, err)
			return nil
		}

		if err != nil {
			return fmt.Errorf("error listing CloudWatch RUM App Monitors (%s): %w", region, err)
		}

		for _, v := range page.AppMonitorSummaries {
			r := resourceAppMonitor()
			d := r.Data(nil)
			d.SetId(aws.ToString(v.Name))

			sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
		}
	}

	err = sweep.SweepOrchestrator(ctx, sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping CloudWatch RUM App Monitors (%s): %w", region, err)
	}

	return nil
}
