// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::MaintenanceWindowTask
 *
 * @deprecated MaintenanceWindowTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class MaintenanceWindowTask extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTask {
        pulumi.log.warn("MaintenanceWindowTask is deprecated: MaintenanceWindowTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new MaintenanceWindowTask(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ssm:MaintenanceWindowTask';

    /**
     * Returns true if the given object is an instance of MaintenanceWindowTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindowTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindowTask.__pulumiType;
    }

    public readonly cutoffBehavior!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly loggingInfo!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskLoggingInfo | undefined>;
    public readonly maxConcurrency!: pulumi.Output<string | undefined>;
    public readonly maxErrors!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly priority!: pulumi.Output<number>;
    public readonly serviceRoleArn!: pulumi.Output<string | undefined>;
    public readonly targets!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTarget[] | undefined>;
    public readonly taskArn!: pulumi.Output<string>;
    public readonly taskInvocationParameters!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTaskInvocationParameters | undefined>;
    public readonly taskParameters!: pulumi.Output<any | undefined>;
    public readonly taskType!: pulumi.Output<string>;
    public readonly windowId!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated MaintenanceWindowTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: MaintenanceWindowTaskArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MaintenanceWindowTask is deprecated: MaintenanceWindowTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.taskArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskArn'");
            }
            if ((!args || args.taskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskType'");
            }
            if ((!args || args.windowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'windowId'");
            }
            resourceInputs["cutoffBehavior"] = args ? args.cutoffBehavior : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["loggingInfo"] = args ? args.loggingInfo : undefined;
            resourceInputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            resourceInputs["maxErrors"] = args ? args.maxErrors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["taskArn"] = args ? args.taskArn : undefined;
            resourceInputs["taskInvocationParameters"] = args ? args.taskInvocationParameters : undefined;
            resourceInputs["taskParameters"] = args ? args.taskParameters : undefined;
            resourceInputs["taskType"] = args ? args.taskType : undefined;
            resourceInputs["windowId"] = args ? args.windowId : undefined;
        } else {
            resourceInputs["cutoffBehavior"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["loggingInfo"] = undefined /*out*/;
            resourceInputs["maxConcurrency"] = undefined /*out*/;
            resourceInputs["maxErrors"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["serviceRoleArn"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
            resourceInputs["taskArn"] = undefined /*out*/;
            resourceInputs["taskInvocationParameters"] = undefined /*out*/;
            resourceInputs["taskParameters"] = undefined /*out*/;
            resourceInputs["taskType"] = undefined /*out*/;
            resourceInputs["windowId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MaintenanceWindowTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MaintenanceWindowTask resource.
 */
export interface MaintenanceWindowTaskArgs {
    cutoffBehavior?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    loggingInfo?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskLoggingInfoArgs>;
    maxConcurrency?: pulumi.Input<string>;
    maxErrors?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    priority: pulumi.Input<number>;
    serviceRoleArn?: pulumi.Input<string>;
    targets?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTargetArgs>[]>;
    taskArn: pulumi.Input<string>;
    taskInvocationParameters?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskInvocationParametersArgs>;
    taskParameters?: any;
    taskType: pulumi.Input<string>;
    windowId: pulumi.Input<string>;
}