// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const IpSetIpAddressVersion = {
    Ipv4: "IPV4",
    Ipv6: "IPV6",
} as const;

/**
 * Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.
 */
export type IpSetIpAddressVersion = (typeof IpSetIpAddressVersion)[keyof typeof IpSetIpAddressVersion];

export const IpSetScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type IpSetScope = (typeof IpSetScope)[keyof typeof IpSetScope];

export const LoggingConfigurationConditionActionConditionPropertiesAction = {
    Allow: "ALLOW",
    Block: "BLOCK",
    Count: "COUNT",
    Captcha: "CAPTCHA",
    Challenge: "CHALLENGE",
    ExcludedAsCount: "EXCLUDED_AS_COUNT",
} as const;

/**
 * Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
 */
export type LoggingConfigurationConditionActionConditionPropertiesAction = (typeof LoggingConfigurationConditionActionConditionPropertiesAction)[keyof typeof LoggingConfigurationConditionActionConditionPropertiesAction];

export const LoggingConfigurationFilterBehavior = {
    Keep: "KEEP",
    Drop: "DROP",
} as const;

/**
 * How to handle logs that satisfy the filter's conditions and requirement. 
 */
export type LoggingConfigurationFilterBehavior = (typeof LoggingConfigurationFilterBehavior)[keyof typeof LoggingConfigurationFilterBehavior];

export const LoggingConfigurationFilterRequirement = {
    MeetsAll: "MEETS_ALL",
    MeetsAny: "MEETS_ANY",
} as const;

/**
 * Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
 */
export type LoggingConfigurationFilterRequirement = (typeof LoggingConfigurationFilterRequirement)[keyof typeof LoggingConfigurationFilterRequirement];

export const LoggingConfigurationLoggingFilterPropertiesDefaultBehavior = {
    Keep: "KEEP",
    Drop: "DROP",
} as const;

/**
 * Default handling for logs that don't match any of the specified filtering conditions.
 */
export type LoggingConfigurationLoggingFilterPropertiesDefaultBehavior = (typeof LoggingConfigurationLoggingFilterPropertiesDefaultBehavior)[keyof typeof LoggingConfigurationLoggingFilterPropertiesDefaultBehavior];

export const RegexPatternSetScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type RegexPatternSetScope = (typeof RegexPatternSetScope)[keyof typeof RegexPatternSetScope];

export const RuleGroupBodyParsingFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
    EvaluateAsString: "EVALUATE_AS_STRING",
} as const;

/**
 * The inspection behavior to fall back to if the JSON in the request body is invalid.
 */
export type RuleGroupBodyParsingFallbackBehavior = (typeof RuleGroupBodyParsingFallbackBehavior)[keyof typeof RuleGroupBodyParsingFallbackBehavior];

export const RuleGroupForwardedIpConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
 *
 * > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all. 
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupForwardedIpConfigurationFallbackBehavior = (typeof RuleGroupForwardedIpConfigurationFallbackBehavior)[keyof typeof RuleGroupForwardedIpConfigurationFallbackBehavior];

export const RuleGroupIpSetForwardedIpConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
 *
 * > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all. 
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupIpSetForwardedIpConfigurationFallbackBehavior = (typeof RuleGroupIpSetForwardedIpConfigurationFallbackBehavior)[keyof typeof RuleGroupIpSetForwardedIpConfigurationFallbackBehavior];

export const RuleGroupIpSetForwardedIpConfigurationPosition = {
    First: "FIRST",
    Last: "LAST",
    Any: "ANY",
} as const;

/**
 * The position in the header to search for the IP address. The header can contain IP addresses of the original client and also of proxies. For example, the header value could be `10.1.1.1, 127.0.0.0, 10.10.10.10` where the first IP address identifies the original client and the rest identify proxies that the request went through.
 *
 * The options for this setting are the following:
 *
 * - FIRST - Inspect the first IP address in the list of IP addresses in the header. This is usually the client's original IP.
 * - LAST - Inspect the last IP address in the list of IP addresses in the header.
 * - ANY - Inspect all IP addresses in the header for a match. If the header contains more than 10 IP addresses, AWS WAF inspects the last 10.
 */
