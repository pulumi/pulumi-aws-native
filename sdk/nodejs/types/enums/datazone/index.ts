// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DataSourceEnableSetting = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Specifies whether the data source is enabled.
 */
export type DataSourceEnableSetting = (typeof DataSourceEnableSetting)[keyof typeof DataSourceEnableSetting];

export const DataSourceFilterExpressionType = {
    Include: "INCLUDE",
    Exclude: "EXCLUDE",
} as const;

/**
 * The search filter expression type.
 */
export type DataSourceFilterExpressionType = (typeof DataSourceFilterExpressionType)[keyof typeof DataSourceFilterExpressionType];

export const DataSourceStatus = {
    Creating: "CREATING",
    FailedCreation: "FAILED_CREATION",
    Ready: "READY",
    Updating: "UPDATING",
    FailedUpdate: "FAILED_UPDATE",
    Running: "RUNNING",
    Deleting: "DELETING",
    FailedDeletion: "FAILED_DELETION",
} as const;

/**
 * The status of the data source.
 */
export type DataSourceStatus = (typeof DataSourceStatus)[keyof typeof DataSourceStatus];

export const DomainAuthType = {
    IamIdc: "IAM_IDC",
    Disabled: "DISABLED",
} as const;

/**
 * The type of single sign-on in Amazon DataZone.
 */
export type DomainAuthType = (typeof DomainAuthType)[keyof typeof DomainAuthType];

export const DomainStatus = {
    Creating: "CREATING",
    Available: "AVAILABLE",
    CreationFailed: "CREATION_FAILED",
    Deleting: "DELETING",
    Deleted: "DELETED",
    DeletionFailed: "DELETION_FAILED",
} as const;

/**
 * The status of the Amazon DataZone domain.
 */
export type DomainStatus = (typeof DomainStatus)[keyof typeof DomainStatus];

export const DomainUserAssignment = {
    Automatic: "AUTOMATIC",
    Manual: "MANUAL",
} as const;

/**
 * The single sign-on user assignment in Amazon DataZone.
 */
export type DomainUserAssignment = (typeof DomainUserAssignment)[keyof typeof DomainUserAssignment];

export const EnvironmentStatus = {
    Active: "ACTIVE",
    Creating: "CREATING",
    Updating: "UPDATING",
    Deleting: "DELETING",
    CreateFailed: "CREATE_FAILED",
    UpdateFailed: "UPDATE_FAILED",
    DeleteFailed: "DELETE_FAILED",
    ValidationFailed: "VALIDATION_FAILED",
    Suspended: "SUSPENDED",
    Disabled: "DISABLED",
    Expired: "EXPIRED",
    Deleted: "DELETED",
    Inaccessible: "INACCESSIBLE",
} as const;

/**
 * The status of the Amazon DataZone environment.
 */
export type EnvironmentStatus = (typeof EnvironmentStatus)[keyof typeof EnvironmentStatus];