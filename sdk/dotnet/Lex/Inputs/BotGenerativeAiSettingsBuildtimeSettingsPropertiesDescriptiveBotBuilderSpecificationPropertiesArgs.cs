// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotGenerativeAiSettingsBuildtimeSettingsPropertiesDescriptiveBotBuilderSpecificationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("bedrockModelSpecification")]
        public Input<Inputs.BotBedrockModelSpecificationArgs>? BedrockModelSpecification { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BotGenerativeAiSettingsBuildtimeSettingsPropertiesDescriptiveBotBuilderSpecificationPropertiesArgs()
        {
        }
        public static new BotGenerativeAiSettingsBuildtimeSettingsPropertiesDescriptiveBotBuilderSpecificationPropertiesArgs Empty => new BotGenerativeAiSettingsBuildtimeSettingsPropertiesDescriptiveBotBuilderSpecificationPropertiesArgs();
    }
}
