// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::GraphQLSchema
 *
 * @deprecated GraphQLSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class GraphQLSchema extends pulumi.CustomResource {
    /**
     * Get an existing GraphQLSchema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GraphQLSchema {
        pulumi.log.warn("GraphQLSchema is deprecated: GraphQLSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new GraphQLSchema(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appsync:GraphQLSchema';

    /**
     * Returns true if the given object is an instance of GraphQLSchema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GraphQLSchema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GraphQLSchema.__pulumiType;
    }

    public readonly apiId!: pulumi.Output<string>;
    public readonly definition!: pulumi.Output<string | undefined>;
    public readonly definitionS3Location!: pulumi.Output<string | undefined>;

    /**
     * Create a GraphQLSchema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated GraphQLSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: GraphQLSchemaArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("GraphQLSchema is deprecated: GraphQLSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["definitionS3Location"] = args ? args.definitionS3Location : undefined;
        } else {
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["definitionS3Location"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GraphQLSchema.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GraphQLSchema resource.
 */
export interface GraphQLSchemaArgs {
    apiId: pulumi.Input<string>;
    definition?: pulumi.Input<string>;
    definitionS3Location?: pulumi.Input<string>;
}