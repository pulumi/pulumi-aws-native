// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// One query argument in a web request, identified by name, for example UserName or SalesRegion. The name can be up to 30 characters long and isn't case sensitive.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupFieldToMatchSingleQueryArgumentProperties
    {
        public readonly string Name;

        [OutputConstructor]
        private RuleGroupFieldToMatchSingleQueryArgumentProperties(string name)
        {
            Name = name;
        }
    }
}
