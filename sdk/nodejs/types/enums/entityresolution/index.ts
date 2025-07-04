// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const IdMappingWorkflowIdMappingRuleBasedPropertiesAttributeMatchingModel = {
    OneToOne: "ONE_TO_ONE",
    ManyToMany: "MANY_TO_MANY",
} as const;

/**
 * The comparison type. You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
 *
 * If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A matches the value of the `BusinessEmail` field of Profile B, the two profiles are matched on the `Email` attribute type.
 *
 * If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
 */
export type IdMappingWorkflowIdMappingRuleBasedPropertiesAttributeMatchingModel = (typeof IdMappingWorkflowIdMappingRuleBasedPropertiesAttributeMatchingModel)[keyof typeof IdMappingWorkflowIdMappingRuleBasedPropertiesAttributeMatchingModel];

export const IdMappingWorkflowIdMappingRuleBasedPropertiesRecordMatchingModel = {
    OneSourceToOneTarget: "ONE_SOURCE_TO_ONE_TARGET",
    ManySourceToOneTarget: "MANY_SOURCE_TO_ONE_TARGET",
} as const;

/**
 * The type of matching record that is allowed to be used in an ID mapping workflow.
 *
 * If the value is set to `ONE_SOURCE_TO_ONE_TARGET` , only one record in the source can be matched to the same record in the target.
 *
 * If the value is set to `MANY_SOURCE_TO_ONE_TARGET` , multiple records in the source can be matched to one record in the target.
 */
export type IdMappingWorkflowIdMappingRuleBasedPropertiesRecordMatchingModel = (typeof IdMappingWorkflowIdMappingRuleBasedPropertiesRecordMatchingModel)[keyof typeof IdMappingWorkflowIdMappingRuleBasedPropertiesRecordMatchingModel];

export const IdMappingWorkflowIdMappingRuleBasedPropertiesRuleDefinitionType = {
    Source: "SOURCE",
    Target: "TARGET",
} as const;

/**
 * The set of rules you can use in an ID mapping workflow. The limitations specified for the source or target to define the match rules must be compatible.
 */
export type IdMappingWorkflowIdMappingRuleBasedPropertiesRuleDefinitionType = (typeof IdMappingWorkflowIdMappingRuleBasedPropertiesRuleDefinitionType)[keyof typeof IdMappingWorkflowIdMappingRuleBasedPropertiesRuleDefinitionType];

export const IdMappingWorkflowIdMappingTechniquesIdMappingType = {
    Provider: "PROVIDER",
    RuleBased: "RULE_BASED",
} as const;

/**
 * The type of ID mapping.
 */
export type IdMappingWorkflowIdMappingTechniquesIdMappingType = (typeof IdMappingWorkflowIdMappingTechniquesIdMappingType)[keyof typeof IdMappingWorkflowIdMappingTechniquesIdMappingType];

export const IdMappingWorkflowInputSourceType = {
    Source: "SOURCE",
    Target: "TARGET",
} as const;

/**
 * The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
 *
 * The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
 *
 * The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
 */
export type IdMappingWorkflowInputSourceType = (typeof IdMappingWorkflowInputSourceType)[keyof typeof IdMappingWorkflowInputSourceType];

export const IdNamespaceIdMappingWorkflowPropertiesIdMappingType = {
    Provider: "PROVIDER",
    RuleBased: "RULE_BASED",
} as const;

/**
 * The type of ID mapping.
 */
export type IdNamespaceIdMappingWorkflowPropertiesIdMappingType = (typeof IdNamespaceIdMappingWorkflowPropertiesIdMappingType)[keyof typeof IdNamespaceIdMappingWorkflowPropertiesIdMappingType];

export const IdNamespaceNamespaceRuleBasedPropertiesAttributeMatchingModel = {
    OneToOne: "ONE_TO_ONE",
    ManyToMany: "MANY_TO_MANY",
} as const;

/**
 * The comparison type. You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
 *
 * If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A matches the value of `BusinessEmail` field of Profile B, the two profiles are matched on the `Email` attribute type.
 *
 * If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
 */
export type IdNamespaceNamespaceRuleBasedPropertiesAttributeMatchingModel = (typeof IdNamespaceNamespaceRuleBasedPropertiesAttributeMatchingModel)[keyof typeof IdNamespaceNamespaceRuleBasedPropertiesAttributeMatchingModel];

export const IdNamespaceRecordMatchingModel = {
    OneSourceToOneTarget: "ONE_SOURCE_TO_ONE_TARGET",
    ManySourceToOneTarget: "MANY_SOURCE_TO_ONE_TARGET",
} as const;

