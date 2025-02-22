// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package lightsail

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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBucket,
			TypeName: "aws_lightsail_bucket",
			Name:     "Bucket",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Bucket",
			},
		},
		{
			Factory:  ResourceBucketAccessKey,
			TypeName: "aws_lightsail_bucket_access_key",
			Name:     "Bucket Access Key",
		},
		{
			Factory:  ResourceBucketResourceAccess,
			TypeName: "aws_lightsail_bucket_resource_access",
			Name:     "Bucket Resource Access",
		},
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_lightsail_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Certificate",
			},
		},
		{
			Factory:  ResourceContainerService,
			TypeName: "aws_lightsail_container_service",
			Name:     "Container Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "ContainerService",
			},
		},
		{
			Factory:  ResourceContainerServiceDeploymentVersion,
			TypeName: "aws_lightsail_container_service_deployment_version",
			Name:     "Container Service Deployment Version",
		},
		{
			Factory:  ResourceDatabase,
			TypeName: "aws_lightsail_database",
			Name:     "Database",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Database",
			},
		},
		{
			Factory:  ResourceDisk,
			TypeName: "aws_lightsail_disk",
			Name:     "Disk",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Disk",
			},
		},
		{
			Factory:  ResourceDiskAttachment,
			TypeName: "aws_lightsail_disk_attachment",
			Name:     "Disk Attachment",
		},
		{
			Factory:  ResourceDistribution,
			TypeName: "aws_lightsail_distribution",
			Name:     "Distribution",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Distribution",
			},
		},
		{
			Factory:  ResourceDomain,
			TypeName: "aws_lightsail_domain",
			Name:     "Domain",
		},
		{
			Factory:  ResourceDomainEntry,
			TypeName: "aws_lightsail_domain_entry",
			Name:     "Domain Entry",
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_lightsail_instance",
			Name:     "Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "Instance",
			},
		},
		{
			Factory:  ResourceInstancePublicPorts,
			TypeName: "aws_lightsail_instance_public_ports",
			Name:     "Instance Public Ports",
		},
		{
			Factory:  ResourceKeyPair,
			TypeName: "aws_lightsail_key_pair",
			Name:     "KeyPair",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "KeyPair",
			},
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_lightsail_lb",
			Name:     "LB",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "LB",
			},
		},
		{
			Factory:  ResourceLoadBalancerAttachment,
			TypeName: "aws_lightsail_lb_attachment",
			Name:     "Load Balancer Attachment",
		},
		{
			Factory:  ResourceLoadBalancerCertificate,
			TypeName: "aws_lightsail_lb_certificate",
			Name:     "Load Balancer Certificate",
		},
		{
			Factory:  ResourceLoadBalancerCertificateAttachment,
			TypeName: "aws_lightsail_lb_certificate_attachment",
			Name:     "Load Balancer Certificate Attachment",
		},
		{
			Factory:  ResourceLoadBalancerHTTPSRedirectionPolicy,
			TypeName: "aws_lightsail_lb_https_redirection_policy",
			Name:     "Load Balancer HTTPS Redirection Policy",
		},
		{
			Factory:  ResourceLoadBalancerStickinessPolicy,
			TypeName: "aws_lightsail_lb_stickiness_policy",
			Name:     "Load Balancer Stickiness Policy",
		},
		{
			Factory:  ResourceStaticIP,
			TypeName: "aws_lightsail_static_ip",
			Name:     "Static IP",
		},
		{
			Factory:  ResourceStaticIPAttachment,
			TypeName: "aws_lightsail_static_ip_attachment",
			Name:     "Static IP Attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Lightsail
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
