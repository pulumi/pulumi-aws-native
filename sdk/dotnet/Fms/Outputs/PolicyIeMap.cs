// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fms.Outputs
{

    /// <summary>
    /// An FMS includeMap or excludeMap.
    /// </summary>
    [OutputType]
    public sealed class PolicyIeMap
    {
        public readonly ImmutableArray<string> Account;
        public readonly ImmutableArray<string> Orgunit;

        [OutputConstructor]
        private PolicyIeMap(
            ImmutableArray<string> account,

            ImmutableArray<string> orgunit)
        {
            Account = account;
            Orgunit = orgunit;
        }
    }
}