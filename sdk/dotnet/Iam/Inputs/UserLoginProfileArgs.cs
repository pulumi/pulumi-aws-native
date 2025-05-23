// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam.Inputs
{

    /// <summary>
    /// Creates a password for the specified user, giving the user the ability to access AWS services through the console. For more information about managing passwords, see [Managing Passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *User Guide*.
    /// </summary>
    public sealed class UserLoginProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user's password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Specifies whether the user is required to set a new password on next sign-in.
        /// </summary>
        [Input("passwordResetRequired")]
        public Input<bool>? PasswordResetRequired { get; set; }

        public UserLoginProfileArgs()
        {
        }
        public static new UserLoginProfileArgs Empty => new UserLoginProfileArgs();
    }
}
