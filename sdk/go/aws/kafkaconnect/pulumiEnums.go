// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafkaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of client authentication used to connect to the Kafka cluster. Value NONE means that no client authentication is used.
type ConnectorKafkaClusterClientAuthenticationType string

const (
	ConnectorKafkaClusterClientAuthenticationTypeNone = ConnectorKafkaClusterClientAuthenticationType("NONE")
	ConnectorKafkaClusterClientAuthenticationTypeIam  = ConnectorKafkaClusterClientAuthenticationType("IAM")
)

func (ConnectorKafkaClusterClientAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorKafkaClusterClientAuthenticationType)(nil)).Elem()
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToConnectorKafkaClusterClientAuthenticationTypeOutput() ConnectorKafkaClusterClientAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(ConnectorKafkaClusterClientAuthenticationTypeOutput)
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToConnectorKafkaClusterClientAuthenticationTypeOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectorKafkaClusterClientAuthenticationTypeOutput)
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToConnectorKafkaClusterClientAuthenticationTypePtrOutput() ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return e.ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return ConnectorKafkaClusterClientAuthenticationType(e).ToConnectorKafkaClusterClientAuthenticationTypeOutputWithContext(ctx).ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(ctx)
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectorKafkaClusterClientAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectorKafkaClusterClientAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (ConnectorKafkaClusterClientAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorKafkaClusterClientAuthenticationType)(nil)).Elem()
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToConnectorKafkaClusterClientAuthenticationTypeOutput() ConnectorKafkaClusterClientAuthenticationTypeOutput {
	return o
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToConnectorKafkaClusterClientAuthenticationTypeOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypeOutput {
	return o
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToConnectorKafkaClusterClientAuthenticationTypePtrOutput() ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return o.ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorKafkaClusterClientAuthenticationType) *ConnectorKafkaClusterClientAuthenticationType {
		return &v
	}).(ConnectorKafkaClusterClientAuthenticationTypePtrOutput)
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorKafkaClusterClientAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterClientAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorKafkaClusterClientAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectorKafkaClusterClientAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectorKafkaClusterClientAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorKafkaClusterClientAuthenticationType)(nil)).Elem()
}

func (o ConnectorKafkaClusterClientAuthenticationTypePtrOutput) ToConnectorKafkaClusterClientAuthenticationTypePtrOutput() ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return o
}

func (o ConnectorKafkaClusterClientAuthenticationTypePtrOutput) ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return o
}

func (o ConnectorKafkaClusterClientAuthenticationTypePtrOutput) Elem() ConnectorKafkaClusterClientAuthenticationTypeOutput {
	return o.ApplyT(func(v *ConnectorKafkaClusterClientAuthenticationType) ConnectorKafkaClusterClientAuthenticationType {
		if v != nil {
			return *v
		}
		var ret ConnectorKafkaClusterClientAuthenticationType
		return ret
	}).(ConnectorKafkaClusterClientAuthenticationTypeOutput)
}

func (o ConnectorKafkaClusterClientAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterClientAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectorKafkaClusterClientAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConnectorKafkaClusterClientAuthenticationTypeInput is an input type that accepts values of the ConnectorKafkaClusterClientAuthenticationType enum
// A concrete instance of `ConnectorKafkaClusterClientAuthenticationTypeInput` can be one of the following:
//
//	ConnectorKafkaClusterClientAuthenticationTypeNone
//	ConnectorKafkaClusterClientAuthenticationTypeIam
type ConnectorKafkaClusterClientAuthenticationTypeInput interface {
	pulumi.Input

	ToConnectorKafkaClusterClientAuthenticationTypeOutput() ConnectorKafkaClusterClientAuthenticationTypeOutput
	ToConnectorKafkaClusterClientAuthenticationTypeOutputWithContext(context.Context) ConnectorKafkaClusterClientAuthenticationTypeOutput
}

var connectorKafkaClusterClientAuthenticationTypePtrType = reflect.TypeOf((**ConnectorKafkaClusterClientAuthenticationType)(nil)).Elem()

type ConnectorKafkaClusterClientAuthenticationTypePtrInput interface {
	pulumi.Input

	ToConnectorKafkaClusterClientAuthenticationTypePtrOutput() ConnectorKafkaClusterClientAuthenticationTypePtrOutput
	ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(context.Context) ConnectorKafkaClusterClientAuthenticationTypePtrOutput
}

type connectorKafkaClusterClientAuthenticationTypePtr string

func ConnectorKafkaClusterClientAuthenticationTypePtr(v string) ConnectorKafkaClusterClientAuthenticationTypePtrInput {
	return (*connectorKafkaClusterClientAuthenticationTypePtr)(&v)
}

func (*connectorKafkaClusterClientAuthenticationTypePtr) ElementType() reflect.Type {
	return connectorKafkaClusterClientAuthenticationTypePtrType
}

func (in *connectorKafkaClusterClientAuthenticationTypePtr) ToConnectorKafkaClusterClientAuthenticationTypePtrOutput() ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectorKafkaClusterClientAuthenticationTypePtrOutput)
}

