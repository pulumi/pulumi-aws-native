// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    public sealed class RestApiEndpointConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("types")]
        private InputList<string>? _types;
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        [Input("vpcEndpointIds")]
        private InputList<string>? _vpcEndpointIds;
        public InputList<string> VpcEndpointIds
        {
            get => _vpcEndpointIds ?? (_vpcEndpointIds = new InputList<string>());
            set => _vpcEndpointIds = value;
        }

        public RestApiEndpointConfigurationArgs()
        {
        }
        public static new RestApiEndpointConfigurationArgs Empty => new RestApiEndpointConfigurationArgs();
    }
}