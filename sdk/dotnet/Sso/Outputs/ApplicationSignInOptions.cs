// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sso.Outputs
{

    /// <summary>
    /// A structure that describes the sign-in options for an application portal
    /// </summary>
    [OutputType]
    public sealed class ApplicationSignInOptions
    {
        /// <summary>
        /// The URL that accepts authentication requests for an application, this is a required parameter if the Origin parameter is APPLICATION
        /// </summary>
        public readonly string? ApplicationUrl;
        /// <summary>
        /// This determines how IAM Identity Center navigates the user to the target application
        /// </summary>
        public readonly Pulumi.AwsNative.Sso.ApplicationSignInOptionsOrigin Origin;

        [OutputConstructor]
        private ApplicationSignInOptions(
            string? applicationUrl,

            Pulumi.AwsNative.Sso.ApplicationSignInOptionsOrigin origin)
        {
            ApplicationUrl = applicationUrl;
            Origin = origin;
        }
    }
}