func (in *connectorKafkaClusterClientAuthenticationTypePtr) ToConnectorKafkaClusterClientAuthenticationTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterClientAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectorKafkaClusterClientAuthenticationTypePtrOutput)
}

// The type of encryption in transit to the Kafka cluster.
type ConnectorKafkaClusterEncryptionInTransitType string

const (
	ConnectorKafkaClusterEncryptionInTransitTypePlaintext = ConnectorKafkaClusterEncryptionInTransitType("PLAINTEXT")
	ConnectorKafkaClusterEncryptionInTransitTypeTls       = ConnectorKafkaClusterEncryptionInTransitType("TLS")
)

func (ConnectorKafkaClusterEncryptionInTransitType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorKafkaClusterEncryptionInTransitType)(nil)).Elem()
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToConnectorKafkaClusterEncryptionInTransitTypeOutput() ConnectorKafkaClusterEncryptionInTransitTypeOutput {
	return pulumi.ToOutput(e).(ConnectorKafkaClusterEncryptionInTransitTypeOutput)
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToConnectorKafkaClusterEncryptionInTransitTypeOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectorKafkaClusterEncryptionInTransitTypeOutput)
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutput() ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return e.ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(context.Background())
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return ConnectorKafkaClusterEncryptionInTransitType(e).ToConnectorKafkaClusterEncryptionInTransitTypeOutputWithContext(ctx).ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(ctx)
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectorKafkaClusterEncryptionInTransitType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectorKafkaClusterEncryptionInTransitTypeOutput struct{ *pulumi.OutputState }

func (ConnectorKafkaClusterEncryptionInTransitTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorKafkaClusterEncryptionInTransitType)(nil)).Elem()
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToConnectorKafkaClusterEncryptionInTransitTypeOutput() ConnectorKafkaClusterEncryptionInTransitTypeOutput {
	return o
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToConnectorKafkaClusterEncryptionInTransitTypeOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypeOutput {
	return o
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutput() ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return o.ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorKafkaClusterEncryptionInTransitType) *ConnectorKafkaClusterEncryptionInTransitType {
		return &v
	}).(ConnectorKafkaClusterEncryptionInTransitTypePtrOutput)
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorKafkaClusterEncryptionInTransitType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterEncryptionInTransitTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorKafkaClusterEncryptionInTransitType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectorKafkaClusterEncryptionInTransitTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorKafkaClusterEncryptionInTransitType)(nil)).Elem()
}

func (o ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutput() ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return o
}

func (o ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return o
}

func (o ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) Elem() ConnectorKafkaClusterEncryptionInTransitTypeOutput {
	return o.ApplyT(func(v *ConnectorKafkaClusterEncryptionInTransitType) ConnectorKafkaClusterEncryptionInTransitType {
		if v != nil {
			return *v
		}
		var ret ConnectorKafkaClusterEncryptionInTransitType
		return ret
	}).(ConnectorKafkaClusterEncryptionInTransitTypeOutput)
}

