// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class DomainCodeRepository
    {
        /// <summary>
        /// A CodeRepository (valid URL) to be used within Jupyter's Git extension.
        /// </summary>
        public readonly string RepositoryUrl;

        [OutputConstructor]
        private DomainCodeRepository(string repositoryUrl)
        {
            RepositoryUrl = repositoryUrl;
        }
    }
}
