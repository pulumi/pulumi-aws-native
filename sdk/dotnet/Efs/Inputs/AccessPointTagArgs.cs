// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Inputs
{

    /// <summary>
    /// A tag is a key-value pair attached to a file system. Allowed characters in the ``Key`` and ``Value`` properties are letters, white space, and numbers that can be represented in UTF-8, and the following characters:``+ - = . _ : /``
    /// </summary>
    public sealed class AccessPointTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tag key (String). The key can't start with ``aws:``.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value of the tag key.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public AccessPointTagArgs()
        {
        }
        public static new AccessPointTagArgs Empty => new AccessPointTagArgs();
    }
}