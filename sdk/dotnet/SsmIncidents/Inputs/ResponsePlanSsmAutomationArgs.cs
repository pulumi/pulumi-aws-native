// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Inputs
{

    /// <summary>
    /// The configuration to use when starting the SSM automation document.
    /// </summary>
    public sealed class ResponsePlanSsmAutomationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The document name to use when starting the SSM automation document.
        /// </summary>
        [Input("documentName", required: true)]
        public Input<string> DocumentName { get; set; } = null!;

        /// <summary>
        /// The version of the document to use when starting the SSM automation document.
        /// </summary>
        [Input("documentVersion")]
        public Input<string>? DocumentVersion { get; set; }

        [Input("dynamicParameters")]
        private InputList<Inputs.ResponsePlanDynamicSsmParameterArgs>? _dynamicParameters;

        /// <summary>
        /// The parameters with dynamic values to set when starting the SSM automation document.
        /// </summary>
        public InputList<Inputs.ResponsePlanDynamicSsmParameterArgs> DynamicParameters
        {
            get => _dynamicParameters ?? (_dynamicParameters = new InputList<Inputs.ResponsePlanDynamicSsmParameterArgs>());
            set => _dynamicParameters = value;
        }

        [Input("parameters")]
        private InputList<Inputs.ResponsePlanSsmParameterArgs>? _parameters;

        /// <summary>
        /// The parameters to set when starting the SSM automation document.
        /// </summary>
        public InputList<Inputs.ResponsePlanSsmParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ResponsePlanSsmParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The role ARN to use when starting the SSM automation document.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The account type to use when starting the SSM automation document.
        /// </summary>
        [Input("targetAccount")]
        public Input<Pulumi.AwsNative.SsmIncidents.ResponsePlanSsmAutomationTargetAccount>? TargetAccount { get; set; }

        public ResponsePlanSsmAutomationArgs()
        {
        }
        public static new ResponsePlanSsmAutomationArgs Empty => new ResponsePlanSsmAutomationArgs();
    }
}