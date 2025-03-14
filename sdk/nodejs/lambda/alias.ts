// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Alias
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    /**
     * Lambda Alias ARN generated by the service.
     */
    public /*out*/ readonly aliasArn!: pulumi.Output<string>;
    /**
     * A description of the alias.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the Lambda function.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * The function version that the alias invokes.
     */
    public readonly functionVersion!: pulumi.Output<string>;
    /**
     * The name of the alias.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies a provisioned concurrency configuration for a function's alias.
     */
    public readonly provisionedConcurrencyConfig!: pulumi.Output<outputs.lambda.AliasProvisionedConcurrencyConfiguration | undefined>;
    /**
     * The routing configuration of the alias.
     */
    public readonly routingConfig!: pulumi.Output<outputs.lambda.AliasRoutingConfiguration | undefined>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            if ((!args || args.functionVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionVersion'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["functionVersion"] = args ? args.functionVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["provisionedConcurrencyConfig"] = args ? args.provisionedConcurrencyConfig : undefined;
            resourceInputs["routingConfig"] = args ? args.routingConfig : undefined;
            resourceInputs["aliasArn"] = undefined /*out*/;
        } else {
            resourceInputs["aliasArn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["functionName"] = undefined /*out*/;
            resourceInputs["functionVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisionedConcurrencyConfig"] = undefined /*out*/;
            resourceInputs["routingConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["functionName", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Alias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    /**
     * A description of the alias.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the Lambda function.
     */
    functionName: pulumi.Input<string>;
    /**
     * The function version that the alias invokes.
     */
    functionVersion: pulumi.Input<string>;
    /**
     * The name of the alias.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a provisioned concurrency configuration for a function's alias.
     */
    provisionedConcurrencyConfig?: pulumi.Input<inputs.lambda.AliasProvisionedConcurrencyConfigurationArgs>;
    /**
     * The routing configuration of the alias.
     */
    routingConfig?: pulumi.Input<inputs.lambda.AliasRoutingConfigurationArgs>;
}
