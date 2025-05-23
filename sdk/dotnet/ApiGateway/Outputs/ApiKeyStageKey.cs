// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    /// <summary>
    /// ``StageKey`` is a property of the [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.
    /// </summary>
    [OutputType]
    public sealed class ApiKeyStageKey
    {
        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        public readonly string? RestApiId;
        /// <summary>
        /// The stage name associated with the stage key.
        /// </summary>
        public readonly string? StageName;

        [OutputConstructor]
        private ApiKeyStageKey(
            string? restApiId,

            string? stageName)
        {
            RestApiId = restApiId;
            StageName = stageName;
        }
    }
}