func (o ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorKafkaClusterEncryptionInTransitTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectorKafkaClusterEncryptionInTransitType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConnectorKafkaClusterEncryptionInTransitTypeInput is an input type that accepts values of the ConnectorKafkaClusterEncryptionInTransitType enum
// A concrete instance of `ConnectorKafkaClusterEncryptionInTransitTypeInput` can be one of the following:
//
//	ConnectorKafkaClusterEncryptionInTransitTypePlaintext
//	ConnectorKafkaClusterEncryptionInTransitTypeTls
type ConnectorKafkaClusterEncryptionInTransitTypeInput interface {
	pulumi.Input

	ToConnectorKafkaClusterEncryptionInTransitTypeOutput() ConnectorKafkaClusterEncryptionInTransitTypeOutput
	ToConnectorKafkaClusterEncryptionInTransitTypeOutputWithContext(context.Context) ConnectorKafkaClusterEncryptionInTransitTypeOutput
}

var connectorKafkaClusterEncryptionInTransitTypePtrType = reflect.TypeOf((**ConnectorKafkaClusterEncryptionInTransitType)(nil)).Elem()

type ConnectorKafkaClusterEncryptionInTransitTypePtrInput interface {
	pulumi.Input

	ToConnectorKafkaClusterEncryptionInTransitTypePtrOutput() ConnectorKafkaClusterEncryptionInTransitTypePtrOutput
	ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(context.Context) ConnectorKafkaClusterEncryptionInTransitTypePtrOutput
}

type connectorKafkaClusterEncryptionInTransitTypePtr string

func ConnectorKafkaClusterEncryptionInTransitTypePtr(v string) ConnectorKafkaClusterEncryptionInTransitTypePtrInput {
	return (*connectorKafkaClusterEncryptionInTransitTypePtr)(&v)
}

func (*connectorKafkaClusterEncryptionInTransitTypePtr) ElementType() reflect.Type {
	return connectorKafkaClusterEncryptionInTransitTypePtrType
}

func (in *connectorKafkaClusterEncryptionInTransitTypePtr) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutput() ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectorKafkaClusterEncryptionInTransitTypePtrOutput)
}

func (in *connectorKafkaClusterEncryptionInTransitTypePtr) ToConnectorKafkaClusterEncryptionInTransitTypePtrOutputWithContext(ctx context.Context) ConnectorKafkaClusterEncryptionInTransitTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectorKafkaClusterEncryptionInTransitTypePtrOutput)
}

// The type of the plugin file.
type CustomPluginContentType string

const (
	CustomPluginContentTypeJar = CustomPluginContentType("JAR")
	CustomPluginContentTypeZip = CustomPluginContentType("ZIP")
)

func (CustomPluginContentType) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomPluginContentType)(nil)).Elem()
}

func (e CustomPluginContentType) ToCustomPluginContentTypeOutput() CustomPluginContentTypeOutput {
	return pulumi.ToOutput(e).(CustomPluginContentTypeOutput)
}

func (e CustomPluginContentType) ToCustomPluginContentTypeOutputWithContext(ctx context.Context) CustomPluginContentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CustomPluginContentTypeOutput)
}

func (e CustomPluginContentType) ToCustomPluginContentTypePtrOutput() CustomPluginContentTypePtrOutput {
	return e.ToCustomPluginContentTypePtrOutputWithContext(context.Background())
}

func (e CustomPluginContentType) ToCustomPluginContentTypePtrOutputWithContext(ctx context.Context) CustomPluginContentTypePtrOutput {
	return CustomPluginContentType(e).ToCustomPluginContentTypeOutputWithContext(ctx).ToCustomPluginContentTypePtrOutputWithContext(ctx)
}

func (e CustomPluginContentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomPluginContentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomPluginContentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CustomPluginContentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CustomPluginContentTypeOutput struct{ *pulumi.OutputState }

func (CustomPluginContentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomPluginContentType)(nil)).Elem()
}

func (o CustomPluginContentTypeOutput) ToCustomPluginContentTypeOutput() CustomPluginContentTypeOutput {
	return o
}

func (o CustomPluginContentTypeOutput) ToCustomPluginContentTypeOutputWithContext(ctx context.Context) CustomPluginContentTypeOutput {
	return o
}

func (o CustomPluginContentTypeOutput) ToCustomPluginContentTypePtrOutput() CustomPluginContentTypePtrOutput {
	return o.ToCustomPluginContentTypePtrOutputWithContext(context.Background())
}

