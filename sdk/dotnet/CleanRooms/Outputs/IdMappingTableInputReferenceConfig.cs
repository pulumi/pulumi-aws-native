// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class IdMappingTableInputReferenceConfig
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the referenced resource in AWS Entity Resolution . Valid values are ID mapping workflow ARNs.
        /// </summary>
        public readonly string InputReferenceArn;
        /// <summary>
        /// When `TRUE` , AWS Clean Rooms manages permissions for the ID mapping table resource.
        /// 
        /// When `FALSE` , the resource owner manages permissions for the ID mapping table resource.
        /// </summary>
        public readonly bool ManageResourcePolicies;

        [OutputConstructor]
        private IdMappingTableInputReferenceConfig(
            string inputReferenceArn,

            bool manageResourcePolicies)
        {
            InputReferenceArn = inputReferenceArn;
            ManageResourcePolicies = manageResourcePolicies;
        }
    }
}