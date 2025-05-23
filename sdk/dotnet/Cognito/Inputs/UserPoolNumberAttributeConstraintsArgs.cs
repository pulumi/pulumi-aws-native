// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolNumberAttributeConstraintsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum length of a number attribute value. Must be a number less than or equal to `2^1023` , represented as a string with a length of 131072 characters or fewer.
        /// </summary>
        [Input("maxValue")]
        public Input<string>? MaxValue { get; set; }

        /// <summary>
        /// The minimum value of an attribute that is of the number data type.
        /// </summary>
        [Input("minValue")]
        public Input<string>? MinValue { get; set; }

        public UserPoolNumberAttributeConstraintsArgs()
        {
        }
        public static new UserPoolNumberAttributeConstraintsArgs Empty => new UserPoolNumberAttributeConstraintsArgs();
    }
}
