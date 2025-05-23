// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The metadata that you apply to a resource to help you categorize and organize them. Each tag consists of a key and an optional value. You define them.
    ///  The following basic restrictions apply to tags:
    ///   +  Maximum number of tags per resource - 50
    ///   +  For each resource, each tag key must be unique, and each tag key can have only one value.
    ///   +  Maximum key length - 128 Unicode characters in UTF-8
    ///   +  Maximum value length - 256 Unicode characters in UTF-8
    ///   +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
    ///   +  Tag keys and values are case-sensitive.
    ///   +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
    ///   
    ///  In order to tag a service that has the following ARN format, you need to migrate the service to the long ARN. You must use the API, CLI or console to migrate the service ARN. For more information, see [Migrate an short service ARN to a long ARN](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-arn-migration.html) in the *Developer Guide*.
    ///   ``arn:aws:ecs:region:aws_account_id:service/service-name`` 
    ///  After the migration is complete, the following are true:
    ///   +   The service ARN is: ``arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name``
    ///   +  You can use CFN to tag the service as you would a service with a long ARN format.
    ///   +  When the ``PhysicalResourceId`` in the CFN stack represents a service, the value does not change and will be the short service ARN.
    /// </summary>
    public sealed class ServiceTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The optional part of a key-value pair that make up a tag. A ``value`` acts as a descriptor within a tag category (key).
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServiceTagArgs()
        {
        }
        public static new ServiceTagArgs Empty => new ServiceTagArgs();
    }
}
