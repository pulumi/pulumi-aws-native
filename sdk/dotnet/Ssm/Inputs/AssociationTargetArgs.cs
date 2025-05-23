// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm.Inputs
{

    public sealed class AssociationTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-defined criteria for sending commands that target managed nodes that meet the criteria.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// User-defined criteria that maps to `Key` . For example, if you specified `tag:ServerRole` , you could specify `value:WebServer` to run a command on instances that include EC2 tags of `ServerRole,WebServer` .
        /// 
        /// Depending on the type of target, the maximum number of values for a key might be lower than the global maximum of 50.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AssociationTargetArgs()
        {
        }
        public static new AssociationTargetArgs Empty => new AssociationTargetArgs();
    }
}
