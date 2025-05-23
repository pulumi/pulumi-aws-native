// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class Ec2FleetAcceleratorCountRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of accelerators. To specify no maximum limit, omit this parameter. To exclude accelerator-enabled instance types, set `Max` to `0` .
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum number of accelerators. To specify no minimum limit, omit this parameter.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public Ec2FleetAcceleratorCountRequestArgs()
        {
        }
        public static new Ec2FleetAcceleratorCountRequestArgs Empty => new Ec2FleetAcceleratorCountRequestArgs();
    }
}
