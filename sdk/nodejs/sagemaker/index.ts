// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AppArgs } from "./app";
export type App = import("./app").App;
export const App: typeof import("./app").App = null as any;
utilities.lazyLoad(exports, ["App"], () => require("./app"));

export { AppImageConfigArgs } from "./appImageConfig";
export type AppImageConfig = import("./appImageConfig").AppImageConfig;
export const AppImageConfig: typeof import("./appImageConfig").AppImageConfig = null as any;
utilities.lazyLoad(exports, ["AppImageConfig"], () => require("./appImageConfig"));

export { CodeRepositoryArgs } from "./codeRepository";
export type CodeRepository = import("./codeRepository").CodeRepository;
export const CodeRepository: typeof import("./codeRepository").CodeRepository = null as any;
utilities.lazyLoad(exports, ["CodeRepository"], () => require("./codeRepository"));

export { DataQualityJobDefinitionArgs } from "./dataQualityJobDefinition";
export type DataQualityJobDefinition = import("./dataQualityJobDefinition").DataQualityJobDefinition;
export const DataQualityJobDefinition: typeof import("./dataQualityJobDefinition").DataQualityJobDefinition = null as any;
utilities.lazyLoad(exports, ["DataQualityJobDefinition"], () => require("./dataQualityJobDefinition"));

export { DeviceArgs } from "./device";
export type Device = import("./device").Device;
export const Device: typeof import("./device").Device = null as any;
utilities.lazyLoad(exports, ["Device"], () => require("./device"));

export { DeviceFleetArgs } from "./deviceFleet";
export type DeviceFleet = import("./deviceFleet").DeviceFleet;
export const DeviceFleet: typeof import("./deviceFleet").DeviceFleet = null as any;
utilities.lazyLoad(exports, ["DeviceFleet"], () => require("./deviceFleet"));

export { DomainArgs } from "./domain";
export type Domain = import("./domain").Domain;
export const Domain: typeof import("./domain").Domain = null as any;
utilities.lazyLoad(exports, ["Domain"], () => require("./domain"));

export { EndpointArgs } from "./endpoint";
export type Endpoint = import("./endpoint").Endpoint;
export const Endpoint: typeof import("./endpoint").Endpoint = null as any;
utilities.lazyLoad(exports, ["Endpoint"], () => require("./endpoint"));

export { EndpointConfigArgs } from "./endpointConfig";
export type EndpointConfig = import("./endpointConfig").EndpointConfig;
export const EndpointConfig: typeof import("./endpointConfig").EndpointConfig = null as any;
utilities.lazyLoad(exports, ["EndpointConfig"], () => require("./endpointConfig"));

export { FeatureGroupArgs } from "./featureGroup";
export type FeatureGroup = import("./featureGroup").FeatureGroup;
export const FeatureGroup: typeof import("./featureGroup").FeatureGroup = null as any;
utilities.lazyLoad(exports, ["FeatureGroup"], () => require("./featureGroup"));

export { GetAppArgs, GetAppResult, GetAppOutputArgs } from "./getApp";
export const getApp: typeof import("./getApp").getApp = null as any;
export const getAppOutput: typeof import("./getApp").getAppOutput = null as any;
utilities.lazyLoad(exports, ["getApp","getAppOutput"], () => require("./getApp"));

export { GetAppImageConfigArgs, GetAppImageConfigResult, GetAppImageConfigOutputArgs } from "./getAppImageConfig";
export const getAppImageConfig: typeof import("./getAppImageConfig").getAppImageConfig = null as any;
export const getAppImageConfigOutput: typeof import("./getAppImageConfig").getAppImageConfigOutput = null as any;
utilities.lazyLoad(exports, ["getAppImageConfig","getAppImageConfigOutput"], () => require("./getAppImageConfig"));

export { GetCodeRepositoryArgs, GetCodeRepositoryResult, GetCodeRepositoryOutputArgs } from "./getCodeRepository";
export const getCodeRepository: typeof import("./getCodeRepository").getCodeRepository = null as any;
export const getCodeRepositoryOutput: typeof import("./getCodeRepository").getCodeRepositoryOutput = null as any;
utilities.lazyLoad(exports, ["getCodeRepository","getCodeRepositoryOutput"], () => require("./getCodeRepository"));

