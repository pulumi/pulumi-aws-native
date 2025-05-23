// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The Linux capabilities to add or remove from the default Docker configuration for a container defined in the task definition. For more detailed information about these Linux capabilities, see the [capabilities(7)](https://docs.aws.amazon.com/http://man7.org/linux/man-pages/man7/capabilities.7.html) Linux manual page.
    ///  The following describes how Docker processes the Linux capabilities specified in the ``add`` and ``drop`` request parameters. For information about the latest behavior, see [Docker Compose: order of cap_drop and cap_add](https://docs.aws.amazon.com/https://forums.docker.com/t/docker-compose-order-of-cap-drop-and-cap-add/97136/1) in the Docker Community Forum.
    ///   +  When the container is a privleged container, the container capabilities are all of the default Docker capabilities. The capabilities specified in the ``add`` request parameter, and the ``drop`` request parameter are ignored.
    ///   +  When the ``add`` request parameter is set to ALL, the container capabilities are all of the default Docker capabilities, excluding those specified in the ``drop`` request parameter.
    ///   +  When the ``drop`` request parameter is set to ALL, the container capabilities are the capabilities specified in the ``add`` request parameter.
    ///   +  When the ``add`` request parameter and the ``drop`` request parameter are both empty, the capabilities the container capabilities are all of the default Docker capabilities.
    ///   +  The default is to first drop the capabilities specified in the ``drop`` request parameter, and then add the capabilities specified in the ``add`` request parameter.
    /// </summary>
    public sealed class TaskDefinitionKernelCapabilitiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("add")]
        private InputList<string>? _add;

        /// <summary>
        /// The Linux capabilities for the container that have been added to the default configuration provided by Docker. This parameter maps to ``CapAdd`` in the docker container create command and the ``--cap-add`` option to docker run.
        ///   Tasks launched on FARGATElong only support adding the ``SYS_PTRACE`` kernel capability.
        ///   Valid values: ``"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"``
        /// </summary>
        public InputList<string> Add
        {
            get => _add ?? (_add = new InputList<string>());
            set => _add = value;
        }

        [Input("drop")]
        private InputList<string>? _drop;

        /// <summary>
        /// The Linux capabilities for the container that have been removed from the default configuration provided by Docker. This parameter maps to ``CapDrop`` in the docker container create command and the ``--cap-drop`` option to docker run.
        ///  Valid values: ``"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"``
        /// </summary>
        public InputList<string> Drop
        {
            get => _drop ?? (_drop = new InputList<string>());
            set => _drop = value;
        }

        public TaskDefinitionKernelCapabilitiesArgs()
        {
        }
        public static new TaskDefinitionKernelCapabilitiesArgs Empty => new TaskDefinitionKernelCapabilitiesArgs();
    }
}
