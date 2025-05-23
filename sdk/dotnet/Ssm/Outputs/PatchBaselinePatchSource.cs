// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm.Outputs
{

    /// <summary>
    /// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
    /// </summary>
    [OutputType]
    public sealed class PatchBaselinePatchSource
    {
        /// <summary>
        /// The value of the yum repo configuration. For example:
        /// 
        /// `[main]`
        /// 
        /// `name=MyCustomRepository`
        /// 
        /// `baseurl=https://my-custom-repository`
        /// 
        /// `enabled=1`
        /// 
        /// &gt; For information about other options available for your yum repository configuration, see [dnf.conf(5)](https://docs.aws.amazon.com/https://man7.org/linux/man-pages/man5/dnf.conf.5.html) .
        /// </summary>
        public readonly string? Configuration;
        /// <summary>
        /// The name specified to identify the patch source.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
        /// </summary>
        public readonly ImmutableArray<string> Products;

        [OutputConstructor]
        private PatchBaselinePatchSource(
            string? configuration,

            string? name,

            ImmutableArray<string> products)
        {
            Configuration = configuration;
            Name = name;
            Products = products;
        }
    }
}
