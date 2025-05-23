// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RedshiftServerless.Inputs
{

    public sealed class WorkgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
        /// </summary>
        [Input("baseCapacity")]
        public Input<int>? BaseCapacity { get; set; }

        [Input("configParameters")]
        private InputList<Inputs.WorkgroupConfigParameterArgs>? _configParameters;

        /// <summary>
        /// An array of parameters to set for advanced control over a database. The options are `auto_mv` , `datestyle` , `enable_case_sensitive_identifier` , `enable_user_activity_logging` , `query_group` , `search_path` , `require_ssl` , `use_fips_ssl` , and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see [Query monitoring metrics for Amazon Redshift Serverless](https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless) .
        /// </summary>
        public InputList<Inputs.WorkgroupConfigParameterArgs> ConfigParameters
        {
            get => _configParameters ?? (_configParameters = new InputList<Inputs.WorkgroupConfigParameterArgs>());
            set => _configParameters = value;
        }

        /// <summary>
        /// The creation date of the workgroup.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The endpoint that is created from the workgroup.
        /// </summary>
        [Input("endpoint")]
        public Input<Inputs.WorkgroupEndpointArgs>? Endpoint { get; set; }

        /// <summary>
        /// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
        /// </summary>
        [Input("enhancedVpcRouting")]
        public Input<bool>? EnhancedVpcRouting { get; set; }

        /// <summary>
        /// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries. The max capacity is specified in RPUs.
        /// </summary>
        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        /// <summary>
        /// The namespace the workgroup is associated with.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// An object that represents the price performance target settings for the workgroup.
        /// </summary>
        [Input("pricePerformanceTarget")]
        public Input<Inputs.WorkgroupPerformanceTargetArgs>? PricePerformanceTarget { get; set; }

        /// <summary>
        /// A value that specifies whether the workgroup can be accessible from a public network.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// An array of security group IDs to associate with the workgroup.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The status of the workgroup.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AwsNative.RedshiftServerless.WorkgroupStatus>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// An array of subnet IDs the workgroup is associated with.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The name of the track for the workgroup.
        /// </summary>
        [Input("trackName")]
        public Input<string>? TrackName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) that links to the workgroup.
        /// </summary>
        [Input("workgroupArn")]
        public Input<string>? WorkgroupArn { get; set; }

        /// <summary>
        /// The unique identifier of the workgroup.
        /// </summary>
        [Input("workgroupId")]
        public Input<string>? WorkgroupId { get; set; }

        /// <summary>
        /// The name of the workgroup.
        /// </summary>
        [Input("workgroupName")]
        public Input<string>? WorkgroupName { get; set; }

        public WorkgroupArgs()
        {
        }
        public static new WorkgroupArgs Empty => new WorkgroupArgs();
    }
}
