// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessGrantGranteeGranteeType = {
    Iam: "IAM",
    DirectoryUser: "DIRECTORY_USER",
    DirectoryGroup: "DIRECTORY_GROUP",
} as const;

/**
 * Configures the transfer acceleration state for an Amazon S3 bucket.
 */
export type AccessGrantGranteeGranteeType = (typeof AccessGrantGranteeGranteeType)[keyof typeof AccessGrantGranteeGranteeType];

export const AccessGrantPermission = {
    Read: "READ",
    Write: "WRITE",
    Readwrite: "READWRITE",
} as const;

/**
 * The level of access to be afforded to the grantee
 */
export type AccessGrantPermission = (typeof AccessGrantPermission)[keyof typeof AccessGrantPermission];

export const AccessGrantS3PrefixType = {
    Object: "Object",
} as const;

/**
 * The type of S3SubPrefix.
 */
export type AccessGrantS3PrefixType = (typeof AccessGrantS3PrefixType)[keyof typeof AccessGrantS3PrefixType];

export const AccessPointNetworkOrigin = {
    Internet: "Internet",
    Vpc: "VPC",
} as const;

/**
 * Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
 */
export type AccessPointNetworkOrigin = (typeof AccessPointNetworkOrigin)[keyof typeof AccessPointNetworkOrigin];

export const BucketAccelerateConfigurationAccelerationStatus = {
    Enabled: "Enabled",
    Suspended: "Suspended",
} as const;

/**
 * Specifies the transfer acceleration status of the bucket.
 */
export type BucketAccelerateConfigurationAccelerationStatus = (typeof BucketAccelerateConfigurationAccelerationStatus)[keyof typeof BucketAccelerateConfigurationAccelerationStatus];

export const BucketAccessControl = {
    AuthenticatedRead: "AuthenticatedRead",
    AwsExecRead: "AwsExecRead",
    BucketOwnerFullControl: "BucketOwnerFullControl",
    BucketOwnerRead: "BucketOwnerRead",
    LogDeliveryWrite: "LogDeliveryWrite",
    Private: "Private",
    PublicRead: "PublicRead",
    PublicReadWrite: "PublicReadWrite",
} as const;

/**
 * This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
 *   A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.
 *   S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.
 *   The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.
 */
export type BucketAccessControl = (typeof BucketAccessControl)[keyof typeof BucketAccessControl];

export const BucketCorsRuleAllowedMethodsItem = {
    Get: "GET",
    Put: "PUT",
    Head: "HEAD",
    Post: "POST",
    Delete: "DELETE",
} as const;

export type BucketCorsRuleAllowedMethodsItem = (typeof BucketCorsRuleAllowedMethodsItem)[keyof typeof BucketCorsRuleAllowedMethodsItem];

export const BucketDefaultRetentionMode = {
    Compliance: "COMPLIANCE",
    Governance: "GOVERNANCE",
} as const;

/**
 * The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
 */
export type BucketDefaultRetentionMode = (typeof BucketDefaultRetentionMode)[keyof typeof BucketDefaultRetentionMode];

export const BucketDeleteMarkerReplicationStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Indicates whether to replicate delete markers. Disabled by default.
 */
export type BucketDeleteMarkerReplicationStatus = (typeof BucketDeleteMarkerReplicationStatus)[keyof typeof BucketDeleteMarkerReplicationStatus];

export const BucketDestinationFormat = {
    Csv: "CSV",
    Orc: "ORC",
    Parquet: "Parquet",
} as const;

/**
 * Specifies the file format used when exporting data to Amazon S3.
 *  *Allowed values*: ``CSV`` | ``ORC`` | ``Parquet``
 */
export type BucketDestinationFormat = (typeof BucketDestinationFormat)[keyof typeof BucketDestinationFormat];

export const BucketIntelligentTieringConfigurationStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies the status of the configuration.
 */
export type BucketIntelligentTieringConfigurationStatus = (typeof BucketIntelligentTieringConfigurationStatus)[keyof typeof BucketIntelligentTieringConfigurationStatus];

export const BucketInventoryConfigurationIncludedObjectVersions = {
    All: "All",
    Current: "Current",
} as const;

/**
 * Object versions to include in the inventory list. If set to ``All``, the list includes all the object versions, which adds the version-related fields ``VersionId``, ``IsLatest``, and ``DeleteMarker`` to the list. If set to ``Current``, the list does not contain these version-related fields.
 */
export type BucketInventoryConfigurationIncludedObjectVersions = (typeof BucketInventoryConfigurationIncludedObjectVersions)[keyof typeof BucketInventoryConfigurationIncludedObjectVersions];

