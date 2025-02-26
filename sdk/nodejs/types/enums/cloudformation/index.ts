// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const GuardHookAction = {
    Create: "CREATE",
    Update: "UPDATE",
    Delete: "DELETE",
} as const;

/**
 * Target actions are the type of operation hooks will be executed at.
 */
export type GuardHookAction = (typeof GuardHookAction)[keyof typeof GuardHookAction];

export const GuardHookFailureMode = {
    Fail: "FAIL",
    Warn: "WARN",
} as const;

/**
 * Attribute to specify CloudFormation behavior on hook failure.
 */
export type GuardHookFailureMode = (typeof GuardHookFailureMode)[keyof typeof GuardHookFailureMode];

export const GuardHookHookStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Attribute to specify which stacks this hook applies to or should get invoked for
 */
export type GuardHookHookStatus = (typeof GuardHookHookStatus)[keyof typeof GuardHookHookStatus];

export const GuardHookInvocationPoint = {
    PreProvision: "PRE_PROVISION",
} as const;

/**
 * Invocation points are the point in provisioning workflow where hooks will be executed.
 */
export type GuardHookInvocationPoint = (typeof GuardHookInvocationPoint)[keyof typeof GuardHookInvocationPoint];

export const GuardHookStackFiltersPropertiesFilteringCriteria = {
    All: "ALL",
    Any: "ANY",
} as const;

/**
 * Attribute to specify the filtering behavior. ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match
 */
export type GuardHookStackFiltersPropertiesFilteringCriteria = (typeof GuardHookStackFiltersPropertiesFilteringCriteria)[keyof typeof GuardHookStackFiltersPropertiesFilteringCriteria];

export const GuardHookTargetOperation = {
    Resource: "RESOURCE",
    Stack: "STACK",
    ChangeSet: "CHANGE_SET",
    CloudControl: "CLOUD_CONTROL",
} as const;

/**
 * Which operations should this Hook run against? Resource changes, stacks or change sets.
 */
export type GuardHookTargetOperation = (typeof GuardHookTargetOperation)[keyof typeof GuardHookTargetOperation];

export const HookTypeConfigConfigurationAlias = {
    Default: "default",
} as const;

/**
 * An alias by which to refer to this extension configuration data.
 */
export type HookTypeConfigConfigurationAlias = (typeof HookTypeConfigConfigurationAlias)[keyof typeof HookTypeConfigConfigurationAlias];

export const HookVersionVisibility = {
    Public: "PUBLIC",
    Private: "PRIVATE",
} as const;

/**
 * The scope at which the type is visible and usable in CloudFormation operations.
 *
 * Valid values include:
 *
 * PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
 *
 * PUBLIC: The type is publically visible and usable within any Amazon account.
 */
export type HookVersionVisibility = (typeof HookVersionVisibility)[keyof typeof HookVersionVisibility];

export const LambdaHookAction = {
    Create: "CREATE",
    Update: "UPDATE",
    Delete: "DELETE",
} as const;

/**
 * Target actions are the type of operation hooks will be executed at.
 */
export type LambdaHookAction = (typeof LambdaHookAction)[keyof typeof LambdaHookAction];

export const LambdaHookFailureMode = {
    Fail: "FAIL",
    Warn: "WARN",
} as const;

/**
 * Attribute to specify CloudFormation behavior on hook failure.
 */
export type LambdaHookFailureMode = (typeof LambdaHookFailureMode)[keyof typeof LambdaHookFailureMode];

export const LambdaHookHookStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Attribute to specify which stacks this hook applies to or should get invoked for
 */
export type LambdaHookHookStatus = (typeof LambdaHookHookStatus)[keyof typeof LambdaHookHookStatus];

export const LambdaHookInvocationPoint = {
    PreProvision: "PRE_PROVISION",
} as const;

/**
 * Invocation points are the point in provisioning workflow where hooks will be executed.
 */
export type LambdaHookInvocationPoint = (typeof LambdaHookInvocationPoint)[keyof typeof LambdaHookInvocationPoint];

