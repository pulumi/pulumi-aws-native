package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <languages> <schema-file> <root-dir>"
		_, err := fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		contract.IgnoreError(err)
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", "", "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		flag.Usage()
		return
	}

	languages, inputFile, version := args[0], args[1], args[2]

	pkgSpec := gatherPackage(readCFNSchema(inputFile), readSupportedResourceTypes())
	pkgSpec.Version = version
	ppkg, err := pschema.ImportSpec(pkgSpec, nil)
	if err != nil {
		panic(fmt.Sprintf("error importing schema: %v", err))
	}

	for _, language := range strings.Split(languages, ",") {
		outdir := path.Join(".", "sdk", language)
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")

		switch language {
		case "nodejs":
			writeNodeJSSDK(ppkg, outdir)
		case "python":
			writePythonSDK(ppkg, outdir)
		case "dotnet":
			writeDotnetSDK(ppkg, outdir)
		case "go":
			writeGoSDK(ppkg, outdir)
		case "schema":
			writePulumiSchema(pkgSpec, providerDir)
		default:
			panic(fmt.Sprintf("Unrecognized language '%s'", language))
		}
	}
}

func readCFNSchema(schemaPath string) schema.CloudFormationSchema {
	// Read in, decode, and import the schema.
	schemaBytes, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		panic(err)
	}

	var sch schema.CloudFormationSchema
	if err = json.Unmarshal(schemaBytes, &sch); err != nil {
		panic(err)
	}

	return sch
}

func readSupportedResourceTypes() []string {
	return strings.Split(supportedResourceTypesCsv, "\n")
}

func writeNodeJSSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := nodejsgen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writePythonSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := pythongen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writeDotnetSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := dotnetgen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writeGoSDK(pkg *pschema.Package, outdir string) {
	files, err := gogen.GeneratePackage("pulumigen", pkg)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writePulumiSchema(pkgSpec pschema.PackageSpec, outdir string) {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		panic(errors.Wrap(err, "marshaling Pulumi schema"))
	}
	mustWriteFile(outdir, "schema.json", schemaJSON)
}

func mustWriteFiles(rootDir string, files map[string][]byte) {
	for filename, contents := range files {
		mustWriteFile(rootDir, filename, contents)
	}
}

func mustWriteFile(rootDir, filename string, contents []byte) {
	outPath := filepath.Join(rootDir, filename)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		panic(err)
	}
	err := ioutil.WriteFile(outPath, contents, 0600)
	if err != nil {
		panic(err)
	}
}

