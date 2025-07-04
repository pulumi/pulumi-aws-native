// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify.Inputs
{

    public sealed class AppCacheConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of cache configuration to use for an Amplify app.
        /// 
        /// The `AMPLIFY_MANAGED` cache configuration automatically applies an optimized cache configuration for your app based on its platform, routing rules, and rewrite rules.
        /// 
        /// The `AMPLIFY_MANAGED_NO_COOKIES` cache configuration type is the same as `AMPLIFY_MANAGED` , except that it excludes all cookies from the cache key. This is the default setting.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.Amplify.AppCacheConfigType>? Type { get; set; }

        public AppCacheConfigArgs()
        {
        }
        public static new AppCacheConfigArgs Empty => new AppCacheConfigArgs();
    }
}