export { GetDataQualityJobDefinitionArgs, GetDataQualityJobDefinitionResult, GetDataQualityJobDefinitionOutputArgs } from "./getDataQualityJobDefinition";
export const getDataQualityJobDefinition: typeof import("./getDataQualityJobDefinition").getDataQualityJobDefinition = null as any;
export const getDataQualityJobDefinitionOutput: typeof import("./getDataQualityJobDefinition").getDataQualityJobDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getDataQualityJobDefinition","getDataQualityJobDefinitionOutput"], () => require("./getDataQualityJobDefinition"));

export { GetDeviceFleetArgs, GetDeviceFleetResult, GetDeviceFleetOutputArgs } from "./getDeviceFleet";
export const getDeviceFleet: typeof import("./getDeviceFleet").getDeviceFleet = null as any;
export const getDeviceFleetOutput: typeof import("./getDeviceFleet").getDeviceFleetOutput = null as any;
utilities.lazyLoad(exports, ["getDeviceFleet","getDeviceFleetOutput"], () => require("./getDeviceFleet"));

export { GetDomainArgs, GetDomainResult, GetDomainOutputArgs } from "./getDomain";
export const getDomain: typeof import("./getDomain").getDomain = null as any;
export const getDomainOutput: typeof import("./getDomain").getDomainOutput = null as any;
utilities.lazyLoad(exports, ["getDomain","getDomainOutput"], () => require("./getDomain"));

export { GetEndpointArgs, GetEndpointResult, GetEndpointOutputArgs } from "./getEndpoint";
export const getEndpoint: typeof import("./getEndpoint").getEndpoint = null as any;
export const getEndpointOutput: typeof import("./getEndpoint").getEndpointOutput = null as any;
utilities.lazyLoad(exports, ["getEndpoint","getEndpointOutput"], () => require("./getEndpoint"));

export { GetEndpointConfigArgs, GetEndpointConfigResult, GetEndpointConfigOutputArgs } from "./getEndpointConfig";
export const getEndpointConfig: typeof import("./getEndpointConfig").getEndpointConfig = null as any;
export const getEndpointConfigOutput: typeof import("./getEndpointConfig").getEndpointConfigOutput = null as any;
utilities.lazyLoad(exports, ["getEndpointConfig","getEndpointConfigOutput"], () => require("./getEndpointConfig"));

export { GetFeatureGroupArgs, GetFeatureGroupResult, GetFeatureGroupOutputArgs } from "./getFeatureGroup";
export const getFeatureGroup: typeof import("./getFeatureGroup").getFeatureGroup = null as any;
export const getFeatureGroupOutput: typeof import("./getFeatureGroup").getFeatureGroupOutput = null as any;
utilities.lazyLoad(exports, ["getFeatureGroup","getFeatureGroupOutput"], () => require("./getFeatureGroup"));

export { GetImageArgs, GetImageResult, GetImageOutputArgs } from "./getImage";
export const getImage: typeof import("./getImage").getImage = null as any;
export const getImageOutput: typeof import("./getImage").getImageOutput = null as any;
utilities.lazyLoad(exports, ["getImage","getImageOutput"], () => require("./getImage"));

export { GetImageVersionArgs, GetImageVersionResult, GetImageVersionOutputArgs } from "./getImageVersion";
export const getImageVersion: typeof import("./getImageVersion").getImageVersion = null as any;
export const getImageVersionOutput: typeof import("./getImageVersion").getImageVersionOutput = null as any;
utilities.lazyLoad(exports, ["getImageVersion","getImageVersionOutput"], () => require("./getImageVersion"));

export { GetModelArgs, GetModelResult, GetModelOutputArgs } from "./getModel";
export const getModel: typeof import("./getModel").getModel = null as any;
export const getModelOutput: typeof import("./getModel").getModelOutput = null as any;
utilities.lazyLoad(exports, ["getModel","getModelOutput"], () => require("./getModel"));

