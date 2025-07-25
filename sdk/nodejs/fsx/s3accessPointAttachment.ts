// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::FSx::S3AccessPointAttachment
 */
export class S3AccessPointAttachment extends pulumi.CustomResource {
    /**
     * Get an existing S3AccessPointAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): S3AccessPointAttachment {
        return new S3AccessPointAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:fsx:S3AccessPointAttachment';

    /**
     * Returns true if the given object is an instance of S3AccessPointAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3AccessPointAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3AccessPointAttachment.__pulumiType;
    }

    /**
     * The Name of the S3AccessPointAttachment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The OpenZFSConfiguration of the S3 access point attachment.
     */
    public readonly openZfsConfiguration!: pulumi.Output<outputs.fsx.S3AccessPointAttachmentS3AccessPointOpenZfsConfiguration>;
    /**
     * The S3 access point configuration of the S3 access point attachment.
     */
    public readonly s3AccessPoint!: pulumi.Output<outputs.fsx.S3AccessPointAttachmentS3AccessPoint | undefined>;
    /**
     * The type of Amazon FSx volume that the S3 access point is attached to.
     */
    public readonly type!: pulumi.Output<enums.fsx.S3AccessPointAttachmentType>;

    /**
     * Create a S3AccessPointAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3AccessPointAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.openZfsConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'openZfsConfiguration'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openZfsConfiguration"] = args ? args.openZfsConfiguration : undefined;
            resourceInputs["s3AccessPoint"] = args ? args.s3AccessPoint : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["openZfsConfiguration"] = undefined /*out*/;
            resourceInputs["s3AccessPoint"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "openZfsConfiguration", "s3AccessPoint", "type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(S3AccessPointAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a S3AccessPointAttachment resource.
 */
export interface S3AccessPointAttachmentArgs {
    /**
     * The Name of the S3AccessPointAttachment
     */
    name?: pulumi.Input<string>;
    /**
     * The OpenZFSConfiguration of the S3 access point attachment.
     */
    openZfsConfiguration: pulumi.Input<inputs.fsx.S3AccessPointAttachmentS3AccessPointOpenZfsConfigurationArgs>;
    /**
     * The S3 access point configuration of the S3 access point attachment.
     */
    s3AccessPoint?: pulumi.Input<inputs.fsx.S3AccessPointAttachmentS3AccessPointArgs>;
    /**
     * The type of Amazon FSx volume that the S3 access point is attached to.
     */
    type: pulumi.Input<enums.fsx.S3AccessPointAttachmentType>;
}
