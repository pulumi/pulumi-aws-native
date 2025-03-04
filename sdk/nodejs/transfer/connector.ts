// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Transfer::Connector
 */
export class Connector extends pulumi.CustomResource {
    /**
     * Get an existing Connector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connector {
        return new Connector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:transfer:Connector';

    /**
     * Returns true if the given object is an instance of Connector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connector.__pulumiType;
    }

    /**
     * Specifies the access role for the connector.
     */
    public readonly accessRole!: pulumi.Output<string>;
    /**
     * Specifies the unique Amazon Resource Name (ARN) for the connector.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration for an AS2 connector.
     */
    public readonly as2Config!: pulumi.Output<outputs.transfer.As2ConfigProperties | undefined>;
    /**
     * A unique identifier for the connector.
     */
    public /*out*/ readonly connectorId!: pulumi.Output<string>;
    /**
     * Specifies the logging role for the connector.
     */
    public readonly loggingRole!: pulumi.Output<string | undefined>;
    /**
     * Security policy for SFTP Connector
     */
    public readonly securityPolicyName!: pulumi.Output<string | undefined>;
    /**
     * The list of egress IP addresses of this connector. These IP addresses are assigned automatically when you create the connector.
     */
    public /*out*/ readonly serviceManagedEgressIpAddresses!: pulumi.Output<string[]>;
    /**
     * Configuration for an SFTP connector.
     */
    public readonly sftpConfig!: pulumi.Output<outputs.transfer.SftpConfigProperties | undefined>;
    /**
     * Key-value pairs that can be used to group and search for connectors. Tags are metadata attached to connectors for any purpose.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * URL for Connector
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Connector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accessRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessRole'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["accessRole"] = args ? args.accessRole : undefined;
            resourceInputs["as2Config"] = args ? args.as2Config : undefined;
            resourceInputs["loggingRole"] = args ? args.loggingRole : undefined;
            resourceInputs["securityPolicyName"] = args ? args.securityPolicyName : undefined;
            resourceInputs["sftpConfig"] = args ? args.sftpConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["connectorId"] = undefined /*out*/;
            resourceInputs["serviceManagedEgressIpAddresses"] = undefined /*out*/;
        } else {
            resourceInputs["accessRole"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["as2Config"] = undefined /*out*/;
            resourceInputs["connectorId"] = undefined /*out*/;
            resourceInputs["loggingRole"] = undefined /*out*/;
            resourceInputs["securityPolicyName"] = undefined /*out*/;
            resourceInputs["serviceManagedEgressIpAddresses"] = undefined /*out*/;
            resourceInputs["sftpConfig"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connector resource.
 */
export interface ConnectorArgs {
    /**
     * Specifies the access role for the connector.
     */
    accessRole: pulumi.Input<string>;
    /**
     * Configuration for an AS2 connector.
     */
    as2Config?: pulumi.Input<inputs.transfer.As2ConfigPropertiesArgs>;
    /**
     * Specifies the logging role for the connector.
     */
    loggingRole?: pulumi.Input<string>;
    /**
     * Security policy for SFTP Connector
     */
    securityPolicyName?: pulumi.Input<string>;
    /**
     * Configuration for an SFTP connector.
     */
    sftpConfig?: pulumi.Input<inputs.transfer.SftpConfigPropertiesArgs>;
    /**
     * Key-value pairs that can be used to group and search for connectors. Tags are metadata attached to connectors for any purpose.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * URL for Connector
     */
    url: pulumi.Input<string>;
}
