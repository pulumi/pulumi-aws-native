// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Comprehend
{
    public static class GetFlywheel
    {
        /// <summary>
        /// The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.
        /// </summary>
        public static Task<GetFlywheelResult> InvokeAsync(GetFlywheelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlywheelResult>("aws-native:comprehend:getFlywheel", args ?? new GetFlywheelArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.
        /// </summary>
        public static Output<GetFlywheelResult> Invoke(GetFlywheelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlywheelResult>("aws-native:comprehend:getFlywheel", args ?? new GetFlywheelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlywheelArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetFlywheelArgs()
        {
        }
        public static new GetFlywheelArgs Empty => new GetFlywheelArgs();
    }

    public sealed class GetFlywheelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetFlywheelInvokeArgs()
        {
        }
        public static new GetFlywheelInvokeArgs Empty => new GetFlywheelInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlywheelResult
    {
        public readonly string? ActiveModelArn;
        public readonly string? Arn;
        public readonly string? DataAccessRoleArn;
        public readonly Outputs.FlywheelDataSecurityConfig? DataSecurityConfig;
        public readonly ImmutableArray<Outputs.FlywheelTag> Tags;

        [OutputConstructor]
        private GetFlywheelResult(
            string? activeModelArn,

            string? arn,

            string? dataAccessRoleArn,

            Outputs.FlywheelDataSecurityConfig? dataSecurityConfig,

            ImmutableArray<Outputs.FlywheelTag> tags)
        {
            ActiveModelArn = activeModelArn;
            Arn = arn;
            DataAccessRoleArn = dataAccessRoleArn;
            DataSecurityConfig = dataSecurityConfig;
            Tags = tags;
        }
    }
}