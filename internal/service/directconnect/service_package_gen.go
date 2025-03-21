// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
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
			Factory:  dataSourceConnection,
			TypeName: "aws_dx_connection",
			Name:     "Connection",
		},
		{
			Factory:  dataSourceGateway,
			TypeName: "aws_dx_gateway",
			Name:     "Gateway",
		},
		{
			Factory:  dataSourceLocation,
			TypeName: "aws_dx_location",
			Name:     "Location",
		},
		{
			Factory:  dataSourceLocations,
			TypeName: "aws_dx_locations",
			Name:     "Locations",
		},
		{
			Factory:  dataSourceRouterConfiguration,
			TypeName: "aws_dx_router_configuration",
			Name:     "Router Configuration",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBGPPeer,
			TypeName: "aws_dx_bgp_peer",
			Name:     "BGP Peer",
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_dx_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnectionAssociation,
			TypeName: "aws_dx_connection_association",
			Name:     "Connection LAG Association",
		},
		{
			Factory:  resourceConnectionConfirmation,
			TypeName: "aws_dx_connection_confirmation",
			Name:     "Connection Confirmation",
		},
		{
			Factory:  resourceGateway,
			TypeName: "aws_dx_gateway",
			Name:     "Gateway",
		},
		{
			Factory:  resourceGatewayAssociation,
			TypeName: "aws_dx_gateway_association",
			Name:     "Gateway Association",
		},
		{
			Factory:  resourceGatewayAssociationProposal,
			TypeName: "aws_dx_gateway_association_proposal",
			Name:     "Gateway Association Proposal",
		},
		{
			Factory:  resourceHostedConnection,
			TypeName: "aws_dx_hosted_connection",
			Name:     "Hosted Connection",
		},
		{
			Factory:  resourceHostedPrivateVirtualInterface,
			TypeName: "aws_dx_hosted_private_virtual_interface",
			Name:     "Hosted Private Virtual Interface",
		},
		{
			Factory:  resourceHostedPrivateVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_private_virtual_interface_accepter",
			Name:     "Hosted Private Virtual Interface Accepter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHostedPublicVirtualInterface,
			TypeName: "aws_dx_hosted_public_virtual_interface",
			Name:     "Hosted Public Virtual Interface",
		},
		{
			Factory:  resourceHostedPublicVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_public_virtual_interface_accepter",
			Name:     "Hosted Public Virtual Interface Accepter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHostedTransitVirtualInterface,
			TypeName: "aws_dx_hosted_transit_virtual_interface",
			Name:     "Hosted Transit Virtual Interface",
		},
		{
			Factory:  resourceHostedTransitVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_transit_virtual_interface_accepter",
			Name:     "Hosted Transit Virtual Interface Accepter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceLag,
			TypeName: "aws_dx_lag",
			Name:     "LAG",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceMacSecKeyAssociation,
			TypeName: "aws_dx_macsec_key_association",
			Name:     "MACSec Key Association",
		},
		{
			Factory:  resourcePrivateVirtualInterface,
			TypeName: "aws_dx_private_virtual_interface",
			Name:     "Private Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourcePublicVirtualInterface,
			TypeName: "aws_dx_public_virtual_interface",
			Name:     "Public Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTransitVirtualInterface,
			TypeName: "aws_dx_transit_virtual_interface",
			Name:     "Transit Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DirectConnect
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*directconnect.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*directconnect.Options){
		directconnect.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return directconnect.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*directconnect.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*directconnect.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *directconnect.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*directconnect.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
