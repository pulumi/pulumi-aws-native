// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states*. For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.
 */
export class DetectorModel extends pulumi.CustomResource {
    /**
     * Get an existing DetectorModel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DetectorModel {
        return new DetectorModel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotevents:DetectorModel';

    /**
     * Returns true if the given object is an instance of DetectorModel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DetectorModel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DetectorModel.__pulumiType;
    }

    public readonly detectorModelDefinition!: pulumi.Output<outputs.iotevents.DetectorModelDefinition>;
    /**
     * A brief description of the detector model.
     */
    public readonly detectorModelDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the detector model.
     */
    public readonly detectorModelName!: pulumi.Output<string | undefined>;
    /**
     * Information about the order in which events are evaluated and how actions are executed.
     */
    public readonly evaluationMethod!: pulumi.Output<enums.iotevents.DetectorModelEvaluationMethod | undefined>;
    /**
     * The value used to identify a detector instance. When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
     *
     * This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the role that grants permission to AWS IoT Events to perform its operations.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    public readonly tags!: pulumi.Output<outputs.iotevents.DetectorModelTag[] | undefined>;

    /**
     * Create a DetectorModel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DetectorModelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.detectorModelDefinition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'detectorModelDefinition'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["detectorModelDefinition"] = args ? args.detectorModelDefinition : undefined;
            resourceInputs["detectorModelDescription"] = args ? args.detectorModelDescription : undefined;
            resourceInputs["detectorModelName"] = args ? args.detectorModelName : undefined;
            resourceInputs["evaluationMethod"] = args ? args.evaluationMethod : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["detectorModelDefinition"] = undefined /*out*/;
            resourceInputs["detectorModelDescription"] = undefined /*out*/;
            resourceInputs["detectorModelName"] = undefined /*out*/;
            resourceInputs["evaluationMethod"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DetectorModel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DetectorModel resource.
 */
export interface DetectorModelArgs {
    detectorModelDefinition: pulumi.Input<inputs.iotevents.DetectorModelDefinitionArgs>;
    /**
     * A brief description of the detector model.
     */
    detectorModelDescription?: pulumi.Input<string>;
    /**
     * The name of the detector model.
     */
    detectorModelName?: pulumi.Input<string>;
    /**
     * Information about the order in which events are evaluated and how actions are executed.
     */
    evaluationMethod?: pulumi.Input<enums.iotevents.DetectorModelEvaluationMethod>;
    /**
     * The value used to identify a detector instance. When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
     *
     * This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
     */
    key?: pulumi.Input<string>;
    /**
     * The ARN of the role that grants permission to AWS IoT Events to perform its operations.
     */
    roleArn: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    tags?: pulumi.Input<pulumi.Input<inputs.iotevents.DetectorModelTagArgs>[]>;
}