// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DataCellsFilterArgs } from "./dataCellsFilter";
export type DataCellsFilter = import("./dataCellsFilter").DataCellsFilter;
export const DataCellsFilter: typeof import("./dataCellsFilter").DataCellsFilter = null as any;

export { DataLakeSettingsArgs } from "./dataLakeSettings";
export type DataLakeSettings = import("./dataLakeSettings").DataLakeSettings;
export const DataLakeSettings: typeof import("./dataLakeSettings").DataLakeSettings = null as any;

export { GetDataLakeSettingsArgs, GetDataLakeSettingsResult, GetDataLakeSettingsOutputArgs } from "./getDataLakeSettings";
export const getDataLakeSettings: typeof import("./getDataLakeSettings").getDataLakeSettings = null as any;
export const getDataLakeSettingsOutput: typeof import("./getDataLakeSettings").getDataLakeSettingsOutput = null as any;

export { GetPermissionsArgs, GetPermissionsResult, GetPermissionsOutputArgs } from "./getPermissions";
export const getPermissions: typeof import("./getPermissions").getPermissions = null as any;
export const getPermissionsOutput: typeof import("./getPermissions").getPermissionsOutput = null as any;

export { GetPrincipalPermissionsArgs, GetPrincipalPermissionsResult, GetPrincipalPermissionsOutputArgs } from "./getPrincipalPermissions";
export const getPrincipalPermissions: typeof import("./getPrincipalPermissions").getPrincipalPermissions = null as any;
export const getPrincipalPermissionsOutput: typeof import("./getPrincipalPermissions").getPrincipalPermissionsOutput = null as any;

export { GetResourceArgs, GetResourceResult, GetResourceOutputArgs } from "./getResource";
export const getResource: typeof import("./getResource").getResource = null as any;
export const getResourceOutput: typeof import("./getResource").getResourceOutput = null as any;

export { GetTagArgs, GetTagResult, GetTagOutputArgs } from "./getTag";
export const getTag: typeof import("./getTag").getTag = null as any;
export const getTagOutput: typeof import("./getTag").getTagOutput = null as any;

export { GetTagAssociationArgs, GetTagAssociationResult, GetTagAssociationOutputArgs } from "./getTagAssociation";
export const getTagAssociation: typeof import("./getTagAssociation").getTagAssociation = null as any;
export const getTagAssociationOutput: typeof import("./getTagAssociation").getTagAssociationOutput = null as any;

export { PermissionsArgs } from "./permissions";
export type Permissions = import("./permissions").Permissions;
export const Permissions: typeof import("./permissions").Permissions = null as any;

export { PrincipalPermissionsArgs } from "./principalPermissions";
export type PrincipalPermissions = import("./principalPermissions").PrincipalPermissions;
export const PrincipalPermissions: typeof import("./principalPermissions").PrincipalPermissions = null as any;

export { ResourceArgs } from "./resource";
export type Resource = import("./resource").Resource;
export const Resource: typeof import("./resource").Resource = null as any;

export { TagArgs } from "./tag";
export type Tag = import("./tag").Tag;
export const Tag: typeof import("./tag").Tag = null as any;

export { TagAssociationArgs } from "./tagAssociation";
export type TagAssociation = import("./tagAssociation").TagAssociation;
export const TagAssociation: typeof import("./tagAssociation").TagAssociation = null as any;

utilities.lazyLoad(exports, ["DataCellsFilter"], () => require("./dataCellsFilter"));
utilities.lazyLoad(exports, ["DataLakeSettings"], () => require("./dataLakeSettings"));
utilities.lazyLoad(exports, ["getDataLakeSettings","getDataLakeSettingsOutput"], () => require("./getDataLakeSettings"));
utilities.lazyLoad(exports, ["getPermissions","getPermissionsOutput"], () => require("./getPermissions"));
utilities.lazyLoad(exports, ["getPrincipalPermissions","getPrincipalPermissionsOutput"], () => require("./getPrincipalPermissions"));
utilities.lazyLoad(exports, ["getResource","getResourceOutput"], () => require("./getResource"));
utilities.lazyLoad(exports, ["getTag","getTagOutput"], () => require("./getTag"));
utilities.lazyLoad(exports, ["getTagAssociation","getTagAssociationOutput"], () => require("./getTagAssociation"));
utilities.lazyLoad(exports, ["Permissions"], () => require("./permissions"));
utilities.lazyLoad(exports, ["PrincipalPermissions"], () => require("./principalPermissions"));
utilities.lazyLoad(exports, ["Resource"], () => require("./resource"));
utilities.lazyLoad(exports, ["Tag"], () => require("./tag"));
utilities.lazyLoad(exports, ["TagAssociation"], () => require("./tagAssociation"));

// Export enums:
export * from "../types/enums/lakeformation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:lakeformation:DataCellsFilter":
                return new DataCellsFilter(name, <any>undefined, { urn })
            case "aws-native:lakeformation:DataLakeSettings":
                return new DataLakeSettings(name, <any>undefined, { urn })
            case "aws-native:lakeformation:Permissions":
                return new Permissions(name, <any>undefined, { urn })
            case "aws-native:lakeformation:PrincipalPermissions":
                return new PrincipalPermissions(name, <any>undefined, { urn })
            case "aws-native:lakeformation:Resource":
                return new Resource(name, <any>undefined, { urn })
            case "aws-native:lakeformation:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "aws-native:lakeformation:TagAssociation":
                return new TagAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "lakeformation", _module)