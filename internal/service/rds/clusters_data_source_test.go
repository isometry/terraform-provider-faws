// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	tfrds "github.com/isometry/terraform-provider-faws/internal/service/rds"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccRDSClustersDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	var dbCluster types.DBCluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_rds_clusters.test"
	resourceName := "aws_rds_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.RDSServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccClustersDataSourceConfig_filter(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &dbCluster),
					resource.TestCheckResourceAttr(dataSourceName, "cluster_arns.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_arns.0", resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(dataSourceName, "cluster_identifiers.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_identifiers.0", resourceName, names.AttrClusterIdentifier),
				),
			},
		},
	})
}

func testAccClustersDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
resource "aws_rds_cluster" "test" {
  cluster_identifier  = %[1]q
  database_name       = "test"
  engine              = %[2]q
  master_username     = "tfacctest"
  master_password     = "avoid-plaintext-passwords"
  skip_final_snapshot = true
}

resource "aws_rds_cluster" "wrong" {
  cluster_identifier  = "wrong-%[1]s"
  database_name       = "test"
  engine              = %[2]q
  master_username     = "tfacctest"
  master_password     = "avoid-plaintext-passwords"
  skip_final_snapshot = true
}

data "aws_rds_clusters" "test" {
  filter {
    name   = "db-cluster-id"
    values = [aws_rds_cluster.test.cluster_identifier]
  }

  depends_on = [aws_rds_cluster.wrong]
}
`, rName, tfrds.ClusterEngineAuroraMySQL)
}
