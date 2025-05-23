// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sso
{
    public static class GetApplication
    {
        /// <summary>
        /// Resource Type definition for Identity Center (SSO) Application
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws-native:sso:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Identity Center (SSO) Application
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:sso:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Identity Center (SSO) Application
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:sso:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Application ARN that is returned upon creation of the Identity Center (SSO) Application
        /// </summary>
        [Input("applicationArn", required: true)]
        public string ApplicationArn { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Application ARN that is returned upon creation of the Identity Center (SSO) Application
        /// </summary>
        [Input("applicationArn", required: true)]
        public Input<string> ApplicationArn { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The Application ARN that is returned upon creation of the Identity Center (SSO) Application
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// The description information for the Identity Center (SSO) Application
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name you want to assign to this Identity Center (SSO) Application
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A structure that describes the options for the portal associated with an application
        /// </summary>
        public readonly Outputs.ApplicationPortalOptionsConfiguration? PortalOptions;
        /// <summary>
        /// Specifies whether the application is enabled or disabled
        /// </summary>
        public readonly Pulumi.AwsNative.Sso.ApplicationStatus? Status;
        /// <summary>
        /// Specifies tags to be attached to the application
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationArn,

            string? description,

            string? name,

            Outputs.ApplicationPortalOptionsConfiguration? portalOptions,

            Pulumi.AwsNative.Sso.ApplicationStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            ApplicationArn = applicationArn;
            Description = description;
            Name = name;
            PortalOptions = portalOptions;
            Status = status;
            Tags = tags;
        }
    }
}
