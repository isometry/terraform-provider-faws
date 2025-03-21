// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/types"
	"github.com/isometry/terraform-provider-faws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_cloudhsm_v2_cluster",
			Name:     "Cluster",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCluster,
			TypeName: "aws_cloudhsm_v2_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceHSM,
			TypeName: "aws_cloudhsm_v2_hsm",
			Name:     "HSM",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudHSMV2
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
