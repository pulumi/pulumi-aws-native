// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Scheduler::Schedule Resource Type
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:scheduler:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the schedule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the schedule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
     */
    public readonly endDate!: pulumi.Output<string | undefined>;
    /**
     * Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
     */
    public readonly flexibleTimeWindow!: pulumi.Output<outputs.scheduler.ScheduleFlexibleTimeWindow>;
    /**
     * The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used.
     */
    public readonly groupName!: pulumi.Output<string | undefined>;
    /**
     * The ARN for a KMS Key that will be used to encrypt customer data.
     */
    public readonly kmsKeyArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the schedule.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The scheduling expression.
     */
    public readonly scheduleExpression!: pulumi.Output<string>;
    /**
     * The timezone in which the scheduling expression is evaluated.
     */
    public readonly scheduleExpressionTimezone!: pulumi.Output<string | undefined>;
    /**
     * The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
     */
    public readonly startDate!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the schedule is enabled or disabled.
     *
     * *Allowed Values* : `ENABLED` | `DISABLED`
     */
    public readonly state!: pulumi.Output<enums.scheduler.ScheduleState | undefined>;
    /**
     * The schedule's target details.
     */
    public readonly target!: pulumi.Output<outputs.scheduler.ScheduleTarget>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.flexibleTimeWindow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flexibleTimeWindow'");
            }
            if ((!args || args.scheduleExpression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduleExpression'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["flexibleTimeWindow"] = args ? args.flexibleTimeWindow : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scheduleExpression"] = args ? args.scheduleExpression : undefined;
            resourceInputs["scheduleExpressionTimezone"] = args ? args.scheduleExpressionTimezone : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endDate"] = undefined /*out*/;
            resourceInputs["flexibleTimeWindow"] = undefined /*out*/;
            resourceInputs["groupName"] = undefined /*out*/;
            resourceInputs["kmsKeyArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["scheduleExpression"] = undefined /*out*/;
            resourceInputs["scheduleExpressionTimezone"] = undefined /*out*/;
            resourceInputs["startDate"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["target"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Schedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    /**
     * The description of the schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
     */
    endDate?: pulumi.Input<string>;
    /**
     * Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
     */
    flexibleTimeWindow: pulumi.Input<inputs.scheduler.ScheduleFlexibleTimeWindowArgs>;
    /**
     * The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The ARN for a KMS Key that will be used to encrypt customer data.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * The name of the schedule.
     */
    name?: pulumi.Input<string>;
    /**
     * The scheduling expression.
     */
    scheduleExpression: pulumi.Input<string>;
    /**
     * The timezone in which the scheduling expression is evaluated.
     */
    scheduleExpressionTimezone?: pulumi.Input<string>;
    /**
     * The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
     */
    startDate?: pulumi.Input<string>;
    /**
     * Specifies whether the schedule is enabled or disabled.
     *
     * *Allowed Values* : `ENABLED` | `DISABLED`
     */
    state?: pulumi.Input<enums.scheduler.ScheduleState>;
    /**
     * The schedule's target details.
     */
    target: pulumi.Input<inputs.scheduler.ScheduleTargetArgs>;
}
