// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Information about level three.
    /// </summary>
    [OutputType]
    public sealed class UserHierarchyStructureLevelThree
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the hierarchy level.
        /// </summary>
        public readonly string? HierarchyLevelArn;
        public readonly string? HierarchyLevelId;
        /// <summary>
        /// The name of the hierarchy level.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private UserHierarchyStructureLevelThree(
            string? hierarchyLevelArn,

            string? hierarchyLevelId,

            string name)
        {
            HierarchyLevelArn = hierarchyLevelArn;
            HierarchyLevelId = hierarchyLevelId;
            Name = name;
        }
    }
}