// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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

        /// <summary>
        /// Resource Type definition for AWS::IoT::Thing
        /// </summary>
        public static Output<GetThingResult> Invoke(GetThingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetThingResult>("aws-native:iot:getThing", args ?? new GetThingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetThingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the thing to update.
        /// 
        /// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
        /// </summary>
        [Input("thingName", required: true)]
        public string ThingName { get; set; } = null!;

        public GetThingArgs()
        {
        }
        public static new GetThingArgs Empty => new GetThingArgs();
    }

    public sealed class GetThingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the thing to update.
        /// 
        /// You can't change a thing's name. To change a thing's name, you must create a new thing, give it the new name, and then delete the old thing.
        /// </summary>
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
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS IoT thing, such as `arn:aws:iot:us-east-2:123456789012:thing/MyThing` .
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A string that contains up to three key value pairs. Maximum length of 800. Duplicates not allowed.
        /// </summary>
        public readonly Outputs.ThingAttributePayload? AttributePayload;
        /// <summary>
        /// The Id of this thing.
        /// </summary>
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
