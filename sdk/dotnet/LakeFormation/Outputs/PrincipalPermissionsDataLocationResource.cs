// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Outputs
{

    [OutputType]
    public sealed class PrincipalPermissionsDataLocationResource
    {
        public readonly string CatalogId;
        public readonly string ResourceArn;

        [OutputConstructor]
        private PrincipalPermissionsDataLocationResource(
            string catalogId,

            string resourceArn)
        {
            CatalogId = catalogId;
            ResourceArn = resourceArn;
        }
    }
}