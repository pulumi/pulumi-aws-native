// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTCoreDeviceAdvisor
{
    public static class GetSuiteDefinition
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetSuiteDefinitionResult> InvokeAsync(GetSuiteDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSuiteDefinitionResult>("aws-native:iotcoredeviceadvisor:getSuiteDefinition", args ?? new GetSuiteDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetSuiteDefinitionResult> Invoke(GetSuiteDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSuiteDefinitionResult>("aws-native:iotcoredeviceadvisor:getSuiteDefinition", args ?? new GetSuiteDefinitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetSuiteDefinitionResult> Invoke(GetSuiteDefinitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSuiteDefinitionResult>("aws-native:iotcoredeviceadvisor:getSuiteDefinition", args ?? new GetSuiteDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSuiteDefinitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for the suite definition.
        /// </summary>
        [Input("suiteDefinitionId", required: true)]
        public string SuiteDefinitionId { get; set; } = null!;

        public GetSuiteDefinitionArgs()
        {
        }
        public static new GetSuiteDefinitionArgs Empty => new GetSuiteDefinitionArgs();
    }

    public sealed class GetSuiteDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for the suite definition.
        /// </summary>
        [Input("suiteDefinitionId", required: true)]
        public Input<string> SuiteDefinitionId { get; set; } = null!;

        public GetSuiteDefinitionInvokeArgs()
        {
        }
        public static new GetSuiteDefinitionInvokeArgs Empty => new GetSuiteDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSuiteDefinitionResult
    {
        /// <summary>
        /// The Amazon Resource name for the suite definition.
        /// </summary>
        public readonly string? SuiteDefinitionArn;
        /// <summary>
        /// The configuration of the Suite Definition. Listed below are the required elements of the `SuiteDefinitionConfiguration` .
        /// 
        /// - ***devicePermissionRoleArn*** - The device permission arn.
        /// 
        /// This is a required element.
        /// 
        /// *Type:* String
        /// - ***devices*** - The list of configured devices under test. For more information on devices under test, see [DeviceUnderTest](https://docs.aws.amazon.com/iot/latest/apireference/API_iotdeviceadvisor_DeviceUnderTest.html)
        /// 
        /// Not a required element.
        /// 
        /// *Type:* List of devices under test
        /// - ***intendedForQualification*** - The tests intended for qualification in a suite.
        /// 
        /// Not a required element.
        /// 
        /// *Type:* Boolean
        /// - ***rootGroup*** - The test suite root group. For more information on creating and using root groups see the [Device Advisor workflow](https://docs.aws.amazon.com/iot/latest/developerguide/device-advisor-workflow.html) .
        /// 
        /// This is a required element.
        /// 
        /// *Type:* String
        /// - ***suiteDefinitionName*** - The Suite Definition Configuration name.
        /// 
        /// This is a required element.
        /// 
        /// *Type:* String
        /// </summary>
        public readonly Outputs.SuiteDefinitionConfigurationProperties? SuiteDefinitionConfiguration;
        /// <summary>
        /// The unique identifier for the suite definition.
        /// </summary>
        public readonly string? SuiteDefinitionId;
        /// <summary>
        /// The suite definition version of a test suite.
        /// </summary>
        public readonly string? SuiteDefinitionVersion;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetSuiteDefinitionResult(
            string? suiteDefinitionArn,

            Outputs.SuiteDefinitionConfigurationProperties? suiteDefinitionConfiguration,

            string? suiteDefinitionId,

            string? suiteDefinitionVersion,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            SuiteDefinitionArn = suiteDefinitionArn;
            SuiteDefinitionConfiguration = suiteDefinitionConfiguration;
            SuiteDefinitionId = suiteDefinitionId;
            SuiteDefinitionVersion = suiteDefinitionVersion;
            Tags = tags;
        }
    }
}
