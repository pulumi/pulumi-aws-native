// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClassifierArgs } from "./classifier";
export type Classifier = import("./classifier").Classifier;
export const Classifier: typeof import("./classifier").Classifier = null as any;
utilities.lazyLoad(exports, ["Classifier"], () => require("./classifier"));

export { ConnectionArgs } from "./connection";
export type Connection = import("./connection").Connection;
export const Connection: typeof import("./connection").Connection = null as any;
utilities.lazyLoad(exports, ["Connection"], () => require("./connection"));

export { CrawlerArgs } from "./crawler";
export type Crawler = import("./crawler").Crawler;
export const Crawler: typeof import("./crawler").Crawler = null as any;
utilities.lazyLoad(exports, ["Crawler"], () => require("./crawler"));

export { DataCatalogEncryptionSettingsArgs } from "./dataCatalogEncryptionSettings";
export type DataCatalogEncryptionSettings = import("./dataCatalogEncryptionSettings").DataCatalogEncryptionSettings;
export const DataCatalogEncryptionSettings: typeof import("./dataCatalogEncryptionSettings").DataCatalogEncryptionSettings = null as any;
utilities.lazyLoad(exports, ["DataCatalogEncryptionSettings"], () => require("./dataCatalogEncryptionSettings"));

export { DatabaseArgs } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DevEndpointArgs } from "./devEndpoint";
export type DevEndpoint = import("./devEndpoint").DevEndpoint;
export const DevEndpoint: typeof import("./devEndpoint").DevEndpoint = null as any;
utilities.lazyLoad(exports, ["DevEndpoint"], () => require("./devEndpoint"));

export { GetClassifierArgs, GetClassifierResult, GetClassifierOutputArgs } from "./getClassifier";
export const getClassifier: typeof import("./getClassifier").getClassifier = null as any;
export const getClassifierOutput: typeof import("./getClassifier").getClassifierOutput = null as any;
utilities.lazyLoad(exports, ["getClassifier","getClassifierOutput"], () => require("./getClassifier"));

export { GetConnectionArgs, GetConnectionResult, GetConnectionOutputArgs } from "./getConnection";
export const getConnection: typeof import("./getConnection").getConnection = null as any;
export const getConnectionOutput: typeof import("./getConnection").getConnectionOutput = null as any;
utilities.lazyLoad(exports, ["getConnection","getConnectionOutput"], () => require("./getConnection"));

export { GetCrawlerArgs, GetCrawlerResult, GetCrawlerOutputArgs } from "./getCrawler";
export const getCrawler: typeof import("./getCrawler").getCrawler = null as any;
export const getCrawlerOutput: typeof import("./getCrawler").getCrawlerOutput = null as any;
utilities.lazyLoad(exports, ["getCrawler","getCrawlerOutput"], () => require("./getCrawler"));

export { GetDataCatalogEncryptionSettingsArgs, GetDataCatalogEncryptionSettingsResult, GetDataCatalogEncryptionSettingsOutputArgs } from "./getDataCatalogEncryptionSettings";
export const getDataCatalogEncryptionSettings: typeof import("./getDataCatalogEncryptionSettings").getDataCatalogEncryptionSettings = null as any;
export const getDataCatalogEncryptionSettingsOutput: typeof import("./getDataCatalogEncryptionSettings").getDataCatalogEncryptionSettingsOutput = null as any;
utilities.lazyLoad(exports, ["getDataCatalogEncryptionSettings","getDataCatalogEncryptionSettingsOutput"], () => require("./getDataCatalogEncryptionSettings"));

export { GetDatabaseArgs, GetDatabaseResult, GetDatabaseOutputArgs } from "./getDatabase";
export const getDatabase: typeof import("./getDatabase").getDatabase = null as any;
export const getDatabaseOutput: typeof import("./getDatabase").getDatabaseOutput = null as any;
utilities.lazyLoad(exports, ["getDatabase","getDatabaseOutput"], () => require("./getDatabase"));

export { GetDevEndpointArgs, GetDevEndpointResult, GetDevEndpointOutputArgs } from "./getDevEndpoint";
export const getDevEndpoint: typeof import("./getDevEndpoint").getDevEndpoint = null as any;
export const getDevEndpointOutput: typeof import("./getDevEndpoint").getDevEndpointOutput = null as any;
utilities.lazyLoad(exports, ["getDevEndpoint","getDevEndpointOutput"], () => require("./getDevEndpoint"));

