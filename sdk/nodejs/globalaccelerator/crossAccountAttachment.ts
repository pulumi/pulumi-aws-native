// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment
 */
export class CrossAccountAttachment extends pulumi.CustomResource {
    /**
     * Get an existing CrossAccountAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CrossAccountAttachment {
        return new CrossAccountAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:globalaccelerator:CrossAccountAttachment';

    /**
     * Returns true if the given object is an instance of CrossAccountAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CrossAccountAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CrossAccountAttachment.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the attachment.
     */
    public /*out*/ readonly attachmentArn!: pulumi.Output<string>;
    /**
     * The Friendly identifier of the attachment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Principals to share the resources with.
     */
    public readonly principals!: pulumi.Output<string[] | undefined>;
    /**
     * Resources shared using the attachment.
     */
    public readonly resources!: pulumi.Output<outputs.globalaccelerator.CrossAccountAttachmentResource[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a CrossAccountAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CrossAccountAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["principals"] = args ? args.principals : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["attachmentArn"] = undefined /*out*/;
        } else {
            resourceInputs["attachmentArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["principals"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CrossAccountAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CrossAccountAttachment resource.
 */
export interface CrossAccountAttachmentArgs {
    /**
     * The Friendly identifier of the attachment.
     */
    name?: pulumi.Input<string>;
    /**
     * Principals to share the resources with.
     */
    principals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resources shared using the attachment.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.CrossAccountAttachmentResourceArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}