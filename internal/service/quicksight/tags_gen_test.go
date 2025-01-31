// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package quicksight_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/isometry/terraform-provider-faws/internal/acctest/statecheck"
	tfquicksight "github.com/isometry/terraform-provider-faws/internal/service/quicksight"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfquicksight.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfquicksight.ServicePackage(context.Background()), resourceAddress, knownValue)
}
