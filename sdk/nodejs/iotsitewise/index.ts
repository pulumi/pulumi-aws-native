// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccessPolicyArgs } from "./accessPolicy";
export type AccessPolicy = import("./accessPolicy").AccessPolicy;
export const AccessPolicy: typeof import("./accessPolicy").AccessPolicy = null as any;
utilities.lazyLoad(exports, ["AccessPolicy"], () => require("./accessPolicy"));

export { AssetArgs } from "./asset";
export type Asset = import("./asset").Asset;
export const Asset: typeof import("./asset").Asset = null as any;
utilities.lazyLoad(exports, ["Asset"], () => require("./asset"));

export { AssetModelArgs } from "./assetModel";
export type AssetModel = import("./assetModel").AssetModel;
export const AssetModel: typeof import("./assetModel").AssetModel = null as any;
utilities.lazyLoad(exports, ["AssetModel"], () => require("./assetModel"));

export { DashboardArgs } from "./dashboard";
export type Dashboard = import("./dashboard").Dashboard;
export const Dashboard: typeof import("./dashboard").Dashboard = null as any;
utilities.lazyLoad(exports, ["Dashboard"], () => require("./dashboard"));

export { GatewayArgs } from "./gateway";
export type Gateway = import("./gateway").Gateway;
export const Gateway: typeof import("./gateway").Gateway = null as any;
utilities.lazyLoad(exports, ["Gateway"], () => require("./gateway"));

export { GetAccessPolicyArgs, GetAccessPolicyResult, GetAccessPolicyOutputArgs } from "./getAccessPolicy";
export const getAccessPolicy: typeof import("./getAccessPolicy").getAccessPolicy = null as any;
export const getAccessPolicyOutput: typeof import("./getAccessPolicy").getAccessPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getAccessPolicy","getAccessPolicyOutput"], () => require("./getAccessPolicy"));

export { GetAssetArgs, GetAssetResult, GetAssetOutputArgs } from "./getAsset";
export const getAsset: typeof import("./getAsset").getAsset = null as any;
export const getAssetOutput: typeof import("./getAsset").getAssetOutput = null as any;
utilities.lazyLoad(exports, ["getAsset","getAssetOutput"], () => require("./getAsset"));

export { GetAssetModelArgs, GetAssetModelResult, GetAssetModelOutputArgs } from "./getAssetModel";
export const getAssetModel: typeof import("./getAssetModel").getAssetModel = null as any;
export const getAssetModelOutput: typeof import("./getAssetModel").getAssetModelOutput = null as any;
utilities.lazyLoad(exports, ["getAssetModel","getAssetModelOutput"], () => require("./getAssetModel"));

export { GetDashboardArgs, GetDashboardResult, GetDashboardOutputArgs } from "./getDashboard";
export const getDashboard: typeof import("./getDashboard").getDashboard = null as any;
export const getDashboardOutput: typeof import("./getDashboard").getDashboardOutput = null as any;
utilities.lazyLoad(exports, ["getDashboard","getDashboardOutput"], () => require("./getDashboard"));

export { GetGatewayArgs, GetGatewayResult, GetGatewayOutputArgs } from "./getGateway";
export const getGateway: typeof import("./getGateway").getGateway = null as any;
export const getGatewayOutput: typeof import("./getGateway").getGatewayOutput = null as any;
utilities.lazyLoad(exports, ["getGateway","getGatewayOutput"], () => require("./getGateway"));

export { GetPortalArgs, GetPortalResult, GetPortalOutputArgs } from "./getPortal";
export const getPortal: typeof import("./getPortal").getPortal = null as any;
export const getPortalOutput: typeof import("./getPortal").getPortalOutput = null as any;
utilities.lazyLoad(exports, ["getPortal","getPortalOutput"], () => require("./getPortal"));

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));

export { PortalArgs } from "./portal";
export type Portal = import("./portal").Portal;
export const Portal: typeof import("./portal").Portal = null as any;
utilities.lazyLoad(exports, ["Portal"], () => require("./portal"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));


// Export enums:
export * from "../types/enums/iotsitewise";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iotsitewise:AccessPolicy":
                return new AccessPolicy(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:Asset":
                return new Asset(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:AssetModel":
                return new AssetModel(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:Dashboard":
                return new Dashboard(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:Portal":
                return new Portal(name, <any>undefined, { urn })
            case "aws-native:iotsitewise:Project":
                return new Project(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iotsitewise", _module)