func (o CustomPluginContentTypeOutput) ToCustomPluginContentTypePtrOutputWithContext(ctx context.Context) CustomPluginContentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomPluginContentType) *CustomPluginContentType {
		return &v
	}).(CustomPluginContentTypePtrOutput)
}

func (o CustomPluginContentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CustomPluginContentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomPluginContentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CustomPluginContentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomPluginContentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomPluginContentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CustomPluginContentTypePtrOutput struct{ *pulumi.OutputState }

func (CustomPluginContentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomPluginContentType)(nil)).Elem()
}

func (o CustomPluginContentTypePtrOutput) ToCustomPluginContentTypePtrOutput() CustomPluginContentTypePtrOutput {
	return o
}

func (o CustomPluginContentTypePtrOutput) ToCustomPluginContentTypePtrOutputWithContext(ctx context.Context) CustomPluginContentTypePtrOutput {
	return o
}

func (o CustomPluginContentTypePtrOutput) Elem() CustomPluginContentTypeOutput {
	return o.ApplyT(func(v *CustomPluginContentType) CustomPluginContentType {
		if v != nil {
			return *v
		}
		var ret CustomPluginContentType
		return ret
	}).(CustomPluginContentTypeOutput)
}

func (o CustomPluginContentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomPluginContentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CustomPluginContentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CustomPluginContentTypeInput is an input type that accepts values of the CustomPluginContentType enum
// A concrete instance of `CustomPluginContentTypeInput` can be one of the following:
//
//	CustomPluginContentTypeJar
//	CustomPluginContentTypeZip
type CustomPluginContentTypeInput interface {
	pulumi.Input

	ToCustomPluginContentTypeOutput() CustomPluginContentTypeOutput
	ToCustomPluginContentTypeOutputWithContext(context.Context) CustomPluginContentTypeOutput
}

var customPluginContentTypePtrType = reflect.TypeOf((**CustomPluginContentType)(nil)).Elem()

type CustomPluginContentTypePtrInput interface {
	pulumi.Input

	ToCustomPluginContentTypePtrOutput() CustomPluginContentTypePtrOutput
	ToCustomPluginContentTypePtrOutputWithContext(context.Context) CustomPluginContentTypePtrOutput
}

type customPluginContentTypePtr string

func CustomPluginContentTypePtr(v string) CustomPluginContentTypePtrInput {
	return (*customPluginContentTypePtr)(&v)
}

func (*customPluginContentTypePtr) ElementType() reflect.Type {
	return customPluginContentTypePtrType
}

func (in *customPluginContentTypePtr) ToCustomPluginContentTypePtrOutput() CustomPluginContentTypePtrOutput {
	return pulumi.ToOutput(in).(CustomPluginContentTypePtrOutput)
}

func (in *customPluginContentTypePtr) ToCustomPluginContentTypePtrOutputWithContext(ctx context.Context) CustomPluginContentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CustomPluginContentTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorKafkaClusterClientAuthenticationTypeInput)(nil)).Elem(), ConnectorKafkaClusterClientAuthenticationType("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorKafkaClusterClientAuthenticationTypePtrInput)(nil)).Elem(), ConnectorKafkaClusterClientAuthenticationType("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorKafkaClusterEncryptionInTransitTypeInput)(nil)).Elem(), ConnectorKafkaClusterEncryptionInTransitType("PLAINTEXT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorKafkaClusterEncryptionInTransitTypePtrInput)(nil)).Elem(), ConnectorKafkaClusterEncryptionInTransitType("PLAINTEXT"))
	pulumi.RegisterInputType(reflect.TypeOf((*CustomPluginContentTypeInput)(nil)).Elem(), CustomPluginContentType("JAR"))
	pulumi.RegisterInputType(reflect.TypeOf((*CustomPluginContentTypePtrInput)(nil)).Elem(), CustomPluginContentType("JAR"))
	pulumi.RegisterOutputType(ConnectorKafkaClusterClientAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(ConnectorKafkaClusterClientAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectorKafkaClusterEncryptionInTransitTypeOutput{})
	pulumi.RegisterOutputType(ConnectorKafkaClusterEncryptionInTransitTypePtrOutput{})
	pulumi.RegisterOutputType(CustomPluginContentTypeOutput{})
	pulumi.RegisterOutputType(CustomPluginContentTypePtrOutput{})
}
