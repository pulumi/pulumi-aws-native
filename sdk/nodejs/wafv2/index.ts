// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetIPSetArgs, GetIPSetResult, GetIPSetOutputArgs } from "./getIPSet";
export const getIPSet: typeof import("./getIPSet").getIPSet = null as any;
export const getIPSetOutput: typeof import("./getIPSet").getIPSetOutput = null as any;

export { GetLoggingConfigurationArgs, GetLoggingConfigurationResult, GetLoggingConfigurationOutputArgs } from "./getLoggingConfiguration";
export const getLoggingConfiguration: typeof import("./getLoggingConfiguration").getLoggingConfiguration = null as any;
export const getLoggingConfigurationOutput: typeof import("./getLoggingConfiguration").getLoggingConfigurationOutput = null as any;

export { GetRegexPatternSetArgs, GetRegexPatternSetResult, GetRegexPatternSetOutputArgs } from "./getRegexPatternSet";
export const getRegexPatternSet: typeof import("./getRegexPatternSet").getRegexPatternSet = null as any;
export const getRegexPatternSetOutput: typeof import("./getRegexPatternSet").getRegexPatternSetOutput = null as any;

export { GetRuleGroupArgs, GetRuleGroupResult, GetRuleGroupOutputArgs } from "./getRuleGroup";
export const getRuleGroup: typeof import("./getRuleGroup").getRuleGroup = null as any;
export const getRuleGroupOutput: typeof import("./getRuleGroup").getRuleGroupOutput = null as any;

export { GetWebACLArgs, GetWebACLResult, GetWebACLOutputArgs } from "./getWebACL";
export const getWebACL: typeof import("./getWebACL").getWebACL = null as any;
export const getWebACLOutput: typeof import("./getWebACL").getWebACLOutput = null as any;

export { IPSetArgs } from "./ipset";
export type IPSet = import("./ipset").IPSet;
export const IPSet: typeof import("./ipset").IPSet = null as any;

export { LoggingConfigurationArgs } from "./loggingConfiguration";
export type LoggingConfiguration = import("./loggingConfiguration").LoggingConfiguration;
export const LoggingConfiguration: typeof import("./loggingConfiguration").LoggingConfiguration = null as any;

export { RegexPatternSetArgs } from "./regexPatternSet";
export type RegexPatternSet = import("./regexPatternSet").RegexPatternSet;
export const RegexPatternSet: typeof import("./regexPatternSet").RegexPatternSet = null as any;

export { RuleGroupArgs } from "./ruleGroup";
export type RuleGroup = import("./ruleGroup").RuleGroup;
export const RuleGroup: typeof import("./ruleGroup").RuleGroup = null as any;

export { WebACLArgs } from "./webACL";
export type WebACL = import("./webACL").WebACL;
export const WebACL: typeof import("./webACL").WebACL = null as any;

export { WebACLAssociationArgs } from "./webACLAssociation";
export type WebACLAssociation = import("./webACLAssociation").WebACLAssociation;
export const WebACLAssociation: typeof import("./webACLAssociation").WebACLAssociation = null as any;

utilities.lazyLoad(exports, ["getIPSet","getIPSetOutput"], () => require("./getIPSet"));
utilities.lazyLoad(exports, ["getLoggingConfiguration","getLoggingConfigurationOutput"], () => require("./getLoggingConfiguration"));
utilities.lazyLoad(exports, ["getRegexPatternSet","getRegexPatternSetOutput"], () => require("./getRegexPatternSet"));
utilities.lazyLoad(exports, ["getRuleGroup","getRuleGroupOutput"], () => require("./getRuleGroup"));
utilities.lazyLoad(exports, ["getWebACL","getWebACLOutput"], () => require("./getWebACL"));
utilities.lazyLoad(exports, ["IPSet"], () => require("./ipset"));
utilities.lazyLoad(exports, ["LoggingConfiguration"], () => require("./loggingConfiguration"));
utilities.lazyLoad(exports, ["RegexPatternSet"], () => require("./regexPatternSet"));
utilities.lazyLoad(exports, ["RuleGroup"], () => require("./ruleGroup"));
utilities.lazyLoad(exports, ["WebACL"], () => require("./webACL"));
utilities.lazyLoad(exports, ["WebACLAssociation"], () => require("./webACLAssociation"));

// Export enums:
export * from "../types/enums/wafv2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:wafv2:IPSet":
                return new IPSet(name, <any>undefined, { urn })
            case "aws-native:wafv2:LoggingConfiguration":
                return new LoggingConfiguration(name, <any>undefined, { urn })
            case "aws-native:wafv2:RegexPatternSet":
                return new RegexPatternSet(name, <any>undefined, { urn })
            case "aws-native:wafv2:RuleGroup":
                return new RuleGroup(name, <any>undefined, { urn })
            case "aws-native:wafv2:WebACL":
                return new WebACL(name, <any>undefined, { urn })
            case "aws-native:wafv2:WebACLAssociation":
                return new WebACLAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "wafv2", _module)