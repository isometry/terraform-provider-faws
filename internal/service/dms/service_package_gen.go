// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package dms

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	databasemigrationservice_sdkv2 "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
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
			Factory:  dataSourceCertificate,
			TypeName: "aws_dms_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrCertificateARN,
			},
		},
		{
			Factory:  dataSourceEndpoint,
			TypeName: "aws_dms_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
		},
		{
			Factory:  dataSourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
			Name:     "Replication Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_instance_arn",
			},
		},
		{
			Factory:  dataSourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
			Name:     "Replication Subnet Group",
		},
		{
			Factory:  dataSourceReplicationTask,
			TypeName: "aws_dms_replication_task",
			Name:     "Replication Task",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCertificate,
			TypeName: "aws_dms_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrCertificateARN,
			},
		},
		{
			Factory:  resourceEndpoint,
			TypeName: "aws_dms_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_dms_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceReplicationConfig,
			TypeName: "aws_dms_replication_config",
			Name:     "Replication Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceReplicationInstance,
			TypeName: "aws_dms_replication_instance",
			Name:     "Replication Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_instance_arn",
			},
		},
		{
			Factory:  resourceReplicationSubnetGroup,
			TypeName: "aws_dms_replication_subnet_group",
			Name:     "Replication Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_subnet_group_arn",
			},
		},
		{
			Factory:  resourceReplicationTask,
			TypeName: "aws_dms_replication_task",
			Name:     "Replication Task",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "replication_task_arn",
			},
		},
		{
			Factory:  resourceS3Endpoint,
			TypeName: "aws_dms_s3_endpoint",
			Name:     "S3 Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "endpoint_arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DMS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*databasemigrationservice_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return databasemigrationservice_sdkv2.NewFromConfig(cfg,
		databasemigrationservice_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
