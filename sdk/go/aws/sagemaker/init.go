// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:sagemaker:App":
		r = &App{}
	case "aws-native:sagemaker:AppImageConfig":
		r = &AppImageConfig{}
	case "aws-native:sagemaker:Cluster":
		r = &Cluster{}
	case "aws-native:sagemaker:DataQualityJobDefinition":
		r = &DataQualityJobDefinition{}
	case "aws-native:sagemaker:Device":
		r = &Device{}
	case "aws-native:sagemaker:DeviceFleet":
		r = &DeviceFleet{}
	case "aws-native:sagemaker:Domain":
		r = &Domain{}
	case "aws-native:sagemaker:Endpoint":
		r = &Endpoint{}
	case "aws-native:sagemaker:FeatureGroup":
		r = &FeatureGroup{}
	case "aws-native:sagemaker:Image":
		r = &Image{}
	case "aws-native:sagemaker:ImageVersion":
		r = &ImageVersion{}
	case "aws-native:sagemaker:InferenceComponent":
		r = &InferenceComponent{}
	case "aws-native:sagemaker:InferenceExperiment":
		r = &InferenceExperiment{}
	case "aws-native:sagemaker:MlflowTrackingServer":
		r = &MlflowTrackingServer{}
	case "aws-native:sagemaker:ModelBiasJobDefinition":
		r = &ModelBiasJobDefinition{}
	case "aws-native:sagemaker:ModelCard":
		r = &ModelCard{}
	case "aws-native:sagemaker:ModelExplainabilityJobDefinition":
		r = &ModelExplainabilityJobDefinition{}
	case "aws-native:sagemaker:ModelPackage":
		r = &ModelPackage{}
	case "aws-native:sagemaker:ModelPackageGroup":
		r = &ModelPackageGroup{}
	case "aws-native:sagemaker:ModelQualityJobDefinition":
		r = &ModelQualityJobDefinition{}
	case "aws-native:sagemaker:MonitoringSchedule":
		r = &MonitoringSchedule{}
	case "aws-native:sagemaker:PartnerApp":
		r = &PartnerApp{}
	case "aws-native:sagemaker:Pipeline":
		r = &Pipeline{}
	case "aws-native:sagemaker:Project":
		r = &Project{}
	case "aws-native:sagemaker:Space":
		r = &Space{}
	case "aws-native:sagemaker:StudioLifecycleConfig":
		r = &StudioLifecycleConfig{}
	case "aws-native:sagemaker:UserProfile":
		r = &UserProfile{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"sagemaker",
		&module{version},
	)
}
