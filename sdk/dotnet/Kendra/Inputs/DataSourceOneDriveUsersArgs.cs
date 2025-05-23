// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceOneDriveUsersArgs : global::Pulumi.ResourceArgs
    {
        [Input("oneDriveUserList")]
        private InputList<string>? _oneDriveUserList;

        /// <summary>
        /// A list of users whose documents should be indexed. Specify the user names in email format, for example, `username@tenantdomain` . If you need to index the documents of more than 10 users, use the `OneDriveUserS3Path` field to specify the location of a file containing a list of users.
        /// </summary>
        public InputList<string> OneDriveUserList
        {
            get => _oneDriveUserList ?? (_oneDriveUserList = new InputList<string>());
            set => _oneDriveUserList = value;
        }

        /// <summary>
        /// The S3 bucket location of a file containing a list of users whose documents should be indexed.
        /// </summary>
        [Input("oneDriveUserS3Path")]
        public Input<Inputs.DataSourceS3PathArgs>? OneDriveUserS3Path { get; set; }

        public DataSourceOneDriveUsersArgs()
        {
        }
        public static new DataSourceOneDriveUsersArgs Empty => new DataSourceOneDriveUsersArgs();
    }
}
