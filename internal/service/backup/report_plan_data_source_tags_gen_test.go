// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package backup_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/isometry/terraform-provider-faws/internal/acctest"
	"github.com/isometry/terraform-provider-faws/names"
)

func TestAccBackupReportPlanDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BackupServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ReportPlan/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtKey1: config.StringVariable(acctest.CtValue1),
					}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
		},
	})
}

func TestAccBackupReportPlanDataSource_tags_NullMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BackupServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ReportPlan/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:        config.StringVariable(rName),
					acctest.CtResourceTags: nil,
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccBackupReportPlanDataSource_tags_EmptyMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BackupServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ReportPlan/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:        config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccBackupReportPlanDataSource_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.BackupServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ReportPlan/data.tags_defaults/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccBackupReportPlanDataSource_tags_IgnoreTags_Overlap_DefaultTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.BackupServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ReportPlan/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtProviderKey1),
					),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
					expectFullDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccBackupReportPlanDataSource_tags_IgnoreTags_Overlap_ResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_backup_report_plan.test"
	rName := randomReportPlanName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.BackupServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ReportPlan/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtResourceKey1),
					),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
					expectFullDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
