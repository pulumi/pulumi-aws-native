// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetUserHierarchyStructure
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::UserHierarchyStructure
        /// </summary>
        public static Task<GetUserHierarchyStructureResult> InvokeAsync(GetUserHierarchyStructureArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserHierarchyStructureResult>("aws-native:connect:getUserHierarchyStructure", args ?? new GetUserHierarchyStructureArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::UserHierarchyStructure
        /// </summary>
        public static Output<GetUserHierarchyStructureResult> Invoke(GetUserHierarchyStructureInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserHierarchyStructureResult>("aws-native:connect:getUserHierarchyStructure", args ?? new GetUserHierarchyStructureInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::UserHierarchyStructure
        /// </summary>
        public static Output<GetUserHierarchyStructureResult> Invoke(GetUserHierarchyStructureInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserHierarchyStructureResult>("aws-native:connect:getUserHierarchyStructure", args ?? new GetUserHierarchyStructureInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserHierarchyStructureArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the User Hierarchy Structure.
        /// </summary>
        [Input("userHierarchyStructureArn", required: true)]
        public string UserHierarchyStructureArn { get; set; } = null!;

        public GetUserHierarchyStructureArgs()
        {
        }
        public static new GetUserHierarchyStructureArgs Empty => new GetUserHierarchyStructureArgs();
    }

    public sealed class GetUserHierarchyStructureInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the User Hierarchy Structure.
        /// </summary>
        [Input("userHierarchyStructureArn", required: true)]
        public Input<string> UserHierarchyStructureArn { get; set; } = null!;

        public GetUserHierarchyStructureInvokeArgs()
        {
        }
        public static new GetUserHierarchyStructureInvokeArgs Empty => new GetUserHierarchyStructureInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserHierarchyStructureResult
    {
        /// <summary>
        /// Information about the hierarchy structure.
        /// </summary>
        public readonly Outputs.UserHierarchyStructureProperties? UserHierarchyStructureValue;
        /// <summary>
        /// The identifier of the User Hierarchy Structure.
        /// </summary>
        public readonly string? UserHierarchyStructureArn;

        [OutputConstructor]
        private GetUserHierarchyStructureResult(
            Outputs.UserHierarchyStructureProperties? userHierarchyStructure,

            string? userHierarchyStructureArn)
        {
            UserHierarchyStructureValue = userHierarchyStructure;
            UserHierarchyStructureArn = userHierarchyStructureArn;
        }
    }
}
