// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetProfilingGroupArgs, GetProfilingGroupResult, GetProfilingGroupOutputArgs } from "./getProfilingGroup";
export const getProfilingGroup: typeof import("./getProfilingGroup").getProfilingGroup = null as any;
export const getProfilingGroupOutput: typeof import("./getProfilingGroup").getProfilingGroupOutput = null as any;

export { ProfilingGroupArgs } from "./profilingGroup";
export type ProfilingGroup = import("./profilingGroup").ProfilingGroup;
export const ProfilingGroup: typeof import("./profilingGroup").ProfilingGroup = null as any;

utilities.lazyLoad(exports, ["getProfilingGroup","getProfilingGroupOutput"], () => require("./getProfilingGroup"));
utilities.lazyLoad(exports, ["ProfilingGroup"], () => require("./profilingGroup"));

// Export enums:
export * from "../types/enums/codeguruprofiler";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:codeguruprofiler:ProfilingGroup":
                return new ProfilingGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "codeguruprofiler", _module)