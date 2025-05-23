// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Information about the hierarchy structure.
    /// </summary>
    [OutputType]
    public sealed class UserHierarchyStructureProperties
    {
        public readonly Outputs.UserHierarchyStructureLevelFive? LevelFive;
        /// <summary>
        /// The update for level four.
        /// </summary>
        public readonly Outputs.UserHierarchyStructureLevelFour? LevelFour;
        /// <summary>
        /// The update for level one.
        /// </summary>
        public readonly Outputs.UserHierarchyStructureLevelOne? LevelOne;
        /// <summary>
        /// The update for level three.
        /// </summary>
        public readonly Outputs.UserHierarchyStructureLevelThree? LevelThree;
        /// <summary>
        /// The update for level two.
        /// </summary>
        public readonly Outputs.UserHierarchyStructureLevelTwo? LevelTwo;

        [OutputConstructor]
        private UserHierarchyStructureProperties(
            Outputs.UserHierarchyStructureLevelFive? levelFive,

            Outputs.UserHierarchyStructureLevelFour? levelFour,

            Outputs.UserHierarchyStructureLevelOne? levelOne,

            Outputs.UserHierarchyStructureLevelThree? levelThree,

            Outputs.UserHierarchyStructureLevelTwo? levelTwo)
        {
            LevelFive = levelFive;
            LevelFour = levelFour;
            LevelOne = levelOne;
            LevelThree = levelThree;
            LevelTwo = levelTwo;
        }
    }
}
