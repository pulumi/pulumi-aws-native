// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApiDestinationArgs } from "./apiDestination";
export type ApiDestination = import("./apiDestination").ApiDestination;
export const ApiDestination: typeof import("./apiDestination").ApiDestination = null as any;

export { ArchiveArgs } from "./archive";
export type Archive = import("./archive").Archive;
export const Archive: typeof import("./archive").Archive = null as any;

export { ConnectionArgs } from "./connection";
export type Connection = import("./connection").Connection;
export const Connection: typeof import("./connection").Connection = null as any;

export { EndpointArgs } from "./endpoint";
export type Endpoint = import("./endpoint").Endpoint;
export const Endpoint: typeof import("./endpoint").Endpoint = null as any;

export { EventBusArgs } from "./eventBus";
export type EventBus = import("./eventBus").EventBus;
export const EventBus: typeof import("./eventBus").EventBus = null as any;

export { EventBusPolicyArgs } from "./eventBusPolicy";
export type EventBusPolicy = import("./eventBusPolicy").EventBusPolicy;
export const EventBusPolicy: typeof import("./eventBusPolicy").EventBusPolicy = null as any;

export { GetApiDestinationArgs, GetApiDestinationResult, GetApiDestinationOutputArgs } from "./getApiDestination";
export const getApiDestination: typeof import("./getApiDestination").getApiDestination = null as any;
export const getApiDestinationOutput: typeof import("./getApiDestination").getApiDestinationOutput = null as any;

export { GetArchiveArgs, GetArchiveResult, GetArchiveOutputArgs } from "./getArchive";
export const getArchive: typeof import("./getArchive").getArchive = null as any;
export const getArchiveOutput: typeof import("./getArchive").getArchiveOutput = null as any;

export { GetConnectionArgs, GetConnectionResult, GetConnectionOutputArgs } from "./getConnection";
export const getConnection: typeof import("./getConnection").getConnection = null as any;
export const getConnectionOutput: typeof import("./getConnection").getConnectionOutput = null as any;

export { GetEndpointArgs, GetEndpointResult, GetEndpointOutputArgs } from "./getEndpoint";
export const getEndpoint: typeof import("./getEndpoint").getEndpoint = null as any;
export const getEndpointOutput: typeof import("./getEndpoint").getEndpointOutput = null as any;

export { GetEventBusArgs, GetEventBusResult, GetEventBusOutputArgs } from "./getEventBus";
export const getEventBus: typeof import("./getEventBus").getEventBus = null as any;
export const getEventBusOutput: typeof import("./getEventBus").getEventBusOutput = null as any;

export { GetEventBusPolicyArgs, GetEventBusPolicyResult, GetEventBusPolicyOutputArgs } from "./getEventBusPolicy";
export const getEventBusPolicy: typeof import("./getEventBusPolicy").getEventBusPolicy = null as any;
export const getEventBusPolicyOutput: typeof import("./getEventBusPolicy").getEventBusPolicyOutput = null as any;

export { GetRuleArgs, GetRuleResult, GetRuleOutputArgs } from "./getRule";
export const getRule: typeof import("./getRule").getRule = null as any;
export const getRuleOutput: typeof import("./getRule").getRuleOutput = null as any;

export { RuleArgs } from "./rule";
export type Rule = import("./rule").Rule;
export const Rule: typeof import("./rule").Rule = null as any;

utilities.lazyLoad(exports, ["ApiDestination"], () => require("./apiDestination"));
utilities.lazyLoad(exports, ["Archive"], () => require("./archive"));
utilities.lazyLoad(exports, ["Connection"], () => require("./connection"));
utilities.lazyLoad(exports, ["Endpoint"], () => require("./endpoint"));
utilities.lazyLoad(exports, ["EventBus"], () => require("./eventBus"));
utilities.lazyLoad(exports, ["EventBusPolicy"], () => require("./eventBusPolicy"));
utilities.lazyLoad(exports, ["getApiDestination","getApiDestinationOutput"], () => require("./getApiDestination"));
utilities.lazyLoad(exports, ["getArchive","getArchiveOutput"], () => require("./getArchive"));
utilities.lazyLoad(exports, ["getConnection","getConnectionOutput"], () => require("./getConnection"));
utilities.lazyLoad(exports, ["getEndpoint","getEndpointOutput"], () => require("./getEndpoint"));
utilities.lazyLoad(exports, ["getEventBus","getEventBusOutput"], () => require("./getEventBus"));
utilities.lazyLoad(exports, ["getEventBusPolicy","getEventBusPolicyOutput"], () => require("./getEventBusPolicy"));
utilities.lazyLoad(exports, ["getRule","getRuleOutput"], () => require("./getRule"));
utilities.lazyLoad(exports, ["Rule"], () => require("./rule"));

// Export enums:
export * from "../types/enums/events";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:events:ApiDestination":
                return new ApiDestination(name, <any>undefined, { urn })
            case "aws-native:events:Archive":
                return new Archive(name, <any>undefined, { urn })
            case "aws-native:events:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "aws-native:events:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "aws-native:events:EventBus":
                return new EventBus(name, <any>undefined, { urn })
            case "aws-native:events:EventBusPolicy":
                return new EventBusPolicy(name, <any>undefined, { urn })
            case "aws-native:events:Rule":
                return new Rule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "events", _module)