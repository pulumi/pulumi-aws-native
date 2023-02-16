// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Outputs
{

    [OutputType]
    public sealed class FunctionDefinitionEnvironment
    {
        public readonly bool? AccessSysfs;
        public readonly Outputs.FunctionDefinitionExecution? Execution;
        public readonly ImmutableArray<Outputs.FunctionDefinitionResourceAccessPolicy> ResourceAccessPolicies;
        public readonly object? Variables;

        [OutputConstructor]
        private FunctionDefinitionEnvironment(
            bool? accessSysfs,

            Outputs.FunctionDefinitionExecution? execution,

            ImmutableArray<Outputs.FunctionDefinitionResourceAccessPolicy> resourceAccessPolicies,

            object? variables)
        {
            AccessSysfs = accessSysfs;
            Execution = execution;
            ResourceAccessPolicies = resourceAccessPolicies;
            Variables = variables;
        }
    }
}