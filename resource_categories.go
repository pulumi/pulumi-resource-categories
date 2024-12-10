// Copyright 2024, Pulumi Corporation.  All rights reserved.

package main

import (
	"slices"
	"strings"
)

type ResourceKind string

const (
	API           ResourceKind = "api"
	Data          ResourceKind = "data"    // databases
	Storage       ResourceKind = "storage" // blob, file, block, object
	Compute       ResourceKind = "compute"
	Container     ResourceKind = "container" // more specific than compute
	Security      ResourceKind = "security"
	Identity      ResourceKind = "identity"
	Management    ResourceKind = "management" // tools for managing resources
	Network       ResourceKind = "network"
	Messaging     ResourceKind = "messaging"
	Timer         ResourceKind = "timer"
	Observability ResourceKind = "observability"   // logs, analytics
	Build         ResourceKind = "build"           // dev tools, deployment
	ML            ResourceKind = "machinelearning" // machine learning, ai
	Media         ResourceKind = "media"           // media services
	// catch-all for resources that don't fit into any of the above categories.
	Misc     ResourceKind = "misc" //
	Unknown  ResourceKind = "unknown"
	NotFound ResourceKind = "not-found" // returned when the resource type is not found in our list
)

// A map of types to kinds.
var typeToKindMap = map[string]ResourceKind{
	// used for @pulumi/aws and @pulumi/aws-native
	"aws:acmpca":                    Security,
	"aws:aps":                       Observability,
	"aws:arczonalshift":             Network,
	"aws:accessanalyzer":            Management,
	"aws:amazonmq":                  Messaging,
	"aws:amplify":                   Build,
	"aws:amplifyuibuilder":          Build,
	"aws:acm":                       Security,
	"aws:apigateway":                Network,
	"aws:apigatewayv2":              Network,
	"aws:appautoscaling":            Compute,
	"aws:applicationloadbalancing":  Network,
	"aws:alb":                       Network,
	"aws:appconfig":                 Build,
	"aws:appflow":                   Management,
	"aws:appintegrations":           Management,
	"aws:appmesh":                   Network,
	"aws:apprunner":                 Compute,
	"aws:appstream":                 Misc,
	"aws:appsync":                   Build,
	"aws:applicationautoscaling":    Network,
	"aws:applicationinsights":       Unknown,
	"aws:applicationsignals":        Unknown,
	"aws:athena":                    Data,
	"aws:auditmanager":              Security,
	"aws:autoscaling":               Network,
	"aws:autoscalingplans":          Unknown,
	"aws:b2bi":                      Unknown,
	"aws:backup":                    Data,
	"aws:backupgateway":             Unknown,
	"aws:batch":                     Compute,
	"aws:bedrock":                   ML,
	"aws:budgets":                   Unknown,
	"aws:ce":                        Unknown,
	"aws:cassandra":                 Data,
	"aws:certificatemanager":        Security,
	"aws:chatbot":                   Management,
	"aws:cleanrooms":                Data,
	"aws:cleanroomsml":              Data,
	"aws:cloud9":                    Build,
	"aws:cloudformation":            Management,
	"aws:cloudfront":                Network,
	"aws:cloudtrail":                Observability,
	"aws:cloudwatch":                Observability,
	"aws:codeartifact":              Build,
	"aws:codebuild":                 Management,
	"aws:codecommit":                Management,
	"aws:codeconnections":           Unknown,
	"aws:codedeploy":                Management,
	"aws:codeguruprofiler":          ML,
	"aws:codegurureviewer":          ML,
	"aws:codepipeline":              Build,
	"aws:codestar":                  Build,
	"aws:codestarconnections":       Build,
	"aws:codestarnotifications":     Build,
	"aws:cognito":                   Identity,
	"aws:comprehend":                Unknown,
	"aws:config":                    Management,
	"aws:connect":                   Unknown,
	"aws:connectcampaigns":          Unknown,
	"aws:controltower":              Management,
	"aws:customerprofiles":          Unknown,
	"aws:dax":                       Data,
	"aws:dlm":                       Unknown,
	"aws:dms":                       Data,
	"aws:databrew":                  Unknown,
	"aws:datapipeline":              Unknown,
	"aws:datasync":                  Management,
	"aws:datazone":                  Unknown,
	"aws:deadline":                  Media,
	"aws:detective":                 Security,
	"aws:devopsguru":                Unknown,
	"aws:devicefarm":                Build,
	"aws:directconnect":             Network,
	"aws:directoryservice":          Identity,
	"aws:docdb":                     Data,
	"aws:docdbelastic":              Data,
	"aws:dynamodb":                  Data,
	"aws:ebs":                       Compute,
	"aws:ec2":                       Compute,
	"aws:ecr":                       Container,
	"aws:ecs":                       Container,
	"aws:efs":                       Data,
	"aws:eks":                       Container,
	"aws:emr":                       Data,
	"aws:emrcontainers":             Data,
	"aws:emrserverless":             Compute,
	"aws:elastictranscoder":         Media,
	"aws:elasticache":               Data,
	"aws:elasticbeanstalk":          Container,
	"aws:elasticloadbalancing":      Network,
	"aws:elasticloadbalancingv2":    Network,
	"aws:lb":                        Network,
	"aws:elasticsearch":             Compute,
	"aws:entityresolution":          Unknown,
	"aws:eventschemas":              Unknown,
	"aws:events":                    Unknown,
	"aws:evidently":                 Unknown,
	"aws:fis":                       Build,
	"aws:fms":                       Unknown,
	"aws:fsx":                       Data,
	"aws:finspace":                  Data,
	"aws:forecast":                  ML,
	"aws:frauddetector":             ML,
	"aws:gamelift":                  Misc, // videogames
	"aws:globalaccelerator":         Network,
	"aws:glacier":                   Storage,
	"aws:glue":                      Data,
	"aws:grafana":                   Observability,
	"aws:greengrass":                Management,
	"aws:greengrassv2":              Unknown,
	"aws:groundstation":             Build,
	"aws:guardduty":                 Security,
	"aws:healthimaging":             ML,
	"aws:healthlake":                ML,
	"aws:iam":                       Identity,
	"aws:ivs":                       Media,
	"aws:ivschat":                   Media,
	"aws:identitystore":             Unknown,
	"aws:imagebuilder":              Unknown,
	"aws:inspector":                 Security,
	"aws:inspectorv2":               Security,
	"aws:internetmonitor":           Unknown,
	"aws:iot1click":                 Unknown,
	"aws:iot":                       Misc,
	"aws:iotanalytics":              Misc,
	"aws:iotcoredeviceadvisor":      Misc,
	"aws:iotevents":                 Misc,
	"aws:iotfleethub":               Misc,
	"aws:iotsitewise":               Misc,
	"aws:iotthingsgraph":            Misc,
	"aws:iottwinmaker":              Misc,
	"aws:iotwireless":               Misc,
	"aws:kms":                       Security,
	"aws:kafkaconnect":              Messaging,
	"aws:kendra":                    ML,
	"aws:kendraranking":             ML,
	"aws:kinesis":                   Messaging,
	"aws:kinesisanalytics":          Messaging,
	"aws:kinesisanalyticsv2":        Messaging,
	"aws:kinesisfirehose":           Messaging,
	"aws:kinesisvideo":              Media,
	"aws:lakeformation":             Storage,
	"aws:lambda":                    Compute,
	"aws:launchwizard":              Unknown,
	"aws:lex":                       ML,
	"aws:licensemanager":            Management,
	"aws:lightsail":                 Compute,
	"aws:location":                  Build,
	"aws:logs":                      Observability,
	"aws:lookoutmetrics":            ML,
	"aws:lookoutvision":             ML,
	"aws:m2":                        Unknown,
	"aws:msk":                       Unknown,
	"aws:mwaa":                      Management, // airflow
	"aws:macie":                     Security,
	"aws:managedblockchain":         Misc,
	"aws:mediaconnect":              Media,
	"aws:mediaconvert":              Media,
	"aws:medialive":                 Media,
	"aws:mediapackage":              Media,
	"aws:mediapackagev2":            Media,
	"aws:mediastore":                Media,
	"aws:mediatailor":               Media,
	"aws:memorydb":                  Data,
	"aws:mq":                        Messaging,
	"aws:neptune":                   Data,
	"aws:neptunegraph":              Data,
	"aws:networkfirewall":           Network,
	"aws:networkmanager":            Network,
	"aws:nimblestudio":              Media,
	"aws:osis":                      Unknown,
	"aws:oam":                       Unknown,
	"aws:omics":                     Unknown,
	"aws:opensearchserverless":      Data,
	"aws:opensearchservice":         Data,
	"aws:opsworks":                  Management,
	"aws:opsworkscm":                Management,
	"aws:organizations":             Management,
	"aws:pcaconnectorad":            Unknown,
	"aws:panorama":                  ML,
	"aws:paymentcryptography":       Security,
	"aws:personalize":               ML,
	"aws:pinpoint":                  Unknown,
	"aws:pinpointemail":             Unknown,
	"aws:pipes":                     Unknown,
	"aws:proton":                    Management,
	"aws:qbusiness":                 ML,
	"aws:qldb":                      Data,
	"aws:quicksight":                Unknown,
	"aws:ram":                       Unknown,
	"aws:rds":                       Data,
	"aws:rum":                       Unknown,
	"aws:redshift":                  Data,
	"aws:redshiftserverless":        Data,
	"aws:refactorspaces":            Unknown,
	"aws:rekognition":               ML,
	"aws:resiliencehub":             Management,
	"aws:resourceexplorer2":         Management,
	"aws:resourcegroups":            Unknown,
	"aws:robomaker":                 Misc,
	"aws:rolesanywhere":             Unknown,
	"aws:route53":                   Network,
	"aws:route53profiles":           Network,
	"aws:route53recoverycontrol":    Network,
	"aws:route53recoveryreadiness":  Network,
	"aws:route53resolver":           Network,
	"aws:s3":                        Storage,
	"aws:s3express":                 Storage,
	"aws:s3objectlambda":            Storage,
	"aws:s3outposts":                Storage,
	"aws:sdb":                       Unknown,
	"aws:ses":                       Management,
	"aws:sns":                       Messaging,
	"aws:sqs":                       Messaging,
	"aws:ssm":                       Management,
	"aws:ssmcontacts":               Management,
	"aws:ssmincidents":              Management,
	"aws:sso":                       Identity,
	"aws:sagemaker":                 ML,
	"aws:scheduler":                 Unknown,
	"aws:serverless":                Compute,
	"aws:secretsmanager":            Security,
	"aws:securityhub":               Security,
	"aws:securitylake":              Security,
	"aws:servicecatalog":            Management,
	"aws:servicecatalogappregistry": Management,
	"aws:servicediscovery":          Management,
	"aws:sfn":                       Compute,
	"aws:shield":                    Unknown,
	"aws:signer":                    Security,
	"aws:simpledb":                  Data,
	"aws:simspaceweaver":            Compute,
	"aws:stepfunctions":             Unknown,
	"aws:supportapp":                Unknown,
	"aws:synthetics":                Unknown,
	"aws:systemsmanagersap":         Management,
	"aws:timestream":                Data,
	"aws:transfer":                  Unknown,
	"aws:verifiedpermissions":       Unknown,
	"aws:voiceid":                   Unknown,
	"aws:vpclattice":                Network,
	"aws:waf":                       Network,
	"aws:wafregional":               Network,
	"aws:wafv2":                     Network,
	"aws:wisdom":                    Unknown,
	"aws:workspaces":                Misc,
	"aws:workspacesthinclient":      Misc,
	"aws:workspacesweb":             Misc,
	"aws:xray":                      Build,
	// -----------------

	// @pulumi/azure
	"azure:appinsights":            Observability,
	"azure:appservice":             Container,
	"azure:appservice:FunctionApp": Compute,
	"azure:automation":             Timer,
	"azure:cdn":                    Network,
	"azure:compute":                Compute,
	"azure:containerservice":       Container,
	"azure:core":                   Management,
	"azure:cosmosdb":               Data,
	"azure:dns":                    Network,
	"azure:eventhub":               Messaging,
	"azure:iot":                    Messaging,
	"azure:keyvault":               Identity,
	"azure:lb":                     Network,
	"azure:managementresource":     Management,
	"azure:monitoring":             Observability,
	"azure:mysql":                  Data,
	"azure:network":                Network,
	"azure:operationalinsights":    Management,
	"azure:policy":                 Management,
	"azure:postgresql":             Data,
	"azure:redis":                  Compute,
	"azure:role":                   Identity,
	"azure:scheduler":              Compute,
	"azure:search":                 Compute,
	"azure:sql":                    Data,
	"azure:storage":                Storage,
	"azure:trafficmanager":         Network,

	// @pulumi/cloud
	"cloud:bucket":       Storage,
	"cloud:http":         API,
	"cloud:function":     Compute,
	"cloud:logCollector": Observability,
	"cloud:service":      Container,
	"cloud:table":        Data,
	"cloud:task":         Container,
	"cloud:timer":        Timer,
	"cloud:topic":        Messaging,

	// @pulumi/gcp
	"gcp:cloudfunctions/function": Compute,
	"gcp:compute:address":         Network,
	"gcp:compute/disk":            Storage,
	"gcp:compute/firewall":        Network,
	"gcp:compute/forwardingRule":  Network,
	"gcp:compute/instance":        Compute,
	"gcp:compute/network":         Network,
	"gcp:compute/regionDisk":      Storage,
	"gcp:compute/router":          Network,
	"gcp:compute/routerInterface": Network,
	"gcp:compute/routerPeer":      Network,
	"gcp:compute/subnetwork":      Network,
	"gcp:compute/vPNTunnel":       Network,
	"gcp:container/cluster":       Container,
	"gcp:serverless:HttpFunction": Compute,
	"gcp:storage":                 Storage,

	// @pulumi/kubernetes
	"kubernetes:admissionregistration": Network,
	"kubernetes:apiextensions":         Management,
	"kubernetes:apiregistration":       API,
	"kubernetes:apps":                  Container,
	"kubernetes:authentication":        Identity,
	"kubernetes:authorization":         Identity,
	"kubernetes:autoscaling":           Container,
	"kubernetes:batch":                 Timer,
	"kubernetes:certificates":          Identity,
	"kubernetes:core":                  Container,
	"kubernetes:events":                Messaging,
	"kubernetes:extensions":            Container,
	"kubernetes:meta":                  Management,
	"kubernetes:networking":            Network,
	"kubernetes:policy":                Management,
	"kubernetes:rbac":                  Identity,
	"kubernetes:scheduling":            Container,
	"kubernetes:settings":              Management,
	"kubernetes:storage":               Storage,
}

// GetResourceKind returns the kind of the given type.
func GetResourceKind(resourceType string) ResourceKind {
	t := strings.ToLower(resourceType)

	// if prefix is aws-native, use aws: prefix
	if strings.HasPrefix(t, "aws-native:") {
		t = "aws:" + t[len("aws-native:"):]
	}

	supportedProviders := []string{
		"aws",
		"azure",
		"kubernetes",
		"cloud",
	}
	isSupportedProvider := slices.ContainsFunc(supportedProviders, func(provider string) bool {
		return strings.HasPrefix(t, provider+":")
	})

	if !isSupportedProvider {
		return "" // Displayed as empty in the UI, otherwise they will fallback to "unknown"
	}

	var kind = NotFound
	for t != "" {
		// See if there's a direct hit for the type we're seeking.
		if k, ok := typeToKindMap[t]; ok {
			kind = k
			break
		}

		// If there's no direct hit, keep peeling off parts until we get a hit or run out of string.
		colonIndex := strings.LastIndex(t, ":")
		slashIndex := strings.LastIndex(t, "/")
		if colonIndex == -1 && slashIndex == -1 {
			break
		}

		// Strip off up to the most specific part.
		t = t[:max(colonIndex, slashIndex)]
	}

	return kind
}