// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies an Object Ownership rule.
    ///  S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable access control lists (ACLs) and take ownership of every object in your bucket, simplifying access management for data stored in Amazon S3. For more information, see [Controlling ownership of objects and disabling ACLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
    /// </summary>
    [OutputType]
    public sealed class BucketOwnershipControlsRule
    {
        /// <summary>
        /// Specifies an object ownership rule.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketOwnershipControlsRuleObjectOwnership? ObjectOwnership;

        [OutputConstructor]
        private BucketOwnershipControlsRule(Pulumi.AwsNative.S3.BucketOwnershipControlsRuleObjectOwnership? objectOwnership)
        {
            ObjectOwnership = objectOwnership;
        }
    }
}
