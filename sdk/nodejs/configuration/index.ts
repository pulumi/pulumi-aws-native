// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AggregationAuthorizationArgs } from "./aggregationAuthorization";
export type AggregationAuthorization = import("./aggregationAuthorization").AggregationAuthorization;
export const AggregationAuthorization: typeof import("./aggregationAuthorization").AggregationAuthorization = null as any;
utilities.lazyLoad(exports, ["AggregationAuthorization"], () => require("./aggregationAuthorization"));

export { ConfigRuleArgs } from "./configRule";
export type ConfigRule = import("./configRule").ConfigRule;
export const ConfigRule: typeof import("./configRule").ConfigRule = null as any;
utilities.lazyLoad(exports, ["ConfigRule"], () => require("./configRule"));

export { ConfigurationAggregatorArgs } from "./configurationAggregator";
export type ConfigurationAggregator = import("./configurationAggregator").ConfigurationAggregator;
export const ConfigurationAggregator: typeof import("./configurationAggregator").ConfigurationAggregator = null as any;
utilities.lazyLoad(exports, ["ConfigurationAggregator"], () => require("./configurationAggregator"));

export { ConformancePackArgs } from "./conformancePack";
export type ConformancePack = import("./conformancePack").ConformancePack;
export const ConformancePack: typeof import("./conformancePack").ConformancePack = null as any;
utilities.lazyLoad(exports, ["ConformancePack"], () => require("./conformancePack"));

export { GetAggregationAuthorizationArgs, GetAggregationAuthorizationResult, GetAggregationAuthorizationOutputArgs } from "./getAggregationAuthorization";
export const getAggregationAuthorization: typeof import("./getAggregationAuthorization").getAggregationAuthorization = null as any;
export const getAggregationAuthorizationOutput: typeof import("./getAggregationAuthorization").getAggregationAuthorizationOutput = null as any;
utilities.lazyLoad(exports, ["getAggregationAuthorization","getAggregationAuthorizationOutput"], () => require("./getAggregationAuthorization"));

export { GetConfigRuleArgs, GetConfigRuleResult, GetConfigRuleOutputArgs } from "./getConfigRule";
export const getConfigRule: typeof import("./getConfigRule").getConfigRule = null as any;
export const getConfigRuleOutput: typeof import("./getConfigRule").getConfigRuleOutput = null as any;
utilities.lazyLoad(exports, ["getConfigRule","getConfigRuleOutput"], () => require("./getConfigRule"));

export { GetConfigurationAggregatorArgs, GetConfigurationAggregatorResult, GetConfigurationAggregatorOutputArgs } from "./getConfigurationAggregator";
export const getConfigurationAggregator: typeof import("./getConfigurationAggregator").getConfigurationAggregator = null as any;
export const getConfigurationAggregatorOutput: typeof import("./getConfigurationAggregator").getConfigurationAggregatorOutput = null as any;
utilities.lazyLoad(exports, ["getConfigurationAggregator","getConfigurationAggregatorOutput"], () => require("./getConfigurationAggregator"));

export { GetConformancePackArgs, GetConformancePackResult, GetConformancePackOutputArgs } from "./getConformancePack";
export const getConformancePack: typeof import("./getConformancePack").getConformancePack = null as any;
export const getConformancePackOutput: typeof import("./getConformancePack").getConformancePackOutput = null as any;
utilities.lazyLoad(exports, ["getConformancePack","getConformancePackOutput"], () => require("./getConformancePack"));

export { GetOrganizationConformancePackArgs, GetOrganizationConformancePackResult, GetOrganizationConformancePackOutputArgs } from "./getOrganizationConformancePack";
export const getOrganizationConformancePack: typeof import("./getOrganizationConformancePack").getOrganizationConformancePack = null as any;
export const getOrganizationConformancePackOutput: typeof import("./getOrganizationConformancePack").getOrganizationConformancePackOutput = null as any;
utilities.lazyLoad(exports, ["getOrganizationConformancePack","getOrganizationConformancePackOutput"], () => require("./getOrganizationConformancePack"));

export { GetStoredQueryArgs, GetStoredQueryResult, GetStoredQueryOutputArgs } from "./getStoredQuery";
export const getStoredQuery: typeof import("./getStoredQuery").getStoredQuery = null as any;
export const getStoredQueryOutput: typeof import("./getStoredQuery").getStoredQueryOutput = null as any;
utilities.lazyLoad(exports, ["getStoredQuery","getStoredQueryOutput"], () => require("./getStoredQuery"));

export { OrganizationConformancePackArgs } from "./organizationConformancePack";
export type OrganizationConformancePack = import("./organizationConformancePack").OrganizationConformancePack;
export const OrganizationConformancePack: typeof import("./organizationConformancePack").OrganizationConformancePack = null as any;
utilities.lazyLoad(exports, ["OrganizationConformancePack"], () => require("./organizationConformancePack"));

export { StoredQueryArgs } from "./storedQuery";
export type StoredQuery = import("./storedQuery").StoredQuery;
export const StoredQuery: typeof import("./storedQuery").StoredQuery = null as any;
utilities.lazyLoad(exports, ["StoredQuery"], () => require("./storedQuery"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:configuration:AggregationAuthorization":
                return new AggregationAuthorization(name, <any>undefined, { urn })
            case "aws-native:configuration:ConfigRule":
                return new ConfigRule(name, <any>undefined, { urn })
            case "aws-native:configuration:ConfigurationAggregator":
                return new ConfigurationAggregator(name, <any>undefined, { urn })
            case "aws-native:configuration:ConformancePack":
                return new ConformancePack(name, <any>undefined, { urn })
            case "aws-native:configuration:OrganizationConformancePack":
                return new OrganizationConformancePack(name, <any>undefined, { urn })
            case "aws-native:configuration:StoredQuery":
                return new StoredQuery(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "configuration", _module)
