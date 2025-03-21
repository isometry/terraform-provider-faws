// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalk_test

import (
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccElasticBeanstalkApplicationDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_elastic_beanstalk_application.test"
	resourceName := "aws_elastic_beanstalk_application.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ElasticBeanstalkServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(dataSourceName, "appversion_lifecycle.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "appversion_lifecycle.0.delete_source_from_s3", dataSourceName, "appversion_lifecycle.0.delete_source_from_s3"),
					resource.TestCheckResourceAttrPair(resourceName, "appversion_lifecycle.0.max_age_in_days", dataSourceName, "appversion_lifecycle.0.max_age_in_days"),
					resource.TestCheckResourceAttrPair(resourceName, "appversion_lifecycle.0.service_role", dataSourceName, "appversion_lifecycle.0.service_role"),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrDescription, dataSourceName, names.AttrDescription),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrName, dataSourceName, names.AttrName),
				),
			},
		},
	})
}

func testAccApplicationDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccApplicationConfig_maxAge(rName), `
data "aws_elastic_beanstalk_application" "test" {
  name = aws_elastic_beanstalk_application.test.name
}
`)
}