export { GetModelBiasJobDefinitionArgs, GetModelBiasJobDefinitionResult, GetModelBiasJobDefinitionOutputArgs } from "./getModelBiasJobDefinition";
export const getModelBiasJobDefinition: typeof import("./getModelBiasJobDefinition").getModelBiasJobDefinition = null as any;
export const getModelBiasJobDefinitionOutput: typeof import("./getModelBiasJobDefinition").getModelBiasJobDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getModelBiasJobDefinition","getModelBiasJobDefinitionOutput"], () => require("./getModelBiasJobDefinition"));

export { GetModelExplainabilityJobDefinitionArgs, GetModelExplainabilityJobDefinitionResult, GetModelExplainabilityJobDefinitionOutputArgs } from "./getModelExplainabilityJobDefinition";
export const getModelExplainabilityJobDefinition: typeof import("./getModelExplainabilityJobDefinition").getModelExplainabilityJobDefinition = null as any;
export const getModelExplainabilityJobDefinitionOutput: typeof import("./getModelExplainabilityJobDefinition").getModelExplainabilityJobDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getModelExplainabilityJobDefinition","getModelExplainabilityJobDefinitionOutput"], () => require("./getModelExplainabilityJobDefinition"));

export { GetModelPackageArgs, GetModelPackageResult, GetModelPackageOutputArgs } from "./getModelPackage";
export const getModelPackage: typeof import("./getModelPackage").getModelPackage = null as any;
export const getModelPackageOutput: typeof import("./getModelPackage").getModelPackageOutput = null as any;
utilities.lazyLoad(exports, ["getModelPackage","getModelPackageOutput"], () => require("./getModelPackage"));

export { GetModelPackageGroupArgs, GetModelPackageGroupResult, GetModelPackageGroupOutputArgs } from "./getModelPackageGroup";
export const getModelPackageGroup: typeof import("./getModelPackageGroup").getModelPackageGroup = null as any;
export const getModelPackageGroupOutput: typeof import("./getModelPackageGroup").getModelPackageGroupOutput = null as any;
utilities.lazyLoad(exports, ["getModelPackageGroup","getModelPackageGroupOutput"], () => require("./getModelPackageGroup"));

export { GetModelQualityJobDefinitionArgs, GetModelQualityJobDefinitionResult, GetModelQualityJobDefinitionOutputArgs } from "./getModelQualityJobDefinition";
export const getModelQualityJobDefinition: typeof import("./getModelQualityJobDefinition").getModelQualityJobDefinition = null as any;
export const getModelQualityJobDefinitionOutput: typeof import("./getModelQualityJobDefinition").getModelQualityJobDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getModelQualityJobDefinition","getModelQualityJobDefinitionOutput"], () => require("./getModelQualityJobDefinition"));

export { GetMonitoringScheduleArgs, GetMonitoringScheduleResult, GetMonitoringScheduleOutputArgs } from "./getMonitoringSchedule";
export const getMonitoringSchedule: typeof import("./getMonitoringSchedule").getMonitoringSchedule = null as any;
export const getMonitoringScheduleOutput: typeof import("./getMonitoringSchedule").getMonitoringScheduleOutput = null as any;
utilities.lazyLoad(exports, ["getMonitoringSchedule","getMonitoringScheduleOutput"], () => require("./getMonitoringSchedule"));

export { GetNotebookInstanceArgs, GetNotebookInstanceResult, GetNotebookInstanceOutputArgs } from "./getNotebookInstance";
export const getNotebookInstance: typeof import("./getNotebookInstance").getNotebookInstance = null as any;
export const getNotebookInstanceOutput: typeof import("./getNotebookInstance").getNotebookInstanceOutput = null as any;
utilities.lazyLoad(exports, ["getNotebookInstance","getNotebookInstanceOutput"], () => require("./getNotebookInstance"));

