// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms
{
    /// <summary>
    /// Represents a privacy budget within a collaboration
    /// </summary>
    [AwsNativeResourceType("aws-native:cleanrooms:PrivacyBudgetTemplate")]
    public partial class PrivacyBudgetTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the privacy budget template.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// How often the privacy budget refreshes.
        /// 
        /// &gt; If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.
        /// </summary>
        [Output("autoRefresh")]
        public Output<Pulumi.AwsNative.CleanRooms.PrivacyBudgetTemplateAutoRefresh> AutoRefresh { get; private set; } = null!;

        /// <summary>
        /// The ARN of the collaboration that contains this privacy budget template.
        /// </summary>
        [Output("collaborationArn")]
        public Output<string> CollaborationArn { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the collaboration that contains this privacy budget template.
        /// </summary>
        [Output("collaborationIdentifier")]
        public Output<string> CollaborationIdentifier { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the member who created the privacy budget template.
        /// </summary>
        [Output("membershipArn")]
        public Output<string> MembershipArn { get; private set; } = null!;

        /// <summary>
        /// The identifier for a membership resource.
        /// </summary>
        [Output("membershipIdentifier")]
        public Output<string> MembershipIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies the epislon and noise parameters for the privacy budget template.
        /// </summary>
        [Output("parameters")]
        public Output<Outputs.ParametersProperties> Parameters { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for one of your memberships for a collaboration. The privacy budget template is created in the collaboration that this membership belongs to. Accepts a membership ID.
        /// </summary>
        [Output("privacyBudgetTemplateIdentifier")]
        public Output<string> PrivacyBudgetTemplateIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of the privacy budget template.
        /// </summary>
        [Output("privacyBudgetType")]
        public Output<Pulumi.AwsNative.CleanRooms.PrivacyBudgetTemplatePrivacyBudgetType> PrivacyBudgetType { get; private set; } = null!;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a PrivacyBudgetTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivacyBudgetTemplate(string name, PrivacyBudgetTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cleanrooms:PrivacyBudgetTemplate", name, args ?? new PrivacyBudgetTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivacyBudgetTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cleanrooms:PrivacyBudgetTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "autoRefresh",
                    "membershipIdentifier",
                    "privacyBudgetType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivacyBudgetTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivacyBudgetTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivacyBudgetTemplate(name, id, options);
        }
    }

    public sealed class PrivacyBudgetTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How often the privacy budget refreshes.
        /// 
        /// &gt; If you plan to regularly bring new data into the collaboration, use `CALENDAR_MONTH` to automatically get a new privacy budget for the collaboration every calendar month. Choosing this option allows arbitrary amounts of information to be revealed about rows of the data when repeatedly queried across refreshes. Avoid choosing this if the same rows will be repeatedly queried between privacy budget refreshes.
        /// </summary>
        [Input("autoRefresh", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.PrivacyBudgetTemplateAutoRefresh> AutoRefresh { get; set; } = null!;

        /// <summary>
        /// The identifier for a membership resource.
        /// </summary>
        [Input("membershipIdentifier", required: true)]
        public Input<string> MembershipIdentifier { get; set; } = null!;

        /// <summary>
        /// Specifies the epislon and noise parameters for the privacy budget template.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<Inputs.ParametersPropertiesArgs> Parameters { get; set; } = null!;

        /// <summary>
        /// Specifies the type of the privacy budget template.
        /// </summary>
        [Input("privacyBudgetType", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.PrivacyBudgetTemplatePrivacyBudgetType> PrivacyBudgetType { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public PrivacyBudgetTemplateArgs()
        {
        }
        public static new PrivacyBudgetTemplateArgs Empty => new PrivacyBudgetTemplateArgs();
    }
}