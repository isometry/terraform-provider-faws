package passes

import (
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT001"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT002"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT003"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT004"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT005"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSAT006"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSR001"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSR002"
	"github.com/isometry/terraform-provider-faws/ci/providerlint/passes/AWSV001"
	"golang.org/x/tools/go/analysis"
)

var AllChecks = []*analysis.Analyzer{
	AWSAT001.Analyzer,
	AWSAT002.Analyzer,
	AWSAT003.Analyzer,
	AWSAT004.Analyzer,
	AWSAT005.Analyzer,
	AWSAT006.Analyzer,
	AWSR001.Analyzer,
	AWSR002.Analyzer,
	AWSV001.Analyzer,
}
