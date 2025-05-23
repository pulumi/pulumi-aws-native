// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Outputs
{

    /// <summary>
    /// Specifies a separate directory for each type of file to store for an AS2 message.
    /// </summary>
    [OutputType]
    public sealed class CustomDirectoriesProperties
    {
        /// <summary>
        /// Specifies a location to store the failed files for an AS2 message.
        /// </summary>
        public readonly string FailedFilesDirectory;
        /// <summary>
        /// Specifies a location to store the MDN file for an AS2 message.
        /// </summary>
        public readonly string MdnFilesDirectory;
        /// <summary>
        /// Specifies a location to store the payload file for an AS2 message.
        /// </summary>
        public readonly string PayloadFilesDirectory;
        /// <summary>
        /// Specifies a location to store the status file for an AS2 message.
        /// </summary>
        public readonly string StatusFilesDirectory;
        /// <summary>
        /// Specifies a location to store the temporary processing file for an AS2 message.
        /// </summary>
        public readonly string TemporaryFilesDirectory;

        [OutputConstructor]
        private CustomDirectoriesProperties(
            string failedFilesDirectory,

            string mdnFilesDirectory,

            string payloadFilesDirectory,

            string statusFilesDirectory,

            string temporaryFilesDirectory)
        {
            FailedFilesDirectory = failedFilesDirectory;
            MdnFilesDirectory = mdnFilesDirectory;
            PayloadFilesDirectory = payloadFilesDirectory;
            StatusFilesDirectory = statusFilesDirectory;
            TemporaryFilesDirectory = temporaryFilesDirectory;
        }
    }
}
