// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::Monitor Resource Type
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:deadline:Monitor';

    /**
     * Returns true if the given object is an instance of Monitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Monitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Monitor.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the monitor.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the monitor that displays on the Deadline Cloud console.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) that the IAM Identity Center assigned to the monitor when it was created.
     */
    public /*out*/ readonly identityCenterApplicationArn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
     */
    public readonly identityCenterInstanceArn!: pulumi.Output<string>;
    /**
     * The unique identifier for the monitor.
     */
    public /*out*/ readonly monitorId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
     */
    public readonly subdomain!: pulumi.Output<string>;
    /**
     * The complete URL of the monitor. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.identityCenterInstanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityCenterInstanceArn'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.subdomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subdomain'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["identityCenterInstanceArn"] = args ? args.identityCenterInstanceArn : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["subdomain"] = args ? args.subdomain : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["identityCenterApplicationArn"] = undefined /*out*/;
            resourceInputs["monitorId"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["identityCenterApplicationArn"] = undefined /*out*/;
            resourceInputs["identityCenterInstanceArn"] = undefined /*out*/;
            resourceInputs["monitorId"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["subdomain"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["identityCenterInstanceArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Monitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * The name of the monitor that displays on the Deadline Cloud console.
     */
    displayName: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible for authenticating monitor users.
     */
    identityCenterInstanceArn: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the monitor use this role to access Deadline Cloud resources.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The subdomain used for the monitor URL. The full URL of the monitor is subdomain.Region.deadlinecloud.amazonaws.com.
     */
    subdomain: pulumi.Input<string>;
}