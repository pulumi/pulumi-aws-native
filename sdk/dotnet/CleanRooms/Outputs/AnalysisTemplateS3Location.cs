// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class AnalysisTemplateS3Location
    {
        public readonly string Bucket;
        public readonly string Key;

        [OutputConstructor]
        private AnalysisTemplateS3Location(
            string bucket,

            string key)
        {
            Bucket = bucket;
            Key = key;
        }
    }
}
