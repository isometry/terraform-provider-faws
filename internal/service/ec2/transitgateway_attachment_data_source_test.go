// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	tfsync "github.com/isometry/terraform-provider-faws/internal/experimental/sync"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccTransitGatewayAttachmentDataSource_Filter(t *testing.T, semaphore tfsync.Semaphore) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_transit_gateway_attachment.test"
	resourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckTransitGatewaySynchronize(t, semaphore)
			acctest.PreCheck(ctx, t)
			testAccPreCheckTransitGateway(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransitGatewayAttachmentDataSourceConfig_filter(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceAttrRegionalARNFormat(ctx, dataSourceName, names.AttrARN, "ec2", "transit-gateway-attachment/{id}"),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrResourceID, resourceName, names.AttrVPCID),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrID, resourceName, names.AttrID),
					acctest.CheckResourceAttrAccountID(ctx, dataSourceName, "resource_owner_id"),
					resource.TestCheckResourceAttr(dataSourceName, names.AttrResourceType, "vpc"),
					resource.TestCheckResourceAttrSet(dataSourceName, names.AttrState),
					resource.TestCheckResourceAttrPair(dataSourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrTransitGatewayAttachmentID, dataSourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrTransitGatewayID, resourceName, names.AttrTransitGatewayID),
					acctest.CheckResourceAttrAccountID(ctx, dataSourceName, "transit_gateway_owner_id"),
					resource.TestCheckResourceAttr(dataSourceName, "association_state", "associated"),
					resource.TestCheckResourceAttrSet(dataSourceName, "association_transit_gateway_route_table_id"),
				),
			},
		},
	})
}

func testAccTransitGatewayAttachmentDataSource_ID(t *testing.T, semaphore tfsync.Semaphore) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_transit_gateway_attachment.test"
	resourceName := "aws_ec2_transit_gateway_vpc_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckTransitGatewaySynchronize(t, semaphore)
			acctest.PreCheck(ctx, t)
			testAccPreCheckTransitGateway(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransitGatewayAttachmentDataSourceConfig_id(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceAttrRegionalARNFormat(ctx, dataSourceName, names.AttrARN, "ec2", "transit-gateway-attachment/{id}"),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrResourceID, resourceName, names.AttrVPCID),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrID, resourceName, names.AttrID),
					acctest.CheckResourceAttrAccountID(ctx, dataSourceName, "resource_owner_id"),
					resource.TestCheckResourceAttr(dataSourceName, names.AttrResourceType, "vpc"),
					resource.TestCheckResourceAttrSet(dataSourceName, names.AttrState),
					resource.TestCheckResourceAttrPair(dataSourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrTransitGatewayAttachmentID, dataSourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrTransitGatewayID, resourceName, names.AttrTransitGatewayID),
					acctest.CheckResourceAttrAccountID(ctx, dataSourceName, "transit_gateway_owner_id"),
					resource.TestCheckResourceAttr(dataSourceName, "association_state", "associated"),
					resource.TestCheckResourceAttrSet(dataSourceName, "association_transit_gateway_route_table_id"),
				),
			},
		},
	})
}

func testAccTransitGatewayAttachmentDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids         = aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_id             = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

data "aws_ec2_transit_gateway_attachment" "test" {
  filter {
    name   = "transit-gateway-id"
    values = [aws_ec2_transit_gateway.test.id]
  }

  filter {
    name   = "resource-type"
    values = ["vpc"]
  }

  filter {
    name   = "resource-id"
    values = [aws_vpc.test.id]
  }

  depends_on = [aws_ec2_transit_gateway_vpc_attachment.test]
}
`, rName))
}

func testAccTransitGatewayAttachmentDataSourceConfig_id(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids         = aws_subnet.test[*].id
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_id             = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

data "aws_ec2_transit_gateway_attachment" "test" {
  transit_gateway_attachment_id = aws_ec2_transit_gateway_vpc_attachment.test.id
}
`, rName))
}
