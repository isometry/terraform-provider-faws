// Code generated by internal/generate/tagresource/main.go; DO NOT EDIT.

package ecs_test

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	tfecs "github.com/isometry/terraform-provider-faws/internal/service/ecs"
	tftags "github.com/isometry/terraform-provider-faws/internal/tags"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/names"
)

func testAccCheckTagDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).ECSClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_ecs_tag" {
				continue
			}

			identifier, key, err := tftags.GetResourceID(rs.Primary.ID)
			if err != nil {
				return err
			}

			_, err = tfecs.FindTag(ctx, conn, identifier, key)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("%s resource (%s) tag (%s) still exists", names.ECS, identifier, key)
		}

		return nil
	}
}

func testAccCheckTagExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)
		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ECSClient(ctx)

		_, err = tfecs.FindTag(ctx, conn, identifier, key)

		return err
	}
}
