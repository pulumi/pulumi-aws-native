// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// A session context that is activated when an intent is fulfilled.
    /// </summary>
    [OutputType]
    public sealed class BotOutputContext
    {
        public readonly string Name;
        public readonly int TimeToLiveInSeconds;
        public readonly int TurnsToLive;

        [OutputConstructor]
        private BotOutputContext(
            string name,

            int timeToLiveInSeconds,

            int turnsToLive)
        {
            Name = name;
            TimeToLiveInSeconds = timeToLiveInSeconds;
            TurnsToLive = turnsToLive;
        }
    }
}