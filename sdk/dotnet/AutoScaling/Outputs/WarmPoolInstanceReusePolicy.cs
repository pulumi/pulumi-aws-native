// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class WarmPoolInstanceReusePolicy
    {
        /// <summary>
        /// Specifies whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
        /// </summary>
        public readonly bool? ReuseOnScaleIn;

        [OutputConstructor]
        private WarmPoolInstanceReusePolicy(bool? reuseOnScaleIn)
        {
            ReuseOnScaleIn = reuseOnScaleIn;
        }
    }
}
