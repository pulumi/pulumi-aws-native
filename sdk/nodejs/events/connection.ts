// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Connection.
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:events:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The arn of the connection resource.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly authParameters!: pulumi.Output<outputs.events.ConnectionAuthParameters>;
    public readonly authorizationType!: pulumi.Output<enums.events.ConnectionAuthorizationType>;
    /**
     * Description of the connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the connection.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The arn of the secrets manager secret created in the customer account.
     */
    public /*out*/ readonly secretArn!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authParameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authParameters'");
            }
            if ((!args || args.authorizationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationType'");
            }
            resourceInputs["authParameters"] = args ? args.authParameters : undefined;
            resourceInputs["authorizationType"] = args ? args.authorizationType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["secretArn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["authParameters"] = undefined /*out*/;
            resourceInputs["authorizationType"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["secretArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    authParameters: pulumi.Input<inputs.events.ConnectionAuthParametersArgs>;
    authorizationType: pulumi.Input<enums.events.ConnectionAuthorizationType>;
    /**
     * Description of the connection.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the connection.
     */
    name?: pulumi.Input<string>;
}