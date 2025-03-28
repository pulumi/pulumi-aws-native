// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Integration from a source AWS service to a Redshift cluster
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:redshift:Integration';

    /**
     * Returns true if the given object is an instance of Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Integration.__pulumiType;
    }

    /**
     * The encryption context for the integration. For more information, see [Encryption context](https://docs.aws.amazon.com/) in the *AWS Key Management Service Developer Guide* .
     */
    public readonly additionalEncryptionContext!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time (UTC) when the integration was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the integration.
     */
    public /*out*/ readonly integrationArn!: pulumi.Output<string>;
    /**
     * The name of the integration.
     */
    public readonly integrationName!: pulumi.Output<string | undefined>;
    /**
     * An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the database to use as the source for replication
     */
    public readonly sourceArn!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication
     */
    public readonly targetArn!: pulumi.Output<string>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceArn'");
            }
            if ((!args || args.targetArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetArn'");
            }
            resourceInputs["additionalEncryptionContext"] = args ? args.additionalEncryptionContext : undefined;
            resourceInputs["integrationName"] = args ? args.integrationName : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["sourceArn"] = args ? args.sourceArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetArn"] = args ? args.targetArn : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["integrationArn"] = undefined /*out*/;
        } else {
            resourceInputs["additionalEncryptionContext"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["integrationArn"] = undefined /*out*/;
            resourceInputs["integrationName"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["sourceArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targetArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["additionalEncryptionContext.*", "kmsKeyId", "sourceArn", "targetArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * The encryption context for the integration. For more information, see [Encryption context](https://docs.aws.amazon.com/) in the *AWS Key Management Service Developer Guide* .
     */
    additionalEncryptionContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the integration.
     */
    integrationName?: pulumi.Input<string>;
    /**
     * An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the database to use as the source for replication
     */
    sourceArn: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication
     */
    targetArn: pulumi.Input<string>;
}
