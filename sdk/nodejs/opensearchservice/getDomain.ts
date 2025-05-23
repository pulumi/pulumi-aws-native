// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:opensearchservice:getDomain", {
        "domainName": args.domainName,
    }, opts);
}

export interface GetDomainArgs {
    /**
     * A name for the OpenSearch Service domain. The name must have a minimum length of 3 and a maximum length of 28. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
     *
     * Required when creating a new domain.
     *
     * > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    domainName: string;
}

export interface GetDomainResult {
    /**
     * An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions. For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guide* .
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::OpenSearchService::Domain` for more information about the expected schema for this property.
     */
    readonly accessPolicies?: any;
    /**
     * Additional options to specify for the OpenSearch Service domain. For more information, see [AdvancedOptions](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_CreateDomain.html#API_CreateDomain_RequestBody) in the OpenSearch Service API reference.
     */
    readonly advancedOptions?: {[key: string]: string};
    /**
     * Specifies options for fine-grained access control and SAML authentication.
     *
     * If you specify advanced security options, you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html) ) and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html) ). You must also enable `EnforceHTTPS` within [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html) , which requires HTTPS for all traffic to the domain.
     */
    readonly advancedSecurityOptions?: outputs.opensearchservice.DomainAdvancedSecurityOptionsInput;
    /**
     * The Amazon Resource Name (ARN) of the CloudFormation stack.
     */
    readonly arn?: string;
    /**
     * Container for the cluster configuration of a domain.
     */
    readonly clusterConfig?: outputs.opensearchservice.DomainClusterConfig;
    /**
     * Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
     */
    readonly cognitoOptions?: outputs.opensearchservice.DomainCognitoOptions;
    /**
     * The Amazon Resource Name (ARN) of the domain. See [Identifiers for IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in *Using AWS Identity and Access Management* for more information.
     */
    readonly domainArn?: string;
    /**
     * The domain-specific endpoint used for requests to the OpenSearch APIs, such as `search-mystack-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com` .
     */
    readonly domainEndpoint?: string;
    /**
     * Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
     */
    readonly domainEndpointOptions?: outputs.opensearchservice.DomainEndpointOptions;
    /**
     * If `IPAddressType` to set to `dualstack` , a version 2 domain endpoint is provisioned. This endpoint functions like a normal endpoint, except that it works with both IPv4 and IPv6 IP addresses. Normal endpoints work only with IPv4 IP addresses.
     */
    readonly domainEndpointV2?: string;
    readonly domainEndpoints?: {[key: string]: string};
    /**
     * The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
     */
    readonly ebsOptions?: outputs.opensearchservice.DomainEbsOptions;
    /**
     * Whether the domain should encrypt data at rest, and if so, the AWS KMS key to use. See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
     *
     * If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template, the domain is deleted and recreated in order to modify the property.
     */
    readonly encryptionAtRestOptions?: outputs.opensearchservice.DomainEncryptionAtRestOptions;
    /**
     * The version of OpenSearch to use. The value must be in the format `OpenSearch_X.Y` or `Elasticsearch_X.Y` . If not specified, the latest version of OpenSearch is used. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
     *
     * If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `EngineVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `EngineVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
     */
    readonly engineVersion?: string;
    /**
     * The resource ID. For example, `123456789012/my-domain` .
     */
    readonly id?: string;
    /**
     * Configuration options for controlling IAM Identity Center integration within a domain.
     */
    readonly identityCenterOptions?: outputs.opensearchservice.DomainIdentityCenterOptions;
    /**
     * Choose either dual stack or IPv4 as your IP address type. Dual stack allows you to share domain resources across IPv4 and IPv6 address types, and is the recommended option. If you set your IP address type to dual stack, you can't change your address type later.
     */
    readonly ipAddressType?: string;
    /**
     * An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish. Each key needs a valid `LogPublishingOption` value. For the full syntax, see the [examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--examples) .
     */
    readonly logPublishingOptions?: {[key: string]: outputs.opensearchservice.DomainLogPublishingOption};
    /**
     * Specifies whether node-to-node encryption is enabled. See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
     */
    readonly nodeToNodeEncryptionOptions?: outputs.opensearchservice.DomainNodeToNodeEncryptionOptions;
    /**
     * Options for a domain's off-peak window, during which OpenSearch Service can perform mandatory configuration changes on the domain.
     */
    readonly offPeakWindowOptions?: outputs.opensearchservice.DomainOffPeakWindowOptions;
    readonly serviceSoftwareOptions?: outputs.opensearchservice.DomainServiceSoftwareOptions;
    readonly skipShardMigrationWait?: boolean;
    /**
     * *DEPRECATED* . The automated snapshot configuration for the OpenSearch Service domain indexes.
     */
    readonly snapshotOptions?: outputs.opensearchservice.DomainSnapshotOptions;
    /**
     * Service software update options for the domain.
     */
    readonly softwareUpdateOptions?: outputs.opensearchservice.DomainSoftwareUpdateOptions;
    /**
     * An arbitrary set of tags (key-value pairs) for this Domain.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The virtual private cloud (VPC) configuration for the OpenSearch Service domain. For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
     *
     * If you remove this entity altogether, along with its associated properties, it causes a replacement. You might encounter this scenario if you're updating your security configuration from a VPC to a public endpoint.
     */
    readonly vpcOptions?: outputs.opensearchservice.DomainVpcOptions;
}
/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:opensearchservice:getDomain", {
        "domainName": args.domainName,
    }, opts);
}

export interface GetDomainOutputArgs {
    /**
     * A name for the OpenSearch Service domain. The name must have a minimum length of 3 and a maximum length of 28. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
     *
     * Required when creating a new domain.
     *
     * > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    domainName: pulumi.Input<string>;
}
