// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const LaunchProfileStreamingClipboardMode = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

export type LaunchProfileStreamingClipboardMode = (typeof LaunchProfileStreamingClipboardMode)[keyof typeof LaunchProfileStreamingClipboardMode];

export const LaunchProfileStreamingInstanceType = {
    G4dnXlarge: "g4dn.xlarge",
    G4dn2xlarge: "g4dn.2xlarge",
    G4dn4xlarge: "g4dn.4xlarge",
    G4dn8xlarge: "g4dn.8xlarge",
    G4dn12xlarge: "g4dn.12xlarge",
    G4dn16xlarge: "g4dn.16xlarge",
    G34xlarge: "g3.4xlarge",
    G3sXlarge: "g3s.xlarge",
    G5Xlarge: "g5.xlarge",
    G52xlarge: "g5.2xlarge",
    G54xlarge: "g5.4xlarge",
    G58xlarge: "g5.8xlarge",
    G516xlarge: "g5.16xlarge",
} as const;

export type LaunchProfileStreamingInstanceType = (typeof LaunchProfileStreamingInstanceType)[keyof typeof LaunchProfileStreamingInstanceType];

export const LaunchProfileStreamingSessionStorageMode = {
    Upload: "UPLOAD",
} as const;

export type LaunchProfileStreamingSessionStorageMode = (typeof LaunchProfileStreamingSessionStorageMode)[keyof typeof LaunchProfileStreamingSessionStorageMode];

export const StreamingImageEncryptionConfigurationKeyType = {
    CustomerManagedKey: "CUSTOMER_MANAGED_KEY",
} as const;

/**
 * <p/>
 */
export type StreamingImageEncryptionConfigurationKeyType = (typeof StreamingImageEncryptionConfigurationKeyType)[keyof typeof StreamingImageEncryptionConfigurationKeyType];

export const StudioComponentInitializationScriptRunContext = {
    SystemInitialization: "SYSTEM_INITIALIZATION",
    UserInitialization: "USER_INITIALIZATION",
} as const;

export type StudioComponentInitializationScriptRunContext = (typeof StudioComponentInitializationScriptRunContext)[keyof typeof StudioComponentInitializationScriptRunContext];

export const StudioComponentLaunchProfilePlatform = {
    Linux: "LINUX",
    Windows: "WINDOWS",
} as const;

export type StudioComponentLaunchProfilePlatform = (typeof StudioComponentLaunchProfilePlatform)[keyof typeof StudioComponentLaunchProfilePlatform];

export const StudioComponentSubtype = {
    AwsManagedMicrosoftAd: "AWS_MANAGED_MICROSOFT_AD",
    AmazonFsxForWindows: "AMAZON_FSX_FOR_WINDOWS",
    AmazonFsxForLustre: "AMAZON_FSX_FOR_LUSTRE",
    Custom: "CUSTOM",
} as const;

export type StudioComponentSubtype = (typeof StudioComponentSubtype)[keyof typeof StudioComponentSubtype];

export const StudioComponentType = {
    ActiveDirectory: "ACTIVE_DIRECTORY",
    SharedFileSystem: "SHARED_FILE_SYSTEM",
    ComputeFarm: "COMPUTE_FARM",
    LicenseService: "LICENSE_SERVICE",
    Custom: "CUSTOM",
} as const;

export type StudioComponentType = (typeof StudioComponentType)[keyof typeof StudioComponentType];

export const StudioEncryptionConfigurationKeyType = {
    AwsOwnedKey: "AWS_OWNED_KEY",
    CustomerManagedKey: "CUSTOMER_MANAGED_KEY",
} as const;

/**
 * <p>The type of KMS key that is used to encrypt studio data.</p>
 */
export type StudioEncryptionConfigurationKeyType = (typeof StudioEncryptionConfigurationKeyType)[keyof typeof StudioEncryptionConfigurationKeyType];