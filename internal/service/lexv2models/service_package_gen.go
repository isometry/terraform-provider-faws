// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package lexv2models

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/types"
	"github.com/isometry/terraform-provider-faws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceBot,
			TypeName: "aws_lexv2models_bot",
			Name:     "Bot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceBotLocale,
			TypeName: "aws_lexv2models_bot_locale",
			Name:     "Bot Locale",
		},
		{
			Factory:  newResourceBotVersion,
			TypeName: "aws_lexv2models_bot_version",
			Name:     "Bot Version",
		},
		{
			Factory:  newResourceIntent,
			TypeName: "aws_lexv2models_intent",
			Name:     "Intent",
		},
		{
			Factory:  newResourceSlot,
			TypeName: "aws_lexv2models_slot",
			Name:     "Slot",
		},
		{
			Factory:  newResourceSlotType,
			TypeName: "aws_lexv2models_slot_type",
			Name:     "Slot Type",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LexV2Models
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*lexmodelsv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*lexmodelsv2.Options){
		lexmodelsv2.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return lexmodelsv2.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*lexmodelsv2.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*lexmodelsv2.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *lexmodelsv2.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*lexmodelsv2.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
