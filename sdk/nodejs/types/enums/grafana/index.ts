// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const WorkspaceAccountAccessType = {
    CurrentAccount: "CURRENT_ACCOUNT",
    Organization: "ORGANIZATION",
} as const;

/**
 * These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.
 */
export type WorkspaceAccountAccessType = (typeof WorkspaceAccountAccessType)[keyof typeof WorkspaceAccountAccessType];

export const WorkspaceAuthenticationProviderTypes = {
    AwsSso: "AWS_SSO",
    Saml: "SAML",
} as const;

/**
 * Valid workspace authentication providers.
 */
export type WorkspaceAuthenticationProviderTypes = (typeof WorkspaceAuthenticationProviderTypes)[keyof typeof WorkspaceAuthenticationProviderTypes];

export const WorkspaceDataSourceType = {
    AmazonOpensearchService: "AMAZON_OPENSEARCH_SERVICE",
    Cloudwatch: "CLOUDWATCH",
    Prometheus: "PROMETHEUS",
    Xray: "XRAY",
    Timestream: "TIMESTREAM",
    Sitewise: "SITEWISE",
    Athena: "ATHENA",
    Redshift: "REDSHIFT",
} as const;

/**
 * These enums represent valid AWS data sources that can be queried via the Grafana workspace. These data sources are primarily used to help customers visualize which data sources have been added to a service managed workspace IAM role.
 */
export type WorkspaceDataSourceType = (typeof WorkspaceDataSourceType)[keyof typeof WorkspaceDataSourceType];

export const WorkspaceNotificationDestinationType = {
    Sns: "SNS",
} as const;

/**
 * These enums represent valid AWS notification destinations that the Grafana workspace has permission to use. These notification destinations are primarily used to help customers visualize which destinations have been added to a service managed IAM role.
 */
export type WorkspaceNotificationDestinationType = (typeof WorkspaceNotificationDestinationType)[keyof typeof WorkspaceNotificationDestinationType];

export const WorkspacePermissionType = {
    CustomerManaged: "CUSTOMER_MANAGED",
    ServiceManaged: "SERVICE_MANAGED",
} as const;

/**
 * These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.
 */
export type WorkspacePermissionType = (typeof WorkspacePermissionType)[keyof typeof WorkspacePermissionType];

export const WorkspaceSamlConfigurationStatus = {
    Configured: "CONFIGURED",
    NotConfigured: "NOT_CONFIGURED",
} as const;

/**
 * Valid SAML configuration statuses.
 */
export type WorkspaceSamlConfigurationStatus = (typeof WorkspaceSamlConfigurationStatus)[keyof typeof WorkspaceSamlConfigurationStatus];

export const WorkspaceStatus = {
    Active: "ACTIVE",
    Creating: "CREATING",
    Deleting: "DELETING",
    Failed: "FAILED",
    Updating: "UPDATING",
    Upgrading: "UPGRADING",
    VersionUpdating: "VERSION_UPDATING",
    DeletionFailed: "DELETION_FAILED",
    CreationFailed: "CREATION_FAILED",
    UpdateFailed: "UPDATE_FAILED",
    UpgradeFailed: "UPGRADE_FAILED",
    LicenseRemovalFailed: "LICENSE_REMOVAL_FAILED",
    VersionUpdateFailed: "VERSION_UPDATE_FAILED",
} as const;

/**
 * These enums represent the status of a workspace.
 */
export type WorkspaceStatus = (typeof WorkspaceStatus)[keyof typeof WorkspaceStatus];
