// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDb
{
    public static class GetDbInstance
    {
        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBInstance
        /// </summary>
        public static Task<GetDbInstanceResult> InvokeAsync(GetDbInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDbInstanceResult>("aws-native:docdb:getDbInstance", args ?? new GetDbInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBInstance
        /// </summary>
        public static Output<GetDbInstanceResult> Invoke(GetDbInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDbInstanceResult>("aws-native:docdb:getDbInstance", args ?? new GetDbInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDbInstanceArgs()
        {
        }
        public static new GetDbInstanceArgs Empty => new GetDbInstanceArgs();
    }

    public sealed class GetDbInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDbInstanceInvokeArgs()
        {
        }
        public static new GetDbInstanceInvokeArgs Empty => new GetDbInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDbInstanceResult
    {
        public readonly bool? AutoMinorVersionUpgrade;
        public readonly string? CaCertificateIdentifier;
        public readonly bool? CertificateRotationRestart;
        public readonly string? DbInstanceClass;
        public readonly bool? EnablePerformanceInsights;
        public readonly string? Endpoint;
        public readonly string? Id;
        public readonly string? Port;
        public readonly string? PreferredMaintenanceWindow;
        public readonly ImmutableArray<Outputs.DbInstanceTag> Tags;

        [OutputConstructor]
        private GetDbInstanceResult(
            bool? autoMinorVersionUpgrade,

            string? caCertificateIdentifier,

            bool? certificateRotationRestart,

            string? dbInstanceClass,

            bool? enablePerformanceInsights,

            string? endpoint,

            string? id,

            string? port,

            string? preferredMaintenanceWindow,

            ImmutableArray<Outputs.DbInstanceTag> tags)
        {
            AutoMinorVersionUpgrade = autoMinorVersionUpgrade;
            CaCertificateIdentifier = caCertificateIdentifier;
            CertificateRotationRestart = certificateRotationRestart;
            DbInstanceClass = dbInstanceClass;
            EnablePerformanceInsights = enablePerformanceInsights;
            Endpoint = endpoint;
            Id = id;
            Port = port;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            Tags = tags;
        }
    }
}