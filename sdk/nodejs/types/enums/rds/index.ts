// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CustomDbEngineVersionStatus = {
    Available: "available",
    Inactive: "inactive",
    InactiveExceptRestore: "inactive-except-restore",
} as const;

/**
 * A value that indicates the status of a custom engine version (CEV).
 */
export type CustomDbEngineVersionStatus = (typeof CustomDbEngineVersionStatus)[keyof typeof CustomDbEngineVersionStatus];

export const DbInstanceProcessorFeatureName = {
    CoreCount: "coreCount",
    ThreadsPerCore: "threadsPerCore",
} as const;

/**
 * The name of the processor feature. Valid names are ``coreCount`` and ``threadsPerCore``.
 */
export type DbInstanceProcessorFeatureName = (typeof DbInstanceProcessorFeatureName)[keyof typeof DbInstanceProcessorFeatureName];

export const DbProxyAuthFormatAuthScheme = {
    Secrets: "SECRETS",
} as const;

/**
 * The type of authentication that the proxy uses for connections from the proxy to the underlying database. 
 */
export type DbProxyAuthFormatAuthScheme = (typeof DbProxyAuthFormatAuthScheme)[keyof typeof DbProxyAuthFormatAuthScheme];

export const DbProxyAuthFormatClientPasswordAuthType = {
    MysqlNativePassword: "MYSQL_NATIVE_PASSWORD",
    PostgresScramSha256: "POSTGRES_SCRAM_SHA_256",
    PostgresMd5: "POSTGRES_MD5",
    SqlServerAuthentication: "SQL_SERVER_AUTHENTICATION",
} as const;

/**
 * The type of authentication the proxy uses for connections from clients.
 */
export type DbProxyAuthFormatClientPasswordAuthType = (typeof DbProxyAuthFormatClientPasswordAuthType)[keyof typeof DbProxyAuthFormatClientPasswordAuthType];

export const DbProxyAuthFormatIamAuth = {
    Disabled: "DISABLED",
    Required: "REQUIRED",
    Enabled: "ENABLED",
} as const;

/**
 * Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.
 */
export type DbProxyAuthFormatIamAuth = (typeof DbProxyAuthFormatIamAuth)[keyof typeof DbProxyAuthFormatIamAuth];

export const DbProxyEndpointTargetRole = {
    ReadWrite: "READ_WRITE",
    ReadOnly: "READ_ONLY",
} as const;

/**
 * A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
 */
export type DbProxyEndpointTargetRole = (typeof DbProxyEndpointTargetRole)[keyof typeof DbProxyEndpointTargetRole];

export const DbProxyEngineFamily = {
    Mysql: "MYSQL",
    Postgresql: "POSTGRESQL",
    Sqlserver: "SQLSERVER",
} as const;

/**
 * The kinds of databases that the proxy can connect to.
 */
export type DbProxyEngineFamily = (typeof DbProxyEngineFamily)[keyof typeof DbProxyEngineFamily];

export const DbProxyTargetGroupTargetGroupName = {
    Default: "default",
} as const;

/**
 * The identifier for the DBProxyTargetGroup
 */
export type DbProxyTargetGroupTargetGroupName = (typeof DbProxyTargetGroupTargetGroupName)[keyof typeof DbProxyTargetGroupTargetGroupName];

export const GlobalClusterEngine = {
    Aurora: "aurora",
    AuroraMysql: "aurora-mysql",
    AuroraPostgresql: "aurora-postgresql",
} as const;

/**
 * The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).
 * If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
 */
export type GlobalClusterEngine = (typeof GlobalClusterEngine)[keyof typeof GlobalClusterEngine];
