// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelArchiveS3Settings
    {
        public readonly string? CannedAcl;

        [OutputConstructor]
        private ChannelArchiveS3Settings(string? cannedAcl)
        {
            CannedAcl = cannedAcl;
        }
    }
}