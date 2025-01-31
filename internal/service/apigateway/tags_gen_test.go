// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package apigateway_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/isometry/terraform-provider-faws/internal/acctest/statecheck"
	tfapigateway "github.com/isometry/terraform-provider-faws/internal/service/apigateway"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfapigateway.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfapigateway.ServicePackage(context.Background()), resourceAddress, knownValue)
}
