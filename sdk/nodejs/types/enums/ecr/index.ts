// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const RegistryScanningConfigurationFilterType = {
    Wildcard: "WILDCARD",
} as const;

/**
 * The type associated with the filter.
 */
export type RegistryScanningConfigurationFilterType = (typeof RegistryScanningConfigurationFilterType)[keyof typeof RegistryScanningConfigurationFilterType];

export const RegistryScanningConfigurationScanFrequency = {
    ScanOnPush: "SCAN_ON_PUSH",
    ContinuousScan: "CONTINUOUS_SCAN",
} as const;

/**
 * The frequency that scans are performed.
 */
export type RegistryScanningConfigurationScanFrequency = (typeof RegistryScanningConfigurationScanFrequency)[keyof typeof RegistryScanningConfigurationScanFrequency];

export const RegistryScanningConfigurationScanType = {
    Basic: "BASIC",
    Enhanced: "ENHANCED",
} as const;

/**
 * The type of scanning configured for the registry.
 */
export type RegistryScanningConfigurationScanType = (typeof RegistryScanningConfigurationScanType)[keyof typeof RegistryScanningConfigurationScanType];

export const ReplicationConfigurationFilterType = {
    PrefixMatch: "PREFIX_MATCH",
} as const;

/**
 * Type of repository filter
 */
export type ReplicationConfigurationFilterType = (typeof ReplicationConfigurationFilterType)[keyof typeof ReplicationConfigurationFilterType];

export const RepositoryCreationTemplateAppliedForItem = {
    Replication: "REPLICATION",
    PullThroughCache: "PULL_THROUGH_CACHE",
} as const;

/**
 * Enumerable Strings representing the repository creation scenarios that the template will apply towards.
 */
export type RepositoryCreationTemplateAppliedForItem = (typeof RepositoryCreationTemplateAppliedForItem)[keyof typeof RepositoryCreationTemplateAppliedForItem];

export const RepositoryCreationTemplateEncryptionType = {
    Aes256: "AES256",
    Kms: "KMS",
    KmsDsse: "KMS_DSSE",
} as const;

/**
 * The encryption type to use.
 */
export type RepositoryCreationTemplateEncryptionType = (typeof RepositoryCreationTemplateEncryptionType)[keyof typeof RepositoryCreationTemplateEncryptionType];

export const RepositoryCreationTemplateImageTagMutability = {
    Mutable: "MUTABLE",
    Immutable: "IMMUTABLE",
} as const;

/**
 * The tag mutability setting for the repository. If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten. If IMMUTABLE is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
 */
export type RepositoryCreationTemplateImageTagMutability = (typeof RepositoryCreationTemplateImageTagMutability)[keyof typeof RepositoryCreationTemplateImageTagMutability];

export const RepositoryEncryptionType = {
    Aes256: "AES256",
    Kms: "KMS",
    KmsDsse: "KMS_DSSE",
} as const;

/**
 * The encryption type to use.
 */
export type RepositoryEncryptionType = (typeof RepositoryEncryptionType)[keyof typeof RepositoryEncryptionType];

export const RepositoryImageTagMutability = {
    Mutable: "MUTABLE",
    Immutable: "IMMUTABLE",
} as const;

/**
 * The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
 */
export type RepositoryImageTagMutability = (typeof RepositoryImageTagMutability)[keyof typeof RepositoryImageTagMutability];