export type RuleGroupIpSetForwardedIpConfigurationPosition = (typeof RuleGroupIpSetForwardedIpConfigurationPosition)[keyof typeof RuleGroupIpSetForwardedIpConfigurationPosition];

export const RuleGroupJa3FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a JA3 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupJa3FingerprintFallbackBehavior = (typeof RuleGroupJa3FingerprintFallbackBehavior)[keyof typeof RuleGroupJa3FingerprintFallbackBehavior];

export const RuleGroupJa4FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a JA4 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupJa4FingerprintFallbackBehavior = (typeof RuleGroupJa4FingerprintFallbackBehavior)[keyof typeof RuleGroupJa4FingerprintFallbackBehavior];

export const RuleGroupJsonMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the JSON to match against using the MatchPattern.
 */
export type RuleGroupJsonMatchScope = (typeof RuleGroupJsonMatchScope)[keyof typeof RuleGroupJsonMatchScope];

export const RuleGroupLabelMatchScope = {
    Label: "LABEL",
    Namespace: "NAMESPACE",
} as const;

export type RuleGroupLabelMatchScope = (typeof RuleGroupLabelMatchScope)[keyof typeof RuleGroupLabelMatchScope];

export const RuleGroupMapMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the request to match against using the MatchPattern.
 */
export type RuleGroupMapMatchScope = (typeof RuleGroupMapMatchScope)[keyof typeof RuleGroupMapMatchScope];