export const BucketInventoryConfigurationOptionalFieldsItem = {
    Size: "Size",
    LastModifiedDate: "LastModifiedDate",
    StorageClass: "StorageClass",
    ETag: "ETag",
    IsMultipartUploaded: "IsMultipartUploaded",
    ReplicationStatus: "ReplicationStatus",
    EncryptionStatus: "EncryptionStatus",
    ObjectLockRetainUntilDate: "ObjectLockRetainUntilDate",
    ObjectLockMode: "ObjectLockMode",
    ObjectLockLegalHoldStatus: "ObjectLockLegalHoldStatus",
    IntelligentTieringAccessTier: "IntelligentTieringAccessTier",
    BucketKeyStatus: "BucketKeyStatus",
    ChecksumAlgorithm: "ChecksumAlgorithm",
    ObjectAccessControlList: "ObjectAccessControlList",
    ObjectOwner: "ObjectOwner",
} as const;

export type BucketInventoryConfigurationOptionalFieldsItem = (typeof BucketInventoryConfigurationOptionalFieldsItem)[keyof typeof BucketInventoryConfigurationOptionalFieldsItem];

export const BucketInventoryConfigurationScheduleFrequency = {
    Daily: "Daily",
    Weekly: "Weekly",
} as const;

/**
 * Specifies the schedule for generating inventory results.
 */
export type BucketInventoryConfigurationScheduleFrequency = (typeof BucketInventoryConfigurationScheduleFrequency)[keyof typeof BucketInventoryConfigurationScheduleFrequency];

export const BucketInventoryTableConfigurationConfigurationState = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Specifies whether inventory table configuration is enabled or disabled.
 */
export type BucketInventoryTableConfigurationConfigurationState = (typeof BucketInventoryTableConfigurationConfigurationState)[keyof typeof BucketInventoryTableConfigurationConfigurationState];

export const BucketLifecycleConfigurationTransitionDefaultMinimumObjectSize = {
    VariesByStorageClass: "varies_by_storage_class",
    AllStorageClasses128k: "all_storage_classes_128K",
} as const;

/**
 * Indicates which default minimum object size behavior is applied to the lifecycle configuration.
 *   This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.
 *    +  ``all_storage_classes_128K`` - Objects smaller than 128 KB will not transition to any storage class by default.
 *   +  ``varies_by_storage_class`` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB. 
 *   
 *  To customize the minimum object size for any transition you can add a filter that specifies a custom ``ObjectSizeGreaterThan`` or ``ObjectSizeLessThan`` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.
 */
export type BucketLifecycleConfigurationTransitionDefaultMinimumObjectSize = (typeof BucketLifecycleConfigurationTransitionDefaultMinimumObjectSize)[keyof typeof BucketLifecycleConfigurationTransitionDefaultMinimumObjectSize];

export const BucketMetadataDestinationTableBucketType = {
    Aws: "aws",
    Customer: "customer",
} as const;

/**
 * The type of the table bucket.
 */
export type BucketMetadataDestinationTableBucketType = (typeof BucketMetadataDestinationTableBucketType)[keyof typeof BucketMetadataDestinationTableBucketType];

export const BucketMetadataTableEncryptionConfigurationSseAlgorithm = {
    Awskms: "aws:kms",
    Aes256: "AES256",
} as const;

/**
 * Specifies the server-side encryption algorithm to use for encrypting tables.
 */
export type BucketMetadataTableEncryptionConfigurationSseAlgorithm = (typeof BucketMetadataTableEncryptionConfigurationSseAlgorithm)[keyof typeof BucketMetadataTableEncryptionConfigurationSseAlgorithm];

export const BucketMetricsStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether the replication metrics are enabled.
 */
export type BucketMetricsStatus = (typeof BucketMetricsStatus)[keyof typeof BucketMetricsStatus];

export const BucketNoncurrentVersionTransitionStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    StandardIa: "STANDARD_IA",
} as const;

/**
 * The class of storage used to store the object.
 */
export type BucketNoncurrentVersionTransitionStorageClass = (typeof BucketNoncurrentVersionTransitionStorageClass)[keyof typeof BucketNoncurrentVersionTransitionStorageClass];

export const BucketOwnershipControlsRuleObjectOwnership = {
    ObjectWriter: "ObjectWriter",
    BucketOwnerPreferred: "BucketOwnerPreferred",
    BucketOwnerEnforced: "BucketOwnerEnforced",
} as const;

/**
 * Specifies an object ownership rule.
 */
export type BucketOwnershipControlsRuleObjectOwnership = (typeof BucketOwnershipControlsRuleObjectOwnership)[keyof typeof BucketOwnershipControlsRuleObjectOwnership];

export const BucketRecordExpirationExpiration = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Specifies whether record expiration is enabled or disabled.
 */
export type BucketRecordExpirationExpiration = (typeof BucketRecordExpirationExpiration)[keyof typeof BucketRecordExpirationExpiration];

export const BucketRedirectAllRequestsToProtocol = {
    Http: "http",
    Https: "https",
} as const;

