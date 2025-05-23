// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLiftStreams
{
    /// <summary>
    /// Definition of AWS::GameLiftStreams::StreamGroup Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:gameliftstreams:StreamGroup")]
    public partial class StreamGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the stream group resource. For example: `arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4` .
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An ID that uniquely identifies the stream group resource. For example: `sg-1AB2C3De4` .
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Object that identifies the Amazon GameLift Streams application to stream with this stream group.
        /// </summary>
        [Output("defaultApplication")]
        public Output<Outputs.StreamGroupDefaultApplication?> DefaultApplication { get; private set; } = null!;

        /// <summary>
        /// A descriptive label for the stream group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A set of one or more locations and the streaming capacity for each location. One of the locations MUST be your primary location, which is the AWS Region where you are specifying this resource.
        /// </summary>
        [Output("locationConfigurations")]
        public Output<ImmutableArray<Outputs.StreamGroupLocationConfiguration>> LocationConfigurations { get; private set; } = null!;

        /// <summary>
        /// The target stream quality for sessions that are hosted in this stream group. Set a stream class that is appropriate to the type of content that you're streaming. Stream class determines the type of computing resources Amazon GameLift Streams uses and impacts the cost of streaming. The following options are available:
        /// 
        /// A stream class can be one of the following:
        /// 
        /// - *`gen5n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen5n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 4 vCPUs, 16 GB RAM, 12 GB VRAM
        /// - Tenancy: Supports up to 2 concurrent stream sessions
        /// - *`gen5n_ultra` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Uses dedicated NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen4n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen4n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 4 vCPUs, 16 GB RAM, 8 GB VRAM
        /// - Tenancy: Supports up to 2 concurrent stream sessions
        /// - *`gen4n_ultra` (NVIDIA, ultra)* Supports applications with high 3D scene complexity. Uses dedicated NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// </summary>
        [Output("streamClass")]
        public Output<string> StreamClass { get; private set; } = null!;

        /// <summary>
        /// A list of labels to assign to the new stream group resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a StreamGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamGroup(string name, StreamGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:gameliftstreams:StreamGroup", name, args ?? new StreamGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gameliftstreams:StreamGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "defaultApplication.id",
                    "streamClass",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamGroup(name, id, options);
        }
    }

    public sealed class StreamGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Object that identifies the Amazon GameLift Streams application to stream with this stream group.
        /// </summary>
        [Input("defaultApplication")]
        public Input<Inputs.StreamGroupDefaultApplicationArgs>? DefaultApplication { get; set; }

        /// <summary>
        /// A descriptive label for the stream group.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("locationConfigurations", required: true)]
        private InputList<Inputs.StreamGroupLocationConfigurationArgs>? _locationConfigurations;

        /// <summary>
        /// A set of one or more locations and the streaming capacity for each location. One of the locations MUST be your primary location, which is the AWS Region where you are specifying this resource.
        /// </summary>
        public InputList<Inputs.StreamGroupLocationConfigurationArgs> LocationConfigurations
        {
            get => _locationConfigurations ?? (_locationConfigurations = new InputList<Inputs.StreamGroupLocationConfigurationArgs>());
            set => _locationConfigurations = value;
        }

        /// <summary>
        /// The target stream quality for sessions that are hosted in this stream group. Set a stream class that is appropriate to the type of content that you're streaming. Stream class determines the type of computing resources Amazon GameLift Streams uses and impacts the cost of streaming. The following options are available:
        /// 
        /// A stream class can be one of the following:
        /// 
        /// - *`gen5n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen5n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 4 vCPUs, 16 GB RAM, 12 GB VRAM
        /// - Tenancy: Supports up to 2 concurrent stream sessions
        /// - *`gen5n_ultra` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Uses dedicated NVIDIA A10G Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen4n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// - *`gen4n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 4 vCPUs, 16 GB RAM, 8 GB VRAM
        /// - Tenancy: Supports up to 2 concurrent stream sessions
        /// - *`gen4n_ultra` (NVIDIA, ultra)* Supports applications with high 3D scene complexity. Uses dedicated NVIDIA T4 Tensor GPU.
        /// 
        /// - Reference resolution: 1080p
        /// - Reference frame rate: 60 fps
        /// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
        /// - Tenancy: Supports 1 concurrent stream session
        /// </summary>
        [Input("streamClass", required: true)]
        public Input<string> StreamClass { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of labels to assign to the new stream group resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public StreamGroupArgs()
        {
        }
        public static new StreamGroupArgs Empty => new StreamGroupArgs();
    }
}
