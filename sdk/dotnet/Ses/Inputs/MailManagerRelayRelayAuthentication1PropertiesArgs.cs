// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRelayRelayAuthentication1PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("noAuthentication", required: true)]
        public Input<Inputs.MailManagerRelayNoAuthenticationArgs> NoAuthentication { get; set; } = null!;

        public MailManagerRelayRelayAuthentication1PropertiesArgs()
        {
        }
        public static new MailManagerRelayRelayAuthentication1PropertiesArgs Empty => new MailManagerRelayRelayAuthentication1PropertiesArgs();
    }
}