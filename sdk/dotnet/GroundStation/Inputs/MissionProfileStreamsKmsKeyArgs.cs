// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class MissionProfileStreamsKmsKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// KMS Alias Arn.
        /// </summary>
        [Input("kmsAliasArn")]
        public Input<string>? KmsAliasArn { get; set; }

        /// <summary>
        /// KMS Alias Name.
        /// </summary>
        [Input("kmsAliasName")]
        public Input<string>? KmsAliasName { get; set; }

        /// <summary>
        /// KMS Key Arn.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public MissionProfileStreamsKmsKeyArgs()
        {
        }
        public static new MissionProfileStreamsKmsKeyArgs Empty => new MissionProfileStreamsKmsKeyArgs();
    }
}
