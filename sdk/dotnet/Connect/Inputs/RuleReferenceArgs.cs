// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// Information about the reference when the ``referenceType`` is ``URL``. Otherwise, null. (Supports variable injection in the ``Value`` field.)
    /// </summary>
    public sealed class RuleReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the reference. ``DATE`` must be of type Epoch timestamp. 
        ///  *Allowed values*: ``URL`` | ``ATTACHMENT`` | ``NUMBER`` | ``STRING`` | ``DATE`` | ``EMAIL``
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Connect.RuleReferenceType> Type { get; set; } = null!;

        /// <summary>
        /// A valid value for the reference. For example, for a URL reference, a formatted URL that is displayed to an agent in the Contact Control Panel (CCP).
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public RuleReferenceArgs()
        {
        }
        public static new RuleReferenceArgs Empty => new RuleReferenceArgs();
    }
}
