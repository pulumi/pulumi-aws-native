// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetThing
    {
        /// <summary>
        /// Resource Type definition for AWS::IoT::Thing
        /// </summary>
        public static Task<GetThingResult> InvokeAsync(GetThingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetThingResult>("aws-native:iot:getThing", args ?? new GetThingArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoT::Thing
        /// </summary>
        public static Output<GetThingResult> Invoke(GetThingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetThingResult>("aws-native:iot:getThing", args ?? new GetThingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetThingArgs : global::Pulumi.InvokeArgs
    {
        [Input("thingName", required: true)]
        public string ThingName { get; set; } = null!;

        public GetThingArgs()
        {
        }
        public static new GetThingArgs Empty => new GetThingArgs();
    }

    public sealed class GetThingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("thingName", required: true)]
        public Input<string> ThingName { get; set; } = null!;

        public GetThingInvokeArgs()
        {
        }
        public static new GetThingInvokeArgs Empty => new GetThingInvokeArgs();
    }


    [OutputType]
    public sealed class GetThingResult
    {
        public readonly string? Arn;
        public readonly Outputs.ThingAttributePayload? AttributePayload;
        public readonly string? Id;

        [OutputConstructor]
        private GetThingResult(
            string? arn,

            Outputs.ThingAttributePayload? attributePayload,

            string? id)
        {
            Arn = arn;
            AttributePayload = attributePayload;
            Id = id;
        }
    }
}