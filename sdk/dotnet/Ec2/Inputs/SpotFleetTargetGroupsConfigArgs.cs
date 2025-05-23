// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class SpotFleetTargetGroupsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetGroups", required: true)]
        private InputList<Inputs.SpotFleetTargetGroupArgs>? _targetGroups;

        /// <summary>
        /// One or more target groups.
        /// </summary>
        public InputList<Inputs.SpotFleetTargetGroupArgs> TargetGroups
        {
            get => _targetGroups ?? (_targetGroups = new InputList<Inputs.SpotFleetTargetGroupArgs>());
            set => _targetGroups = value;
        }

        public SpotFleetTargetGroupsConfigArgs()
        {
        }
        public static new SpotFleetTargetGroupsConfigArgs Empty => new SpotFleetTargetGroupsConfigArgs();
    }
}
