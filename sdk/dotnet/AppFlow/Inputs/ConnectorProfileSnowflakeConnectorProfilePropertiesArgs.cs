// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileSnowflakeConnectorProfilePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The name of the Amazon S3 bucket associated with Snowﬂake.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The bucket prefix that refers to the Amazon S3 bucket associated with Snowﬂake.
        /// </summary>
        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        /// <summary>
        /// The Snowﬂake Private Link service name to be used for private data transfers.
        /// </summary>
        [Input("privateLinkServiceName")]
        public Input<string>? PrivateLinkServiceName { get; set; }

        /// <summary>
        /// The region of the Snowﬂake account.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the
        /// Snowﬂake account. This is written in the following format: &lt; Database&gt;&lt; Schema&gt;&lt;Stage Name&gt;.
        /// </summary>
        [Input("stage", required: true)]
        public Input<string> Stage { get; set; } = null!;

        /// <summary>
        /// The name of the Snowﬂake warehouse.
        /// </summary>
        [Input("warehouse", required: true)]
        public Input<string> Warehouse { get; set; } = null!;

        public ConnectorProfileSnowflakeConnectorProfilePropertiesArgs()
        {
        }
        public static new ConnectorProfileSnowflakeConnectorProfilePropertiesArgs Empty => new ConnectorProfileSnowflakeConnectorProfilePropertiesArgs();
    }
}