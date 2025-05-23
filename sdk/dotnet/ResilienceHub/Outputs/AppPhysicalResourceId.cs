// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResilienceHub.Outputs
{

    [OutputType]
    public sealed class AppPhysicalResourceId
    {
        /// <summary>
        /// The AWS account that owns the physical resource.
        /// </summary>
        public readonly string? AwsAccountId;
        /// <summary>
        /// The AWS Region that the physical resource is located in.
        /// </summary>
        public readonly string? AwsRegion;
        /// <summary>
        /// Identifier of the physical resource.
        /// </summary>
        public readonly string Identifier;
        /// <summary>
        /// Specifies the type of physical resource identifier.
        /// 
        /// - **Arn** - The resource identifier is an Amazon Resource Name (ARN) and it can identify the following list of resources:
        /// 
        /// - `AWS::ECS::Service`
        /// - `AWS::EFS::FileSystem`
        /// - `AWS::ElasticLoadBalancingV2::LoadBalancer`
        /// - `AWS::Lambda::Function`
        /// - `AWS::SNS::Topic`
        /// - **Native** - The resource identifier is an AWS Resilience Hub -native identifier and it can identify the following list of resources:
        /// 
        /// - `AWS::ApiGateway::RestApi`
        /// - `AWS::ApiGatewayV2::Api`
        /// - `AWS::AutoScaling::AutoScalingGroup`
        /// - `AWS::DocDB::DBCluster`
        /// - `AWS::DocDB::DBGlobalCluster`
        /// - `AWS::DocDB::DBInstance`
        /// - `AWS::DynamoDB::GlobalTable`
        /// - `AWS::DynamoDB::Table`
        /// - `AWS::EC2::EC2Fleet`
        /// - `AWS::EC2::Instance`
        /// - `AWS::EC2::NatGateway`
        /// - `AWS::EC2::Volume`
        /// - `AWS::ElasticLoadBalancing::LoadBalancer`
        /// - `AWS::RDS::DBCluster`
        /// - `AWS::RDS::DBInstance`
        /// - `AWS::RDS::GlobalCluster`
        /// - `AWS::Route53::RecordSet`
        /// - `AWS::S3::Bucket`
        /// - `AWS::SQS::Queue`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AppPhysicalResourceId(
            string? awsAccountId,

            string? awsRegion,

            string identifier,

            string type)
        {
            AwsAccountId = awsAccountId;
            AwsRegion = awsRegion;
            Identifier = identifier;
            Type = type;
        }
    }
}
