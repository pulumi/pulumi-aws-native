// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class IdentityPoolCognitoStreamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        [Input("streamingStatus")]
        public Input<string>? StreamingStatus { get; set; }

        public IdentityPoolCognitoStreamsArgs()
        {
        }
        public static new IdentityPoolCognitoStreamsArgs Empty => new IdentityPoolCognitoStreamsArgs();
    }
}