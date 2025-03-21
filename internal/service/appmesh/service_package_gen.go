// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
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
			Factory:  dataSourceGatewayRoute,
			TypeName: "aws_appmesh_gateway_route",
			Name:     "Gateway Route",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceMesh,
			TypeName: "aws_appmesh_mesh",
			Name:     "Service Mesh",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceRoute,
			TypeName: "aws_appmesh_route",
			Name:     "Route",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceVirtualGateway,
			TypeName: "aws_appmesh_virtual_gateway",
			Name:     "Virtual Gateway",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceVirtualNode,
			TypeName: "aws_appmesh_virtual_node",
			Name:     "Virtual Node",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceVirtualRouter,
			TypeName: "aws_appmesh_virtual_router",
			Name:     "Virtual Router",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceVirtualService,
			TypeName: "aws_appmesh_virtual_service",
			Name:     "Virtual Service",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceGatewayRoute,
			TypeName: "aws_appmesh_gateway_route",
			Name:     "Gateway Route",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceMesh,
			TypeName: "aws_appmesh_mesh",
			Name:     "Service Mesh",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRoute,
			TypeName: "aws_appmesh_route",
			Name:     "Route",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVirtualGateway,
			TypeName: "aws_appmesh_virtual_gateway",
			Name:     "Virtual Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVirtualNode,
			TypeName: "aws_appmesh_virtual_node",
			Name:     "Virtual Node",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVirtualRouter,
			TypeName: "aws_appmesh_virtual_router",
			Name:     "Virtual Router",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVirtualService,
			TypeName: "aws_appmesh_virtual_service",
			Name:     "Virtual Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppMesh
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appmesh.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*appmesh.Options){
		appmesh.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return appmesh.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*appmesh.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*appmesh.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *appmesh.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*appmesh.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
