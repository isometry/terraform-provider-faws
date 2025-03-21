// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccRDSEventCategoriesDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_db_event_categories.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.RDSServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEventCategoriesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					// These checks are not meant to be exhaustive, as regions have different support.
					// Instead these are generally to indicate that filtering works as expected.
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "availability"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "backup"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "configuration change"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "creation"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "deletion"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "failover"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "failure"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "low storage"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "maintenance"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "notification"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "recovery"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "restoration"),
				),
			},
		},
	})
}

func TestAccRDSEventCategoriesDataSource_sourceType(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_db_event_categories.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.RDSServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEventCategoriesDataSourceConfig_sourceType("db-snapshot"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// These checks are not meant to be exhaustive, as regions have different support.
					// Instead these are generally to indicate that filtering works as expected.
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "creation"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "deletion"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "notification"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "event_categories.*", "restoration"),
				),
			},
		},
	})
}

func testAccEventCategoriesDataSourceConfig_basic() string {
	return `
data "aws_db_event_categories" "test" {}
`
}

func testAccEventCategoriesDataSourceConfig_sourceType(sourceType string) string {
	return fmt.Sprintf(`
data "aws_db_event_categories" "test" {
  source_type = %[1]q
}
`, sourceType)
}