export { GetNotebookInstanceLifecycleConfigArgs, GetNotebookInstanceLifecycleConfigResult, GetNotebookInstanceLifecycleConfigOutputArgs } from "./getNotebookInstanceLifecycleConfig";
export const getNotebookInstanceLifecycleConfig: typeof import("./getNotebookInstanceLifecycleConfig").getNotebookInstanceLifecycleConfig = null as any;
export const getNotebookInstanceLifecycleConfigOutput: typeof import("./getNotebookInstanceLifecycleConfig").getNotebookInstanceLifecycleConfigOutput = null as any;
utilities.lazyLoad(exports, ["getNotebookInstanceLifecycleConfig","getNotebookInstanceLifecycleConfigOutput"], () => require("./getNotebookInstanceLifecycleConfig"));

export { GetPipelineArgs, GetPipelineResult, GetPipelineOutputArgs } from "./getPipeline";
export const getPipeline: typeof import("./getPipeline").getPipeline = null as any;
export const getPipelineOutput: typeof import("./getPipeline").getPipelineOutput = null as any;
utilities.lazyLoad(exports, ["getPipeline","getPipelineOutput"], () => require("./getPipeline"));

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));

export { GetUserProfileArgs, GetUserProfileResult, GetUserProfileOutputArgs } from "./getUserProfile";
export const getUserProfile: typeof import("./getUserProfile").getUserProfile = null as any;
export const getUserProfileOutput: typeof import("./getUserProfile").getUserProfileOutput = null as any;
utilities.lazyLoad(exports, ["getUserProfile","getUserProfileOutput"], () => require("./getUserProfile"));

export { GetWorkteamArgs, GetWorkteamResult, GetWorkteamOutputArgs } from "./getWorkteam";
export const getWorkteam: typeof import("./getWorkteam").getWorkteam = null as any;
export const getWorkteamOutput: typeof import("./getWorkteam").getWorkteamOutput = null as any;
utilities.lazyLoad(exports, ["getWorkteam","getWorkteamOutput"], () => require("./getWorkteam"));

export { ImageArgs } from "./image";
export type Image = import("./image").Image;
export const Image: typeof import("./image").Image = null as any;
utilities.lazyLoad(exports, ["Image"], () => require("./image"));

export { ImageVersionArgs } from "./imageVersion";
export type ImageVersion = import("./imageVersion").ImageVersion;
export const ImageVersion: typeof import("./imageVersion").ImageVersion = null as any;
utilities.lazyLoad(exports, ["ImageVersion"], () => require("./imageVersion"));

export { ModelArgs } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));

export { ModelBiasJobDefinitionArgs } from "./modelBiasJobDefinition";
export type ModelBiasJobDefinition = import("./modelBiasJobDefinition").ModelBiasJobDefinition;
export const ModelBiasJobDefinition: typeof import("./modelBiasJobDefinition").ModelBiasJobDefinition = null as any;
utilities.lazyLoad(exports, ["ModelBiasJobDefinition"], () => require("./modelBiasJobDefinition"));

export { ModelExplainabilityJobDefinitionArgs } from "./modelExplainabilityJobDefinition";
export type ModelExplainabilityJobDefinition = import("./modelExplainabilityJobDefinition").ModelExplainabilityJobDefinition;
export const ModelExplainabilityJobDefinition: typeof import("./modelExplainabilityJobDefinition").ModelExplainabilityJobDefinition = null as any;
utilities.lazyLoad(exports, ["ModelExplainabilityJobDefinition"], () => require("./modelExplainabilityJobDefinition"));

export { ModelPackageArgs } from "./modelPackage";
export type ModelPackage = import("./modelPackage").ModelPackage;
export const ModelPackage: typeof import("./modelPackage").ModelPackage = null as any;
utilities.lazyLoad(exports, ["ModelPackage"], () => require("./modelPackage"));

export { ModelPackageGroupArgs } from "./modelPackageGroup";
export type ModelPackageGroup = import("./modelPackageGroup").ModelPackageGroup;
export const ModelPackageGroup: typeof import("./modelPackageGroup").ModelPackageGroup = null as any;
utilities.lazyLoad(exports, ["ModelPackageGroup"], () => require("./modelPackageGroup"));

