// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    public static class GetUser
    {
        /// <summary>
        /// Creates a new IAM user for your AWS-account.
        ///   For information about quotas for the number of IAM users you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("aws-native:iam:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a new IAM user for your AWS-account.
        ///   For information about quotas for the number of IAM users you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("aws-native:iam:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a new IAM user for your AWS-account.
        ///   For information about quotas for the number of IAM users you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("aws-native:iam:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the user to create. Do not include the path in this value.
        ///  This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The user name must be unique within the account. User names are not distinguished by case. For example, you cannot create users named both "John" and "john".
        ///  If you don't specify a name, CFN generates a unique physical ID and uses that ID for the user name.
        ///  If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).
        ///   Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using ``Fn::Join`` and ``AWS::Region`` to create a Region-specific name, as in the following example: ``{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}``.
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the user to create. Do not include the path in this value.
        ///  This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The user name must be unique within the account. User names are not distinguished by case. For example, you cannot create users named both "John" and "john".
        ///  If you don't specify a name, CFN generates a unique physical ID and uses that ID for the user name.
        ///  If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).
        ///   Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using ``Fn::Join`` and ``AWS::Region`` to create a Region-specific name, as in the following example: ``{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref": "MyResourceName"}]]}``.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) for the specified `AWS::IAM::User` resource. For example: `arn:aws:iam::123456789012:user/mystack-myuser-1CCXAFG2H2U4D` .
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A list of group names to which you want to add the user.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// Creates a password for the specified IAM user. A password allows an IAM user to access AWS services through the console.
        ///  You can use the CLI, the AWS API, or the *Users* page in the IAM console to create a password for any IAM user. Use [ChangePassword](https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html) to update your own existing password in the *My Security Credentials* page in the console.
        ///  For more information about managing passwords, see [Managing passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *User Guide*.
        /// </summary>
        public readonly Outputs.UserLoginProfile? LoginProfile;
        /// <summary>
        /// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the user.
        ///  For more information about ARNs, see [Amazon Resource Names (ARNs) and Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.
        /// </summary>
        public readonly ImmutableArray<string> ManagedPolicyArns;
        /// <summary>
        /// The path for the user name. For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
        ///  This parameter is optional. If it is not included, it defaults to a slash (/).
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The ARN of the managed policy that is used to set the permissions boundary for the user.
        ///  A permissions boundary policy defines the maximum permissions that identity-based policies can grant to an entity, but does not grant permissions. Permissions boundaries do not define the maximum permissions that a resource-based policy can grant to an entity. To learn more, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide*.
        ///  For more information about policy types, see [Policy types](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types) in the *IAM User Guide*.
        /// </summary>
        public readonly string? PermissionsBoundary;
        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM user. To view AWS::IAM::User snippets, see [Declaring an User Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-user).
        ///   The name of each policy for a role, user, or group must be unique. If you don't choose unique names, updates to the IAM identity will fail. 
        ///   For information about limits on the number of inline policies that you can embed in a user, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPolicy> Policies;
        /// <summary>
        /// A list of tags that you want to attach to the new user. Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.
        ///   If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetUserResult(
            string? arn,

            ImmutableArray<string> groups,

            Outputs.UserLoginProfile? loginProfile,

            ImmutableArray<string> managedPolicyArns,

            string? path,

            string? permissionsBoundary,

            ImmutableArray<Outputs.UserPolicy> policies,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Groups = groups;
            LoginProfile = loginProfile;
            ManagedPolicyArns = managedPolicyArns;
            Path = path;
            PermissionsBoundary = permissionsBoundary;
            Policies = policies;
            Tags = tags;
        }
    }
}
