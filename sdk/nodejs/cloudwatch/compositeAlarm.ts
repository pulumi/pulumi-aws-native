// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
 */
export class CompositeAlarm extends pulumi.CustomResource {
    /**
     * Get an existing CompositeAlarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CompositeAlarm {
        return new CompositeAlarm(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudwatch:CompositeAlarm';

    /**
     * Returns true if the given object is an instance of CompositeAlarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CompositeAlarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CompositeAlarm.__pulumiType;
    }

    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    public readonly actionsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. 
     */
    public readonly actionsSuppressor!: pulumi.Output<string | undefined>;
    /**
     * Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.
     */
    public readonly actionsSuppressorExtensionPeriod!: pulumi.Output<number | undefined>;
    /**
     * Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.
     */
    public readonly actionsSuppressorWaitPeriod!: pulumi.Output<number | undefined>;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
     */
    public readonly alarmActions!: pulumi.Output<string[] | undefined>;
    /**
     * The description of the alarm
     */
    public readonly alarmDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the Composite Alarm
     */
    public readonly alarmName!: pulumi.Output<string>;
    /**
     * Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
     */
    public readonly alarmRule!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the alarm
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    public readonly insufficientDataActions!: pulumi.Output<string[] | undefined>;
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    public readonly oKActions!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CompositeAlarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CompositeAlarmArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.alarmName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alarmName'");
            }
            if ((!args || args.alarmRule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alarmRule'");
            }
            resourceInputs["actionsEnabled"] = args ? args.actionsEnabled : undefined;
            resourceInputs["actionsSuppressor"] = args ? args.actionsSuppressor : undefined;
            resourceInputs["actionsSuppressorExtensionPeriod"] = args ? args.actionsSuppressorExtensionPeriod : undefined;
            resourceInputs["actionsSuppressorWaitPeriod"] = args ? args.actionsSuppressorWaitPeriod : undefined;
            resourceInputs["alarmActions"] = args ? args.alarmActions : undefined;
            resourceInputs["alarmDescription"] = args ? args.alarmDescription : undefined;
            resourceInputs["alarmName"] = args ? args.alarmName : undefined;
            resourceInputs["alarmRule"] = args ? args.alarmRule : undefined;
            resourceInputs["insufficientDataActions"] = args ? args.insufficientDataActions : undefined;
            resourceInputs["oKActions"] = args ? args.oKActions : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["actionsEnabled"] = undefined /*out*/;
            resourceInputs["actionsSuppressor"] = undefined /*out*/;
            resourceInputs["actionsSuppressorExtensionPeriod"] = undefined /*out*/;
            resourceInputs["actionsSuppressorWaitPeriod"] = undefined /*out*/;
            resourceInputs["alarmActions"] = undefined /*out*/;
            resourceInputs["alarmDescription"] = undefined /*out*/;
            resourceInputs["alarmName"] = undefined /*out*/;
            resourceInputs["alarmRule"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["insufficientDataActions"] = undefined /*out*/;
            resourceInputs["oKActions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CompositeAlarm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CompositeAlarm resource.
 */
export interface CompositeAlarmArgs {
    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    actionsEnabled?: pulumi.Input<boolean>;
    /**
     * Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. 
     */
    actionsSuppressor?: pulumi.Input<string>;
    /**
     * Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.
     */
    actionsSuppressorExtensionPeriod?: pulumi.Input<number>;
    /**
     * Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.
     */
    actionsSuppressorWaitPeriod?: pulumi.Input<number>;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
     */
    alarmActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the alarm
     */
    alarmDescription?: pulumi.Input<string>;
    /**
     * The name of the Composite Alarm
     */
    alarmName: pulumi.Input<string>;
    /**
     * Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
     */
    alarmRule: pulumi.Input<string>;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    insufficientDataActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    oKActions?: pulumi.Input<pulumi.Input<string>[]>;
}