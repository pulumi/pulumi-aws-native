// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Ground Station DataflowEndpointGroup schema for CloudFormation
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myDataflowEndpointGroup = new aws_native.groundstation.DataflowEndpointGroup("myDataflowEndpointGroup", {endpointDetails: [{
 *     securityDetails: {
 *         subnetIds: ["subnet-6782e71e"],
 *         securityGroupIds: ["sg-6979fe18"],
 *         roleArn: "arn:aws:iam::012345678910:role/groundstation-service-role-AWSServiceRoleForAmazonGroundStation-EXAMPLEBQ4PI",
 *     },
 *     endpoint: {
 *         name: "myEndpoint",
 *         address: {
 *             name: "172.10.0.2",
 *             port: 44720,
 *         },
 *         mtu: 1500,
 *     },
 * }]});
 *
 * ```
 */
export class DataflowEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing DataflowEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataflowEndpointGroup {
        return new DataflowEndpointGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:groundstation:DataflowEndpointGroup';

    /**
     * Returns true if the given object is an instance of DataflowEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataflowEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataflowEndpointGroup.__pulumiType;
    }

    /**
     * The ARN of the dataflow endpoint group, such as `arn:aws:groundstation:us-east-2:1234567890:dataflow-endpoint-group/9940bf3b-d2ba-427e-9906-842b5e5d2296` .
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * UUID of a dataflow endpoint group.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.
     */
    public readonly contactPostPassDurationSeconds!: pulumi.Output<number | undefined>;
    /**
     * Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.
     */
    public readonly contactPrePassDurationSeconds!: pulumi.Output<number | undefined>;
    /**
     * List of Endpoint Details, containing address and port for each endpoint. All dataflow endpoints within a single dataflow endpoint group must be of the same type. You cannot mix AWS Ground Station Agent endpoints with Dataflow endpoints in the same group. If your use case requires both types of endpoints, you must create separate dataflow endpoint groups for each type.
     */
    public readonly endpointDetails!: pulumi.Output<outputs.groundstation.DataflowEndpointGroupEndpointDetails[]>;
    /**
     * Tags assigned to a resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DataflowEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataflowEndpointGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.endpointDetails === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointDetails'");
            }
            resourceInputs["contactPostPassDurationSeconds"] = args ? args.contactPostPassDurationSeconds : undefined;
            resourceInputs["contactPrePassDurationSeconds"] = args ? args.contactPrePassDurationSeconds : undefined;
            resourceInputs["endpointDetails"] = args ? args.endpointDetails : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["contactPostPassDurationSeconds"] = undefined /*out*/;
            resourceInputs["contactPrePassDurationSeconds"] = undefined /*out*/;
            resourceInputs["endpointDetails"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["contactPostPassDurationSeconds", "contactPrePassDurationSeconds", "endpointDetails[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataflowEndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataflowEndpointGroup resource.
 */
export interface DataflowEndpointGroupArgs {
    /**
     * Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.
     */
    contactPostPassDurationSeconds?: pulumi.Input<number>;
    /**
     * Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.
     */
    contactPrePassDurationSeconds?: pulumi.Input<number>;
    /**
     * List of Endpoint Details, containing address and port for each endpoint. All dataflow endpoints within a single dataflow endpoint group must be of the same type. You cannot mix AWS Ground Station Agent endpoints with Dataflow endpoints in the same group. If your use case requires both types of endpoints, you must create separate dataflow endpoint groups for each type.
     */
    endpointDetails: pulumi.Input<pulumi.Input<inputs.groundstation.DataflowEndpointGroupEndpointDetailsArgs>[]>;
    /**
     * Tags assigned to a resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
