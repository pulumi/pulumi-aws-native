// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class SpaceAppLifecycleManagement
    {
        public readonly Outputs.SpaceIdleSettings? IdleSettings;

        [OutputConstructor]
        private SpaceAppLifecycleManagement(Outputs.SpaceIdleSettings? idleSettings)
        {
            IdleSettings = idleSettings;
        }
    }
}