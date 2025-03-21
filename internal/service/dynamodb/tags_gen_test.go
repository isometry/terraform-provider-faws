// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package dynamodb_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/isometry/terraform-provider-faws/internal/acctest/statecheck"
	tfdynamodb "github.com/isometry/terraform-provider-faws/internal/service/dynamodb"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfdynamodb.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfdynamodb.ServicePackage(context.Background()), resourceAddress, knownValue)
}
