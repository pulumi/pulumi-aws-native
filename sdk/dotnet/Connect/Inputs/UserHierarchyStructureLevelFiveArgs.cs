// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// Information about level five.
    /// </summary>
    public sealed class UserHierarchyStructureLevelFiveArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the hierarchy level.
        /// </summary>
        [Input("hierarchyLevelArn")]
        public Input<string>? HierarchyLevelArn { get; set; }

        /// <summary>
        /// The identifier of the hierarchy level.
        /// </summary>
        [Input("hierarchyLevelId")]
        public Input<string>? HierarchyLevelId { get; set; }

        /// <summary>
        /// The name of the hierarchy level.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public UserHierarchyStructureLevelFiveArgs()
        {
        }
        public static new UserHierarchyStructureLevelFiveArgs Empty => new UserHierarchyStructureLevelFiveArgs();
    }
}
