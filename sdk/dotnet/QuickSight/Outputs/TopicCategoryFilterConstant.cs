// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicCategoryFilterConstant
    {
        /// <summary>
        /// A collective constant used in a category filter. This element is used to specify a list of values for the constant.
        /// </summary>
        public readonly Outputs.TopicCollectiveConstant? CollectiveConstant;
        /// <summary>
        /// The type of category filter constant. This element is used to specify whether a constant is a singular or collective. Valid values are `SINGULAR` and `COLLECTIVE` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TopicConstantType? ConstantType;
        /// <summary>
        /// A singular constant used in a category filter. This element is used to specify a single value for the constant.
        /// </summary>
        public readonly string? SingularConstant;

        [OutputConstructor]
        private TopicCategoryFilterConstant(
            Outputs.TopicCollectiveConstant? collectiveConstant,

            Pulumi.AwsNative.QuickSight.TopicConstantType? constantType,

            string? singularConstant)
        {
            CollectiveConstant = collectiveConstant;
            ConstantType = constantType;
            SingularConstant = singularConstant;
        }
    }
}
