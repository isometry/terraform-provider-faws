// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
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
			Factory:  dataSourcePrebuiltECRImage,
			TypeName: "aws_sagemaker_prebuilt_ecr_image",
			Name:     "Prebuilt ECR Image",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceApp,
			TypeName: "aws_sagemaker_app",
			Name:     "App",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceAppImageConfig,
			TypeName: "aws_sagemaker_app_image_config",
			Name:     "App Image Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCodeRepository,
			TypeName: "aws_sagemaker_code_repository",
			Name:     "Code Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDataQualityJobDefinition,
			TypeName: "aws_sagemaker_data_quality_job_definition",
			Name:     "Data Quality Job Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDevice,
			TypeName: "aws_sagemaker_device",
			Name:     "Device",
		},
		{
			Factory:  resourceDeviceFleet,
			TypeName: "aws_sagemaker_device_fleet",
			Name:     "Device Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDomain,
			TypeName: "aws_sagemaker_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEndpoint,
			TypeName: "aws_sagemaker_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEndpointConfiguration,
			TypeName: "aws_sagemaker_endpoint_configuration",
			Name:     "Endpoint Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFeatureGroup,
			TypeName: "aws_sagemaker_feature_group",
			Name:     "Feature Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFlowDefinition,
			TypeName: "aws_sagemaker_flow_definition",
			Name:     "Flow Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHub,
			TypeName: "aws_sagemaker_hub",
			Name:     "Hub",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHumanTaskUI,
			TypeName: "aws_sagemaker_human_task_ui",
			Name:     "Human Task UI",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceImage,
			TypeName: "aws_sagemaker_image",
			Name:     "Image",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceImageVersion,
			TypeName: "aws_sagemaker_image_version",
			Name:     "Image Version",
		},
		{
			Factory:  resourceMlflowTrackingServer,
			TypeName: "aws_sagemaker_mlflow_tracking_server",
			Name:     "Mlflow Tracking Server",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceModel,
			TypeName: "aws_sagemaker_model",
			Name:     "Model",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceModelPackageGroup,
			TypeName: "aws_sagemaker_model_package_group",
			Name:     "Model Package Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceModelPackageGroupPolicy,
			TypeName: "aws_sagemaker_model_package_group_policy",
			Name:     "Model Package Group Policy",
		},
		{
			Factory:  resourceMonitoringSchedule,
			TypeName: "aws_sagemaker_monitoring_schedule",
			Name:     "Monitoring Schedule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceNotebookInstance,
			TypeName: "aws_sagemaker_notebook_instance",
			Name:     "Notebook Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceNotebookInstanceLifeCycleConfiguration,
			TypeName: "aws_sagemaker_notebook_instance_lifecycle_configuration",
			Name:     "Notebook Instance Lifecycle Configuration",
		},
		{
			Factory:  resourcePipeline,
			TypeName: "aws_sagemaker_pipeline",
			Name:     "Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceProject,
			TypeName: "aws_sagemaker_project",
			Name:     "Project",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceServicecatalogPortfolioStatus,
			TypeName: "aws_sagemaker_servicecatalog_portfolio_status",
			Name:     "Servicecatalog Portfolio Status",
		},
		{
			Factory:  resourceSpace,
			TypeName: "aws_sagemaker_space",
			Name:     "Space",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceStudioLifecycleConfig,
			TypeName: "aws_sagemaker_studio_lifecycle_config",
			Name:     "Studio Lifecycle Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUserProfile,
			TypeName: "aws_sagemaker_user_profile",
			Name:     "User Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceWorkforce,
			TypeName: "aws_sagemaker_workforce",
			Name:     "Workforce",
		},
		{
			Factory:  resourceWorkteam,
			TypeName: "aws_sagemaker_workteam",
			Name:     "Workteam",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SageMaker
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sagemaker.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*sagemaker.Options){
		sagemaker.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return sagemaker.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*sagemaker.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*sagemaker.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *sagemaker.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*sagemaker.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
