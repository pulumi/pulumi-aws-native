// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatasetAction
    {
        /// <summary>
        /// The name of the data set action by which data set contents are automatically created.
        /// </summary>
        public readonly string ActionName;
        /// <summary>
        /// Information which allows the system to run a containerized application in order to create the data set contents. The application must be in a Docker container along with any needed support libraries.
        /// </summary>
        public readonly Outputs.DatasetContainerAction? ContainerAction;
        /// <summary>
        /// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
        /// </summary>
        public readonly Outputs.DatasetQueryAction? QueryAction;

        [OutputConstructor]
        private DatasetAction(
            string actionName,

            Outputs.DatasetContainerAction? containerAction,

            Outputs.DatasetQueryAction? queryAction)
        {
            ActionName = actionName;
            ContainerAction = containerAction;
            QueryAction = queryAction;
        }
    }
}
