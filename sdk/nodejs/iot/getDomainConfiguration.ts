// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create and manage a Domain Configuration
 */
export function getDomainConfiguration(args: GetDomainConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getDomainConfiguration", {
        "domainConfigurationName": args.domainConfigurationName,
    }, opts);
}

export interface GetDomainConfigurationArgs {
    /**
     * The name of the domain configuration. This value must be unique to a region.
     */
    domainConfigurationName: string;
}

export interface GetDomainConfigurationResult {
    /**
     * An enumerated string that speciﬁes the application-layer protocol.
     */
    readonly applicationProtocol?: enums.iot.DomainConfigurationApplicationProtocol;
    /**
     * The Amazon Resource Name (ARN) of the domain configuration.
     */
    readonly arn?: string;
    /**
     * An enumerated string that speciﬁes the authentication type.
     */
    readonly authenticationType?: enums.iot.DomainConfigurationAuthenticationType;
    /**
     * An object that specifies the authorization service for a domain.
     */
    readonly authorizerConfig?: outputs.iot.DomainConfigurationAuthorizerConfig;
    /**
     * An object that speciﬁes the client certificate conﬁguration for a domain.
     */
    readonly clientCertificateConfig?: outputs.iot.DomainConfigurationClientCertificateConfig;
    /**
     * The status to which the domain configuration should be updated.
     *
     * Valid values: `ENABLED` | `DISABLED`
     */
    readonly domainConfigurationStatus?: enums.iot.DomainConfigurationStatus;
    /**
     * The type of service delivered by the domain.
     */
    readonly domainType?: enums.iot.DomainConfigurationDomainType;
    /**
     * The server certificate configuration.
     *
     * For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
     */
    readonly serverCertificateConfig?: outputs.iot.DomainConfigurationServerCertificateConfig;
    /**
     * The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
     */
    readonly serverCertificates?: outputs.iot.DomainConfigurationServerCertificateSummary[];
    /**
     * Metadata which can be used to manage the domain configuration.
     *
     * > For URI Request parameters use format: ...key1=value1&key2=value2...
     * > 
     * > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
     * > 
     * > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
     */
    readonly tags?: outputs.Tag[];
    /**
     * An object that specifies the TLS configuration for a domain.
     */
    readonly tlsConfig?: outputs.iot.DomainConfigurationTlsConfig;
}
/**
 * Create and manage a Domain Configuration
 */
export function getDomainConfigurationOutput(args: GetDomainConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDomainConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:iot:getDomainConfiguration", {
        "domainConfigurationName": args.domainConfigurationName,
    }, opts);
}

export interface GetDomainConfigurationOutputArgs {
    /**
     * The name of the domain configuration. This value must be unique to a region.
     */
    domainConfigurationName: pulumi.Input<string>;
}
