// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for SSO PermissionSet
 */
export function getPermissionSet(args: GetPermissionSetArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sso:getPermissionSet", {
        "instanceArn": args.instanceArn,
        "permissionSetArn": args.permissionSetArn,
    }, opts);
}

export interface GetPermissionSetArgs {
    /**
     * The sso instance arn that the permission set is owned.
     */
    instanceArn: string;
    /**
     * The permission set that the policy will be attached to
     */
    permissionSetArn: string;
}

export interface GetPermissionSetResult {
    /**
     * Specifies the names and paths of the customer managed policies that you have attached to your permission set.
     */
    readonly customerManagedPolicyReferences?: outputs.sso.PermissionSetCustomerManagedPolicyReference[];
    /**
     * The permission set description.
     */
    readonly description?: string;
    /**
     * The inline policy to put in permission set.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SSO::PermissionSet` for more information about the expected schema for this property.
     */
    readonly inlinePolicy?: any;
    /**
     * A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
     */
    readonly managedPolicies?: string[];
    /**
     * The permission set that the policy will be attached to
     */
    readonly permissionSetArn?: string;
    /**
     * Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary. Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
     *
     * > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
     */
    readonly permissionsBoundary?: outputs.sso.PermissionSetPermissionsBoundary;
    /**
     * The relay state URL that redirect links to any service in the AWS Management Console.
     */
    readonly relayStateType?: string;
    /**
     * The length of time that a user can be signed in to an AWS account.
     */
    readonly sessionDuration?: string;
    /**
     * The tags to attach to the new `PermissionSet` .
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource Type definition for SSO PermissionSet
 */
export function getPermissionSetOutput(args: GetPermissionSetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPermissionSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:sso:getPermissionSet", {
        "instanceArn": args.instanceArn,
        "permissionSetArn": args.permissionSetArn,
    }, opts);
}

export interface GetPermissionSetOutputArgs {
    /**
     * The sso instance arn that the permission set is owned.
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The permission set that the policy will be attached to
     */
    permissionSetArn: pulumi.Input<string>;
}
