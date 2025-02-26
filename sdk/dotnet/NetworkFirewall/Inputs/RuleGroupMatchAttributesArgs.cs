// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupMatchAttributesArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationPorts")]
        private InputList<Inputs.RuleGroupPortRangeArgs>? _destinationPorts;

        /// <summary>
        /// The destination ports to inspect for. If not specified, this matches with any destination port. This setting is only used for protocols 6 (TCP) and 17 (UDP).
        /// 
        /// You can specify individual ports, for example `1994` and you can specify port ranges, for example `1990:1994` .
        /// </summary>
        public InputList<Inputs.RuleGroupPortRangeArgs> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<Inputs.RuleGroupPortRangeArgs>());
            set => _destinationPorts = value;
        }

        [Input("destinations")]
        private InputList<Inputs.RuleGroupAddressArgs>? _destinations;

        /// <summary>
        /// The destination IP addresses and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address.
        /// </summary>
        public InputList<Inputs.RuleGroupAddressArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.RuleGroupAddressArgs>());
            set => _destinations = value;
        }

        [Input("protocols")]
        private InputList<int>? _protocols;

        /// <summary>
        /// The protocols to inspect for, specified using each protocol's assigned internet protocol number (IANA). If not specified, this matches with any protocol.
        /// </summary>
        public InputList<int> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<int>());
            set => _protocols = value;
        }

        [Input("sourcePorts")]
        private InputList<Inputs.RuleGroupPortRangeArgs>? _sourcePorts;

        /// <summary>
        /// The source ports to inspect for. If not specified, this matches with any source port. This setting is only used for protocols 6 (TCP) and 17 (UDP).
        /// 
        /// You can specify individual ports, for example `1994` and you can specify port ranges, for example `1990:1994` .
        /// </summary>
        public InputList<Inputs.RuleGroupPortRangeArgs> SourcePorts
        {
            get => _sourcePorts ?? (_sourcePorts = new InputList<Inputs.RuleGroupPortRangeArgs>());
            set => _sourcePorts = value;
        }

        [Input("sources")]
        private InputList<Inputs.RuleGroupAddressArgs>? _sources;

        /// <summary>
        /// The source IP addresses and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address.
        /// </summary>
        public InputList<Inputs.RuleGroupAddressArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.RuleGroupAddressArgs>());
            set => _sources = value;
        }

        [Input("tcpFlags")]
        private InputList<Inputs.RuleGroupTcpFlagFieldArgs>? _tcpFlags;

        /// <summary>
        /// The TCP flags and masks to inspect for. If not specified, this matches with any settings. This setting is only used for protocol 6 (TCP).
        /// </summary>
        public InputList<Inputs.RuleGroupTcpFlagFieldArgs> TcpFlags
        {
            get => _tcpFlags ?? (_tcpFlags = new InputList<Inputs.RuleGroupTcpFlagFieldArgs>());
            set => _tcpFlags = value;
        }

        public RuleGroupMatchAttributesArgs()
        {
        }
        public static new RuleGroupMatchAttributesArgs Empty => new RuleGroupMatchAttributesArgs();
    }
}
