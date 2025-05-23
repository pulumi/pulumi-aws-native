// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    public sealed class StudioComponentScriptParameterKeyValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A script parameter key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A script parameter value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public StudioComponentScriptParameterKeyValueArgs()
        {
        }
        public static new StudioComponentScriptParameterKeyValueArgs Empty => new StudioComponentScriptParameterKeyValueArgs();
    }
}
