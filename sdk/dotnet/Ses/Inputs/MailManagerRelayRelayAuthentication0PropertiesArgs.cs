// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRelayRelayAuthentication0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("secretArn", required: true)]
        public Input<string> SecretArn { get; set; } = null!;

        public MailManagerRelayRelayAuthentication0PropertiesArgs()
        {
        }
        public static new MailManagerRelayRelayAuthentication0PropertiesArgs Empty => new MailManagerRelayRelayAuthentication0PropertiesArgs();
    }
}
