// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    /// <summary>
    /// An object representing a remote access configuration specification for AWS EKS Nodegroup.
    /// </summary>
    public sealed class NodegroupRemoteAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon EC2 SSH key name that provides access for SSH communication with the nodes in the managed node group. For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Linux Instances* . For Windows, an Amazon EC2 SSH key is used to obtain the RDP password. For more information, see [Amazon EC2 key pairs and Windows instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Windows Instances* .
        /// </summary>
        [Input("ec2SshKey", required: true)]
        public Input<string> Ec2SshKey { get; set; } = null!;

        [Input("sourceSecurityGroups")]
        private InputList<string>? _sourceSecurityGroups;

        /// <summary>
        /// The security group IDs that are allowed SSH access (port 22) to the nodes. For Windows, the port is 3389. If you specify an Amazon EC2 SSH key but don't specify a source security group when you create a managed node group, then the port on the nodes is opened to the internet ( `0.0.0.0/0` ). For more information, see [Security Groups for Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
        /// </summary>
        public InputList<string> SourceSecurityGroups
        {
            get => _sourceSecurityGroups ?? (_sourceSecurityGroups = new InputList<string>());
            set => _sourceSecurityGroups = value;
        }

        public NodegroupRemoteAccessArgs()
        {
        }
        public static new NodegroupRemoteAccessArgs Empty => new NodegroupRemoteAccessArgs();
    }
}
