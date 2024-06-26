// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.
 */
export class ExtensionResource extends pulumi.CustomResource {
    /**
     * Get an existing ExtensionResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExtensionResource {
        return new ExtensionResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:index:ExtensionResource';

    /**
     * Returns true if the given object is an instance of ExtensionResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtensionResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtensionResource.__pulumiType;
    }

    /**
     * Dictionary of the extension resource attributes.
     */
    public /*out*/ readonly outputs!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a ExtensionResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["autoNaming"] = args ? args.autoNaming : undefined;
            resourceInputs["createOnly"] = args ? args.createOnly : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["tagsProperty"] = args ? args.tagsProperty : undefined;
            resourceInputs["tagsStyle"] = args ? args.tagsStyle : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["writeOnly"] = args ? args.writeOnly : undefined;
            resourceInputs["outputs"] = undefined /*out*/;
        } else {
            resourceInputs["outputs"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExtensionResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExtensionResource resource.
 */
export interface ExtensionResourceArgs {
    /**
     * Optional auto-naming specification for the resource.
     * If provided and the name is not specified manually, the provider will automatically generate a name based on the Pulumi resource name and a random suffix.
     */
    autoNaming?: pulumi.Input<inputs.AutoNamingArgs>;
    /**
     * Property names as defined by `createOnlyProperties` in the CloudFormation schema. Create-only properties can't be set during updates, so will not be included in patches even if they are also marked as write-only, and will cause an error if attempted to be updated. Therefore any property here should also be included in the `replaceOnChanges` resource option too.
     * In the CloudFormation schema these are fully qualified property paths (e.g. `/properties/AccessToken`) whereas here we only include the top-level property name (e.g. `AccessToken`).
     */
    createOnly?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Property bag containing the properties for the resource. These should be defined using the casing expected by the CloudControl API as these values are sent exact as provided.
     */
    properties: pulumi.Input<{[key: string]: any}>;
    /**
     * Optional name of the property containing the tags. Defaults to "Tags" if the `tagsStyle` is set to either "stringMap" or "keyValueArray". This is used to apply default tags to the resource and can be ignored if not using default tags.
     */
    tagsProperty?: pulumi.Input<string>;
    /**
     * Optional style of tags this resource uses. Valid values are "stringMap", "keyValueArray" or "none". Defaults to `keyValueArray` if `tagsProperty` is set. This is used to apply default tags to the resource and can be ignored if not using default tags.
     */
    tagsStyle?: pulumi.Input<string>;
    /**
     * CloudFormation type name. This has three parts, each separated by two colons. For AWS resources this starts with `AWS::` e.g. `AWS::Logs::LogGroup`. Third party resources should use a namespace prefix e.g. `MyCompany::MyService::MyResource`.
     */
    type: pulumi.Input<string>;
    /**
     * Property names as defined by `writeOnlyProperties` in the CloudFormation schema. Write-only properties are not returned during read operations and have to be included in all update operations as CloudControl itself can't read their previous values.
     * In the CloudFormation schema these are fully qualified property paths (e.g. `/properties/AccessToken`) whereas here we only include the top-level property name (e.g. `AccessToken`).
     */
    writeOnly?: pulumi.Input<pulumi.Input<string>[]>;
}
