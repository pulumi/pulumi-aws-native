// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetIdMappingWorkflowArgs, GetIdMappingWorkflowResult, GetIdMappingWorkflowOutputArgs } from "./getIdMappingWorkflow";
export const getIdMappingWorkflow: typeof import("./getIdMappingWorkflow").getIdMappingWorkflow = null as any;
export const getIdMappingWorkflowOutput: typeof import("./getIdMappingWorkflow").getIdMappingWorkflowOutput = null as any;
utilities.lazyLoad(exports, ["getIdMappingWorkflow","getIdMappingWorkflowOutput"], () => require("./getIdMappingWorkflow"));

export { GetIdNamespaceArgs, GetIdNamespaceResult, GetIdNamespaceOutputArgs } from "./getIdNamespace";
export const getIdNamespace: typeof import("./getIdNamespace").getIdNamespace = null as any;
export const getIdNamespaceOutput: typeof import("./getIdNamespace").getIdNamespaceOutput = null as any;
utilities.lazyLoad(exports, ["getIdNamespace","getIdNamespaceOutput"], () => require("./getIdNamespace"));

export { GetMatchingWorkflowArgs, GetMatchingWorkflowResult, GetMatchingWorkflowOutputArgs } from "./getMatchingWorkflow";
export const getMatchingWorkflow: typeof import("./getMatchingWorkflow").getMatchingWorkflow = null as any;
export const getMatchingWorkflowOutput: typeof import("./getMatchingWorkflow").getMatchingWorkflowOutput = null as any;
utilities.lazyLoad(exports, ["getMatchingWorkflow","getMatchingWorkflowOutput"], () => require("./getMatchingWorkflow"));

export { GetPolicyStatementArgs, GetPolicyStatementResult, GetPolicyStatementOutputArgs } from "./getPolicyStatement";
export const getPolicyStatement: typeof import("./getPolicyStatement").getPolicyStatement = null as any;
export const getPolicyStatementOutput: typeof import("./getPolicyStatement").getPolicyStatementOutput = null as any;
utilities.lazyLoad(exports, ["getPolicyStatement","getPolicyStatementOutput"], () => require("./getPolicyStatement"));

export { GetSchemaMappingArgs, GetSchemaMappingResult, GetSchemaMappingOutputArgs } from "./getSchemaMapping";
export const getSchemaMapping: typeof import("./getSchemaMapping").getSchemaMapping = null as any;
export const getSchemaMappingOutput: typeof import("./getSchemaMapping").getSchemaMappingOutput = null as any;
utilities.lazyLoad(exports, ["getSchemaMapping","getSchemaMappingOutput"], () => require("./getSchemaMapping"));

export { IdMappingWorkflowArgs } from "./idMappingWorkflow";
export type IdMappingWorkflow = import("./idMappingWorkflow").IdMappingWorkflow;
export const IdMappingWorkflow: typeof import("./idMappingWorkflow").IdMappingWorkflow = null as any;
utilities.lazyLoad(exports, ["IdMappingWorkflow"], () => require("./idMappingWorkflow"));

export { IdNamespaceArgs } from "./idNamespace";
export type IdNamespace = import("./idNamespace").IdNamespace;
export const IdNamespace: typeof import("./idNamespace").IdNamespace = null as any;
utilities.lazyLoad(exports, ["IdNamespace"], () => require("./idNamespace"));

export { MatchingWorkflowArgs } from "./matchingWorkflow";
export type MatchingWorkflow = import("./matchingWorkflow").MatchingWorkflow;
export const MatchingWorkflow: typeof import("./matchingWorkflow").MatchingWorkflow = null as any;
utilities.lazyLoad(exports, ["MatchingWorkflow"], () => require("./matchingWorkflow"));

export { PolicyStatementArgs } from "./policyStatement";
export type PolicyStatement = import("./policyStatement").PolicyStatement;
export const PolicyStatement: typeof import("./policyStatement").PolicyStatement = null as any;
utilities.lazyLoad(exports, ["PolicyStatement"], () => require("./policyStatement"));

export { SchemaMappingArgs } from "./schemaMapping";
export type SchemaMapping = import("./schemaMapping").SchemaMapping;
export const SchemaMapping: typeof import("./schemaMapping").SchemaMapping = null as any;
utilities.lazyLoad(exports, ["SchemaMapping"], () => require("./schemaMapping"));


// Export enums:
export * from "../types/enums/entityresolution";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:entityresolution:IdMappingWorkflow":
                return new IdMappingWorkflow(name, <any>undefined, { urn })
            case "aws-native:entityresolution:IdNamespace":
                return new IdNamespace(name, <any>undefined, { urn })
            case "aws-native:entityresolution:MatchingWorkflow":
                return new MatchingWorkflow(name, <any>undefined, { urn })
            case "aws-native:entityresolution:PolicyStatement":
                return new PolicyStatement(name, <any>undefined, { urn })
            case "aws-native:entityresolution:SchemaMapping":
                return new SchemaMapping(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "entityresolution", _module)
