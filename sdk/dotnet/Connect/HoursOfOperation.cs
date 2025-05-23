// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Resource Type definition for AWS::Connect::HoursOfOperation
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hoursOfOperation = new AwsNative.Connect.HoursOfOperation("hoursOfOperation", new()
    ///     {
    ///         Name = "ExampleHoursOfOperation",
    ///         Description = "hours of operation created using cfn",
    ///         InstanceArn = "arn:aws:connect:region-name:aws-account-id:instance/instance-arn",
    ///         TimeZone = "Pacific/Midway",
    ///         Config = new[]
    ///         {
    ///             new AwsNative.Connect.Inputs.HoursOfOperationConfigArgs
    ///             {
    ///                 Day = AwsNative.Connect.HoursOfOperationConfigDay.Sunday,
    ///                 EndTime = new AwsNative.Connect.Inputs.HoursOfOperationTimeSliceArgs
    ///                 {
    ///                     Hours = 11,
    ///                     Minutes = 59,
    ///                 },
    ///                 StartTime = new AwsNative.Connect.Inputs.HoursOfOperationTimeSliceArgs
    ///                 {
    ///                     Hours = 10,
    ///                     Minutes = 1,
    ///                 },
    ///             },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "tagKey",
    ///                 Value = "tagValue",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:HoursOfOperation")]
    public partial class HoursOfOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration information for the hours of operation: day, start time, and end time.
        /// </summary>
        [Output("config")]
        public Output<ImmutableArray<Outputs.HoursOfOperationConfig>> Config { get; private set; } = null!;

        /// <summary>
        /// The description of the hours of operation.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the hours of operation.
        /// </summary>
        [Output("hoursOfOperationArn")]
        public Output<string> HoursOfOperationArn { get; private set; } = null!;

        /// <summary>
        /// One or more hours of operation overrides assigned to an hour of operation.
        /// </summary>
        [Output("hoursOfOperationOverrides")]
        public Output<ImmutableArray<Outputs.HoursOfOperationOverride>> HoursOfOperationOverrides { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the hours of operation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The time zone of the hours of operation.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a HoursOfOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HoursOfOperation(string name, HoursOfOperationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:HoursOfOperation", name, args ?? new HoursOfOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HoursOfOperation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:HoursOfOperation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HoursOfOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HoursOfOperation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HoursOfOperation(name, id, options);
        }
    }

    public sealed class HoursOfOperationArgs : global::Pulumi.ResourceArgs
    {
        [Input("config", required: true)]
        private InputList<Inputs.HoursOfOperationConfigArgs>? _config;

        /// <summary>
        /// Configuration information for the hours of operation: day, start time, and end time.
        /// </summary>
        public InputList<Inputs.HoursOfOperationConfigArgs> Config
        {
            get => _config ?? (_config = new InputList<Inputs.HoursOfOperationConfigArgs>());
            set => _config = value;
        }

        /// <summary>
        /// The description of the hours of operation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hoursOfOperationOverrides")]
        private InputList<Inputs.HoursOfOperationOverrideArgs>? _hoursOfOperationOverrides;

        /// <summary>
        /// One or more hours of operation overrides assigned to an hour of operation.
        /// </summary>
        public InputList<Inputs.HoursOfOperationOverrideArgs> HoursOfOperationOverrides
        {
            get => _hoursOfOperationOverrides ?? (_hoursOfOperationOverrides = new InputList<Inputs.HoursOfOperationOverrideArgs>());
            set => _hoursOfOperationOverrides = value;
        }

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        /// <summary>
        /// The name of the hours of operation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// One or more tags.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The time zone of the hours of operation.
        /// </summary>
        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public HoursOfOperationArgs()
        {
        }
        public static new HoursOfOperationArgs Empty => new HoursOfOperationArgs();
    }
}
