// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorksCM.Outputs
{

    [OutputType]
    public sealed class ServerEngineAttribute
    {
        public readonly string? Name;
        public readonly string? Value;

        [OutputConstructor]
        private ServerEngineAttribute(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}