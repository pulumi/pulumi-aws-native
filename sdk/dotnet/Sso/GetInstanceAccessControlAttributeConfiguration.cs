// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sso
{
    public static class GetInstanceAccessControlAttributeConfiguration
    {
        /// <summary>
        /// Resource Type definition for SSO InstanceAccessControlAttributeConfiguration
        /// </summary>
        public static Task<GetInstanceAccessControlAttributeConfigurationResult> InvokeAsync(GetInstanceAccessControlAttributeConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceAccessControlAttributeConfigurationResult>("aws-native:sso:getInstanceAccessControlAttributeConfiguration", args ?? new GetInstanceAccessControlAttributeConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for SSO InstanceAccessControlAttributeConfiguration
        /// </summary>
        public static Output<GetInstanceAccessControlAttributeConfigurationResult> Invoke(GetInstanceAccessControlAttributeConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceAccessControlAttributeConfigurationResult>("aws-native:sso:getInstanceAccessControlAttributeConfiguration", args ?? new GetInstanceAccessControlAttributeConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceAccessControlAttributeConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AWS SSO instance under which the operation will be executed.
        /// </summary>
        [Input("instanceArn", required: true)]
        public string InstanceArn { get; set; } = null!;

        public GetInstanceAccessControlAttributeConfigurationArgs()
        {
        }
        public static new GetInstanceAccessControlAttributeConfigurationArgs Empty => new GetInstanceAccessControlAttributeConfigurationArgs();
    }

    public sealed class GetInstanceAccessControlAttributeConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AWS SSO instance under which the operation will be executed.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        public GetInstanceAccessControlAttributeConfigurationInvokeArgs()
        {
        }
        public static new GetInstanceAccessControlAttributeConfigurationInvokeArgs Empty => new GetInstanceAccessControlAttributeConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceAccessControlAttributeConfigurationResult
    {
        public readonly ImmutableArray<Outputs.InstanceAccessControlAttributeConfigurationAccessControlAttribute> AccessControlAttributes;
        /// <summary>
        /// The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.
        /// </summary>
        public readonly Outputs.InstanceAccessControlAttributeConfigurationProperties? InstanceAccessControlAttributeConfigurationValue;

        [OutputConstructor]
        private GetInstanceAccessControlAttributeConfigurationResult(
            ImmutableArray<Outputs.InstanceAccessControlAttributeConfigurationAccessControlAttribute> accessControlAttributes,

            Outputs.InstanceAccessControlAttributeConfigurationProperties? instanceAccessControlAttributeConfiguration)
        {
            AccessControlAttributes = accessControlAttributes;
            InstanceAccessControlAttributeConfigurationValue = instanceAccessControlAttributeConfiguration;
        }
    }
}