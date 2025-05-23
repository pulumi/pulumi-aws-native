// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    public static class GetRolePolicy
    {
        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role.
        ///  When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role, using [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html). You can update a role's trust policy using [UpdateAssumeRolePolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html). For information about roles, see [roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html) in the *IAM User Guide*.
        ///  A role can also have a managed policy attached to it. To attach a managed policy to a role, use [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
        ///  For information about the maximum number of inline policies that you can embed with a role, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Task<GetRolePolicyResult> InvokeAsync(GetRolePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRolePolicyResult>("aws-native:iam:getRolePolicy", args ?? new GetRolePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role.
        ///  When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role, using [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html). You can update a role's trust policy using [UpdateAssumeRolePolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html). For information about roles, see [roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html) in the *IAM User Guide*.
        ///  A role can also have a managed policy attached to it. To attach a managed policy to a role, use [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
        ///  For information about the maximum number of inline policies that you can embed with a role, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Output<GetRolePolicyResult> Invoke(GetRolePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolePolicyResult>("aws-native:iam:getRolePolicy", args ?? new GetRolePolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role.
        ///  When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role, using [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html). You can update a role's trust policy using [UpdateAssumeRolePolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html). For information about roles, see [roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html) in the *IAM User Guide*.
        ///  A role can also have a managed policy attached to it. To attach a managed policy to a role, use [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
        ///  For information about the maximum number of inline policies that you can embed with a role, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Output<GetRolePolicyResult> Invoke(GetRolePolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolePolicyResult>("aws-native:iam:getRolePolicy", args ?? new GetRolePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRolePolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy document.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the role to associate the policy with.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        public GetRolePolicyArgs()
        {
        }
        public static new GetRolePolicyArgs Empty => new GetRolePolicyArgs();
    }

    public sealed class GetRolePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy document.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the role to associate the policy with.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public GetRolePolicyInvokeArgs()
        {
        }
        public static new GetRolePolicyInvokeArgs Empty => new GetRolePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetRolePolicyResult
    {
        /// <summary>
        /// The policy document.
        ///  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
        ///  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
        ///   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
        ///   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
        ///   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::RolePolicy` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? PolicyDocument;

        [OutputConstructor]
        private GetRolePolicyResult(object? policyDocument)
        {
            PolicyDocument = policyDocument;
        }
    }
}
