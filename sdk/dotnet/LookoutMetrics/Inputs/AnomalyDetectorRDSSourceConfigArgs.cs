// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorRDSSourceConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("dBInstanceIdentifier", required: true)]
        public Input<string> DBInstanceIdentifier { get; set; } = null!;

        [Input("databaseHost", required: true)]
        public Input<string> DatabaseHost { get; set; } = null!;

        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("databasePort", required: true)]
        public Input<int> DatabasePort { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("secretManagerArn", required: true)]
        public Input<string> SecretManagerArn { get; set; } = null!;

        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("vpcConfiguration", required: true)]
        public Input<Inputs.AnomalyDetectorVpcConfigurationArgs> VpcConfiguration { get; set; } = null!;

        public AnomalyDetectorRDSSourceConfigArgs()
        {
        }
        public static new AnomalyDetectorRDSSourceConfigArgs Empty => new AnomalyDetectorRDSSourceConfigArgs();
    }
}