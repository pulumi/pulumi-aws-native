// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * IdNamespace defined in AWS Entity Resolution service
 */
export class IdNamespace extends pulumi.CustomResource {
    /**
     * Get an existing IdNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IdNamespace {
        return new IdNamespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:entityresolution:IdNamespace';

    /**
     * Returns true if the given object is an instance of IdNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdNamespace.__pulumiType;
    }

    /**
     * The date and time when the IdNamespace was created
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly idMappingWorkflowProperties!: pulumi.Output<outputs.entityresolution.IdNamespaceIdMappingWorkflowProperties[] | undefined>;
    /**
     * The arn associated with the IdNamespace
     */
    public /*out*/ readonly idNamespaceArn!: pulumi.Output<string>;
    public readonly idNamespaceName!: pulumi.Output<string>;
    public readonly inputSourceConfig!: pulumi.Output<outputs.entityresolution.IdNamespaceInputSource[] | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly type!: pulumi.Output<enums.entityresolution.IdNamespaceType>;
    /**
     * The date and time when the IdNamespace was updated
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IdNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdNamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["idMappingWorkflowProperties"] = args ? args.idMappingWorkflowProperties : undefined;
            resourceInputs["idNamespaceName"] = args ? args.idNamespaceName : undefined;
            resourceInputs["inputSourceConfig"] = args ? args.inputSourceConfig : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["idNamespaceArn"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["idMappingWorkflowProperties"] = undefined /*out*/;
            resourceInputs["idNamespaceArn"] = undefined /*out*/;
            resourceInputs["idNamespaceName"] = undefined /*out*/;
            resourceInputs["inputSourceConfig"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["idNamespaceName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(IdNamespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IdNamespace resource.
 */
export interface IdNamespaceArgs {
    description?: pulumi.Input<string>;
    idMappingWorkflowProperties?: pulumi.Input<pulumi.Input<inputs.entityresolution.IdNamespaceIdMappingWorkflowPropertiesArgs>[]>;
    idNamespaceName?: pulumi.Input<string>;
    inputSourceConfig?: pulumi.Input<pulumi.Input<inputs.entityresolution.IdNamespaceInputSourceArgs>[]>;
    roleArn?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    type: pulumi.Input<enums.entityresolution.IdNamespaceType>;
}