// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CampaignArgs } from "./campaign";
export type Campaign = import("./campaign").Campaign;
export const Campaign: typeof import("./campaign").Campaign = null as any;

export { DecoderManifestArgs } from "./decoderManifest";
export type DecoderManifest = import("./decoderManifest").DecoderManifest;
export const DecoderManifest: typeof import("./decoderManifest").DecoderManifest = null as any;

export { FleetArgs } from "./fleet";
export type Fleet = import("./fleet").Fleet;
export const Fleet: typeof import("./fleet").Fleet = null as any;

export { GetCampaignArgs, GetCampaignResult, GetCampaignOutputArgs } from "./getCampaign";
export const getCampaign: typeof import("./getCampaign").getCampaign = null as any;
export const getCampaignOutput: typeof import("./getCampaign").getCampaignOutput = null as any;

export { GetDecoderManifestArgs, GetDecoderManifestResult, GetDecoderManifestOutputArgs } from "./getDecoderManifest";
export const getDecoderManifest: typeof import("./getDecoderManifest").getDecoderManifest = null as any;
export const getDecoderManifestOutput: typeof import("./getDecoderManifest").getDecoderManifestOutput = null as any;

export { GetFleetArgs, GetFleetResult, GetFleetOutputArgs } from "./getFleet";
export const getFleet: typeof import("./getFleet").getFleet = null as any;
export const getFleetOutput: typeof import("./getFleet").getFleetOutput = null as any;

export { GetModelManifestArgs, GetModelManifestResult, GetModelManifestOutputArgs } from "./getModelManifest";
export const getModelManifest: typeof import("./getModelManifest").getModelManifest = null as any;
export const getModelManifestOutput: typeof import("./getModelManifest").getModelManifestOutput = null as any;

export { GetSignalCatalogArgs, GetSignalCatalogResult, GetSignalCatalogOutputArgs } from "./getSignalCatalog";
export const getSignalCatalog: typeof import("./getSignalCatalog").getSignalCatalog = null as any;
export const getSignalCatalogOutput: typeof import("./getSignalCatalog").getSignalCatalogOutput = null as any;

export { GetVehicleArgs, GetVehicleResult, GetVehicleOutputArgs } from "./getVehicle";
export const getVehicle: typeof import("./getVehicle").getVehicle = null as any;
export const getVehicleOutput: typeof import("./getVehicle").getVehicleOutput = null as any;

export { ModelManifestArgs } from "./modelManifest";
export type ModelManifest = import("./modelManifest").ModelManifest;
export const ModelManifest: typeof import("./modelManifest").ModelManifest = null as any;

export { SignalCatalogArgs } from "./signalCatalog";
export type SignalCatalog = import("./signalCatalog").SignalCatalog;
export const SignalCatalog: typeof import("./signalCatalog").SignalCatalog = null as any;

export { VehicleArgs } from "./vehicle";
export type Vehicle = import("./vehicle").Vehicle;
export const Vehicle: typeof import("./vehicle").Vehicle = null as any;

utilities.lazyLoad(exports, ["Campaign"], () => require("./campaign"));
utilities.lazyLoad(exports, ["DecoderManifest"], () => require("./decoderManifest"));
utilities.lazyLoad(exports, ["Fleet"], () => require("./fleet"));
utilities.lazyLoad(exports, ["getCampaign","getCampaignOutput"], () => require("./getCampaign"));
utilities.lazyLoad(exports, ["getDecoderManifest","getDecoderManifestOutput"], () => require("./getDecoderManifest"));
utilities.lazyLoad(exports, ["getFleet","getFleetOutput"], () => require("./getFleet"));
utilities.lazyLoad(exports, ["getModelManifest","getModelManifestOutput"], () => require("./getModelManifest"));
utilities.lazyLoad(exports, ["getSignalCatalog","getSignalCatalogOutput"], () => require("./getSignalCatalog"));
utilities.lazyLoad(exports, ["getVehicle","getVehicleOutput"], () => require("./getVehicle"));
utilities.lazyLoad(exports, ["ModelManifest"], () => require("./modelManifest"));
utilities.lazyLoad(exports, ["SignalCatalog"], () => require("./signalCatalog"));
utilities.lazyLoad(exports, ["Vehicle"], () => require("./vehicle"));

// Export enums:
export * from "../types/enums/iotfleetwise";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iotfleetwise:Campaign":
                return new Campaign(name, <any>undefined, { urn })
            case "aws-native:iotfleetwise:DecoderManifest":
                return new DecoderManifest(name, <any>undefined, { urn })
            case "aws-native:iotfleetwise:Fleet":
                return new Fleet(name, <any>undefined, { urn })
            case "aws-native:iotfleetwise:ModelManifest":
                return new ModelManifest(name, <any>undefined, { urn })
            case "aws-native:iotfleetwise:SignalCatalog":
                return new SignalCatalog(name, <any>undefined, { urn })
            case "aws-native:iotfleetwise:Vehicle":
                return new Vehicle(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iotfleetwise", _module)