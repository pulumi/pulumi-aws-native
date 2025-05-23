// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize
{
    public static class GetSolution
    {
        /// <summary>
        /// Resource schema for AWS::Personalize::Solution.
        /// </summary>
        public static Task<GetSolutionResult> InvokeAsync(GetSolutionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSolutionResult>("aws-native:personalize:getSolution", args ?? new GetSolutionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Personalize::Solution.
        /// </summary>
        public static Output<GetSolutionResult> Invoke(GetSolutionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSolutionResult>("aws-native:personalize:getSolution", args ?? new GetSolutionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Personalize::Solution.
        /// </summary>
        public static Output<GetSolutionResult> Invoke(GetSolutionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSolutionResult>("aws-native:personalize:getSolution", args ?? new GetSolutionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSolutionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the solution.
        /// </summary>
        [Input("solutionArn", required: true)]
        public string SolutionArn { get; set; } = null!;

        public GetSolutionArgs()
        {
        }
        public static new GetSolutionArgs Empty => new GetSolutionArgs();
    }

    public sealed class GetSolutionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the solution.
        /// </summary>
        [Input("solutionArn", required: true)]
        public Input<string> SolutionArn { get; set; } = null!;

        public GetSolutionInvokeArgs()
        {
        }
        public static new GetSolutionInvokeArgs Empty => new GetSolutionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSolutionResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the solution.
        /// </summary>
        public readonly string? SolutionArn;

        [OutputConstructor]
        private GetSolutionResult(string? solutionArn)
        {
            SolutionArn = solutionArn;
        }
    }
}
