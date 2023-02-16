// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AmplifyUIBuilder::Theme Resource Type
 */
export class Theme extends pulumi.CustomResource {
    /**
     * Get an existing Theme resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Theme {
        return new Theme(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:amplifyuibuilder:Theme';

    /**
     * Returns true if the given object is an instance of Theme.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Theme {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Theme.__pulumiType;
    }

    public readonly appId!: pulumi.Output<string | undefined>;
    public readonly environmentName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly overrides!: pulumi.Output<outputs.amplifyuibuilder.ThemeValues[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.amplifyuibuilder.ThemeTags | undefined>;
    public readonly values!: pulumi.Output<outputs.amplifyuibuilder.ThemeValues[]>;

    /**
     * Create a Theme resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ThemeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.values === undefined) && !opts.urn) {
                throw new Error("Missing required property 'values'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrides"] = args ? args.overrides : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["environmentName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["overrides"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["values"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Theme.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Theme resource.
 */
export interface ThemeArgs {
    appId?: pulumi.Input<string>;
    environmentName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrides?: pulumi.Input<pulumi.Input<inputs.amplifyuibuilder.ThemeValuesArgs>[]>;
    tags?: pulumi.Input<inputs.amplifyuibuilder.ThemeTagsArgs>;
    values: pulumi.Input<pulumi.Input<inputs.amplifyuibuilder.ThemeValuesArgs>[]>;
}