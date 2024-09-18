// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ses

// Exports for use in tests only.
var (
	ResourceReceiptRule    = resourceReceiptRule
	ResourceReceiptRuleSet = resourceReceiptRuleSet
	ResourceTemplate       = resourceTemplate

	FindReceiptRuleByTwoPartKey = findReceiptRuleByTwoPartKey
	FindReceiptRuleSetByName    = findReceiptRuleSetByName
	FindTemplateByName          = findTemplateByName
)
