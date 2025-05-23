// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryControl
{
    public static class GetControlPanel
    {
        /// <summary>
        /// AWS Route53 Recovery Control Control Panel resource schema .
        /// </summary>
        public static Task<GetControlPanelResult> InvokeAsync(GetControlPanelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetControlPanelResult>("aws-native:route53recoverycontrol:getControlPanel", args ?? new GetControlPanelArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Route53 Recovery Control Control Panel resource schema .
        /// </summary>
        public static Output<GetControlPanelResult> Invoke(GetControlPanelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetControlPanelResult>("aws-native:route53recoverycontrol:getControlPanel", args ?? new GetControlPanelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Route53 Recovery Control Control Panel resource schema .
        /// </summary>
        public static Output<GetControlPanelResult> Invoke(GetControlPanelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetControlPanelResult>("aws-native:route53recoverycontrol:getControlPanel", args ?? new GetControlPanelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetControlPanelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Input("controlPanelArn", required: true)]
        public string ControlPanelArn { get; set; } = null!;

        public GetControlPanelArgs()
        {
        }
        public static new GetControlPanelArgs Empty => new GetControlPanelArgs();
    }

    public sealed class GetControlPanelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Input("controlPanelArn", required: true)]
        public Input<string> ControlPanelArn { get; set; } = null!;

        public GetControlPanelInvokeArgs()
        {
        }
        public static new GetControlPanelInvokeArgs Empty => new GetControlPanelInvokeArgs();
    }


    [OutputType]
    public sealed class GetControlPanelResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        public readonly string? ControlPanelArn;
        /// <summary>
        /// A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.
        /// </summary>
        public readonly bool? DefaultControlPanel;
        /// <summary>
        /// The name of the control panel. You can use any non-white space character in the name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Count of associated routing controls
        /// </summary>
        public readonly int? RoutingControlCount;
        /// <summary>
        /// The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53RecoveryControl.ControlPanelStatus? Status;

        [OutputConstructor]
        private GetControlPanelResult(
            string? controlPanelArn,

            bool? defaultControlPanel,

            string? name,

            int? routingControlCount,

            Pulumi.AwsNative.Route53RecoveryControl.ControlPanelStatus? status)
        {
            ControlPanelArn = controlPanelArn;
            DefaultControlPanel = defaultControlPanel;
            Name = name;
            RoutingControlCount = routingControlCount;
            Status = status;
        }
    }
}
