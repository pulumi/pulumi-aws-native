// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk
{
    public static class GetApplication
    {
        /// <summary>
        /// The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws-native:elasticbeanstalk:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:elasticbeanstalk:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// Your description of the application.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
        /// </summary>
        public readonly Outputs.ApplicationResourceLifecycleConfig? ResourceLifecycleConfig;

        [OutputConstructor]
        private GetApplicationResult(
            string? description,

            Outputs.ApplicationResourceLifecycleConfig? resourceLifecycleConfig)
        {
            Description = description;
            ResourceLifecycleConfig = resourceLifecycleConfig;
        }
    }
}