// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ChannelArgs } from "./channel";
export type Channel = import("./channel").Channel;
export const Channel: typeof import("./channel").Channel = null as any;

export { DatasetArgs } from "./dataset";
export type Dataset = import("./dataset").Dataset;
export const Dataset: typeof import("./dataset").Dataset = null as any;

export { DatastoreArgs } from "./datastore";
export type Datastore = import("./datastore").Datastore;
export const Datastore: typeof import("./datastore").Datastore = null as any;

export { GetChannelArgs, GetChannelResult, GetChannelOutputArgs } from "./getChannel";
export const getChannel: typeof import("./getChannel").getChannel = null as any;
export const getChannelOutput: typeof import("./getChannel").getChannelOutput = null as any;

export { GetDatasetArgs, GetDatasetResult, GetDatasetOutputArgs } from "./getDataset";
export const getDataset: typeof import("./getDataset").getDataset = null as any;
export const getDatasetOutput: typeof import("./getDataset").getDatasetOutput = null as any;

export { GetDatastoreArgs, GetDatastoreResult, GetDatastoreOutputArgs } from "./getDatastore";
export const getDatastore: typeof import("./getDatastore").getDatastore = null as any;
export const getDatastoreOutput: typeof import("./getDatastore").getDatastoreOutput = null as any;

export { GetPipelineArgs, GetPipelineResult, GetPipelineOutputArgs } from "./getPipeline";
export const getPipeline: typeof import("./getPipeline").getPipeline = null as any;
export const getPipelineOutput: typeof import("./getPipeline").getPipelineOutput = null as any;

export { PipelineArgs } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;

utilities.lazyLoad(exports, ["Channel"], () => require("./channel"));
utilities.lazyLoad(exports, ["Dataset"], () => require("./dataset"));
utilities.lazyLoad(exports, ["Datastore"], () => require("./datastore"));
utilities.lazyLoad(exports, ["getChannel","getChannelOutput"], () => require("./getChannel"));
utilities.lazyLoad(exports, ["getDataset","getDatasetOutput"], () => require("./getDataset"));
utilities.lazyLoad(exports, ["getDatastore","getDatastoreOutput"], () => require("./getDatastore"));
utilities.lazyLoad(exports, ["getPipeline","getPipelineOutput"], () => require("./getPipeline"));
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));

// Export enums:
export * from "../types/enums/iotanalytics";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iotanalytics:Channel":
                return new Channel(name, <any>undefined, { urn })
            case "aws-native:iotanalytics:Dataset":
                return new Dataset(name, <any>undefined, { urn })
            case "aws-native:iotanalytics:Datastore":
                return new Datastore(name, <any>undefined, { urn })
            case "aws-native:iotanalytics:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iotanalytics", _module)