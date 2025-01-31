// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sweep_test

import (
	"context"
	"slices"

	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/service/accessanalyzer"
	"github.com/isometry/terraform-provider-faws/internal/service/account"
	"github.com/isometry/terraform-provider-faws/internal/service/acm"
	"github.com/isometry/terraform-provider-faws/internal/service/acmpca"
	"github.com/isometry/terraform-provider-faws/internal/service/amp"
	"github.com/isometry/terraform-provider-faws/internal/service/amplify"
	"github.com/isometry/terraform-provider-faws/internal/service/apigateway"
	"github.com/isometry/terraform-provider-faws/internal/service/apigatewayv2"
	"github.com/isometry/terraform-provider-faws/internal/service/appautoscaling"
	"github.com/isometry/terraform-provider-faws/internal/service/appconfig"
	"github.com/isometry/terraform-provider-faws/internal/service/appfabric"
	"github.com/isometry/terraform-provider-faws/internal/service/appflow"
	"github.com/isometry/terraform-provider-faws/internal/service/appintegrations"
	"github.com/isometry/terraform-provider-faws/internal/service/applicationinsights"
	"github.com/isometry/terraform-provider-faws/internal/service/applicationsignals"
	"github.com/isometry/terraform-provider-faws/internal/service/appmesh"
	"github.com/isometry/terraform-provider-faws/internal/service/apprunner"
	"github.com/isometry/terraform-provider-faws/internal/service/appstream"
	"github.com/isometry/terraform-provider-faws/internal/service/appsync"
	"github.com/isometry/terraform-provider-faws/internal/service/athena"
	"github.com/isometry/terraform-provider-faws/internal/service/auditmanager"
	"github.com/isometry/terraform-provider-faws/internal/service/autoscaling"
	"github.com/isometry/terraform-provider-faws/internal/service/autoscalingplans"
	"github.com/isometry/terraform-provider-faws/internal/service/backup"
	"github.com/isometry/terraform-provider-faws/internal/service/batch"
	"github.com/isometry/terraform-provider-faws/internal/service/bcmdataexports"
	"github.com/isometry/terraform-provider-faws/internal/service/bedrock"
	"github.com/isometry/terraform-provider-faws/internal/service/bedrockagent"
	"github.com/isometry/terraform-provider-faws/internal/service/billing"
	"github.com/isometry/terraform-provider-faws/internal/service/budgets"
	"github.com/isometry/terraform-provider-faws/internal/service/ce"
	"github.com/isometry/terraform-provider-faws/internal/service/chatbot"
	"github.com/isometry/terraform-provider-faws/internal/service/chime"
	"github.com/isometry/terraform-provider-faws/internal/service/chimesdkmediapipelines"
	"github.com/isometry/terraform-provider-faws/internal/service/chimesdkvoice"
	"github.com/isometry/terraform-provider-faws/internal/service/cleanrooms"
	"github.com/isometry/terraform-provider-faws/internal/service/cloud9"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudcontrol"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudformation"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudfront"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudfrontkeyvaluestore"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudhsmv2"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudsearch"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudtrail"
	"github.com/isometry/terraform-provider-faws/internal/service/cloudwatch"
	"github.com/isometry/terraform-provider-faws/internal/service/codeartifact"
	"github.com/isometry/terraform-provider-faws/internal/service/codebuild"
	"github.com/isometry/terraform-provider-faws/internal/service/codecatalyst"
	"github.com/isometry/terraform-provider-faws/internal/service/codecommit"
	"github.com/isometry/terraform-provider-faws/internal/service/codeconnections"
	"github.com/isometry/terraform-provider-faws/internal/service/codeguruprofiler"
	"github.com/isometry/terraform-provider-faws/internal/service/codegurureviewer"
	"github.com/isometry/terraform-provider-faws/internal/service/codepipeline"
	"github.com/isometry/terraform-provider-faws/internal/service/codestarconnections"
	"github.com/isometry/terraform-provider-faws/internal/service/codestarnotifications"
	"github.com/isometry/terraform-provider-faws/internal/service/cognitoidentity"
	"github.com/isometry/terraform-provider-faws/internal/service/cognitoidp"
	"github.com/isometry/terraform-provider-faws/internal/service/comprehend"
	"github.com/isometry/terraform-provider-faws/internal/service/computeoptimizer"
	"github.com/isometry/terraform-provider-faws/internal/service/configservice"
	"github.com/isometry/terraform-provider-faws/internal/service/connect"
	"github.com/isometry/terraform-provider-faws/internal/service/connectcases"
	"github.com/isometry/terraform-provider-faws/internal/service/controltower"
	"github.com/isometry/terraform-provider-faws/internal/service/costoptimizationhub"
	"github.com/isometry/terraform-provider-faws/internal/service/cur"
	"github.com/isometry/terraform-provider-faws/internal/service/customerprofiles"
	"github.com/isometry/terraform-provider-faws/internal/service/databrew"
	"github.com/isometry/terraform-provider-faws/internal/service/dataexchange"
	"github.com/isometry/terraform-provider-faws/internal/service/datapipeline"
	"github.com/isometry/terraform-provider-faws/internal/service/datasync"
	"github.com/isometry/terraform-provider-faws/internal/service/datazone"
	"github.com/isometry/terraform-provider-faws/internal/service/dax"
	"github.com/isometry/terraform-provider-faws/internal/service/deploy"
	"github.com/isometry/terraform-provider-faws/internal/service/detective"
	"github.com/isometry/terraform-provider-faws/internal/service/devicefarm"
	"github.com/isometry/terraform-provider-faws/internal/service/devopsguru"
	"github.com/isometry/terraform-provider-faws/internal/service/directconnect"
	"github.com/isometry/terraform-provider-faws/internal/service/dlm"
	"github.com/isometry/terraform-provider-faws/internal/service/dms"
	"github.com/isometry/terraform-provider-faws/internal/service/docdb"
	"github.com/isometry/terraform-provider-faws/internal/service/docdbelastic"
	"github.com/isometry/terraform-provider-faws/internal/service/drs"
	"github.com/isometry/terraform-provider-faws/internal/service/ds"
	"github.com/isometry/terraform-provider-faws/internal/service/dynamodb"
	"github.com/isometry/terraform-provider-faws/internal/service/ec2"
	"github.com/isometry/terraform-provider-faws/internal/service/ecr"
	"github.com/isometry/terraform-provider-faws/internal/service/ecrpublic"
	"github.com/isometry/terraform-provider-faws/internal/service/ecs"
	"github.com/isometry/terraform-provider-faws/internal/service/efs"
	"github.com/isometry/terraform-provider-faws/internal/service/eks"
	"github.com/isometry/terraform-provider-faws/internal/service/elasticache"
	"github.com/isometry/terraform-provider-faws/internal/service/elasticbeanstalk"
	"github.com/isometry/terraform-provider-faws/internal/service/elasticsearch"
	"github.com/isometry/terraform-provider-faws/internal/service/elastictranscoder"
	"github.com/isometry/terraform-provider-faws/internal/service/elb"
	"github.com/isometry/terraform-provider-faws/internal/service/elbv2"
	"github.com/isometry/terraform-provider-faws/internal/service/emr"
	"github.com/isometry/terraform-provider-faws/internal/service/emrcontainers"
	"github.com/isometry/terraform-provider-faws/internal/service/emrserverless"
	"github.com/isometry/terraform-provider-faws/internal/service/events"
	"github.com/isometry/terraform-provider-faws/internal/service/evidently"
	"github.com/isometry/terraform-provider-faws/internal/service/finspace"
	"github.com/isometry/terraform-provider-faws/internal/service/firehose"
	"github.com/isometry/terraform-provider-faws/internal/service/fis"
	"github.com/isometry/terraform-provider-faws/internal/service/fms"
	"github.com/isometry/terraform-provider-faws/internal/service/fsx"
	"github.com/isometry/terraform-provider-faws/internal/service/gamelift"
	"github.com/isometry/terraform-provider-faws/internal/service/glacier"
	"github.com/isometry/terraform-provider-faws/internal/service/globalaccelerator"
	"github.com/isometry/terraform-provider-faws/internal/service/glue"
	"github.com/isometry/terraform-provider-faws/internal/service/grafana"
	"github.com/isometry/terraform-provider-faws/internal/service/greengrass"
	"github.com/isometry/terraform-provider-faws/internal/service/groundstation"
	"github.com/isometry/terraform-provider-faws/internal/service/guardduty"
	"github.com/isometry/terraform-provider-faws/internal/service/healthlake"
	"github.com/isometry/terraform-provider-faws/internal/service/iam"
	"github.com/isometry/terraform-provider-faws/internal/service/identitystore"
	"github.com/isometry/terraform-provider-faws/internal/service/imagebuilder"
	"github.com/isometry/terraform-provider-faws/internal/service/inspector"
	"github.com/isometry/terraform-provider-faws/internal/service/inspector2"
	"github.com/isometry/terraform-provider-faws/internal/service/internetmonitor"
	"github.com/isometry/terraform-provider-faws/internal/service/invoicing"
	"github.com/isometry/terraform-provider-faws/internal/service/iot"
	"github.com/isometry/terraform-provider-faws/internal/service/iotanalytics"
	"github.com/isometry/terraform-provider-faws/internal/service/iotevents"
	"github.com/isometry/terraform-provider-faws/internal/service/ivs"
	"github.com/isometry/terraform-provider-faws/internal/service/ivschat"
	"github.com/isometry/terraform-provider-faws/internal/service/kafka"
	"github.com/isometry/terraform-provider-faws/internal/service/kafkaconnect"
	"github.com/isometry/terraform-provider-faws/internal/service/kendra"
	"github.com/isometry/terraform-provider-faws/internal/service/keyspaces"
	"github.com/isometry/terraform-provider-faws/internal/service/kinesis"
	"github.com/isometry/terraform-provider-faws/internal/service/kinesisanalytics"
	"github.com/isometry/terraform-provider-faws/internal/service/kinesisanalyticsv2"
	"github.com/isometry/terraform-provider-faws/internal/service/kinesisvideo"
	"github.com/isometry/terraform-provider-faws/internal/service/kms"
	"github.com/isometry/terraform-provider-faws/internal/service/lakeformation"
	"github.com/isometry/terraform-provider-faws/internal/service/lambda"
	"github.com/isometry/terraform-provider-faws/internal/service/launchwizard"
	"github.com/isometry/terraform-provider-faws/internal/service/lexmodels"
	"github.com/isometry/terraform-provider-faws/internal/service/lexv2models"
	"github.com/isometry/terraform-provider-faws/internal/service/licensemanager"
	"github.com/isometry/terraform-provider-faws/internal/service/lightsail"
	"github.com/isometry/terraform-provider-faws/internal/service/location"
	"github.com/isometry/terraform-provider-faws/internal/service/logs"
	"github.com/isometry/terraform-provider-faws/internal/service/lookoutmetrics"
	"github.com/isometry/terraform-provider-faws/internal/service/m2"
	"github.com/isometry/terraform-provider-faws/internal/service/macie2"
	"github.com/isometry/terraform-provider-faws/internal/service/mediaconnect"
	"github.com/isometry/terraform-provider-faws/internal/service/mediaconvert"
	"github.com/isometry/terraform-provider-faws/internal/service/medialive"
	"github.com/isometry/terraform-provider-faws/internal/service/mediapackage"
	"github.com/isometry/terraform-provider-faws/internal/service/mediapackagev2"
	"github.com/isometry/terraform-provider-faws/internal/service/mediastore"
	"github.com/isometry/terraform-provider-faws/internal/service/memorydb"
	"github.com/isometry/terraform-provider-faws/internal/service/meta"
	"github.com/isometry/terraform-provider-faws/internal/service/mgn"
	"github.com/isometry/terraform-provider-faws/internal/service/mq"
	"github.com/isometry/terraform-provider-faws/internal/service/mwaa"
	"github.com/isometry/terraform-provider-faws/internal/service/neptune"
	"github.com/isometry/terraform-provider-faws/internal/service/neptunegraph"
	"github.com/isometry/terraform-provider-faws/internal/service/networkfirewall"
	"github.com/isometry/terraform-provider-faws/internal/service/networkmanager"
	"github.com/isometry/terraform-provider-faws/internal/service/networkmonitor"
	"github.com/isometry/terraform-provider-faws/internal/service/oam"
	"github.com/isometry/terraform-provider-faws/internal/service/opensearch"
	"github.com/isometry/terraform-provider-faws/internal/service/opensearchserverless"
	"github.com/isometry/terraform-provider-faws/internal/service/opsworks"
	"github.com/isometry/terraform-provider-faws/internal/service/organizations"
	"github.com/isometry/terraform-provider-faws/internal/service/osis"
	"github.com/isometry/terraform-provider-faws/internal/service/outposts"
	"github.com/isometry/terraform-provider-faws/internal/service/paymentcryptography"
	"github.com/isometry/terraform-provider-faws/internal/service/pcaconnectorad"
	"github.com/isometry/terraform-provider-faws/internal/service/pcs"
	"github.com/isometry/terraform-provider-faws/internal/service/pinpoint"
	"github.com/isometry/terraform-provider-faws/internal/service/pinpointsmsvoicev2"
	"github.com/isometry/terraform-provider-faws/internal/service/pipes"
	"github.com/isometry/terraform-provider-faws/internal/service/polly"
	"github.com/isometry/terraform-provider-faws/internal/service/pricing"
	"github.com/isometry/terraform-provider-faws/internal/service/qbusiness"
	"github.com/isometry/terraform-provider-faws/internal/service/qldb"
	"github.com/isometry/terraform-provider-faws/internal/service/quicksight"
	"github.com/isometry/terraform-provider-faws/internal/service/ram"
	"github.com/isometry/terraform-provider-faws/internal/service/rbin"
	"github.com/isometry/terraform-provider-faws/internal/service/rds"
	"github.com/isometry/terraform-provider-faws/internal/service/redshift"
	"github.com/isometry/terraform-provider-faws/internal/service/redshiftdata"
	"github.com/isometry/terraform-provider-faws/internal/service/redshiftserverless"
	"github.com/isometry/terraform-provider-faws/internal/service/rekognition"
	"github.com/isometry/terraform-provider-faws/internal/service/resiliencehub"
	"github.com/isometry/terraform-provider-faws/internal/service/resourceexplorer2"
	"github.com/isometry/terraform-provider-faws/internal/service/resourcegroups"
	"github.com/isometry/terraform-provider-faws/internal/service/resourcegroupstaggingapi"
	"github.com/isometry/terraform-provider-faws/internal/service/rolesanywhere"
	"github.com/isometry/terraform-provider-faws/internal/service/route53"
	"github.com/isometry/terraform-provider-faws/internal/service/route53domains"
	"github.com/isometry/terraform-provider-faws/internal/service/route53profiles"
	"github.com/isometry/terraform-provider-faws/internal/service/route53recoverycontrolconfig"
	"github.com/isometry/terraform-provider-faws/internal/service/route53recoveryreadiness"
	"github.com/isometry/terraform-provider-faws/internal/service/route53resolver"
	"github.com/isometry/terraform-provider-faws/internal/service/rum"
	"github.com/isometry/terraform-provider-faws/internal/service/s3"
	"github.com/isometry/terraform-provider-faws/internal/service/s3control"
	"github.com/isometry/terraform-provider-faws/internal/service/s3outposts"
	"github.com/isometry/terraform-provider-faws/internal/service/s3tables"
	"github.com/isometry/terraform-provider-faws/internal/service/sagemaker"
	"github.com/isometry/terraform-provider-faws/internal/service/scheduler"
	"github.com/isometry/terraform-provider-faws/internal/service/schemas"
	"github.com/isometry/terraform-provider-faws/internal/service/secretsmanager"
	"github.com/isometry/terraform-provider-faws/internal/service/securityhub"
	"github.com/isometry/terraform-provider-faws/internal/service/securitylake"
	"github.com/isometry/terraform-provider-faws/internal/service/serverlessrepo"
	"github.com/isometry/terraform-provider-faws/internal/service/servicecatalog"
	"github.com/isometry/terraform-provider-faws/internal/service/servicecatalogappregistry"
	"github.com/isometry/terraform-provider-faws/internal/service/servicediscovery"
	"github.com/isometry/terraform-provider-faws/internal/service/servicequotas"
	"github.com/isometry/terraform-provider-faws/internal/service/ses"
	"github.com/isometry/terraform-provider-faws/internal/service/sesv2"
	"github.com/isometry/terraform-provider-faws/internal/service/sfn"
	"github.com/isometry/terraform-provider-faws/internal/service/shield"
	"github.com/isometry/terraform-provider-faws/internal/service/signer"
	"github.com/isometry/terraform-provider-faws/internal/service/simpledb"
	"github.com/isometry/terraform-provider-faws/internal/service/sns"
	"github.com/isometry/terraform-provider-faws/internal/service/sqs"
	"github.com/isometry/terraform-provider-faws/internal/service/ssm"
	"github.com/isometry/terraform-provider-faws/internal/service/ssmcontacts"
	"github.com/isometry/terraform-provider-faws/internal/service/ssmincidents"
	"github.com/isometry/terraform-provider-faws/internal/service/ssmquicksetup"
	"github.com/isometry/terraform-provider-faws/internal/service/ssmsap"
	"github.com/isometry/terraform-provider-faws/internal/service/sso"
	"github.com/isometry/terraform-provider-faws/internal/service/ssoadmin"
	"github.com/isometry/terraform-provider-faws/internal/service/storagegateway"
	"github.com/isometry/terraform-provider-faws/internal/service/sts"
	"github.com/isometry/terraform-provider-faws/internal/service/swf"
	"github.com/isometry/terraform-provider-faws/internal/service/synthetics"
	"github.com/isometry/terraform-provider-faws/internal/service/taxsettings"
	"github.com/isometry/terraform-provider-faws/internal/service/timestreaminfluxdb"
	"github.com/isometry/terraform-provider-faws/internal/service/timestreamquery"
	"github.com/isometry/terraform-provider-faws/internal/service/timestreamwrite"
	"github.com/isometry/terraform-provider-faws/internal/service/transcribe"
	"github.com/isometry/terraform-provider-faws/internal/service/transfer"
	"github.com/isometry/terraform-provider-faws/internal/service/verifiedpermissions"
	"github.com/isometry/terraform-provider-faws/internal/service/vpclattice"
	"github.com/isometry/terraform-provider-faws/internal/service/waf"
	"github.com/isometry/terraform-provider-faws/internal/service/wafregional"
	"github.com/isometry/terraform-provider-faws/internal/service/wafv2"
	"github.com/isometry/terraform-provider-faws/internal/service/wellarchitected"
	"github.com/isometry/terraform-provider-faws/internal/service/worklink"
	"github.com/isometry/terraform-provider-faws/internal/service/workspaces"
	"github.com/isometry/terraform-provider-faws/internal/service/workspacesweb"
	"github.com/isometry/terraform-provider-faws/internal/service/xray"
)

