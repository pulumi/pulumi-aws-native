// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dlm
{
    public static class GetLifecyclePolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::DLM::LifecyclePolicy
        /// </summary>
        public static Task<GetLifecyclePolicyResult> InvokeAsync(GetLifecyclePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLifecyclePolicyResult>("aws-native:dlm:getLifecyclePolicy", args ?? new GetLifecyclePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DLM::LifecyclePolicy
        /// </summary>
        public static Output<GetLifecyclePolicyResult> Invoke(GetLifecyclePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLifecyclePolicyResult>("aws-native:dlm:getLifecyclePolicy", args ?? new GetLifecyclePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLifecyclePolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLifecyclePolicyArgs()
        {
        }
        public static new GetLifecyclePolicyArgs Empty => new GetLifecyclePolicyArgs();
    }

    public sealed class GetLifecyclePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLifecyclePolicyInvokeArgs()
        {
        }
        public static new GetLifecyclePolicyInvokeArgs Empty => new GetLifecyclePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetLifecyclePolicyResult
    {
        public readonly string? Arn;
        public readonly bool? CopyTags;
        public readonly int? CreateInterval;
        public readonly Outputs.LifecyclePolicyCrossRegionCopyTargets? CrossRegionCopyTargets;
        public readonly string? DefaultPolicy;
        public readonly string? Description;
        public readonly Outputs.LifecyclePolicyExclusions? Exclusions;
        public readonly string? ExecutionRoleArn;
        public readonly bool? ExtendDeletion;
        public readonly string? Id;
        public readonly Outputs.LifecyclePolicyPolicyDetails? PolicyDetails;
        public readonly int? RetainInterval;
        public readonly string? State;
        public readonly ImmutableArray<Outputs.LifecyclePolicyTag> Tags;

        [OutputConstructor]
        private GetLifecyclePolicyResult(
            string? arn,

            bool? copyTags,

            int? createInterval,

            Outputs.LifecyclePolicyCrossRegionCopyTargets? crossRegionCopyTargets,

            string? defaultPolicy,

            string? description,

            Outputs.LifecyclePolicyExclusions? exclusions,

            string? executionRoleArn,

            bool? extendDeletion,

            string? id,

            Outputs.LifecyclePolicyPolicyDetails? policyDetails,

            int? retainInterval,

            string? state,

            ImmutableArray<Outputs.LifecyclePolicyTag> tags)
        {
            Arn = arn;
            CopyTags = copyTags;
            CreateInterval = createInterval;
            CrossRegionCopyTargets = crossRegionCopyTargets;
            DefaultPolicy = defaultPolicy;
            Description = description;
            Exclusions = exclusions;
            ExecutionRoleArn = executionRoleArn;
            ExtendDeletion = extendDeletion;
            Id = id;
            PolicyDetails = policyDetails;
            RetainInterval = retainInterval;
            State = state;
            Tags = tags;
        }
    }
}