// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceExplorer2
{
    public static class GetView
    {
        /// <summary>
        /// Definition of AWS::ResourceExplorer2::View Resource Type
        /// </summary>
        public static Task<GetViewResult> InvokeAsync(GetViewArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetViewResult>("aws-native:resourceexplorer2:getView", args ?? new GetViewArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::ResourceExplorer2::View Resource Type
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("aws-native:resourceexplorer2:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::ResourceExplorer2::View Resource Type
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("aws-native:resourceexplorer2:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetViewArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the new view. For example:
        /// 
        /// `arn:aws:resource-explorer-2:us-east-1:123456789012:view/MyView/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`
        /// </summary>
        [Input("viewArn", required: true)]
        public string ViewArn { get; set; } = null!;

        public GetViewArgs()
        {
        }
        public static new GetViewArgs Empty => new GetViewArgs();
    }

    public sealed class GetViewInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the new view. For example:
        /// 
        /// `arn:aws:resource-explorer-2:us-east-1:123456789012:view/MyView/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`
        /// </summary>
        [Input("viewArn", required: true)]
        public Input<string> ViewArn { get; set; } = null!;

        public GetViewInvokeArgs()
        {
        }
        public static new GetViewInvokeArgs Empty => new GetViewInvokeArgs();
    }


    [OutputType]
    public sealed class GetViewResult
    {
        /// <summary>
        /// An array of strings that include search keywords, prefixes, and operators that filter the results that are returned for queries made using this view. When you use this view in a [Search](https://docs.aws.amazon.com/resource-explorer/latest/apireference/API_Search.html) operation, the filter string is combined with the search's `QueryString` parameter using a logical `AND` operator.
        /// 
        /// For information about the supported syntax, see [Search query reference for Resource Explorer](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html) in the *AWS Resource Explorer User Guide* .
        /// 
        /// &gt; This query string in the context of this operation supports only [filter prefixes](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters) with optional [operators](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators) . It doesn't support free-form text. For example, the string `region:us* service:ec2 -tag:stage=prod` includes all Amazon EC2 resources in any AWS Region that begin with the letters `us` and are *not* tagged with a key `Stage` that has the value `prod` .
        /// </summary>
        public readonly Outputs.ViewSearchFilter? Filters;
        /// <summary>
        /// A list of fields that provide additional information about the view.
        /// </summary>
        public readonly ImmutableArray<Outputs.ViewIncludedProperty> IncludedProperties;
        /// <summary>
        /// Tag key and value pairs that are attached to the view.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The ARN of the new view. For example:
        /// 
        /// `arn:aws:resource-explorer-2:us-east-1:123456789012:view/MyView/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`
        /// </summary>
        public readonly string? ViewArn;

        [OutputConstructor]
        private GetViewResult(
            Outputs.ViewSearchFilter? filters,

            ImmutableArray<Outputs.ViewIncludedProperty> includedProperties,

            ImmutableDictionary<string, string>? tags,

            string? viewArn)
        {
            Filters = filters;
            IncludedProperties = includedProperties;
            Tags = tags;
            ViewArn = viewArn;
        }
    }
}
