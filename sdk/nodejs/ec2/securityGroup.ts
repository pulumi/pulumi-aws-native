// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SecurityGroup
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * The group name or group ID depending on whether the SG is created in default or specific VPC
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * A description for the security group.
     */
    public readonly groupDescription!: pulumi.Output<string>;
    /**
     * The group ID of the specified security group.
     */
    public /*out*/ readonly groupId!: pulumi.Output<string>;
    /**
     * The name of the security group.
     */
    public readonly groupName!: pulumi.Output<string | undefined>;
    /**
     * [VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.
     */
    public readonly securityGroupEgress!: pulumi.Output<outputs.ec2.SecurityGroupEgress[] | undefined>;
    /**
     * The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.
     */
    public readonly securityGroupIngress!: pulumi.Output<outputs.ec2.SecurityGroupIngress[] | undefined>;
    /**
     * Any tags assigned to the security group.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The ID of the VPC for the security group.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupDescription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupDescription'");
            }
            resourceInputs["groupDescription"] = args ? args.groupDescription : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["securityGroupEgress"] = args ? args.securityGroupEgress : undefined;
            resourceInputs["securityGroupIngress"] = args ? args.securityGroupIngress : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["groupDescription"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
            resourceInputs["groupName"] = undefined /*out*/;
            resourceInputs["securityGroupEgress"] = undefined /*out*/;
            resourceInputs["securityGroupIngress"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["groupDescription", "groupName", "vpcId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * A description for the security group.
     */
    groupDescription: pulumi.Input<string>;
    /**
     * The name of the security group.
     */
    groupName?: pulumi.Input<string>;
    /**
     * [VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.
     */
    securityGroupEgress?: pulumi.Input<pulumi.Input<inputs.ec2.SecurityGroupEgressArgs>[]>;
    /**
     * The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.
     */
    securityGroupIngress?: pulumi.Input<pulumi.Input<inputs.ec2.SecurityGroupIngressArgs>[]>;
    /**
     * Any tags assigned to the security group.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The ID of the VPC for the security group.
     */
    vpcId?: pulumi.Input<string>;
}
