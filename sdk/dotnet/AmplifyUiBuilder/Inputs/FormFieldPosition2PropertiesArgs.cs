// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormFieldPosition2PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("below", required: true)]
        public Input<string> Below { get; set; } = null!;

        public FormFieldPosition2PropertiesArgs()
        {
        }
        public static new FormFieldPosition2PropertiesArgs Empty => new FormFieldPosition2PropertiesArgs();
    }
}