/**
 * Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
 */
export type BucketRedirectAllRequestsToProtocol = (typeof BucketRedirectAllRequestsToProtocol)[keyof typeof BucketRedirectAllRequestsToProtocol];

export const BucketRedirectRuleProtocol = {
    Http: "http",
    Https: "https",
} as const;

/**
 * Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
 */
export type BucketRedirectRuleProtocol = (typeof BucketRedirectRuleProtocol)[keyof typeof BucketRedirectRuleProtocol];

export const BucketReplicaModificationsStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Specifies whether Amazon S3 replicates modifications on replicas.
 *  *Allowed values*: ``Enabled`` | ``Disabled``
 */
export type BucketReplicaModificationsStatus = (typeof BucketReplicaModificationsStatus)[keyof typeof BucketReplicaModificationsStatus];

export const BucketReplicationDestinationStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    ReducedRedundancy: "REDUCED_REDUNDANCY",
    Standard: "STANDARD",
    StandardIa: "STANDARD_IA",
} as const;

/**
 * The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By default, Amazon S3 uses the storage class of the source object to create the object replica. 
 *  For valid values, see the ``StorageClass`` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference*.
 */
export type BucketReplicationDestinationStorageClass = (typeof BucketReplicationDestinationStorageClass)[keyof typeof BucketReplicationDestinationStorageClass];

export const BucketReplicationRuleStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether the rule is enabled.
 */
export type BucketReplicationRuleStatus = (typeof BucketReplicationRuleStatus)[keyof typeof BucketReplicationRuleStatus];

export const BucketReplicationTimeStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether the replication time is enabled.
 */
export type BucketReplicationTimeStatus = (typeof BucketReplicationTimeStatus)[keyof typeof BucketReplicationTimeStatus];

export const BucketRuleStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * If ``Enabled``, the rule is currently being applied. If ``Disabled``, the rule is not currently being applied.
 */
export type BucketRuleStatus = (typeof BucketRuleStatus)[keyof typeof BucketRuleStatus];

export const BucketServerSideEncryptionByDefaultSseAlgorithm = {
    Awskms: "aws:kms",
    Aes256: "AES256",
    Awskmsdsse: "aws:kms:dsse",
} as const;

/**
 * Server-side encryption algorithm to use for the default encryption.
 *   For directory buckets, there are only two supported values for server-side encryption: ``AES256`` and ``aws:kms``.
 */
export type BucketServerSideEncryptionByDefaultSseAlgorithm = (typeof BucketServerSideEncryptionByDefaultSseAlgorithm)[keyof typeof BucketServerSideEncryptionByDefaultSseAlgorithm];

export const BucketSseKmsEncryptedObjectsStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
 */
export type BucketSseKmsEncryptedObjectsStatus = (typeof BucketSseKmsEncryptedObjectsStatus)[keyof typeof BucketSseKmsEncryptedObjectsStatus];

export const BucketTieringAccessTier = {
    ArchiveAccess: "ARCHIVE_ACCESS",
    DeepArchiveAccess: "DEEP_ARCHIVE_ACCESS",
} as const;

/**
 * S3 Intelligent-Tiering access tier. See [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access tiers in the S3 Intelligent-Tiering storage class.
 */
export type BucketTieringAccessTier = (typeof BucketTieringAccessTier)[keyof typeof BucketTieringAccessTier];

export const BucketTransitionStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    StandardIa: "STANDARD_IA",
} as const;

/**
 * The storage class to which you want the object to transition.
 */
export type BucketTransitionStorageClass = (typeof BucketTransitionStorageClass)[keyof typeof BucketTransitionStorageClass];

export const BucketVersioningConfigurationStatus = {
    Enabled: "Enabled",
    Suspended: "Suspended",
} as const;

/**
 * The versioning state of the bucket.
 */
export type BucketVersioningConfigurationStatus = (typeof BucketVersioningConfigurationStatus)[keyof typeof BucketVersioningConfigurationStatus];

export const MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic = {
    True: "true",
    False: "false",
} as const;

/**
 * Specifies whether the policy is public or not.
 */
export type MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic = (typeof MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic)[keyof typeof MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic];

export const StorageLensS3BucketDestinationFormat = {
    Csv: "CSV",
    Parquet: "Parquet",
} as const;

/**
 * Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.
 */
export type StorageLensS3BucketDestinationFormat = (typeof StorageLensS3BucketDestinationFormat)[keyof typeof StorageLensS3BucketDestinationFormat];

export const StorageLensS3BucketDestinationOutputSchemaVersion = {
    V1: "V_1",
} as const;

/**
 * The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.
 */
export type StorageLensS3BucketDestinationOutputSchemaVersion = (typeof StorageLensS3BucketDestinationOutputSchemaVersion)[keyof typeof StorageLensS3BucketDestinationOutputSchemaVersion];
