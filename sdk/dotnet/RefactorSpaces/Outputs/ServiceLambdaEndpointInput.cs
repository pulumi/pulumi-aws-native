// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RefactorSpaces.Outputs
{

    [OutputType]
    public sealed class ServiceLambdaEndpointInput
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lambda function or alias.
        /// </summary>
        public readonly string Arn;

        [OutputConstructor]
        private ServiceLambdaEndpointInput(string arn)
        {
            Arn = arn;
        }
    }
}
