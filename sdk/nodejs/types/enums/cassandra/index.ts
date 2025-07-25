// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const KeyspaceRegionListItem = {
    AfSouth1: "af-south-1",
    ApNortheast1: "ap-northeast-1",
    ApNortheast2: "ap-northeast-2",
    ApSouth1: "ap-south-1",
    ApSoutheast1: "ap-southeast-1",
    ApSoutheast2: "ap-southeast-2",
    CaCentral1: "ca-central-1",
    EuCentral1: "eu-central-1",
    EuNorth1: "eu-north-1",
    EuWest1: "eu-west-1",
    EuWest2: "eu-west-2",
    EuWest3: "eu-west-3",
    SaEast1: "sa-east-1",
    UsEast1: "us-east-1",
    UsEast2: "us-east-2",
    UsWest1: "us-west-1",
    UsWest2: "us-west-2",
} as const;

export type KeyspaceRegionListItem = (typeof KeyspaceRegionListItem)[keyof typeof KeyspaceRegionListItem];

export const KeyspaceReplicationSpecificationReplicationStrategy = {
    SingleRegion: "SINGLE_REGION",
    MultiRegion: "MULTI_REGION",
} as const;

/**
 * The options are:
 *
 * - `SINGLE_REGION` (optional)
 * - `MULTI_REGION`
 *
 * If no value is specified, the default is `SINGLE_REGION` . If `MULTI_REGION` is specified, `RegionList` is required.
 */
export type KeyspaceReplicationSpecificationReplicationStrategy = (typeof KeyspaceReplicationSpecificationReplicationStrategy)[keyof typeof KeyspaceReplicationSpecificationReplicationStrategy];

export const TableCdcStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Indicates whether CDC is enabled or disabled for the table
 */
export type TableCdcStatus = (typeof TableCdcStatus)[keyof typeof TableCdcStatus];

export const TableCdcViewType = {
    NewImage: "NEW_IMAGE",
    OldImage: "OLD_IMAGE",
    KeysOnly: "KEYS_ONLY",
    NewAndOldImages: "NEW_AND_OLD_IMAGES",
} as const;

/**
 * Specifies what data should be captured in the change data stream
 */
export type TableCdcViewType = (typeof TableCdcViewType)[keyof typeof TableCdcViewType];

export const TableClusteringKeyColumnOrderBy = {
    Asc: "ASC",
    Desc: "DESC",
} as const;

/**
 * The order in which this column's data is stored:
 *
 * - `ASC` (default) - The column's data is stored in ascending order.
 * - `DESC` - The column's data is stored in descending order.
 */
export type TableClusteringKeyColumnOrderBy = (typeof TableClusteringKeyColumnOrderBy)[keyof typeof TableClusteringKeyColumnOrderBy];

export const TableEncryptionType = {
    AwsOwnedKmsKey: "AWS_OWNED_KMS_KEY",
    CustomerManagedKmsKey: "CUSTOMER_MANAGED_KMS_KEY",
} as const;

/**
 * Server-side encryption type
 */
export type TableEncryptionType = (typeof TableEncryptionType)[keyof typeof TableEncryptionType];

export const TableMode = {
    Provisioned: "PROVISIONED",
    OnDemand: "ON_DEMAND",
} as const;

/**
 * Capacity mode for the specified table
 */
export type TableMode = (typeof TableMode)[keyof typeof TableMode];
