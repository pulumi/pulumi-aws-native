// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApplicationArchitecture = {
    Arm64: "ARM64",
    X8664: "X86_64",
} as const;

/**
 * The cpu architecture of an application.
 */
export type ApplicationArchitecture = (typeof ApplicationArchitecture)[keyof typeof ApplicationArchitecture];
