// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firehose

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/sweep"
	"github.com/isometry/terraform-provider-faws/internal/sweep/awsv2"
	"github.com/isometry/terraform-provider-faws/names"
)

func RegisterSweepers() {
	resource.AddTestSweepers("aws_kinesis_firehose_delivery_stream", &resource.Sweeper{
		Name: "aws_kinesis_firehose_delivery_stream",
		F:    sweepDeliveryStreams,
	})
}

func sweepDeliveryStreams(region string) error {
	ctx := sweep.Context(region)
	client, err := sweep.SharedRegionalSweepClient(ctx, region)
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}
	conn := client.FirehoseClient(ctx)
	input := &firehose.ListDeliveryStreamsInput{}
	sweepResources := make([]sweep.Sweepable, 0)

	err = listDeliveryStreamsPages(ctx, conn, input, func(page *firehose.ListDeliveryStreamsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, name := range page.DeliveryStreamNames {
			r := resourceDeliveryStream()
			d := r.Data(nil)
			d.SetId(client.RegionalARN(ctx, "firehose", fmt.Sprintf("deliverystream/%s", name)))
			d.Set(names.AttrName, name)

			sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
		}

		return !lastPage
	})

	if awsv2.SkipSweepError(err) {
		log.Printf("[WARN] Skipping Kinesis Firehose Delivery Stream sweep for %s: %s", region, err)
		return nil
	}

	if err != nil {
		return fmt.Errorf("error listing Kinesis Firehose Delivery Streams (%s): %w", region, err)
	}

	err = sweep.SweepOrchestrator(ctx, sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping Kinesis Firehose Delivery Streams (%s): %w", region, err)
	}

	return nil
}