export type IdNamespaceRecordMatchingModel = (typeof IdNamespaceRecordMatchingModel)[keyof typeof IdNamespaceRecordMatchingModel];

export const IdNamespaceRuleDefinitionType = {
    Source: "SOURCE",
    Target: "TARGET",
} as const;

export type IdNamespaceRuleDefinitionType = (typeof IdNamespaceRuleDefinitionType)[keyof typeof IdNamespaceRuleDefinitionType];

export const IdNamespaceType = {
    Source: "SOURCE",
    Target: "TARGET",
} as const;

/**
 * The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
 *
 * The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
 *
 * The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
 */
export type IdNamespaceType = (typeof IdNamespaceType)[keyof typeof IdNamespaceType];

export const MatchingWorkflowIncrementalRunConfigIncrementalRunType = {
    Immediate: "IMMEDIATE",
} as const;

/**
 * The type of incremental run. The only valid value is `IMMEDIATE` . This appears as "Automatic" in the console.
 *
 * > For workflows where `resolutionType` is `ML_MATCHING` , incremental processing is not supported.
 */
export type MatchingWorkflowIncrementalRunConfigIncrementalRunType = (typeof MatchingWorkflowIncrementalRunConfigIncrementalRunType)[keyof typeof MatchingWorkflowIncrementalRunConfigIncrementalRunType];

export const MatchingWorkflowResolutionType = {
    RuleMatching: "RULE_MATCHING",
    MlMatching: "ML_MATCHING",
    Provider: "PROVIDER",
} as const;

export type MatchingWorkflowResolutionType = (typeof MatchingWorkflowResolutionType)[keyof typeof MatchingWorkflowResolutionType];

export const MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel = {
    OneToOne: "ONE_TO_ONE",
    ManyToMany: "MANY_TO_MANY",
} as const;

/**
 * The comparison type. You can either choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel` .
 *
 * If you choose `MANY_TO_MANY` , the system can match attributes across the sub-types of an attribute type. For example, if the value of the `Email` field of Profile A and the value of `BusinessEmail` field of Profile B matches, the two profiles are matched on the `Email` attribute type.
 *
 * If you choose `ONE_TO_ONE` , the system can only match attributes if the sub-types are an exact match. For example, for the `Email` attribute type, the system will only consider it a match if the value of the `Email` field of Profile A matches the value of the `Email` field of Profile B.
 */
export type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel = (typeof MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel)[keyof typeof MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel];

export const MatchingWorkflowRuleBasedPropertiesMatchPurpose = {
    IdentifierGeneration: "IDENTIFIER_GENERATION",
    Indexing: "INDEXING",
} as const;

/**
 * An indicator of whether to generate IDs and index the data or not.
 *
 * If you choose `IDENTIFIER_GENERATION` , the process generates IDs and indexes the data.
 *
 * If you choose `INDEXING` , the process indexes the data without generating IDs.
 */
export type MatchingWorkflowRuleBasedPropertiesMatchPurpose = (typeof MatchingWorkflowRuleBasedPropertiesMatchPurpose)[keyof typeof MatchingWorkflowRuleBasedPropertiesMatchPurpose];

export const PolicyStatementStatementEffect = {
    Allow: "Allow",
    Deny: "Deny",
} as const;

export type PolicyStatementStatementEffect = (typeof PolicyStatementStatementEffect)[keyof typeof PolicyStatementStatementEffect];

export const SchemaMappingSchemaAttributeType = {
    Name: "NAME",
    NameFirst: "NAME_FIRST",
    NameMiddle: "NAME_MIDDLE",
    NameLast: "NAME_LAST",
    Address: "ADDRESS",
    AddressStreet1: "ADDRESS_STREET1",
    AddressStreet2: "ADDRESS_STREET2",
    AddressStreet3: "ADDRESS_STREET3",
    AddressCity: "ADDRESS_CITY",
    AddressState: "ADDRESS_STATE",
    AddressCountry: "ADDRESS_COUNTRY",
    AddressPostalcode: "ADDRESS_POSTALCODE",
    Phone: "PHONE",
    PhoneNumber: "PHONE_NUMBER",
    PhoneCountrycode: "PHONE_COUNTRYCODE",
    EmailAddress: "EMAIL_ADDRESS",
    UniqueId: "UNIQUE_ID",
    Date: "DATE",
    String: "STRING",
    ProviderId: "PROVIDER_ID",
} as const;

export type SchemaMappingSchemaAttributeType = (typeof SchemaMappingSchemaAttributeType)[keyof typeof SchemaMappingSchemaAttributeType];