export const RuleGroupOversizeHandling = {
    Continue: "CONTINUE",
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * Handling of requests containing oversize fields
 */
export type RuleGroupOversizeHandling = (typeof RuleGroupOversizeHandling)[keyof typeof RuleGroupOversizeHandling];

export const RuleGroupPositionalConstraint = {
    Exactly: "EXACTLY",
    StartsWith: "STARTS_WITH",
    EndsWith: "ENDS_WITH",
    Contains: "CONTAINS",
    ContainsWord: "CONTAINS_WORD",
} as const;

/**
 * Position of the evaluation in the FieldToMatch of request.
 */
export type RuleGroupPositionalConstraint = (typeof RuleGroupPositionalConstraint)[keyof typeof RuleGroupPositionalConstraint];

export const RuleGroupRateBasedStatementAggregateKeyType = {
    Ip: "IP",
    ForwardedIp: "FORWARDED_IP",
    Constant: "CONSTANT",
    CustomKeys: "CUSTOM_KEYS",
} as const;

/**
 * Setting that indicates how to aggregate the request counts.
 *
 * > Web requests that are missing any of the components specified in the aggregation keys are omitted from the rate-based rule evaluation and handling. 
 *
 * - `CONSTANT` - Count and limit the requests that match the rate-based rule's scope-down statement. With this option, the counted requests aren't further aggregated. The scope-down statement is the only specification used. When the count of all requests that satisfy the scope-down statement goes over the limit, AWS WAF applies the rule action to all requests that satisfy the scope-down statement.
 *
 * With this option, you must configure the `ScopeDownStatement` property.
 * - `CUSTOM_KEYS` - Aggregate the request counts using one or more web request components as the aggregate keys.
 *
 * With this option, you must specify the aggregate keys in the `CustomKeys` property.
 *
 * To aggregate on only the IP address or only the forwarded IP address, don't use custom keys. Instead, set the aggregate key type to `IP` or `FORWARDED_IP` .
 * - `FORWARDED_IP` - Aggregate the request counts on the first IP address in an HTTP header.
 *
 * With this option, you must specify the header to use in the `ForwardedIPConfig` property.
 *
 * To aggregate on a combination of the forwarded IP address with other aggregate keys, use `CUSTOM_KEYS` .
 * - `IP` - Aggregate the request counts on the IP address from the web request origin.
 *
 * To aggregate on a combination of the IP address with other aggregate keys, use `CUSTOM_KEYS` .
 */
export type RuleGroupRateBasedStatementAggregateKeyType = (typeof RuleGroupRateBasedStatementAggregateKeyType)[keyof typeof RuleGroupRateBasedStatementAggregateKeyType];

export const RuleGroupRateLimitJa3FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA3 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupRateLimitJa3FingerprintFallbackBehavior = (typeof RuleGroupRateLimitJa3FingerprintFallbackBehavior)[keyof typeof RuleGroupRateLimitJa3FingerprintFallbackBehavior];

export const RuleGroupRateLimitJa4FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA4 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type RuleGroupRateLimitJa4FingerprintFallbackBehavior = (typeof RuleGroupRateLimitJa4FingerprintFallbackBehavior)[keyof typeof RuleGroupRateLimitJa4FingerprintFallbackBehavior];

export const RuleGroupResponseContentType = {
    TextPlain: "TEXT_PLAIN",
    TextHtml: "TEXT_HTML",
    ApplicationJson: "APPLICATION_JSON",
} as const;

/**
 * Valid values are TEXT_PLAIN, TEXT_HTML, and APPLICATION_JSON.
 */
export type RuleGroupResponseContentType = (typeof RuleGroupResponseContentType)[keyof typeof RuleGroupResponseContentType];

export const RuleGroupScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront RuleGroup, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type RuleGroupScope = (typeof RuleGroupScope)[keyof typeof RuleGroupScope];

export const RuleGroupSensitivityLevel = {
    Low: "LOW",
    High: "HIGH",
} as const;

/**
 * Sensitivity Level current only used for sqli match statements.
 */
export type RuleGroupSensitivityLevel = (typeof RuleGroupSensitivityLevel)[keyof typeof RuleGroupSensitivityLevel];

export const RuleGroupSizeConstraintStatementComparisonOperator = {
    Eq: "EQ",
    Ne: "NE",
    Le: "LE",
    Lt: "LT",
    Ge: "GE",
    Gt: "GT",
} as const;

/**
 * The operator to use to compare the request part to the size setting.
 */
export type RuleGroupSizeConstraintStatementComparisonOperator = (typeof RuleGroupSizeConstraintStatementComparisonOperator)[keyof typeof RuleGroupSizeConstraintStatementComparisonOperator];

export const RuleGroupTextTransformationType = {
    None: "NONE",
    CompressWhiteSpace: "COMPRESS_WHITE_SPACE",
    HtmlEntityDecode: "HTML_ENTITY_DECODE",
    Lowercase: "LOWERCASE",
    CmdLine: "CMD_LINE",
    UrlDecode: "URL_DECODE",
    Base64Decode: "BASE64_DECODE",
    HexDecode: "HEX_DECODE",
    Md5: "MD5",
    ReplaceComments: "REPLACE_COMMENTS",
    EscapeSeqDecode: "ESCAPE_SEQ_DECODE",
    SqlHexDecode: "SQL_HEX_DECODE",
    CssDecode: "CSS_DECODE",
    JsDecode: "JS_DECODE",
    NormalizePath: "NORMALIZE_PATH",
    NormalizePathWin: "NORMALIZE_PATH_WIN",
    RemoveNulls: "REMOVE_NULLS",
    ReplaceNulls: "REPLACE_NULLS",
    Base64DecodeExt: "BASE64_DECODE_EXT",
    UrlDecodeUni: "URL_DECODE_UNI",
    Utf8ToUnicode: "UTF8_TO_UNICODE",
} as const;

/**
 * Type of text transformation.
 */
export type RuleGroupTextTransformationType = (typeof RuleGroupTextTransformationType)[keyof typeof RuleGroupTextTransformationType];

export const RuleGroupUriFragmentFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * What AWS WAF should do if it fails to completely parse the JSON body. The options are the following:
 *
 * - `EVALUATE_AS_STRING` - Inspect the body as plain text. AWS WAF applies the text transformations and inspection criteria that you defined for the JSON inspection to the body text string.
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 *
 * If you don't provide this setting, AWS WAF parses and evaluates the content only up to the first parsing failure that it encounters.
 *
 * Example JSON: `{ "UriFragment": { "FallbackBehavior": "MATCH"} }`
 *
 * > AWS WAF parsing doesn't fully validate the input JSON string, so parsing can succeed even for invalid JSON. When parsing succeeds, AWS WAF doesn't apply the fallback behavior. For more information, see [JSON body](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-fields-list.html#waf-rule-statement-request-component-json-body) in the *AWS WAF Developer Guide* .
 */
export type RuleGroupUriFragmentFallbackBehavior = (typeof RuleGroupUriFragmentFallbackBehavior)[keyof typeof RuleGroupUriFragmentFallbackBehavior];

export const WebAclAwsManagedRulesBotControlRuleSetInspectionLevel = {
    Common: "COMMON",
    Targeted: "TARGETED",
} as const;

/**
 * The inspection level to use for the Bot Control rule group. The common level is the least expensive. The targeted level includes all common level rules and adds rules with more advanced inspection criteria. For details, see [AWS WAF Bot Control rule group](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html) in the *AWS WAF Developer Guide* .
 */
export type WebAclAwsManagedRulesBotControlRuleSetInspectionLevel = (typeof WebAclAwsManagedRulesBotControlRuleSetInspectionLevel)[keyof typeof WebAclAwsManagedRulesBotControlRuleSetInspectionLevel];

export const WebAclBodyParsingFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
    EvaluateAsString: "EVALUATE_AS_STRING",
} as const;

/**
 * The inspection behavior to fall back to if the JSON in the request body is invalid.
 */
export type WebAclBodyParsingFallbackBehavior = (typeof WebAclBodyParsingFallbackBehavior)[keyof typeof WebAclBodyParsingFallbackBehavior];

export const WebAclDataProtectionAction = {
    Substitution: "SUBSTITUTION",
    Hash: "HASH",
} as const;

export type WebAclDataProtectionAction = (typeof WebAclDataProtectionAction)[keyof typeof WebAclDataProtectionAction];

export const WebAclFieldToProtectFieldType = {
    SingleHeader: "SINGLE_HEADER",
    SingleCookie: "SINGLE_COOKIE",
    SingleQueryArgument: "SINGLE_QUERY_ARGUMENT",
    QueryString: "QUERY_STRING",
    Body: "BODY",
} as const;

/**
 * Field type to protect
 */
export type WebAclFieldToProtectFieldType = (typeof WebAclFieldToProtectFieldType)[keyof typeof WebAclFieldToProtectFieldType];

export const WebAclForwardedIpConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
 *
 * > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all. 
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclForwardedIpConfigurationFallbackBehavior = (typeof WebAclForwardedIpConfigurationFallbackBehavior)[keyof typeof WebAclForwardedIpConfigurationFallbackBehavior];

export const WebAclIpSetForwardedIpConfigurationFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
 *
 * > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all. 
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclIpSetForwardedIpConfigurationFallbackBehavior = (typeof WebAclIpSetForwardedIpConfigurationFallbackBehavior)[keyof typeof WebAclIpSetForwardedIpConfigurationFallbackBehavior];

export const WebAclIpSetForwardedIpConfigurationPosition = {
    First: "FIRST",
    Last: "LAST",
    Any: "ANY",
} as const;

/**
 * The position in the header to search for the IP address. The header can contain IP addresses of the original client and also of proxies. For example, the header value could be `10.1.1.1, 127.0.0.0, 10.10.10.10` where the first IP address identifies the original client and the rest identify proxies that the request went through.
 *
 * The options for this setting are the following:
 *
 * - FIRST - Inspect the first IP address in the list of IP addresses in the header. This is usually the client's original IP.
 * - LAST - Inspect the last IP address in the list of IP addresses in the header.
 * - ANY - Inspect all IP addresses in the header for a match. If the header contains more than 10 IP addresses, AWS WAF inspects the last 10.
 */
export type WebAclIpSetForwardedIpConfigurationPosition = (typeof WebAclIpSetForwardedIpConfigurationPosition)[keyof typeof WebAclIpSetForwardedIpConfigurationPosition];

export const WebAclJa3FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a JA3 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclJa3FingerprintFallbackBehavior = (typeof WebAclJa3FingerprintFallbackBehavior)[keyof typeof WebAclJa3FingerprintFallbackBehavior];

export const WebAclJa4FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if the request doesn't have a JA4 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclJa4FingerprintFallbackBehavior = (typeof WebAclJa4FingerprintFallbackBehavior)[keyof typeof WebAclJa4FingerprintFallbackBehavior];

export const WebAclJsonMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the JSON to match against using the MatchPattern.
 */
export type WebAclJsonMatchScope = (typeof WebAclJsonMatchScope)[keyof typeof WebAclJsonMatchScope];

export const WebAclLabelMatchScope = {
    Label: "LABEL",
    Namespace: "NAMESPACE",
} as const;

export type WebAclLabelMatchScope = (typeof WebAclLabelMatchScope)[keyof typeof WebAclLabelMatchScope];

export const WebAclManagedRuleGroupConfigPayloadType = {
    Json: "JSON",
    FormEncoded: "FORM_ENCODED",
} as const;

/**
 * > Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet` .
 */
export type WebAclManagedRuleGroupConfigPayloadType = (typeof WebAclManagedRuleGroupConfigPayloadType)[keyof typeof WebAclManagedRuleGroupConfigPayloadType];

export const WebAclMapMatchScope = {
    All: "ALL",
    Key: "KEY",
    Value: "VALUE",
} as const;

/**
 * The parts of the request to match against using the MatchPattern.
 */
export type WebAclMapMatchScope = (typeof WebAclMapMatchScope)[keyof typeof WebAclMapMatchScope];

export const WebAclOnSourceDDoSProtectionConfigAlbLowReputationMode = {
    ActiveUnderDdos: "ACTIVE_UNDER_DDOS",
    AlwaysOn: "ALWAYS_ON",
} as const;

/**
 * The level of DDoS protection that applies to web ACLs associated with Application Load Balancers. `ACTIVE_UNDER_DDOS` protection is enabled by default whenever a web ACL is associated with an Application Load Balancer. In the event that an Application Load Balancer experiences high-load conditions or suspected DDoS attacks, the `ACTIVE_UNDER_DDOS` protection automatically rate limits traffic from known low reputation sources without disrupting Application Load Balancer availability. `ALWAYS_ON` protection provides constant, always-on monitoring of known low reputation sources for suspected DDoS attacks. While this provides a higher level of protection, there may be potential impacts on legitimate traffic.
 */
export type WebAclOnSourceDDoSProtectionConfigAlbLowReputationMode = (typeof WebAclOnSourceDDoSProtectionConfigAlbLowReputationMode)[keyof typeof WebAclOnSourceDDoSProtectionConfigAlbLowReputationMode];

export const WebAclOversizeHandling = {
    Continue: "CONTINUE",
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * Handling of requests containing oversize fields
 */
export type WebAclOversizeHandling = (typeof WebAclOversizeHandling)[keyof typeof WebAclOversizeHandling];

export const WebAclPositionalConstraint = {
    Exactly: "EXACTLY",
    StartsWith: "STARTS_WITH",
    EndsWith: "ENDS_WITH",
    Contains: "CONTAINS",
    ContainsWord: "CONTAINS_WORD",
} as const;

/**
 * Position of the evaluation in the FieldToMatch of request.
 */
export type WebAclPositionalConstraint = (typeof WebAclPositionalConstraint)[keyof typeof WebAclPositionalConstraint];

export const WebAclRateBasedStatementAggregateKeyType = {
    Constant: "CONSTANT",
    Ip: "IP",
    ForwardedIp: "FORWARDED_IP",
    CustomKeys: "CUSTOM_KEYS",
} as const;

/**
 * Setting that indicates how to aggregate the request counts.
 *
 * > Web requests that are missing any of the components specified in the aggregation keys are omitted from the rate-based rule evaluation and handling. 
 *
 * - `CONSTANT` - Count and limit the requests that match the rate-based rule's scope-down statement. With this option, the counted requests aren't further aggregated. The scope-down statement is the only specification used. When the count of all requests that satisfy the scope-down statement goes over the limit, AWS WAF applies the rule action to all requests that satisfy the scope-down statement.
 *
 * With this option, you must configure the `ScopeDownStatement` property.
 * - `CUSTOM_KEYS` - Aggregate the request counts using one or more web request components as the aggregate keys.
 *
 * With this option, you must specify the aggregate keys in the `CustomKeys` property.
 *
 * To aggregate on only the IP address or only the forwarded IP address, don't use custom keys. Instead, set the aggregate key type to `IP` or `FORWARDED_IP` .
 * - `FORWARDED_IP` - Aggregate the request counts on the first IP address in an HTTP header.
 *
 * With this option, you must specify the header to use in the `ForwardedIPConfig` property.
 *
 * To aggregate on a combination of the forwarded IP address with other aggregate keys, use `CUSTOM_KEYS` .
 * - `IP` - Aggregate the request counts on the IP address from the web request origin.
 *
 * To aggregate on a combination of the IP address with other aggregate keys, use `CUSTOM_KEYS` .
 */
export type WebAclRateBasedStatementAggregateKeyType = (typeof WebAclRateBasedStatementAggregateKeyType)[keyof typeof WebAclRateBasedStatementAggregateKeyType];

export const WebAclRateLimitJa3FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA3 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclRateLimitJa3FingerprintFallbackBehavior = (typeof WebAclRateLimitJa3FingerprintFallbackBehavior)[keyof typeof WebAclRateLimitJa3FingerprintFallbackBehavior];

export const WebAclRateLimitJa4FingerprintFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA4 fingerprint.
 *
 * You can specify the following fallback behaviors:
 *
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 */
export type WebAclRateLimitJa4FingerprintFallbackBehavior = (typeof WebAclRateLimitJa4FingerprintFallbackBehavior)[keyof typeof WebAclRateLimitJa4FingerprintFallbackBehavior];

export const WebAclRequestInspectionAcfpPayloadType = {
    Json: "JSON",
    FormEncoded: "FORM_ENCODED",
} as const;

/**
 * The payload type for your account creation endpoint, either JSON or form encoded.
 */
export type WebAclRequestInspectionAcfpPayloadType = (typeof WebAclRequestInspectionAcfpPayloadType)[keyof typeof WebAclRequestInspectionAcfpPayloadType];

export const WebAclRequestInspectionPayloadType = {
    Json: "JSON",
    FormEncoded: "FORM_ENCODED",
} as const;

/**
 * The payload type for your login endpoint, either JSON or form encoded.
 */
export type WebAclRequestInspectionPayloadType = (typeof WebAclRequestInspectionPayloadType)[keyof typeof WebAclRequestInspectionPayloadType];

export const WebAclResponseContentType = {
    TextPlain: "TEXT_PLAIN",
    TextHtml: "TEXT_HTML",
    ApplicationJson: "APPLICATION_JSON",
} as const;

/**
 * Valid values are TEXT_PLAIN, TEXT_HTML, and APPLICATION_JSON.
 */
export type WebAclResponseContentType = (typeof WebAclResponseContentType)[keyof typeof WebAclResponseContentType];

export const WebAclScope = {
    Cloudfront: "CLOUDFRONT",
    Regional: "REGIONAL",
} as const;

/**
 * Use CLOUDFRONT for CloudFront WebACL, use REGIONAL for Application Load Balancer and API Gateway.
 */
export type WebAclScope = (typeof WebAclScope)[keyof typeof WebAclScope];

export const WebAclSensitivityLevel = {
    Low: "LOW",
    High: "HIGH",
} as const;

/**
 * Sensitivity Level current only used for sqli match statements.
 */
export type WebAclSensitivityLevel = (typeof WebAclSensitivityLevel)[keyof typeof WebAclSensitivityLevel];

export const WebAclSensitivityToAct = {
    Low: "LOW",
    Medium: "MEDIUM",
    High: "HIGH",
} as const;

export type WebAclSensitivityToAct = (typeof WebAclSensitivityToAct)[keyof typeof WebAclSensitivityToAct];

export const WebAclSizeConstraintStatementComparisonOperator = {
    Eq: "EQ",
    Ne: "NE",
    Le: "LE",
    Lt: "LT",
    Ge: "GE",
    Gt: "GT",
} as const;

/**
 * The operator to use to compare the request part to the size setting.
 */
export type WebAclSizeConstraintStatementComparisonOperator = (typeof WebAclSizeConstraintStatementComparisonOperator)[keyof typeof WebAclSizeConstraintStatementComparisonOperator];

export const WebAclSizeInspectionLimit = {
    Kb16: "KB_16",
    Kb32: "KB_32",
    Kb48: "KB_48",
    Kb64: "KB_64",
} as const;

export type WebAclSizeInspectionLimit = (typeof WebAclSizeInspectionLimit)[keyof typeof WebAclSizeInspectionLimit];

export const WebAclTextTransformationType = {
    None: "NONE",
    CompressWhiteSpace: "COMPRESS_WHITE_SPACE",
    HtmlEntityDecode: "HTML_ENTITY_DECODE",
    Lowercase: "LOWERCASE",
    CmdLine: "CMD_LINE",
    UrlDecode: "URL_DECODE",
    Base64Decode: "BASE64_DECODE",
    HexDecode: "HEX_DECODE",
    Md5: "MD5",
    ReplaceComments: "REPLACE_COMMENTS",
    EscapeSeqDecode: "ESCAPE_SEQ_DECODE",
    SqlHexDecode: "SQL_HEX_DECODE",
    CssDecode: "CSS_DECODE",
    JsDecode: "JS_DECODE",
    NormalizePath: "NORMALIZE_PATH",
    NormalizePathWin: "NORMALIZE_PATH_WIN",
    RemoveNulls: "REMOVE_NULLS",
    ReplaceNulls: "REPLACE_NULLS",
    Base64DecodeExt: "BASE64_DECODE_EXT",
    UrlDecodeUni: "URL_DECODE_UNI",
    Utf8ToUnicode: "UTF8_TO_UNICODE",
} as const;

/**
 * Type of text transformation.
 */
export type WebAclTextTransformationType = (typeof WebAclTextTransformationType)[keyof typeof WebAclTextTransformationType];

export const WebAclUriFragmentFallbackBehavior = {
    Match: "MATCH",
    NoMatch: "NO_MATCH",
} as const;

/**
 * What AWS WAF should do if it fails to completely parse the JSON body. The options are the following:
 *
 * - `EVALUATE_AS_STRING` - Inspect the body as plain text. AWS WAF applies the text transformations and inspection criteria that you defined for the JSON inspection to the body text string.
 * - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
 * - `NO_MATCH` - Treat the web request as not matching the rule statement.
 *
 * If you don't provide this setting, AWS WAF parses and evaluates the content only up to the first parsing failure that it encounters.
 *
 * Example JSON: `{ "UriFragment": { "FallbackBehavior": "MATCH"} }`
 *
 * > AWS WAF parsing doesn't fully validate the input JSON string, so parsing can succeed even for invalid JSON. When parsing succeeds, AWS WAF doesn't apply the fallback behavior. For more information, see [JSON body](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-fields-list.html#waf-rule-statement-request-component-json-body) in the *AWS WAF Developer Guide* .
 */
export type WebAclUriFragmentFallbackBehavior = (typeof WebAclUriFragmentFallbackBehavior)[keyof typeof WebAclUriFragmentFallbackBehavior];

export const WebAclUsageOfAction = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

export type WebAclUsageOfAction = (typeof WebAclUsageOfAction)[keyof typeof WebAclUsageOfAction];
