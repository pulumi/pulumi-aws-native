// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class MembershipProtectedQueryS3OutputConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 bucket to unload the protected query results.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The S3 prefix to unload the protected query results.
        /// </summary>
        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        /// <summary>
        /// Intended file format of the result.
        /// </summary>
        [Input("resultFormat", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.MembershipResultFormat> ResultFormat { get; set; } = null!;

        /// <summary>
        /// Indicates whether files should be output as a single file ( `TRUE` ) or output as multiple files ( `FALSE` ). This parameter is only supported for analyses with the Spark analytics engine.
        /// </summary>
        [Input("singleFileOutput")]
        public Input<bool>? SingleFileOutput { get; set; }

        public MembershipProtectedQueryS3OutputConfigurationArgs()
        {
        }
        public static new MembershipProtectedQueryS3OutputConfigurationArgs Empty => new MembershipProtectedQueryS3OutputConfigurationArgs();
    }
}
