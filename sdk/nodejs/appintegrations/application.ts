// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS:AppIntegrations::Application
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appintegrations:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the application.
     */
    public /*out*/ readonly applicationArn!: pulumi.Output<string>;
    /**
     * Application source config
     */
    public readonly applicationSourceConfig!: pulumi.Output<outputs.appintegrations.ApplicationSourceConfigProperties>;
    /**
     * The id of the application.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The application description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name of the application.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace of the application.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The tags (keys and values) associated with the application.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationSourceConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationSourceConfig'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            resourceInputs["applicationSourceConfig"] = args ? args.applicationSourceConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["applicationArn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationArn"] = undefined /*out*/;
            resourceInputs["applicationSourceConfig"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespace"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * Application source config
     */
    applicationSourceConfig: pulumi.Input<inputs.appintegrations.ApplicationSourceConfigPropertiesArgs>;
    /**
     * The application description.
     */
    description: pulumi.Input<string>;
    /**
     * The name of the application.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace of the application.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The tags (keys and values) associated with the application.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}