// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class MlTransformMlUserDataEncryptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("mlUserDataEncryptionMode", required: true)]
        public Input<string> MlUserDataEncryptionMode { get; set; } = null!;

        public MlTransformMlUserDataEncryptionArgs()
        {
        }
        public static new MlTransformMlUserDataEncryptionArgs Empty => new MlTransformMlUserDataEncryptionArgs();
    }
}