// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::Endpoint
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("aws-native:dms:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	Id string `pulumi:"id"`
}

type LookupEndpointResult struct {
	CertificateArn             *string                             `pulumi:"certificateArn"`
	DatabaseName               *string                             `pulumi:"databaseName"`
	DocDbSettings              *EndpointDocDbSettings              `pulumi:"docDbSettings"`
	DynamoDbSettings           *EndpointDynamoDbSettings           `pulumi:"dynamoDbSettings"`
	ElasticsearchSettings      *EndpointElasticsearchSettings      `pulumi:"elasticsearchSettings"`
	EndpointIdentifier         *string                             `pulumi:"endpointIdentifier"`
	EndpointType               *string                             `pulumi:"endpointType"`
	EngineName                 *string                             `pulumi:"engineName"`
	ExternalId                 *string                             `pulumi:"externalId"`
	ExtraConnectionAttributes  *string                             `pulumi:"extraConnectionAttributes"`
	GcpMySQLSettings           *EndpointGcpMySQLSettings           `pulumi:"gcpMySQLSettings"`
	IbmDb2Settings             *EndpointIbmDb2Settings             `pulumi:"ibmDb2Settings"`
	Id                         *string                             `pulumi:"id"`
	KafkaSettings              *EndpointKafkaSettings              `pulumi:"kafkaSettings"`
	KinesisSettings            *EndpointKinesisSettings            `pulumi:"kinesisSettings"`
	MicrosoftSqlServerSettings *EndpointMicrosoftSqlServerSettings `pulumi:"microsoftSqlServerSettings"`
	MongoDbSettings            *EndpointMongoDbSettings            `pulumi:"mongoDbSettings"`
	MySqlSettings              *EndpointMySqlSettings              `pulumi:"mySqlSettings"`
	NeptuneSettings            *EndpointNeptuneSettings            `pulumi:"neptuneSettings"`
	OracleSettings             *EndpointOracleSettings             `pulumi:"oracleSettings"`
	Password                   *string                             `pulumi:"password"`
	Port                       *int                                `pulumi:"port"`
	PostgreSqlSettings         *EndpointPostgreSqlSettings         `pulumi:"postgreSqlSettings"`
	RedisSettings              *EndpointRedisSettings              `pulumi:"redisSettings"`
	RedshiftSettings           *EndpointRedshiftSettings           `pulumi:"redshiftSettings"`
	S3Settings                 *EndpointS3Settings                 `pulumi:"s3Settings"`
	ServerName                 *string                             `pulumi:"serverName"`
	SslMode                    *string                             `pulumi:"sslMode"`
	SybaseSettings             *EndpointSybaseSettings             `pulumi:"sybaseSettings"`
	Tags                       []EndpointTag                       `pulumi:"tags"`
	Username                   *string                             `pulumi:"username"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) CertificateArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.CertificateArn }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) DocDbSettings() EndpointDocDbSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointDocDbSettings { return v.DocDbSettings }).(EndpointDocDbSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) DynamoDbSettings() EndpointDynamoDbSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointDynamoDbSettings { return v.DynamoDbSettings }).(EndpointDynamoDbSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) ElasticsearchSettings() EndpointElasticsearchSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointElasticsearchSettings { return v.ElasticsearchSettings }).(EndpointElasticsearchSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EngineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EngineName }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) ExtraConnectionAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ExtraConnectionAttributes }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) GcpMySQLSettings() EndpointGcpMySQLSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointGcpMySQLSettings { return v.GcpMySQLSettings }).(EndpointGcpMySQLSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) IbmDb2Settings() EndpointIbmDb2SettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointIbmDb2Settings { return v.IbmDb2Settings }).(EndpointIbmDb2SettingsPtrOutput)
}

func (o LookupEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) KafkaSettings() EndpointKafkaSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointKafkaSettings { return v.KafkaSettings }).(EndpointKafkaSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) KinesisSettings() EndpointKinesisSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointKinesisSettings { return v.KinesisSettings }).(EndpointKinesisSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) MicrosoftSqlServerSettings() EndpointMicrosoftSqlServerSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointMicrosoftSqlServerSettings { return v.MicrosoftSqlServerSettings }).(EndpointMicrosoftSqlServerSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) MongoDbSettings() EndpointMongoDbSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointMongoDbSettings { return v.MongoDbSettings }).(EndpointMongoDbSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) MySqlSettings() EndpointMySqlSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointMySqlSettings { return v.MySqlSettings }).(EndpointMySqlSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) NeptuneSettings() EndpointNeptuneSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointNeptuneSettings { return v.NeptuneSettings }).(EndpointNeptuneSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) OracleSettings() EndpointOracleSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointOracleSettings { return v.OracleSettings }).(EndpointOracleSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupEndpointResultOutput) PostgreSqlSettings() EndpointPostgreSqlSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointPostgreSqlSettings { return v.PostgreSqlSettings }).(EndpointPostgreSqlSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) RedisSettings() EndpointRedisSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointRedisSettings { return v.RedisSettings }).(EndpointRedisSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) RedshiftSettings() EndpointRedshiftSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointRedshiftSettings { return v.RedshiftSettings }).(EndpointRedshiftSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) S3Settings() EndpointS3SettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointS3Settings { return v.S3Settings }).(EndpointS3SettingsPtrOutput)
}

func (o LookupEndpointResultOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) SslMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.SslMode }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) SybaseSettings() EndpointSybaseSettingsPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointSybaseSettings { return v.SybaseSettings }).(EndpointSybaseSettingsPtrOutput)
}

func (o LookupEndpointResultOutput) Tags() EndpointTagArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointTag { return v.Tags }).(EndpointTagArrayOutput)
}

func (o LookupEndpointResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}