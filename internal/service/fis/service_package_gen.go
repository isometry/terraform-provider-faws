// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package fis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fis"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceExperimentTemplate,
			TypeName: "aws_fis_experiment_template",
			Name:     "Experiment Template",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.FIS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*fis.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*fis.Options){
		fis.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return fis.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*fis.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*fis.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *fis.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*fis.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
