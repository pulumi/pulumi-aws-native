// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Inputs
{

    /// <summary>
    /// CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.
    /// </summary>
    public sealed class TrailDataResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public TrailDataResourceArgs()
        {
        }
        public static new TrailDataResourceArgs Empty => new TrailDataResourceArgs();
    }
}
