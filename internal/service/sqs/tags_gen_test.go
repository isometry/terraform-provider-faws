// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package sqs_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/isometry/terraform-provider-faws/internal/acctest/statecheck"
	tfsqs "github.com/isometry/terraform-provider-faws/internal/service/sqs"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfsqs.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfsqs.ServicePackage(context.Background()), resourceAddress, knownValue)
}
