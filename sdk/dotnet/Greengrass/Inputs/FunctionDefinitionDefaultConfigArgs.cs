// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Inputs
{

    public sealed class FunctionDefinitionDefaultConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("execution", required: true)]
        public Input<Inputs.FunctionDefinitionExecutionArgs> Execution { get; set; } = null!;

        public FunctionDefinitionDefaultConfigArgs()
        {
        }
        public static new FunctionDefinitionDefaultConfigArgs Empty => new FunctionDefinitionDefaultConfigArgs();
    }
}