export { ModelQualityJobDefinitionArgs } from "./modelQualityJobDefinition";
export type ModelQualityJobDefinition = import("./modelQualityJobDefinition").ModelQualityJobDefinition;
export const ModelQualityJobDefinition: typeof import("./modelQualityJobDefinition").ModelQualityJobDefinition = null as any;
utilities.lazyLoad(exports, ["ModelQualityJobDefinition"], () => require("./modelQualityJobDefinition"));

export { MonitoringScheduleArgs } from "./monitoringSchedule";
export type MonitoringSchedule = import("./monitoringSchedule").MonitoringSchedule;
export const MonitoringSchedule: typeof import("./monitoringSchedule").MonitoringSchedule = null as any;
utilities.lazyLoad(exports, ["MonitoringSchedule"], () => require("./monitoringSchedule"));

export { NotebookInstanceArgs } from "./notebookInstance";
export type NotebookInstance = import("./notebookInstance").NotebookInstance;
export const NotebookInstance: typeof import("./notebookInstance").NotebookInstance = null as any;
utilities.lazyLoad(exports, ["NotebookInstance"], () => require("./notebookInstance"));

export { NotebookInstanceLifecycleConfigArgs } from "./notebookInstanceLifecycleConfig";
export type NotebookInstanceLifecycleConfig = import("./notebookInstanceLifecycleConfig").NotebookInstanceLifecycleConfig;
export const NotebookInstanceLifecycleConfig: typeof import("./notebookInstanceLifecycleConfig").NotebookInstanceLifecycleConfig = null as any;
utilities.lazyLoad(exports, ["NotebookInstanceLifecycleConfig"], () => require("./notebookInstanceLifecycleConfig"));

export { PipelineArgs } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { UserProfileArgs } from "./userProfile";
export type UserProfile = import("./userProfile").UserProfile;
export const UserProfile: typeof import("./userProfile").UserProfile = null as any;
utilities.lazyLoad(exports, ["UserProfile"], () => require("./userProfile"));

export { WorkteamArgs } from "./workteam";
export type Workteam = import("./workteam").Workteam;
export const Workteam: typeof import("./workteam").Workteam = null as any;
utilities.lazyLoad(exports, ["Workteam"], () => require("./workteam"));


// Export enums:
export * from "../types/enums/sagemaker";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:sagemaker:App":
                return new App(name, <any>undefined, { urn })
            case "aws-native:sagemaker:AppImageConfig":
                return new AppImageConfig(name, <any>undefined, { urn })
            case "aws-native:sagemaker:CodeRepository":
                return new CodeRepository(name, <any>undefined, { urn })
            case "aws-native:sagemaker:DataQualityJobDefinition":
                return new DataQualityJobDefinition(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Device":
                return new Device(name, <any>undefined, { urn })
            case "aws-native:sagemaker:DeviceFleet":
                return new DeviceFleet(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "aws-native:sagemaker:EndpointConfig":
                return new EndpointConfig(name, <any>undefined, { urn })
            case "aws-native:sagemaker:FeatureGroup":
                return new FeatureGroup(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Image":
                return new Image(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ImageVersion":
                return new ImageVersion(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Model":
                return new Model(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ModelBiasJobDefinition":
                return new ModelBiasJobDefinition(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ModelExplainabilityJobDefinition":
                return new ModelExplainabilityJobDefinition(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ModelPackage":
                return new ModelPackage(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ModelPackageGroup":
                return new ModelPackageGroup(name, <any>undefined, { urn })
            case "aws-native:sagemaker:ModelQualityJobDefinition":
                return new ModelQualityJobDefinition(name, <any>undefined, { urn })
            case "aws-native:sagemaker:MonitoringSchedule":
                return new MonitoringSchedule(name, <any>undefined, { urn })
            case "aws-native:sagemaker:NotebookInstance":
                return new NotebookInstance(name, <any>undefined, { urn })
            case "aws-native:sagemaker:NotebookInstanceLifecycleConfig":
                return new NotebookInstanceLifecycleConfig(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Project":
                return new Project(name, <any>undefined, { urn })
            case "aws-native:sagemaker:UserProfile":
                return new UserProfile(name, <any>undefined, { urn })
            case "aws-native:sagemaker:Workteam":
                return new Workteam(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "sagemaker", _module)