// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ProfilingGroupComputePlatform = {
    Default: "Default",
    AWSLambda: "AWSLambda",
} as const;

/**
 * The compute platform of the profiling group.
 */
export type ProfilingGroupComputePlatform = (typeof ProfilingGroupComputePlatform)[keyof typeof ProfilingGroupComputePlatform];