// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafkaconnect_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccKafkaConnectWorkerConfigurationDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_mskconnect_worker_configuration.test"
	dataSourceName := "data.aws_mskconnect_worker_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.KafkaConnectEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.KafkaConnectServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkerConfigurationDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, names.AttrARN, dataSourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrDescription, dataSourceName, names.AttrDescription),
					resource.TestCheckResourceAttrPair(resourceName, "latest_revision", dataSourceName, "latest_revision"),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrName, dataSourceName, names.AttrName),
					resource.TestCheckResourceAttrPair(resourceName, "properties_file_content", dataSourceName, "properties_file_content"),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrTags, dataSourceName, names.AttrTags),
				),
			},
		},
	})
}

func testAccWorkerConfigurationDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_mskconnect_worker_configuration" "test" {
  name = %[1]q

  properties_file_content = <<EOF
key.converter=org.apache.kafka.connect.storage.StringConverter
value.converter=org.apache.kafka.connect.storage.StringConverter
EOF

  tags = {
    key1 = "value1"
  }
}

data "aws_mskconnect_worker_configuration" "test" {
  name = aws_mskconnect_worker_configuration.test.name
}
`, rName)
}
