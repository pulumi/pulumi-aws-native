// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Adds or updates an inline policy document that is embedded in the specified IAM group.
 *  A group can also have managed policies attached to it. To attach a managed policy to a group, use [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
 *  For information about the maximum number of inline policies that you can embed in a group, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
 */
export class GroupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GroupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GroupPolicy {
        return new GroupPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:GroupPolicy';

    /**
     * Returns true if the given object is an instance of GroupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupPolicy.__pulumiType;
    }

    /**
     * The name of the group to associate the policy with.
     *  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The policy document.
     *  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
     *  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
     *   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
     *   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
     *   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
     */
    public readonly policyDocument!: pulumi.Output<any | undefined>;
    /**
     * The name of the policy document.
     *  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
     */
    public readonly policyName!: pulumi.Output<string>;

    /**
     * Create a GroupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
        } else {
            resourceInputs["groupName"] = undefined /*out*/;
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["policyName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["groupName", "policyName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(GroupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GroupPolicy resource.
 */
export interface GroupPolicyArgs {
    /**
     * The name of the group to associate the policy with.
     *  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
     */
    groupName: pulumi.Input<string>;
    /**
     * The policy document.
     *  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
     *  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
     *   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
     *   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
     *   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
     */
    policyDocument?: any;
    /**
     * The name of the policy document.
     *  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
     */
    policyName?: pulumi.Input<string>;
}
