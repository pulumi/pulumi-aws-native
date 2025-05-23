// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    public static class GetManagedLoginBranding
    {
        /// <summary>
        /// Resource Type definition for AWS::Cognito::ManagedLoginBranding
        /// </summary>
        public static Task<GetManagedLoginBrandingResult> InvokeAsync(GetManagedLoginBrandingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedLoginBrandingResult>("aws-native:cognito:getManagedLoginBranding", args ?? new GetManagedLoginBrandingArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::ManagedLoginBranding
        /// </summary>
        public static Output<GetManagedLoginBrandingResult> Invoke(GetManagedLoginBrandingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedLoginBrandingResult>("aws-native:cognito:getManagedLoginBranding", args ?? new GetManagedLoginBrandingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::ManagedLoginBranding
        /// </summary>
        public static Output<GetManagedLoginBrandingResult> Invoke(GetManagedLoginBrandingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedLoginBrandingResult>("aws-native:cognito:getManagedLoginBranding", args ?? new GetManagedLoginBrandingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedLoginBrandingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the managed login branding style.
        /// </summary>
        [Input("managedLoginBrandingId", required: true)]
        public string ManagedLoginBrandingId { get; set; } = null!;

        /// <summary>
        /// The user pool where the branding style is assigned.
        /// </summary>
        [Input("userPoolId", required: true)]
        public string UserPoolId { get; set; } = null!;

        public GetManagedLoginBrandingArgs()
        {
        }
        public static new GetManagedLoginBrandingArgs Empty => new GetManagedLoginBrandingArgs();
    }

    public sealed class GetManagedLoginBrandingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the managed login branding style.
        /// </summary>
        [Input("managedLoginBrandingId", required: true)]
        public Input<string> ManagedLoginBrandingId { get; set; } = null!;

        /// <summary>
        /// The user pool where the branding style is assigned.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public GetManagedLoginBrandingInvokeArgs()
        {
        }
        public static new GetManagedLoginBrandingInvokeArgs Empty => new GetManagedLoginBrandingInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedLoginBrandingResult
    {
        /// <summary>
        /// An array of image files that you want to apply to roles like backgrounds, logos, and icons. Each object must also indicate whether it is for dark mode, light mode, or browser-adaptive mode.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedLoginBrandingAssetType> Assets;
        /// <summary>
        /// The ID of the managed login branding style.
        /// </summary>
        public readonly string? ManagedLoginBrandingId;
        /// <summary>
        /// A JSON file, encoded as a `Document` type, with the the settings that you want to apply to your style.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::ManagedLoginBranding` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? Settings;
        /// <summary>
        /// When true, applies the default branding style options. This option reverts to default style options that are managed by Amazon Cognito. You can modify them later in the branding editor.
        /// 
        /// When you specify `true` for this option, you must also omit values for `Settings` and `Assets` in the request.
        /// </summary>
        public readonly bool? UseCognitoProvidedValues;

        [OutputConstructor]
        private GetManagedLoginBrandingResult(
            ImmutableArray<Outputs.ManagedLoginBrandingAssetType> assets,

            string? managedLoginBrandingId,

            object? settings,

            bool? useCognitoProvidedValues)
        {
            Assets = assets;
            ManagedLoginBrandingId = managedLoginBrandingId;
            Settings = settings;
            UseCognitoProvidedValues = useCognitoProvidedValues;
        }
    }
}
