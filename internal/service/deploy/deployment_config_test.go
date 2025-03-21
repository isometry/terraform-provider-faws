// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploy_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfcodedeploy "github.com/isometry/terraform-provider-faws/internal/service/deploy"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccDeployDeploymentConfig_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var config types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_fleet(rName, 75),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config),
					acctest.CheckResourceAttrRegionalARNFormat(ctx, resourceName, names.AttrARN, "codedeploy", "deploymentconfig:{deployment_config_name}"),
					resource.TestCheckResourceAttr(resourceName, "deployment_config_name", rName),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.#", "0"),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrID, resourceName, "deployment_config_name"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var config types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_fleet(rName, 75),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfcodedeploy.ResourceDeploymentConfig(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_fleetPercent(t *testing.T) {
	ctx := acctest.Context(t)
	var config1, config2 types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_fleet(rName, 75),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config1),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.type", "FLEET_PERCENT"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.value", "75"),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "0"),
				),
			},
			{
				Config: testAccDeploymentConfigConfig_fleet(rName, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config2),
					testAccCheckDeploymentConfigRecreated(&config1, &config2),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.type", "FLEET_PERCENT"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_hostCount(t *testing.T) {
	ctx := acctest.Context(t)
	var config1, config2 types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_hostCount(rName, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config1),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.type", "HOST_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "0"),
				),
			},
			{
				Config: testAccDeploymentConfigConfig_hostCount(rName, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config2),
					testAccCheckDeploymentConfigRecreated(&config1, &config2),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.type", "HOST_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.0.value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_trafficCanary(t *testing.T) {
	ctx := acctest.Context(t)
	var config1, config2 types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_trafficCanary(rName, 10, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config1),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Lambda"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.type", "TimeBasedCanary"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.0.interval", "10"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.0.percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "0"),
				),
			},
			{
				Config: testAccDeploymentConfigConfig_trafficCanary(rName, 3, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config2),
					testAccCheckDeploymentConfigRecreated(&config1, &config2),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Lambda"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.type", "TimeBasedCanary"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.0.interval", "3"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.0.percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_trafficLinear(t *testing.T) {
	ctx := acctest.Context(t)
	var config1, config2 types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_trafficLinear(rName, 10, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config1),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Lambda"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.type", "TimeBasedLinear"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.0.interval", "10"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.0.percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "0"),
				),
			},
			{
				Config: testAccDeploymentConfigConfig_trafficLinear(rName, 3, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config2),
					testAccCheckDeploymentConfigRecreated(&config1, &config2),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Lambda"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.type", "TimeBasedLinear"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.0.interval", "3"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_linear.0.percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "traffic_routing_config.0.time_based_canary.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "minimum_healthy_hosts.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDeployDeploymentConfig_zonalConfig(t *testing.T) {
	ctx := acctest.Context(t)
	var config1, config2 types.DeploymentConfigInfo
	resourceName := "aws_codedeploy_deployment_config.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.DeployServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeploymentConfigDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigConfig_zonalConfig(rName, 10, "FLEET_PERCENT", 20, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config1),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.first_zone_monitor_duration_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.0.type", "FLEET_PERCENT"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.0.value", "20"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.monitor_duration_in_seconds", "10"),
				),
			},
			{
				Config: testAccDeploymentConfigConfig_zonalConfig(rName, 20, "HOST_COUNT", 2, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentConfigExists(ctx, resourceName, &config2),
					resource.TestCheckResourceAttr(resourceName, "compute_platform", "Server"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.first_zone_monitor_duration_in_seconds", "20"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.0.type", "HOST_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.minimum_healthy_hosts_per_zone.0.value", "2"),
					resource.TestCheckResourceAttr(resourceName, "zonal_config.0.monitor_duration_in_seconds", "20"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeploymentConfigDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).DeployClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_codedeploy_deployment_config" {
				continue
			}

			_, err := tfcodedeploy.FindDeploymentConfigByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("CodeDeploy Deployment Config %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckDeploymentConfigExists(ctx context.Context, n string, v *types.DeploymentConfigInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).DeployClient(ctx)

		output, err := tfcodedeploy.FindDeploymentConfigByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckDeploymentConfigRecreated(i, j *types.DeploymentConfigInfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if aws.ToTime(i.CreateTime).Equal(aws.ToTime(j.CreateTime)) {
			return errors.New("CodeDeploy Deployment Config was not recreated")
		}

		return nil
	}
}

func testAccDeploymentConfigConfig_fleet(rName string, value int) string {
	return fmt.Sprintf(`
resource "aws_codedeploy_deployment_config" "test" {
  deployment_config_name = %[1]q

  minimum_healthy_hosts {
    type  = "FLEET_PERCENT"
    value = %[2]d
  }
}
`, rName, value)
}

func testAccDeploymentConfigConfig_hostCount(rName string, value int) string {
	return fmt.Sprintf(`
resource "aws_codedeploy_deployment_config" "test" {
  deployment_config_name = %[1]q

  minimum_healthy_hosts {
    type  = "HOST_COUNT"
    value = %[2]d
  }
}
`, rName, value)
}

func testAccDeploymentConfigConfig_trafficCanary(rName string, interval, percentage int) string {
	return fmt.Sprintf(`
resource "aws_codedeploy_deployment_config" "test" {
  deployment_config_name = %[1]q
  compute_platform       = "Lambda"

  traffic_routing_config {
    type = "TimeBasedCanary"

    time_based_canary {
      interval   = %[2]d
      percentage = %[3]d
    }
  }
}
`, rName, interval, percentage)
}

func testAccDeploymentConfigConfig_trafficLinear(rName string, interval, percentage int) string {
	return fmt.Sprintf(`
resource "aws_codedeploy_deployment_config" "test" {
  deployment_config_name = %[1]q
  compute_platform       = "Lambda"

  traffic_routing_config {
    type = "TimeBasedLinear"

    time_based_linear {
      interval   = %[2]d
      percentage = %[3]d
    }
  }
}
`, rName, interval, percentage)
}

func testAccDeploymentConfigConfig_zonalConfig(rName string, first_zone_monitor_duration int, minimum_healthy_host_type string, minimum_healthy_host_value int, monitor_duration int) string {
	return fmt.Sprintf(`
resource "aws_codedeploy_deployment_config" "test" {
  deployment_config_name = %[1]q
  compute_platform       = "Server"

  minimum_healthy_hosts {
    type  = "HOST_COUNT"
    value = 3
  }

  zonal_config {
    first_zone_monitor_duration_in_seconds = %[2]d
    minimum_healthy_hosts_per_zone {
      type  = %[3]q
      value = %[4]d
    }
    monitor_duration_in_seconds = %[5]d
  }
}
`, rName, first_zone_monitor_duration, minimum_healthy_host_type, minimum_healthy_host_value, monitor_duration)
}
