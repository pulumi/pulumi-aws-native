// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Describes whether an Amazon Lightsail content delivery network (CDN) distribution forwards cookies to the origin and, if so, which ones.
    /// </summary>
    public sealed class DistributionCookieObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("cookiesAllowList")]
        private InputList<string>? _cookiesAllowList;

        /// <summary>
        /// The specific cookies to forward to your distribution's origin.
        /// </summary>
        public InputList<string> CookiesAllowList
        {
            get => _cookiesAllowList ?? (_cookiesAllowList = new InputList<string>());
            set => _cookiesAllowList = value;
        }

        /// <summary>
        /// Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.
        /// </summary>
        [Input("option")]
        public Input<string>? Option { get; set; }

        public DistributionCookieObjectArgs()
        {
        }
        public static new DistributionCookieObjectArgs Empty => new DistributionCookieObjectArgs();
    }
}