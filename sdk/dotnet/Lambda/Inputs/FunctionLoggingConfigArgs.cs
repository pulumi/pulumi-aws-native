// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// The function's Amazon CloudWatch Logs configuration settings.
    /// </summary>
    public sealed class FunctionLoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set this property to filter the application logs for your function that Lambda sends to CloudWatch. Lambda only sends application logs at the selected level of detail and lower, where ``TRACE`` is the highest level and ``FATAL`` is the lowest.
        /// </summary>
        [Input("applicationLogLevel")]
        public Input<Pulumi.AwsNative.Lambda.FunctionLoggingConfigApplicationLogLevel>? ApplicationLogLevel { get; set; }

        /// <summary>
        /// The format in which Lambda sends your function's application and system logs to CloudWatch. Select between plain text and structured JSON.
        /// </summary>
        [Input("logFormat")]
        public Input<Pulumi.AwsNative.Lambda.FunctionLoggingConfigLogFormat>? LogFormat { get; set; }

        /// <summary>
        /// The name of the Amazon CloudWatch log group the function sends logs to. By default, Lambda functions send logs to a default log group named ``/aws/lambda/&lt;function name&gt;``. To use a different log group, enter an existing log group or enter a new log group name.
        /// </summary>
        [Input("logGroup")]
        public Input<string>? LogGroup { get; set; }

        /// <summary>
        /// Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
        /// </summary>
        [Input("systemLogLevel")]
        public Input<Pulumi.AwsNative.Lambda.FunctionLoggingConfigSystemLogLevel>? SystemLogLevel { get; set; }

        public FunctionLoggingConfigArgs()
        {
        }
        public static new FunctionLoggingConfigArgs Empty => new FunctionLoggingConfigArgs();
    }
}
