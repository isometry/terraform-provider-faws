// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
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
			Factory:  dataSourceTopic,
			TypeName: "aws_sns_topic",
			Name:     "Topic",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourcePlatformApplication,
			TypeName: "aws_sns_platform_application",
			Name:     "Platform Application",
		},
		{
			Factory:  resourceSMSPreferences,
			TypeName: "aws_sns_sms_preferences",
			Name:     "SMS Preferences",
		},
		{
			Factory:  resourceTopic,
			TypeName: "aws_sns_topic",
			Name:     "Topic",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTopicDataProtectionPolicy,
			TypeName: "aws_sns_topic_data_protection_policy",
			Name:     "Topic Data Protection Policy",
		},
		{
			Factory:  resourceTopicPolicy,
			TypeName: "aws_sns_topic_policy",
			Name:     "Topic Policy",
		},
		{
			Factory:  resourceTopicSubscription,
			TypeName: "aws_sns_topic_subscription",
			Name:     "Topic Subscription",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SNS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sns.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*sns.Options){
		sns.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return sns.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*sns.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*sns.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *sns.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*sns.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
