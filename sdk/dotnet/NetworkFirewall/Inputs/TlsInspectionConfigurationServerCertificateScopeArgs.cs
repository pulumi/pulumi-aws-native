// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class TlsInspectionConfigurationServerCertificateScopeArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationPorts")]
        private InputList<Inputs.TlsInspectionConfigurationPortRangeArgs>? _destinationPorts;

        /// <summary>
        /// The destination ports to decrypt for inspection, in Transmission Control Protocol (TCP) format. If not specified, this matches with any destination port.
        /// 
        /// You can specify individual ports, for example `1994` , and you can specify port ranges, such as `1990:1994` .
        /// </summary>
        public InputList<Inputs.TlsInspectionConfigurationPortRangeArgs> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<Inputs.TlsInspectionConfigurationPortRangeArgs>());
            set => _destinationPorts = value;
        }

        [Input("destinations")]
        private InputList<Inputs.TlsInspectionConfigurationAddressArgs>? _destinations;

        /// <summary>
        /// The destination IP addresses and address ranges to decrypt for inspection, in CIDR notation. If not specified, this
        /// matches with any destination address.
        /// </summary>
        public InputList<Inputs.TlsInspectionConfigurationAddressArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.TlsInspectionConfigurationAddressArgs>());
            set => _destinations = value;
        }

        [Input("protocols")]
        private InputList<int>? _protocols;

        /// <summary>
        /// The protocols to inspect for, specified using the assigned internet protocol number (IANA) for each protocol. If not specified, this matches with any protocol.
        /// 
        /// Network Firewall currently supports only TCP.
        /// </summary>
        public InputList<int> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<int>());
            set => _protocols = value;
        }

        [Input("sourcePorts")]
        private InputList<Inputs.TlsInspectionConfigurationPortRangeArgs>? _sourcePorts;

        /// <summary>
        /// The source ports to decrypt for inspection, in Transmission Control Protocol (TCP) format. If not specified, this matches with any source port.
        /// 
        /// You can specify individual ports, for example `1994` , and you can specify port ranges, such as `1990:1994` .
        /// </summary>
        public InputList<Inputs.TlsInspectionConfigurationPortRangeArgs> SourcePorts
        {
            get => _sourcePorts ?? (_sourcePorts = new InputList<Inputs.TlsInspectionConfigurationPortRangeArgs>());
            set => _sourcePorts = value;
        }

        [Input("sources")]
        private InputList<Inputs.TlsInspectionConfigurationAddressArgs>? _sources;

        /// <summary>
        /// The source IP addresses and address ranges to decrypt for inspection, in CIDR notation. If not specified, this
        /// matches with any source address.
        /// </summary>
        public InputList<Inputs.TlsInspectionConfigurationAddressArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.TlsInspectionConfigurationAddressArgs>());
            set => _sources = value;
        }

        public TlsInspectionConfigurationServerCertificateScopeArgs()
        {
        }
        public static new TlsInspectionConfigurationServerCertificateScopeArgs Empty => new TlsInspectionConfigurationServerCertificateScopeArgs();
    }
}
