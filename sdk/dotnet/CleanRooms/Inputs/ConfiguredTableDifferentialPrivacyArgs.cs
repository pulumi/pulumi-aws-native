// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class ConfiguredTableDifferentialPrivacyArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns", required: true)]
        private InputList<Inputs.ConfiguredTableDifferentialPrivacyColumnArgs>? _columns;
        public InputList<Inputs.ConfiguredTableDifferentialPrivacyColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.ConfiguredTableDifferentialPrivacyColumnArgs>());
            set => _columns = value;
        }

        public ConfiguredTableDifferentialPrivacyArgs()
        {
        }
        public static new ConfiguredTableDifferentialPrivacyArgs Empty => new ConfiguredTableDifferentialPrivacyArgs();
    }
}
