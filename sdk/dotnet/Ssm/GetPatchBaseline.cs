// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm
{
    public static class GetPatchBaseline
    {
        /// <summary>
        /// Resource Type definition for AWS::SSM::PatchBaseline
        /// </summary>
        public static Task<GetPatchBaselineResult> InvokeAsync(GetPatchBaselineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPatchBaselineResult>("aws-native:ssm:getPatchBaseline", args ?? new GetPatchBaselineArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SSM::PatchBaseline
        /// </summary>
        public static Output<GetPatchBaselineResult> Invoke(GetPatchBaselineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPatchBaselineResult>("aws-native:ssm:getPatchBaseline", args ?? new GetPatchBaselineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPatchBaselineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the patch baseline.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPatchBaselineArgs()
        {
        }
        public static new GetPatchBaselineArgs Empty => new GetPatchBaselineArgs();
    }

    public sealed class GetPatchBaselineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the patch baseline.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPatchBaselineInvokeArgs()
        {
        }
        public static new GetPatchBaselineInvokeArgs Empty => new GetPatchBaselineInvokeArgs();
    }


    [OutputType]
    public sealed class GetPatchBaselineResult
    {
        public readonly Outputs.PatchBaselineRuleGroup? ApprovalRules;
        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// </summary>
        public readonly ImmutableArray<string> ApprovedPatches;
        /// <summary>
        /// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
        /// </summary>
        public readonly Pulumi.AwsNative.Ssm.PatchBaselineApprovedPatchesComplianceLevel? ApprovedPatchesComplianceLevel;
        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
        /// </summary>
        public readonly bool? ApprovedPatchesEnableNonSecurity;
        /// <summary>
        /// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
        /// </summary>
        public readonly bool? DefaultBaseline;
        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A set of global filters used to include patches in the baseline.
        /// </summary>
        public readonly Outputs.PatchBaselinePatchFilterGroup? GlobalFilters;
        /// <summary>
        /// The ID of the patch baseline.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// PatchGroups is used to associate instances with a specific patch baseline
        /// </summary>
        public readonly ImmutableArray<string> PatchGroups;
        /// <summary>
        /// A list of explicitly rejected patches for the baseline.
        /// </summary>
        public readonly ImmutableArray<string> RejectedPatches;
        /// <summary>
        /// The action for Patch Manager to take on patches included in the RejectedPackages list.
        /// </summary>
        public readonly Pulumi.AwsNative.Ssm.PatchBaselineRejectedPatchesAction? RejectedPatchesAction;
        /// <summary>
        /// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
        /// </summary>
        public readonly ImmutableArray<Outputs.PatchBaselinePatchSource> Sources;
        /// <summary>
        /// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
        /// </summary>
        public readonly ImmutableArray<Outputs.PatchBaselineTag> Tags;

        [OutputConstructor]
        private GetPatchBaselineResult(
            Outputs.PatchBaselineRuleGroup? approvalRules,

            ImmutableArray<string> approvedPatches,

            Pulumi.AwsNative.Ssm.PatchBaselineApprovedPatchesComplianceLevel? approvedPatchesComplianceLevel,

            bool? approvedPatchesEnableNonSecurity,

            bool? defaultBaseline,

            string? description,

            Outputs.PatchBaselinePatchFilterGroup? globalFilters,

            string? id,

            string? name,

            ImmutableArray<string> patchGroups,

            ImmutableArray<string> rejectedPatches,

            Pulumi.AwsNative.Ssm.PatchBaselineRejectedPatchesAction? rejectedPatchesAction,

            ImmutableArray<Outputs.PatchBaselinePatchSource> sources,

            ImmutableArray<Outputs.PatchBaselineTag> tags)
        {
            ApprovalRules = approvalRules;
            ApprovedPatches = approvedPatches;
            ApprovedPatchesComplianceLevel = approvedPatchesComplianceLevel;
            ApprovedPatchesEnableNonSecurity = approvedPatchesEnableNonSecurity;
            DefaultBaseline = defaultBaseline;
            Description = description;
            GlobalFilters = globalFilters;
            Id = id;
            Name = name;
            PatchGroups = patchGroups;
            RejectedPatches = rejectedPatches;
            RejectedPatchesAction = rejectedPatchesAction;
            Sources = sources;
            Tags = tags;
        }
    }
}