export { GetJobArgs, GetJobResult, GetJobOutputArgs } from "./getJob";
export const getJob: typeof import("./getJob").getJob = null as any;
export const getJobOutput: typeof import("./getJob").getJobOutput = null as any;
utilities.lazyLoad(exports, ["getJob","getJobOutput"], () => require("./getJob"));

export { GetMLTransformArgs, GetMLTransformResult, GetMLTransformOutputArgs } from "./getMLTransform";
export const getMLTransform: typeof import("./getMLTransform").getMLTransform = null as any;
export const getMLTransformOutput: typeof import("./getMLTransform").getMLTransformOutput = null as any;
utilities.lazyLoad(exports, ["getMLTransform","getMLTransformOutput"], () => require("./getMLTransform"));

export { GetPartitionArgs, GetPartitionResult, GetPartitionOutputArgs } from "./getPartition";
export const getPartition: typeof import("./getPartition").getPartition = null as any;
export const getPartitionOutput: typeof import("./getPartition").getPartitionOutput = null as any;
utilities.lazyLoad(exports, ["getPartition","getPartitionOutput"], () => require("./getPartition"));

export { GetRegistryArgs, GetRegistryResult, GetRegistryOutputArgs } from "./getRegistry";
export const getRegistry: typeof import("./getRegistry").getRegistry = null as any;
export const getRegistryOutput: typeof import("./getRegistry").getRegistryOutput = null as any;
utilities.lazyLoad(exports, ["getRegistry","getRegistryOutput"], () => require("./getRegistry"));

export { GetSchemaArgs, GetSchemaResult, GetSchemaOutputArgs } from "./getSchema";
export const getSchema: typeof import("./getSchema").getSchema = null as any;
export const getSchemaOutput: typeof import("./getSchema").getSchemaOutput = null as any;
utilities.lazyLoad(exports, ["getSchema","getSchemaOutput"], () => require("./getSchema"));

export { GetSchemaVersionArgs, GetSchemaVersionResult, GetSchemaVersionOutputArgs } from "./getSchemaVersion";
export const getSchemaVersion: typeof import("./getSchemaVersion").getSchemaVersion = null as any;
export const getSchemaVersionOutput: typeof import("./getSchemaVersion").getSchemaVersionOutput = null as any;
utilities.lazyLoad(exports, ["getSchemaVersion","getSchemaVersionOutput"], () => require("./getSchemaVersion"));

export { GetSecurityConfigurationArgs, GetSecurityConfigurationResult, GetSecurityConfigurationOutputArgs } from "./getSecurityConfiguration";
export const getSecurityConfiguration: typeof import("./getSecurityConfiguration").getSecurityConfiguration = null as any;
export const getSecurityConfigurationOutput: typeof import("./getSecurityConfiguration").getSecurityConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityConfiguration","getSecurityConfigurationOutput"], () => require("./getSecurityConfiguration"));

export { GetTableArgs, GetTableResult, GetTableOutputArgs } from "./getTable";
export const getTable: typeof import("./getTable").getTable = null as any;
export const getTableOutput: typeof import("./getTable").getTableOutput = null as any;
utilities.lazyLoad(exports, ["getTable","getTableOutput"], () => require("./getTable"));

export { GetTriggerArgs, GetTriggerResult, GetTriggerOutputArgs } from "./getTrigger";
export const getTrigger: typeof import("./getTrigger").getTrigger = null as any;
export const getTriggerOutput: typeof import("./getTrigger").getTriggerOutput = null as any;
utilities.lazyLoad(exports, ["getTrigger","getTriggerOutput"], () => require("./getTrigger"));

export { GetWorkflowArgs, GetWorkflowResult, GetWorkflowOutputArgs } from "./getWorkflow";
export const getWorkflow: typeof import("./getWorkflow").getWorkflow = null as any;
export const getWorkflowOutput: typeof import("./getWorkflow").getWorkflowOutput = null as any;
utilities.lazyLoad(exports, ["getWorkflow","getWorkflowOutput"], () => require("./getWorkflow"));

export { JobArgs } from "./job";
export type Job = import("./job").Job;
export const Job: typeof import("./job").Job = null as any;
utilities.lazyLoad(exports, ["Job"], () => require("./job"));

