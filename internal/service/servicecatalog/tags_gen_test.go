// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package servicecatalog_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/isometry/terraform-provider-faws/internal/acctest/statecheck"
	tfservicecatalog "github.com/isometry/terraform-provider-faws/internal/service/servicecatalog"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfservicecatalog.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfservicecatalog.ServicePackage(context.Background()), resourceAddress, knownValue)
}
