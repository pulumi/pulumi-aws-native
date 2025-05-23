// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Outputs
{

    [OutputType]
    public sealed class FormSectionalElement
    {
        /// <summary>
        /// Excludes a sectional element that was generated by default for a specified data model.
        /// </summary>
        public readonly bool? Excluded;
        /// <summary>
        /// Specifies the size of the font for a `Heading` sectional element. Valid values are `1 | 2 | 3 | 4 | 5 | 6` .
        /// </summary>
        public readonly double? Level;
        /// <summary>
        /// Specifies the orientation for a `Divider` sectional element. Valid values are `horizontal` or `vertical` .
        /// </summary>
        public readonly string? Orientation;
        /// <summary>
        /// Specifies the position of the text in a field for a `Text` sectional element.
        /// </summary>
        public readonly object? Position;
        /// <summary>
        /// The text for a `Text` sectional element.
        /// </summary>
        public readonly string? Text;
        /// <summary>
        /// The type of sectional element. Valid values are `Heading` , `Text` , and `Divider` .
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FormSectionalElement(
            bool? excluded,

            double? level,

            string? orientation,

            object? position,

            string? text,

            string type)
        {
            Excluded = excluded;
            Level = level;
            Orientation = orientation;
            Position = position;
            Text = text;
            Type = type;
        }
    }
}
