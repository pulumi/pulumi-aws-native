// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterStepConfig
    {
        public readonly string? ActionOnFailure;
        public readonly Outputs.ClusterHadoopJarStepConfig HadoopJarStep;
        public readonly string Name;

        [OutputConstructor]
        private ClusterStepConfig(
            string? actionOnFailure,

            Outputs.ClusterHadoopJarStepConfig hadoopJarStep,

            string name)
        {
            ActionOnFailure = actionOnFailure;
            HadoopJarStep = hadoopJarStep;
            Name = name;
        }
    }
}