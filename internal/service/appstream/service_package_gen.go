// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/types"
	"github.com/isometry/terraform-provider-faws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceImage,
			TypeName: "aws_appstream_image",
			Name:     "Image",
		},
	}
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
			Factory:  ResourceDirectoryConfig,
			TypeName: "aws_appstream_directory_config",
			Name:     "Directory Config",
		},
		{
			Factory:  ResourceFleet,
			TypeName: "aws_appstream_fleet",
			Name:     "Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceFleetStackAssociation,
			TypeName: "aws_appstream_fleet_stack_association",
			Name:     "Fleet Stack Association",
		},
		{
			Factory:  ResourceImageBuilder,
			TypeName: "aws_appstream_image_builder",
			Name:     "Image Builder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceStack,
			TypeName: "aws_appstream_stack",
			Name:     "Stack",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_appstream_user",
			Name:     "User",
		},
		{
			Factory:  ResourceUserStackAssociation,
			TypeName: "aws_appstream_user_stack_association",
			Name:     "User Stack Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppStream
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appstream.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*appstream.Options){
		appstream.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return appstream.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*appstream.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*appstream.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *appstream.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*appstream.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
