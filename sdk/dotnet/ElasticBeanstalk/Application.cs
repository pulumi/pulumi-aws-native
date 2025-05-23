// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk
{
    /// <summary>
    /// The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
    /// </summary>
    [AwsNativeResourceType("aws-native:elasticbeanstalk:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
        /// </summary>
        [Output("applicationName")]
        public Output<string?> ApplicationName { get; private set; } = null!;

        /// <summary>
        /// Your description of the application.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
        /// </summary>
        [Output("resourceLifecycleConfig")]
        public Output<Outputs.ApplicationResourceLifecycleConfig?> ResourceLifecycleConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:elasticbeanstalk:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:elasticbeanstalk:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "applicationName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// Your description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
        /// </summary>
        [Input("resourceLifecycleConfig")]
        public Input<Inputs.ApplicationResourceLifecycleConfigArgs>? ResourceLifecycleConfig { get; set; }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
