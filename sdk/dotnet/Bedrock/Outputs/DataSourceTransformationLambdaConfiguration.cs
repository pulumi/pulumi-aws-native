// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// A Lambda function that processes documents.
    /// </summary>
    [OutputType]
    public sealed class DataSourceTransformationLambdaConfiguration
    {
        /// <summary>
        /// The function's ARN identifier.
        /// </summary>
        public readonly string LambdaArn;

        [OutputConstructor]
        private DataSourceTransformationLambdaConfiguration(string lambdaArn)
        {
            LambdaArn = lambdaArn;
        }
    }
}
