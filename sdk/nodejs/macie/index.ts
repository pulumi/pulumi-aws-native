// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AllowListArgs } from "./allowList";
export type AllowList = import("./allowList").AllowList;
export const AllowList: typeof import("./allowList").AllowList = null as any;
utilities.lazyLoad(exports, ["AllowList"], () => require("./allowList"));

export { CustomDataIdentifierArgs } from "./customDataIdentifier";
export type CustomDataIdentifier = import("./customDataIdentifier").CustomDataIdentifier;
export const CustomDataIdentifier: typeof import("./customDataIdentifier").CustomDataIdentifier = null as any;
utilities.lazyLoad(exports, ["CustomDataIdentifier"], () => require("./customDataIdentifier"));

export { FindingsFilterArgs } from "./findingsFilter";
export type FindingsFilter = import("./findingsFilter").FindingsFilter;
export const FindingsFilter: typeof import("./findingsFilter").FindingsFilter = null as any;
utilities.lazyLoad(exports, ["FindingsFilter"], () => require("./findingsFilter"));

export { GetAllowListArgs, GetAllowListResult, GetAllowListOutputArgs } from "./getAllowList";
export const getAllowList: typeof import("./getAllowList").getAllowList = null as any;
export const getAllowListOutput: typeof import("./getAllowList").getAllowListOutput = null as any;
utilities.lazyLoad(exports, ["getAllowList","getAllowListOutput"], () => require("./getAllowList"));

export { GetCustomDataIdentifierArgs, GetCustomDataIdentifierResult, GetCustomDataIdentifierOutputArgs } from "./getCustomDataIdentifier";
export const getCustomDataIdentifier: typeof import("./getCustomDataIdentifier").getCustomDataIdentifier = null as any;
export const getCustomDataIdentifierOutput: typeof import("./getCustomDataIdentifier").getCustomDataIdentifierOutput = null as any;
utilities.lazyLoad(exports, ["getCustomDataIdentifier","getCustomDataIdentifierOutput"], () => require("./getCustomDataIdentifier"));

export { GetFindingsFilterArgs, GetFindingsFilterResult, GetFindingsFilterOutputArgs } from "./getFindingsFilter";
export const getFindingsFilter: typeof import("./getFindingsFilter").getFindingsFilter = null as any;
export const getFindingsFilterOutput: typeof import("./getFindingsFilter").getFindingsFilterOutput = null as any;
utilities.lazyLoad(exports, ["getFindingsFilter","getFindingsFilterOutput"], () => require("./getFindingsFilter"));

export { GetSessionArgs, GetSessionResult, GetSessionOutputArgs } from "./getSession";
export const getSession: typeof import("./getSession").getSession = null as any;
export const getSessionOutput: typeof import("./getSession").getSessionOutput = null as any;
utilities.lazyLoad(exports, ["getSession","getSessionOutput"], () => require("./getSession"));

export { SessionArgs } from "./session";
export type Session = import("./session").Session;
export const Session: typeof import("./session").Session = null as any;
utilities.lazyLoad(exports, ["Session"], () => require("./session"));


// Export enums:
export * from "../types/enums/macie";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:macie:AllowList":
                return new AllowList(name, <any>undefined, { urn })
            case "aws-native:macie:CustomDataIdentifier":
                return new CustomDataIdentifier(name, <any>undefined, { urn })
            case "aws-native:macie:FindingsFilter":
                return new FindingsFilter(name, <any>undefined, { urn })
            case "aws-native:macie:Session":
                return new Session(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "macie", _module)