// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class DataSourceVpcConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of identifiers of security groups within your Amazon VPC. The security groups should enable Amazon Q Business to connect to the data source.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of identifiers for subnets within your Amazon VPC. The subnets should be able to connect to each other in the VPC, and they should have outgoing access to the Internet through a NAT device.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public DataSourceVpcConfigurationArgs()
        {
        }
        public static new DataSourceVpcConfigurationArgs Empty => new DataSourceVpcConfigurationArgs();
    }
}
