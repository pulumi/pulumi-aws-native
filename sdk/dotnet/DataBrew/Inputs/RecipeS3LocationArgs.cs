// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// Input location
    /// </summary>
    public sealed class RecipeS3LocationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("key")]
        public Input<string>? Key { get; set; }

        public RecipeS3LocationArgs()
        {
        }
        public static new RecipeS3LocationArgs Empty => new RecipeS3LocationArgs();
    }
}