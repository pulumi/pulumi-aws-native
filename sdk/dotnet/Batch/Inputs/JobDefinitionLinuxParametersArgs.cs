// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionLinuxParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.JobDefinitionDeviceArgs>? _devices;

        /// <summary>
        /// Any of the host devices to expose to the container. This parameter maps to `Devices` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--device` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
        /// 
        /// &gt; This parameter isn't applicable to jobs that are running on Fargate resources. Don't provide it for these jobs.
        /// </summary>
        public InputList<Inputs.JobDefinitionDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.JobDefinitionDeviceArgs>());
            set => _devices = value;
        }

        /// <summary>
        /// If true, run an `init` process inside the container that forwards signals and reaps processes. This parameter maps to the `--init` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version | grep "Server API version"`
        /// </summary>
        [Input("initProcessEnabled")]
        public Input<bool>? InitProcessEnabled { get; set; }

        /// <summary>
        /// The total amount of swap memory (in MiB) a container can use. This parameter is translated to the `--memory-swap` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) where the value is the sum of the container memory plus the `maxSwap` value. For more information, see [`--memory-swap` details](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/resource_constraints/#--memory-swap-details) in the Docker documentation.
        /// 
        /// If a `maxSwap` value of `0` is specified, the container doesn't use swap. Accepted values are `0` or any positive integer. If the `maxSwap` parameter is omitted, the container doesn't use the swap configuration for the container instance on which it runs. A `maxSwap` value must be set for the `swappiness` parameter to be used.
        /// 
        /// &gt; This parameter isn't applicable to jobs that are running on Fargate resources. Don't provide it for these jobs.
        /// </summary>
        [Input("maxSwap")]
        public Input<int>? MaxSwap { get; set; }

        /// <summary>
        /// The value for the size (in MiB) of the `/dev/shm` volume. This parameter maps to the `--shm-size` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
        /// 
        /// &gt; This parameter isn't applicable to jobs that are running on Fargate resources. Don't provide it for these jobs.
        /// </summary>
        [Input("sharedMemorySize")]
        public Input<int>? SharedMemorySize { get; set; }

        /// <summary>
        /// You can use this parameter to tune a container's memory swappiness behavior. A `swappiness` value of `0` causes swapping to not occur unless absolutely necessary. A `swappiness` value of `100` causes pages to be swapped aggressively. Valid values are whole numbers between `0` and `100` . If the `swappiness` parameter isn't specified, a default value of `60` is used. If a value isn't specified for `maxSwap` , then this parameter is ignored. If `maxSwap` is set to 0, the container doesn't use swap. This parameter maps to the `--memory-swappiness` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
        /// 
        /// Consider the following when you use a per-container swap configuration.
        /// 
        /// - Swap space must be enabled and allocated on the container instance for the containers to use.
        /// 
        /// &gt; By default, the Amazon ECS optimized AMIs don't have swap enabled. You must enable swap on the instance to use this feature. For more information, see [Instance store swap volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-store-swap-volumes.html) in the *Amazon EC2 User Guide for Linux Instances* or [How do I allocate memory to work as swap space in an Amazon EC2 instance by using a swap file?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/ec2-memory-swap-file/)
        /// - The swap space parameters are only supported for job definitions using EC2 resources.
        /// - If the `maxSwap` and `swappiness` parameters are omitted from a job definition, each container has a default `swappiness` value of 60. Moreover, the total swap usage is limited to two times the memory reservation of the container.
        /// 
        /// &gt; This parameter isn't applicable to jobs that are running on Fargate resources. Don't provide it for these jobs.
        /// </summary>
        [Input("swappiness")]
        public Input<int>? Swappiness { get; set; }

        [Input("tmpfs")]
        private InputList<Inputs.JobDefinitionTmpfsArgs>? _tmpfs;

        /// <summary>
        /// The container path, mount options, and size (in MiB) of the `tmpfs` mount. This parameter maps to the `--tmpfs` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
        /// 
        /// &gt; This parameter isn't applicable to jobs that are running on Fargate resources. Don't provide this parameter for this resource type.
        /// </summary>
        public InputList<Inputs.JobDefinitionTmpfsArgs> Tmpfs
        {
            get => _tmpfs ?? (_tmpfs = new InputList<Inputs.JobDefinitionTmpfsArgs>());
            set => _tmpfs = value;
        }

        public JobDefinitionLinuxParametersArgs()
        {
        }
        public static new JobDefinitionLinuxParametersArgs Empty => new JobDefinitionLinuxParametersArgs();
    }
}