export const LambdaHookStackFiltersPropertiesFilteringCriteria = {
    All: "ALL",
    Any: "ANY",
} as const;

/**
 * Attribute to specify the filtering behavior. ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match
 */
export type LambdaHookStackFiltersPropertiesFilteringCriteria = (typeof LambdaHookStackFiltersPropertiesFilteringCriteria)[keyof typeof LambdaHookStackFiltersPropertiesFilteringCriteria];

export const LambdaHookTargetOperation = {
    Resource: "RESOURCE",
    Stack: "STACK",
    ChangeSet: "CHANGE_SET",
    CloudControl: "CLOUD_CONTROL",
} as const;

/**
 * Which operations should this Hook run against? Resource changes, stacks or change sets.
 */
export type LambdaHookTargetOperation = (typeof LambdaHookTargetOperation)[keyof typeof LambdaHookTargetOperation];

export const ModuleVersionVisibility = {
    Private: "PRIVATE",
} as const;

/**
 * The scope at which the type is visible and usable in CloudFormation operations.
 *
 * The only allowed value at present is:
 *
 * PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
 */
export type ModuleVersionVisibility = (typeof ModuleVersionVisibility)[keyof typeof ModuleVersionVisibility];

export const PublicTypeVersionType = {
    Resource: "RESOURCE",
    Module: "MODULE",
    Hook: "HOOK",
} as const;

/**
 * The kind of extension
 */
export type PublicTypeVersionType = (typeof PublicTypeVersionType)[keyof typeof PublicTypeVersionType];

export const PublisherIdentityProvider = {
    AwsMarketplace: "AWS_Marketplace",
    GitHub: "GitHub",
    Bitbucket: "Bitbucket",
} as const;

/**
 * The type of account used as the identity provider when registering this publisher with CloudFormation.
 */
export type PublisherIdentityProvider = (typeof PublisherIdentityProvider)[keyof typeof PublisherIdentityProvider];

export const PublisherStatus = {
    Verified: "VERIFIED",
    Unverified: "UNVERIFIED",
} as const;

/**
 * Whether the publisher is verified.
 */
export type PublisherStatus = (typeof PublisherStatus)[keyof typeof PublisherStatus];

export const ResourceVersionProvisioningType = {
    NonProvisionable: "NON_PROVISIONABLE",
    Immutable: "IMMUTABLE",
    FullyMutable: "FULLY_MUTABLE",
} as const;

/**
 * The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.
 */
export type ResourceVersionProvisioningType = (typeof ResourceVersionProvisioningType)[keyof typeof ResourceVersionProvisioningType];

export const ResourceVersionVisibility = {
    Public: "PUBLIC",
    Private: "PRIVATE",
} as const;

/**
 * The scope at which the type is visible and usable in CloudFormation operations.
 *
 * Valid values include:
 *
 * PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
 *
 * PUBLIC: The type is publically visible and usable within any Amazon account.
 */
export type ResourceVersionVisibility = (typeof ResourceVersionVisibility)[keyof typeof ResourceVersionVisibility];

export const StackCapabilitiesItem = {
    CapabilityIam: "CAPABILITY_IAM",
    CapabilityNamedIam: "CAPABILITY_NAMED_IAM",
    CapabilityAutoExpand: "CAPABILITY_AUTO_EXPAND",
} as const;

export type StackCapabilitiesItem = (typeof StackCapabilitiesItem)[keyof typeof StackCapabilitiesItem];

export const StackSetCallAs = {
    Self: "SELF",
    DelegatedAdmin: "DELEGATED_ADMIN",
} as const;

/**
 * Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.
 */
export type StackSetCallAs = (typeof StackSetCallAs)[keyof typeof StackSetCallAs];

export const StackSetCapability = {
    CapabilityIam: "CAPABILITY_IAM",
    CapabilityNamedIam: "CAPABILITY_NAMED_IAM",
    CapabilityAutoExpand: "CAPABILITY_AUTO_EXPAND",
} as const;

