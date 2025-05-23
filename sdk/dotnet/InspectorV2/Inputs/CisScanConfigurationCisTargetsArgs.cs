// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Inputs
{

    public sealed class CisScanConfigurationCisTargetsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountIds", required: true)]
        private InputList<string>? _accountIds;
        public InputList<string> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<string>());
            set => _accountIds = value;
        }

        [Input("targetResourceTags", required: true)]
        private InputMap<object>? _targetResourceTags;
        public InputMap<object> TargetResourceTags
        {
            get => _targetResourceTags ?? (_targetResourceTags = new InputMap<object>());
            set => _targetResourceTags = value;
        }

        public CisScanConfigurationCisTargetsArgs()
        {
        }
        public static new CisScanConfigurationCisTargetsArgs Empty => new CisScanConfigurationCisTargetsArgs();
    }
}
