// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkspacesInstances.Outputs
{

    [OutputType]
    public sealed class WorkspaceInstanceIamInstanceProfileSpecification
    {
        public readonly string? Name;

        [OutputConstructor]
        private WorkspaceInstanceIamInstanceProfileSpecification(string? name)
        {
            Name = name;
        }
    }
}