const supportedResourceTypesCsv = `AWS::ACMPCA::Certificate
AWS::ACMPCA::CertificateAuthority
AWS::ACMPCA::CertificateAuthorityActivation
AWS::AccessAnalyzer::Analyzer
AWS::ApiGateway::ApiKey
AWS::ApiGateway::ClientCertificate
AWS::ApiGateway::DocumentationVersion
AWS::ApiGateway::DomainName
AWS::AppFlow::ConnectorProfile
AWS::AppFlow::Flow
AWS::AppIntegrations::EventIntegration
AWS::ApplicationInsights::Application
AWS::Athena::DataCatalog
AWS::Athena::NamedQuery
AWS::Athena::WorkGroup
AWS::AuditManager::Assessment
AWS::AutoScaling::WarmPool
AWS::Backup::BackupPlan
AWS::Backup::BackupVault
AWS::Budgets::BudgetsAction
AWS::CE::AnomalyMonitor
AWS::CE::AnomalySubscription
AWS::CE::CostCategory
AWS::Cassandra::Keyspace
AWS::Cassandra::Table
AWS::CertificateManager::Account
AWS::Backup::BackupSelection
AWS::Chatbot::SlackChannelConfiguration
AWS::CloudFormation::ResourceDefaultVersion
AWS::CloudFormation::StackSet
AWS::CloudFront::CachePolicy
AWS::CloudFront::CloudFrontOriginAccessIdentity
AWS::CloudFront::KeyGroup
AWS::CloudFront::OriginRequestPolicy
AWS::CloudFront::PublicKey
AWS::CloudFront::RealtimeLogConfig
AWS::CloudWatch::CompositeAlarm
AWS::CloudWatch::MetricStream
AWS::CodeArtifact::Domain
AWS::CodeArtifact::Repository
AWS::CodeGuruProfiler::ProfilingGroup
AWS::CodeStarConnections::Connection
AWS::Config::ConformancePack
AWS::Config::OrganizationConformancePack
AWS::Config::StoredQuery
AWS::CustomerProfiles::Domain
AWS::CustomerProfiles::Integration
AWS::CustomerProfiles::ObjectType
AWS::DataBrew::Dataset
AWS::DataBrew::Job
AWS::DataBrew::Project
AWS::DataBrew::Recipe
AWS::DataBrew::Schedule
AWS::DataSync::Agent
AWS::DataSync::LocationEFS
AWS::DataSync::LocationFSxWindows
AWS::DataSync::LocationNFS
AWS::DataSync::LocationObjectStorage
AWS::DataSync::LocationS3
AWS::DataSync::LocationSMB
AWS::DataSync::Task
AWS::Detective::Graph
AWS::Detective::MemberInvitation
AWS::DevOpsGuru::ResourceCollection
AWS::CloudFormation::ModuleDefaultVersion
AWS::CloudFormation::ModuleVersion
AWS::CloudFormation::ResourceVersion
AWS::CodeGuruReviewer::RepositoryAssociation
AWS::DevOpsGuru::NotificationChannel
AWS::EC2::CarrierGateway
AWS::EC2::FlowLog
AWS::EC2::GatewayRouteTableAssociation
AWS::EC2::LocalGatewayRoute
AWS::EC2::LocalGatewayRouteTableVPCAssociation
AWS::EC2::NetworkInsightsAnalysis
AWS::EC2::NetworkInsightsPath
AWS::EC2::PrefixList
AWS::EC2::TransitGatewayConnect
AWS::EC2::TransitGatewayMulticastDomain
AWS::ECR::PublicRepository
AWS::ECR::RegistryPolicy
AWS::ECR::ReplicationConfiguration
AWS::ECR::Repository
AWS::ECS::CapacityProvider
AWS::ECS::Cluster
AWS::ECS::ClusterCapacityProviderAssociations
AWS::ECS::PrimaryTaskSet
AWS::ECS::Service
AWS::ECS::TaskDefinition
AWS::ECS::TaskSet
AWS::EFS::AccessPoint
AWS::EFS::FileSystem
AWS::EKS::Addon
AWS::EKS::FargateProfile
AWS::EMR::Studio
AWS::EC2::EnclaveCertificateIamRoleAssociation
AWS::EC2::TransitGatewayMulticastDomainAssociation
AWS::EC2::TransitGatewayMulticastGroupMember
AWS::EC2::TransitGatewayMulticastGroupSource
AWS::EMR::StudioSessionMapping
AWS::EMRContainers::VirtualCluster
AWS::ElastiCache::GlobalReplicationGroup
AWS::ElastiCache::User
AWS::ElastiCache::UserGroup
AWS::ElasticLoadBalancingV2::Listener
AWS::ElasticLoadBalancingV2::ListenerRule
AWS::EventSchemas::RegistryPolicy
AWS::Events::ApiDestination
AWS::Events::Archive
AWS::Events::Connection
AWS::FIS::ExperimentTemplate
AWS::FMS::NotificationChannel
AWS::FMS::Policy
AWS::GameLift::Alias
AWS::GameLift::Fleet
AWS::GameLift::GameServerGroup
AWS::GlobalAccelerator::Accelerator
AWS::GlobalAccelerator::EndpointGroup
AWS::GlobalAccelerator::Listener
AWS::Glue::Registry
AWS::Glue::Schema
AWS::GreengrassV2::ComponentVersion
AWS::GroundStation::Config
AWS::GroundStation::MissionProfile
AWS::IAM::OIDCProvider
AWS::IAM::SAMLProvider
AWS::IAM::ServerCertificate
AWS::Glue::SchemaVersion
AWS::Glue::SchemaVersionMetadata
AWS::GroundStation::DataflowEndpointGroup
AWS::IAM::VirtualMFADevice
AWS::IVS::Channel
AWS::IVS::PlaybackKeyPair
AWS::IVS::RecordingConfiguration
AWS::IVS::StreamKey
AWS::ImageBuilder::Component
AWS::ImageBuilder::ContainerRecipe
AWS::ImageBuilder::DistributionConfiguration
AWS::ImageBuilder::Image
AWS::ImageBuilder::ImagePipeline
AWS::ImageBuilder::ImageRecipe
AWS::ImageBuilder::InfrastructureConfiguration
AWS::IoT::AccountAuditConfiguration
AWS::IoT::Authorizer
AWS::IoT::Certificate
AWS::IoT::CustomMetric
AWS::IoT::Dimension
AWS::IoT::DomainConfiguration
AWS::IoT::MitigationAction
AWS::IoT::ProvisioningTemplate
AWS::IoT::ScheduledAudit
AWS::IoT::SecurityProfile
AWS::IoT::TopicRule
AWS::IoT::TopicRuleDestination
AWS::IoTEvents::DetectorModel
AWS::IoTEvents::Input
AWS::IoTFleetHub::Application
AWS::IoTSiteWise::AccessPolicy
AWS::IoTSiteWise::Asset
AWS::IoTSiteWise::AssetModel
AWS::IoTSiteWise::Dashboard
AWS::IoTSiteWise::Gateway
AWS::IoTSiteWise::Portal
AWS::IoTSiteWise::Project
AWS::IoTWireless::Destination
AWS::IoTWireless::PartnerAccount
AWS::IoTWireless::WirelessDevice
AWS::IoTWireless::WirelessGateway
AWS::KMS::Alias
AWS::KMS::Key
AWS::Kendra::DataSource
AWS::Kendra::Faq
AWS::Kendra::Index
AWS::Kinesis::Stream
AWS::KinesisFirehose::DeliveryStream
AWS::Lambda::CodeSigningConfig
AWS::Lambda::EventSourceMapping
AWS::LicenseManager::Grant
AWS::LicenseManager::License
AWS::Logs::LogGroup
AWS::Logs::QueryDefinition
AWS::LookoutMetrics::AnomalyDetector
AWS::LookoutVision::Project
AWS::MWAA::Environment
AWS::Macie::CustomDataIdentifier
AWS::Macie::FindingsFilter
AWS::IoTWireless::DeviceProfile
AWS::IoTWireless::ServiceProfile
AWS::IoTWireless::TaskDefinition
AWS::LookoutMetrics::Alert
AWS::Macie::Session
AWS::MediaConnect::Flow
AWS::MediaConnect::FlowEntitlement
AWS::MediaConnect::FlowOutput
AWS::MediaConnect::FlowSource
AWS::MediaConnect::FlowVpcInterface
AWS::MediaPackage::Channel
AWS::MediaPackage::OriginEndpoint
AWS::MediaPackage::PackagingGroup
AWS::NetworkFirewall::Firewall
AWS::NetworkFirewall::FirewallPolicy
AWS::NetworkFirewall::LoggingConfiguration
AWS::NetworkFirewall::RuleGroup
AWS::NetworkManager::Device
AWS::NetworkManager::GlobalNetwork
AWS::NetworkManager::Link
AWS::NetworkManager::Site
AWS::NimbleStudio::LaunchProfile
AWS::NimbleStudio::StreamingImage
AWS::NimbleStudio::Studio
AWS::NimbleStudio::StudioComponent
AWS::OpsWorksCM::Server
AWS::QLDB::Stream
AWS::QuickSight::Analysis
AWS::QuickSight::Dashboard
AWS::QuickSight::DataSet
AWS::QuickSight::DataSource
AWS::QuickSight::Template
AWS::QuickSight::Theme
AWS::RDS::DBProxy
AWS::RDS::DBProxyEndpoint
AWS::RDS::DBProxyTargetGroup
AWS::RDS::GlobalCluster
AWS::ResourceGroups::Group
AWS::MediaPackage::Asset
AWS::MediaPackage::PackagingConfiguration
AWS::NetworkManager::CustomerGatewayAssociation
AWS::NetworkManager::LinkAssociation
AWS::NetworkManager::TransitGatewayRegistration
AWS::Route53::HealthCheck
AWS::Route53::HostedZone
AWS::Route53::KeySigningKey
AWS::Route53Resolver::FirewallDomainList
AWS::Route53Resolver::FirewallRuleGroup
AWS::Route53Resolver::FirewallRuleGroupAssociation
AWS::Route53Resolver::ResolverQueryLoggingConfig
AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation
AWS::S3::AccessPoint
AWS::S3::StorageLens
AWS::S3ObjectLambda::AccessPoint
AWS::S3ObjectLambda::AccessPointPolicy
AWS::S3Outposts::AccessPoint
AWS::S3Outposts::Bucket
AWS::S3Outposts::BucketPolicy
AWS::SES::ContactList
AWS::SSM::Association
AWS::SSM::Document
AWS::SSM::ResourceDataSync
AWS::SSO::Assignment
AWS::SSO::InstanceAccessControlAttributeConfiguration
AWS::SSO::PermissionSet
AWS::SageMaker::AppImageConfig
AWS::SageMaker::Device
AWS::SageMaker::DeviceFleet
AWS::SageMaker::Domain
AWS::SageMaker::Image
AWS::SageMaker::ModelPackageGroup
AWS::SageMaker::MonitoringSchedule
AWS::SageMaker::Pipeline
AWS::SageMaker::Project
AWS::SageMaker::UserProfile
AWS::ServiceCatalog::CloudFormationProvisionedProduct
AWS::ServiceCatalog::ServiceAction
AWS::ServiceCatalogAppRegistry::Application
AWS::ServiceCatalogAppRegistry::AttributeGroup
AWS::Route53::DNSSEC
AWS::Route53Resolver::ResolverDNSSECConfig
AWS::S3Outposts::Endpoint
AWS::SES::ConfigurationSet
AWS::SageMaker::App
AWS::SageMaker::DataQualityJobDefinition
AWS::SageMaker::FeatureGroup
AWS::SageMaker::ImageVersion
AWS::SageMaker::ModelBiasJobDefinition
AWS::SageMaker::ModelExplainabilityJobDefinition
AWS::SageMaker::ModelQualityJobDefinition
AWS::ServiceCatalog::ServiceActionAssociation
AWS::Signer::SigningProfile
AWS::StepFunctions::StateMachine
AWS::Synthetics::Canary
AWS::Timestream::Database
AWS::Timestream::Table
AWS::WAFv2::IPSet
AWS::WAFv2::RegexPatternSet
AWS::WAFv2::RuleGroup
AWS::WAFv2::WebACL
AWS::WAFv2::WebACLAssociation
AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation
AWS::ServiceCatalogAppRegistry::ResourceAssociation
AWS::Signer::ProfilePermission
AWS::WorkSpaces::ConnectionAlias
`