func servicePackages(ctx context.Context) []conns.ServicePackage {
	v := []conns.ServicePackage{
		accessanalyzer.ServicePackage(ctx),
		account.ServicePackage(ctx),
		acm.ServicePackage(ctx),
		acmpca.ServicePackage(ctx),
		amp.ServicePackage(ctx),
		amplify.ServicePackage(ctx),
		apigateway.ServicePackage(ctx),
		apigatewayv2.ServicePackage(ctx),
		appautoscaling.ServicePackage(ctx),
		appconfig.ServicePackage(ctx),
		appfabric.ServicePackage(ctx),
		appflow.ServicePackage(ctx),
		appintegrations.ServicePackage(ctx),
		applicationinsights.ServicePackage(ctx),
		applicationsignals.ServicePackage(ctx),
		appmesh.ServicePackage(ctx),
		apprunner.ServicePackage(ctx),
		appstream.ServicePackage(ctx),
		appsync.ServicePackage(ctx),
		athena.ServicePackage(ctx),
		auditmanager.ServicePackage(ctx),
		autoscaling.ServicePackage(ctx),
		autoscalingplans.ServicePackage(ctx),
		backup.ServicePackage(ctx),
		batch.ServicePackage(ctx),
		bcmdataexports.ServicePackage(ctx),
		bedrock.ServicePackage(ctx),
		bedrockagent.ServicePackage(ctx),
		billing.ServicePackage(ctx),
		budgets.ServicePackage(ctx),
		ce.ServicePackage(ctx),
		chatbot.ServicePackage(ctx),
		chime.ServicePackage(ctx),
		chimesdkmediapipelines.ServicePackage(ctx),
		chimesdkvoice.ServicePackage(ctx),
		cleanrooms.ServicePackage(ctx),
		cloud9.ServicePackage(ctx),
		cloudcontrol.ServicePackage(ctx),
		cloudformation.ServicePackage(ctx),
		cloudfront.ServicePackage(ctx),
		cloudfrontkeyvaluestore.ServicePackage(ctx),
		cloudhsmv2.ServicePackage(ctx),
		cloudsearch.ServicePackage(ctx),
		cloudtrail.ServicePackage(ctx),
		cloudwatch.ServicePackage(ctx),
		codeartifact.ServicePackage(ctx),
		codebuild.ServicePackage(ctx),
		codecatalyst.ServicePackage(ctx),
		codecommit.ServicePackage(ctx),
		codeconnections.ServicePackage(ctx),
		codeguruprofiler.ServicePackage(ctx),
		codegurureviewer.ServicePackage(ctx),
		codepipeline.ServicePackage(ctx),
		codestarconnections.ServicePackage(ctx),
		codestarnotifications.ServicePackage(ctx),
		cognitoidentity.ServicePackage(ctx),
		cognitoidp.ServicePackage(ctx),
		comprehend.ServicePackage(ctx),
		computeoptimizer.ServicePackage(ctx),
		configservice.ServicePackage(ctx),
		connect.ServicePackage(ctx),
		connectcases.ServicePackage(ctx),
		controltower.ServicePackage(ctx),
		costoptimizationhub.ServicePackage(ctx),
		cur.ServicePackage(ctx),
		customerprofiles.ServicePackage(ctx),
		databrew.ServicePackage(ctx),
		dataexchange.ServicePackage(ctx),
		datapipeline.ServicePackage(ctx),
		datasync.ServicePackage(ctx),
		datazone.ServicePackage(ctx),
		dax.ServicePackage(ctx),
		deploy.ServicePackage(ctx),
		detective.ServicePackage(ctx),
		devicefarm.ServicePackage(ctx),
		devopsguru.ServicePackage(ctx),
		directconnect.ServicePackage(ctx),
		dlm.ServicePackage(ctx),
		dms.ServicePackage(ctx),
		docdb.ServicePackage(ctx),
		docdbelastic.ServicePackage(ctx),
		drs.ServicePackage(ctx),
		ds.ServicePackage(ctx),
		dynamodb.ServicePackage(ctx),
		ec2.ServicePackage(ctx),
		ecr.ServicePackage(ctx),
		ecrpublic.ServicePackage(ctx),
		ecs.ServicePackage(ctx),
		efs.ServicePackage(ctx),
		eks.ServicePackage(ctx),
		elasticache.ServicePackage(ctx),
		elasticbeanstalk.ServicePackage(ctx),
		elasticsearch.ServicePackage(ctx),
		elastictranscoder.ServicePackage(ctx),
		elb.ServicePackage(ctx),
		elbv2.ServicePackage(ctx),
		emr.ServicePackage(ctx),
		emrcontainers.ServicePackage(ctx),
		emrserverless.ServicePackage(ctx),
		events.ServicePackage(ctx),
		evidently.ServicePackage(ctx),
		finspace.ServicePackage(ctx),
		firehose.ServicePackage(ctx),
		fis.ServicePackage(ctx),
		fms.ServicePackage(ctx),
		fsx.ServicePackage(ctx),
		gamelift.ServicePackage(ctx),
		glacier.ServicePackage(ctx),
		globalaccelerator.ServicePackage(ctx),
		glue.ServicePackage(ctx),
		grafana.ServicePackage(ctx),
		greengrass.ServicePackage(ctx),
		groundstation.ServicePackage(ctx),
		guardduty.ServicePackage(ctx),
		healthlake.ServicePackage(ctx),
		iam.ServicePackage(ctx),
		identitystore.ServicePackage(ctx),
		imagebuilder.ServicePackage(ctx),
		inspector.ServicePackage(ctx),
		inspector2.ServicePackage(ctx),
		internetmonitor.ServicePackage(ctx),
		invoicing.ServicePackage(ctx),
		iot.ServicePackage(ctx),
		iotanalytics.ServicePackage(ctx),
		iotevents.ServicePackage(ctx),
		ivs.ServicePackage(ctx),
		ivschat.ServicePackage(ctx),
		kafka.ServicePackage(ctx),
		kafkaconnect.ServicePackage(ctx),
		kendra.ServicePackage(ctx),
		keyspaces.ServicePackage(ctx),
		kinesis.ServicePackage(ctx),
		kinesisanalytics.ServicePackage(ctx),
		kinesisanalyticsv2.ServicePackage(ctx),
		kinesisvideo.ServicePackage(ctx),
		kms.ServicePackage(ctx),
		lakeformation.ServicePackage(ctx),
		lambda.ServicePackage(ctx),
		launchwizard.ServicePackage(ctx),
		lexmodels.ServicePackage(ctx),
		lexv2models.ServicePackage(ctx),
		licensemanager.ServicePackage(ctx),
		lightsail.ServicePackage(ctx),
		location.ServicePackage(ctx),
		logs.ServicePackage(ctx),
		lookoutmetrics.ServicePackage(ctx),
		m2.ServicePackage(ctx),
		macie2.ServicePackage(ctx),
		mediaconnect.ServicePackage(ctx),
		mediaconvert.ServicePackage(ctx),
		medialive.ServicePackage(ctx),
		mediapackage.ServicePackage(ctx),
		mediapackagev2.ServicePackage(ctx),
		mediastore.ServicePackage(ctx),
		memorydb.ServicePackage(ctx),
		meta.ServicePackage(ctx),
		mgn.ServicePackage(ctx),
		mq.ServicePackage(ctx),
		mwaa.ServicePackage(ctx),
		neptune.ServicePackage(ctx),
		neptunegraph.ServicePackage(ctx),
		networkfirewall.ServicePackage(ctx),
		networkmanager.ServicePackage(ctx),
		networkmonitor.ServicePackage(ctx),
		oam.ServicePackage(ctx),
		opensearch.ServicePackage(ctx),
		opensearchserverless.ServicePackage(ctx),
		opsworks.ServicePackage(ctx),
		organizations.ServicePackage(ctx),
		osis.ServicePackage(ctx),
		outposts.ServicePackage(ctx),
		paymentcryptography.ServicePackage(ctx),
		pcaconnectorad.ServicePackage(ctx),
		pcs.ServicePackage(ctx),
		pinpoint.ServicePackage(ctx),
		pinpointsmsvoicev2.ServicePackage(ctx),
		pipes.ServicePackage(ctx),
		polly.ServicePackage(ctx),
		pricing.ServicePackage(ctx),
		qbusiness.ServicePackage(ctx),
		qldb.ServicePackage(ctx),
		quicksight.ServicePackage(ctx),
		ram.ServicePackage(ctx),
		rbin.ServicePackage(ctx),
		rds.ServicePackage(ctx),
		redshift.ServicePackage(ctx),
		redshiftdata.ServicePackage(ctx),
		redshiftserverless.ServicePackage(ctx),
		rekognition.ServicePackage(ctx),
		resiliencehub.ServicePackage(ctx),
		resourceexplorer2.ServicePackage(ctx),
		resourcegroups.ServicePackage(ctx),
		resourcegroupstaggingapi.ServicePackage(ctx),
		rolesanywhere.ServicePackage(ctx),
		route53.ServicePackage(ctx),
		route53domains.ServicePackage(ctx),
		route53profiles.ServicePackage(ctx),
		route53recoverycontrolconfig.ServicePackage(ctx),
		route53recoveryreadiness.ServicePackage(ctx),
		route53resolver.ServicePackage(ctx),
		rum.ServicePackage(ctx),
		s3.ServicePackage(ctx),
		s3control.ServicePackage(ctx),
		s3outposts.ServicePackage(ctx),
		s3tables.ServicePackage(ctx),
		sagemaker.ServicePackage(ctx),
		scheduler.ServicePackage(ctx),
		schemas.ServicePackage(ctx),
		secretsmanager.ServicePackage(ctx),
		securityhub.ServicePackage(ctx),
		securitylake.ServicePackage(ctx),
		serverlessrepo.ServicePackage(ctx),
		servicecatalog.ServicePackage(ctx),
		servicecatalogappregistry.ServicePackage(ctx),
		servicediscovery.ServicePackage(ctx),
		servicequotas.ServicePackage(ctx),
		ses.ServicePackage(ctx),
		sesv2.ServicePackage(ctx),
		sfn.ServicePackage(ctx),
		shield.ServicePackage(ctx),
		signer.ServicePackage(ctx),
		simpledb.ServicePackage(ctx),
		sns.ServicePackage(ctx),
		sqs.ServicePackage(ctx),
		ssm.ServicePackage(ctx),
		ssmcontacts.ServicePackage(ctx),
		ssmincidents.ServicePackage(ctx),
		ssmquicksetup.ServicePackage(ctx),
		ssmsap.ServicePackage(ctx),
		sso.ServicePackage(ctx),
		ssoadmin.ServicePackage(ctx),
		storagegateway.ServicePackage(ctx),
		sts.ServicePackage(ctx),
		swf.ServicePackage(ctx),
		synthetics.ServicePackage(ctx),
		taxsettings.ServicePackage(ctx),
		timestreaminfluxdb.ServicePackage(ctx),
		timestreamquery.ServicePackage(ctx),
		timestreamwrite.ServicePackage(ctx),
		transcribe.ServicePackage(ctx),
		transfer.ServicePackage(ctx),
		verifiedpermissions.ServicePackage(ctx),
		vpclattice.ServicePackage(ctx),
		waf.ServicePackage(ctx),
		wafregional.ServicePackage(ctx),
		wafv2.ServicePackage(ctx),
		wellarchitected.ServicePackage(ctx),
		worklink.ServicePackage(ctx),
		workspaces.ServicePackage(ctx),
		workspacesweb.ServicePackage(ctx),
		xray.ServicePackage(ctx),
	}

	return slices.Clone(v)
}
