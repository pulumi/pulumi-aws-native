// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceAclConfiguration
    {
        /// <summary>
        /// A list of groups, separated by semi-colons, that filters a query response based on user context. The document is only returned to users that are in one of the groups specified in the `UserContext` field of the [Query](https://docs.aws.amazon.com/kendra/latest/dg/API_Query.html) operation.
        /// </summary>
        public readonly string AllowedGroupsColumnName;

        [OutputConstructor]
        private DataSourceAclConfiguration(string allowedGroupsColumnName)
        {
            AllowedGroupsColumnName = allowedGroupsColumnName;
        }
    }
}
