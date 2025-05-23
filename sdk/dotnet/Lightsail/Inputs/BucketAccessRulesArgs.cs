// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// An object that sets the public accessibility of objects in the specified bucket.
    /// </summary>
    public sealed class BucketAccessRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.
        /// </summary>
        [Input("allowPublicOverrides")]
        public Input<bool>? AllowPublicOverrides { get; set; }

        /// <summary>
        /// Specifies the anonymous access to all objects in a bucket.
        /// </summary>
        [Input("getObject")]
        public Input<string>? GetObject { get; set; }

        public BucketAccessRulesArgs()
        {
        }
        public static new BucketAccessRulesArgs Empty => new BucketAccessRulesArgs();
    }
}
