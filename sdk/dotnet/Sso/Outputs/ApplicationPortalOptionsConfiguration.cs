// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sso.Outputs
{

    /// <summary>
    /// A structure that describes the options for the access portal associated with an application
    /// </summary>
    [OutputType]
    public sealed class ApplicationPortalOptionsConfiguration
    {
        /// <summary>
        /// A structure that describes the sign-in options for the access portal
        /// </summary>
        public readonly Outputs.ApplicationSignInOptions? SignInOptions;
        /// <summary>
        /// Indicates whether this application is visible in the access portal
        /// </summary>
        public readonly Pulumi.AwsNative.Sso.ApplicationPortalOptionsConfigurationVisibility? Visibility;

        [OutputConstructor]
        private ApplicationPortalOptionsConfiguration(
            Outputs.ApplicationSignInOptions? signInOptions,

            Pulumi.AwsNative.Sso.ApplicationPortalOptionsConfigurationVisibility? visibility)
        {
            SignInOptions = signInOptions;
            Visibility = visibility;
        }
    }
}
