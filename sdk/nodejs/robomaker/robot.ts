// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
 */
export class Robot extends pulumi.CustomResource {
    /**
     * Get an existing Robot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Robot {
        return new Robot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:robomaker:Robot';

    /**
     * Returns true if the given object is an instance of Robot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Robot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Robot.__pulumiType;
    }

    /**
     * The target architecture of the robot.
     */
    public readonly architecture!: pulumi.Output<enums.robomaker.RobotArchitecture>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the fleet.
     */
    public readonly fleet!: pulumi.Output<string | undefined>;
    /**
     * The Greengrass group id.
     */
    public readonly greengrassGroupId!: pulumi.Output<string>;
    /**
     * The name for the robot.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.robomaker.RobotTags | undefined>;

    /**
     * Create a Robot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RobotArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.architecture === undefined) && !opts.urn) {
                throw new Error("Missing required property 'architecture'");
            }
            if ((!args || args.greengrassGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'greengrassGroupId'");
            }
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["fleet"] = args ? args.fleet : undefined;
            resourceInputs["greengrassGroupId"] = args ? args.greengrassGroupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["fleet"] = undefined /*out*/;
            resourceInputs["greengrassGroupId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Robot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Robot resource.
 */
export interface RobotArgs {
    /**
     * The target architecture of the robot.
     */
    architecture: pulumi.Input<enums.robomaker.RobotArchitecture>;
    /**
     * The Amazon Resource Name (ARN) of the fleet.
     */
    fleet?: pulumi.Input<string>;
    /**
     * The Greengrass group id.
     */
    greengrassGroupId: pulumi.Input<string>;
    /**
     * The name for the robot.
     */
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<inputs.robomaker.RobotTagsArgs>;
}