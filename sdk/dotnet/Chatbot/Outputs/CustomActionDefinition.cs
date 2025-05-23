// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Chatbot.Outputs
{

    [OutputType]
    public sealed class CustomActionDefinition
    {
        /// <summary>
        /// The command string to run which may include variables by prefixing with a dollar sign ($).
        /// </summary>
        public readonly string CommandText;

        [OutputConstructor]
        private CustomActionDefinition(string commandText)
        {
            CommandText = commandText;
        }
    }
}