export type StackSetCapability = (typeof StackSetCapability)[keyof typeof StackSetCapability];

export const StackSetConcurrencyMode = {
    StrictFailureTolerance: "STRICT_FAILURE_TOLERANCE",
    SoftFailureTolerance: "SOFT_FAILURE_TOLERANCE",
} as const;

/**
 * Specifies how the concurrency level behaves during the operation execution.
 */
export type StackSetConcurrencyMode = (typeof StackSetConcurrencyMode)[keyof typeof StackSetConcurrencyMode];

export const StackSetDeploymentTargetsAccountFilterType = {
    None: "NONE",
    Union: "UNION",
    Intersection: "INTERSECTION",
    Difference: "DIFFERENCE",
} as const;

/**
 * The filter type you want to apply on organizational units and accounts.
 */
export type StackSetDeploymentTargetsAccountFilterType = (typeof StackSetDeploymentTargetsAccountFilterType)[keyof typeof StackSetDeploymentTargetsAccountFilterType];

export const StackSetPermissionModel = {
    ServiceManaged: "SERVICE_MANAGED",
    SelfManaged: "SELF_MANAGED",
} as const;

/**
 * Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.
 */
export type StackSetPermissionModel = (typeof StackSetPermissionModel)[keyof typeof StackSetPermissionModel];

export const StackSetRegionConcurrencyType = {
    Sequential: "SEQUENTIAL",
    Parallel: "PARALLEL",
} as const;

/**
 * The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time
 */
export type StackSetRegionConcurrencyType = (typeof StackSetRegionConcurrencyType)[keyof typeof StackSetRegionConcurrencyType];

export const StackStatus = {
    CreateInProgress: "CREATE_IN_PROGRESS",
    CreateFailed: "CREATE_FAILED",
    CreateComplete: "CREATE_COMPLETE",
    RollbackInProgress: "ROLLBACK_IN_PROGRESS",
    RollbackFailed: "ROLLBACK_FAILED",
    RollbackComplete: "ROLLBACK_COMPLETE",
    DeleteInProgress: "DELETE_IN_PROGRESS",
    DeleteFailed: "DELETE_FAILED",
    DeleteComplete: "DELETE_COMPLETE",
    UpdateInProgress: "UPDATE_IN_PROGRESS",
    UpdateCompleteCleanupInProgress: "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS",
    UpdateComplete: "UPDATE_COMPLETE",
    UpdateFailed: "UPDATE_FAILED",
    UpdateRollbackInProgress: "UPDATE_ROLLBACK_IN_PROGRESS",
    UpdateRollbackFailed: "UPDATE_ROLLBACK_FAILED",
    UpdateRollbackCompleteCleanupInProgress: "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS",
    UpdateRollbackComplete: "UPDATE_ROLLBACK_COMPLETE",
    ReviewInProgress: "REVIEW_IN_PROGRESS",
    ImportInProgress: "IMPORT_IN_PROGRESS",
    ImportComplete: "IMPORT_COMPLETE",
    ImportRollbackInProgress: "IMPORT_ROLLBACK_IN_PROGRESS",
    ImportRollbackFailed: "IMPORT_ROLLBACK_FAILED",
    ImportRollbackComplete: "IMPORT_ROLLBACK_COMPLETE",
} as const;

/**
 * Current status of the stack.
 */
export type StackStatus = (typeof StackStatus)[keyof typeof StackStatus];

export const TypeActivationType = {
    Resource: "RESOURCE",
    Module: "MODULE",
    Hook: "HOOK",
} as const;

/**
 * The kind of extension
 */
export type TypeActivationType = (typeof TypeActivationType)[keyof typeof TypeActivationType];

export const TypeActivationVersionBump = {
    Major: "MAJOR",
    Minor: "MINOR",
} as const;

/**
 * Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
 */
export type TypeActivationVersionBump = (typeof TypeActivationVersionBump)[keyof typeof TypeActivationVersionBump];
