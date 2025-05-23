// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs.Outputs
{

    [OutputType]
    public sealed class TransformerRenameKeyEntry
    {
        public readonly string Key;
        public readonly bool? OverwriteIfExists;
        public readonly string RenameTo;

        [OutputConstructor]
        private TransformerRenameKeyEntry(
            string key,

            bool? overwriteIfExists,

            string renameTo)
        {
            Key = key;
            OverwriteIfExists = overwriteIfExists;
            RenameTo = renameTo;
        }
    }
}
