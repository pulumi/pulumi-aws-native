// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class CollaborationMlMemberAbilities
    {
        /// <summary>
        /// The custom ML member abilities for a collaboration member.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.CleanRooms.CollaborationCustomMlMemberAbility> CustomMlMemberAbilities;

        [OutputConstructor]
        private CollaborationMlMemberAbilities(ImmutableArray<Pulumi.AwsNative.CleanRooms.CollaborationCustomMlMemberAbility> customMlMemberAbilities)
        {
            CustomMlMemberAbilities = customMlMemberAbilities;
        }
    }
}
