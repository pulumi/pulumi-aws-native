// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DirectoryBucketDataRedundancy = {
    SingleAvailabilityZone: "SingleAvailabilityZone",
} as const;

/**
 * Specifies the number of Avilability Zone that's used for redundancy for the bucket.
 */
export type DirectoryBucketDataRedundancy = (typeof DirectoryBucketDataRedundancy)[keyof typeof DirectoryBucketDataRedundancy];