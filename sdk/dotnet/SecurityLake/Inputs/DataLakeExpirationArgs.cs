// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    /// <summary>
    /// Provides data expiration details of Amazon Security Lake object.
    /// </summary>
    public sealed class DataLakeExpirationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days before data expires in the Amazon Security Lake object.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        public DataLakeExpirationArgs()
        {
        }
        public static new DataLakeExpirationArgs Empty => new DataLakeExpirationArgs();
    }
}
