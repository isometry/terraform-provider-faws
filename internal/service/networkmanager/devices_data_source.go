// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/errs/sdkdiag"
	tftags "github.com/isometry/terraform-provider-faws/internal/tags"
	"github.com/isometry/terraform-provider-faws/names"
)

// @SDKDataSource("aws_networkmanager_devices", name="Devices")
func dataSourceDevices() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceDevicesRead,

		Schema: map[string]*schema.Schema{
			"global_network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			names.AttrIDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"site_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			names.AttrTags: tftags.TagsSchema(),
		},
	}
}

func dataSourceDevicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).NetworkManagerClient(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig(ctx)
	tagsToMatch := tftags.New(ctx, d.Get(names.AttrTags).(map[string]interface{})).IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	input := &networkmanager.GetDevicesInput{
		GlobalNetworkId: aws.String(d.Get("global_network_id").(string)),
	}

	if v, ok := d.GetOk("site_id"); ok {
		input.SiteId = aws.String(v.(string))
	}

	output, err := findDevices(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "listing Network Manager Devices: %s", err)
	}

	var deviceIDs []string

	for _, v := range output {
		if len(tagsToMatch) > 0 {
			if !KeyValueTags(ctx, v.Tags).ContainsAll(tagsToMatch) {
				continue
			}
		}

		deviceIDs = append(deviceIDs, aws.ToString(v.DeviceId))
	}

	d.SetId(meta.(*conns.AWSClient).Region(ctx))
	d.Set(names.AttrIDs, deviceIDs)

	return diags
}