export { MLTransformArgs } from "./mltransform";
export type MLTransform = import("./mltransform").MLTransform;
export const MLTransform: typeof import("./mltransform").MLTransform = null as any;
utilities.lazyLoad(exports, ["MLTransform"], () => require("./mltransform"));

export { PartitionArgs } from "./partition";
export type Partition = import("./partition").Partition;
export const Partition: typeof import("./partition").Partition = null as any;
utilities.lazyLoad(exports, ["Partition"], () => require("./partition"));

export { RegistryArgs } from "./registry";
export type Registry = import("./registry").Registry;
export const Registry: typeof import("./registry").Registry = null as any;
utilities.lazyLoad(exports, ["Registry"], () => require("./registry"));

export { SchemaArgs } from "./schema";
export type Schema = import("./schema").Schema;
export const Schema: typeof import("./schema").Schema = null as any;
utilities.lazyLoad(exports, ["Schema"], () => require("./schema"));

export { SchemaVersionArgs } from "./schemaVersion";
export type SchemaVersion = import("./schemaVersion").SchemaVersion;
export const SchemaVersion: typeof import("./schemaVersion").SchemaVersion = null as any;
utilities.lazyLoad(exports, ["SchemaVersion"], () => require("./schemaVersion"));

export { SchemaVersionMetadataArgs } from "./schemaVersionMetadata";
export type SchemaVersionMetadata = import("./schemaVersionMetadata").SchemaVersionMetadata;
export const SchemaVersionMetadata: typeof import("./schemaVersionMetadata").SchemaVersionMetadata = null as any;
utilities.lazyLoad(exports, ["SchemaVersionMetadata"], () => require("./schemaVersionMetadata"));

export { SecurityConfigurationArgs } from "./securityConfiguration";
export type SecurityConfiguration = import("./securityConfiguration").SecurityConfiguration;
export const SecurityConfiguration: typeof import("./securityConfiguration").SecurityConfiguration = null as any;
utilities.lazyLoad(exports, ["SecurityConfiguration"], () => require("./securityConfiguration"));

export { TableArgs } from "./table";
export type Table = import("./table").Table;
export const Table: typeof import("./table").Table = null as any;
utilities.lazyLoad(exports, ["Table"], () => require("./table"));

export { TriggerArgs } from "./trigger";
export type Trigger = import("./trigger").Trigger;
export const Trigger: typeof import("./trigger").Trigger = null as any;
utilities.lazyLoad(exports, ["Trigger"], () => require("./trigger"));

export { WorkflowArgs } from "./workflow";
export type Workflow = import("./workflow").Workflow;
export const Workflow: typeof import("./workflow").Workflow = null as any;
utilities.lazyLoad(exports, ["Workflow"], () => require("./workflow"));


// Export enums:
export * from "../types/enums/glue";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:glue:Classifier":
                return new Classifier(name, <any>undefined, { urn })
            case "aws-native:glue:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "aws-native:glue:Crawler":
                return new Crawler(name, <any>undefined, { urn })
            case "aws-native:glue:DataCatalogEncryptionSettings":
                return new DataCatalogEncryptionSettings(name, <any>undefined, { urn })
            case "aws-native:glue:Database":
                return new Database(name, <any>undefined, { urn })
            case "aws-native:glue:DevEndpoint":
                return new DevEndpoint(name, <any>undefined, { urn })
            case "aws-native:glue:Job":
                return new Job(name, <any>undefined, { urn })
            case "aws-native:glue:MLTransform":
                return new MLTransform(name, <any>undefined, { urn })
            case "aws-native:glue:Partition":
                return new Partition(name, <any>undefined, { urn })
            case "aws-native:glue:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "aws-native:glue:Schema":
                return new Schema(name, <any>undefined, { urn })
            case "aws-native:glue:SchemaVersion":
                return new SchemaVersion(name, <any>undefined, { urn })
            case "aws-native:glue:SchemaVersionMetadata":
                return new SchemaVersionMetadata(name, <any>undefined, { urn })
            case "aws-native:glue:SecurityConfiguration":
                return new SecurityConfiguration(name, <any>undefined, { urn })
            case "aws-native:glue:Table":
                return new Table(name, <any>undefined, { urn })
            case "aws-native:glue:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            case "aws-native:glue:Workflow":
                return new Workflow(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "glue", _module)