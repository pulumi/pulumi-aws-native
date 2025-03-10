// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device type, which specify notification channel. Currently supported values: "SMS", "VOICE", "EMAIL", "CHATBOT.
type ContactChannelChannelType string

const (
	ContactChannelChannelTypeSms   = ContactChannelChannelType("SMS")
	ContactChannelChannelTypeVoice = ContactChannelChannelType("VOICE")
	ContactChannelChannelTypeEmail = ContactChannelChannelType("EMAIL")
)

func (ContactChannelChannelType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactChannelChannelType)(nil)).Elem()
}

func (e ContactChannelChannelType) ToContactChannelChannelTypeOutput() ContactChannelChannelTypeOutput {
	return pulumi.ToOutput(e).(ContactChannelChannelTypeOutput)
}

func (e ContactChannelChannelType) ToContactChannelChannelTypeOutputWithContext(ctx context.Context) ContactChannelChannelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContactChannelChannelTypeOutput)
}

func (e ContactChannelChannelType) ToContactChannelChannelTypePtrOutput() ContactChannelChannelTypePtrOutput {
	return e.ToContactChannelChannelTypePtrOutputWithContext(context.Background())
}

func (e ContactChannelChannelType) ToContactChannelChannelTypePtrOutputWithContext(ctx context.Context) ContactChannelChannelTypePtrOutput {
	return ContactChannelChannelType(e).ToContactChannelChannelTypeOutputWithContext(ctx).ToContactChannelChannelTypePtrOutputWithContext(ctx)
}

func (e ContactChannelChannelType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContactChannelChannelType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContactChannelChannelType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContactChannelChannelType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContactChannelChannelTypeOutput struct{ *pulumi.OutputState }

func (ContactChannelChannelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactChannelChannelType)(nil)).Elem()
}

func (o ContactChannelChannelTypeOutput) ToContactChannelChannelTypeOutput() ContactChannelChannelTypeOutput {
	return o
}

func (o ContactChannelChannelTypeOutput) ToContactChannelChannelTypeOutputWithContext(ctx context.Context) ContactChannelChannelTypeOutput {
	return o
}

func (o ContactChannelChannelTypeOutput) ToContactChannelChannelTypePtrOutput() ContactChannelChannelTypePtrOutput {
	return o.ToContactChannelChannelTypePtrOutputWithContext(context.Background())
}

func (o ContactChannelChannelTypeOutput) ToContactChannelChannelTypePtrOutputWithContext(ctx context.Context) ContactChannelChannelTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactChannelChannelType) *ContactChannelChannelType {
		return &v
	}).(ContactChannelChannelTypePtrOutput)
}

func (o ContactChannelChannelTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContactChannelChannelTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContactChannelChannelType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContactChannelChannelTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContactChannelChannelTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContactChannelChannelType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContactChannelChannelTypePtrOutput struct{ *pulumi.OutputState }

func (ContactChannelChannelTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactChannelChannelType)(nil)).Elem()
}

func (o ContactChannelChannelTypePtrOutput) ToContactChannelChannelTypePtrOutput() ContactChannelChannelTypePtrOutput {
	return o
}

func (o ContactChannelChannelTypePtrOutput) ToContactChannelChannelTypePtrOutputWithContext(ctx context.Context) ContactChannelChannelTypePtrOutput {
	return o
}

func (o ContactChannelChannelTypePtrOutput) Elem() ContactChannelChannelTypeOutput {
	return o.ApplyT(func(v *ContactChannelChannelType) ContactChannelChannelType {
		if v != nil {
			return *v
		}
		var ret ContactChannelChannelType
		return ret
	}).(ContactChannelChannelTypeOutput)
}

func (o ContactChannelChannelTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContactChannelChannelTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContactChannelChannelType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ContactChannelChannelTypeInput is an input type that accepts values of the ContactChannelChannelType enum
// A concrete instance of `ContactChannelChannelTypeInput` can be one of the following:
//
//	ContactChannelChannelTypeSms
//	ContactChannelChannelTypeVoice
//	ContactChannelChannelTypeEmail
type ContactChannelChannelTypeInput interface {
	pulumi.Input

	ToContactChannelChannelTypeOutput() ContactChannelChannelTypeOutput
	ToContactChannelChannelTypeOutputWithContext(context.Context) ContactChannelChannelTypeOutput
}

var contactChannelChannelTypePtrType = reflect.TypeOf((**ContactChannelChannelType)(nil)).Elem()

type ContactChannelChannelTypePtrInput interface {
	pulumi.Input

	ToContactChannelChannelTypePtrOutput() ContactChannelChannelTypePtrOutput
	ToContactChannelChannelTypePtrOutputWithContext(context.Context) ContactChannelChannelTypePtrOutput
}

type contactChannelChannelTypePtr string

func ContactChannelChannelTypePtr(v string) ContactChannelChannelTypePtrInput {
	return (*contactChannelChannelTypePtr)(&v)
}

func (*contactChannelChannelTypePtr) ElementType() reflect.Type {
	return contactChannelChannelTypePtrType
}

func (in *contactChannelChannelTypePtr) ToContactChannelChannelTypePtrOutput() ContactChannelChannelTypePtrOutput {
	return pulumi.ToOutput(in).(ContactChannelChannelTypePtrOutput)
}

func (in *contactChannelChannelTypePtr) ToContactChannelChannelTypePtrOutputWithContext(ctx context.Context) ContactChannelChannelTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContactChannelChannelTypePtrOutput)
}

// Contact type, which specify type of contact. Currently supported values: "PERSONAL", "SHARED", "OTHER".
type ContactType string

const (
	ContactTypePersonal       = ContactType("PERSONAL")
	ContactTypeEscalation     = ContactType("ESCALATION")
	ContactTypeOncallSchedule = ContactType("ONCALL_SCHEDULE")
)

func (ContactType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactType)(nil)).Elem()
}

func (e ContactType) ToContactTypeOutput() ContactTypeOutput {
	return pulumi.ToOutput(e).(ContactTypeOutput)
}

func (e ContactType) ToContactTypeOutputWithContext(ctx context.Context) ContactTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContactTypeOutput)
}

func (e ContactType) ToContactTypePtrOutput() ContactTypePtrOutput {
	return e.ToContactTypePtrOutputWithContext(context.Background())
}

func (e ContactType) ToContactTypePtrOutputWithContext(ctx context.Context) ContactTypePtrOutput {
	return ContactType(e).ToContactTypeOutputWithContext(ctx).ToContactTypePtrOutputWithContext(ctx)
}

func (e ContactType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContactType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContactType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContactType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContactTypeOutput struct{ *pulumi.OutputState }

func (ContactTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactType)(nil)).Elem()
}

func (o ContactTypeOutput) ToContactTypeOutput() ContactTypeOutput {
	return o
}

func (o ContactTypeOutput) ToContactTypeOutputWithContext(ctx context.Context) ContactTypeOutput {
	return o
}

func (o ContactTypeOutput) ToContactTypePtrOutput() ContactTypePtrOutput {
	return o.ToContactTypePtrOutputWithContext(context.Background())
}

func (o ContactTypeOutput) ToContactTypePtrOutputWithContext(ctx context.Context) ContactTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactType) *ContactType {
		return &v
	}).(ContactTypePtrOutput)
}

func (o ContactTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContactTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContactType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContactTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContactTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContactType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContactTypePtrOutput struct{ *pulumi.OutputState }

func (ContactTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactType)(nil)).Elem()
}

func (o ContactTypePtrOutput) ToContactTypePtrOutput() ContactTypePtrOutput {
	return o
}

func (o ContactTypePtrOutput) ToContactTypePtrOutputWithContext(ctx context.Context) ContactTypePtrOutput {
	return o
}

func (o ContactTypePtrOutput) Elem() ContactTypeOutput {
	return o.ApplyT(func(v *ContactType) ContactType {
		if v != nil {
			return *v
		}
		var ret ContactType
		return ret
	}).(ContactTypeOutput)
}

func (o ContactTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContactTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContactType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ContactTypeInput is an input type that accepts values of the ContactType enum
// A concrete instance of `ContactTypeInput` can be one of the following:
//
//	ContactTypePersonal
//	ContactTypeEscalation
//	ContactTypeOncallSchedule
type ContactTypeInput interface {
	pulumi.Input

	ToContactTypeOutput() ContactTypeOutput
	ToContactTypeOutputWithContext(context.Context) ContactTypeOutput
}

var contactTypePtrType = reflect.TypeOf((**ContactType)(nil)).Elem()

type ContactTypePtrInput interface {
	pulumi.Input

	ToContactTypePtrOutput() ContactTypePtrOutput
	ToContactTypePtrOutputWithContext(context.Context) ContactTypePtrOutput
}

type contactTypePtr string

func ContactTypePtr(v string) ContactTypePtrInput {
	return (*contactTypePtr)(&v)
}

func (*contactTypePtr) ElementType() reflect.Type {
	return contactTypePtrType
}

func (in *contactTypePtr) ToContactTypePtrOutput() ContactTypePtrOutput {
	return pulumi.ToOutput(in).(ContactTypePtrOutput)
}

func (in *contactTypePtr) ToContactTypePtrOutputWithContext(ctx context.Context) ContactTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContactTypePtrOutput)
}

// The day of the week when weekly recurring on-call shift rotations begin.
type RotationDayOfWeek string

const (
	RotationDayOfWeekMon = RotationDayOfWeek("MON")
	RotationDayOfWeekTue = RotationDayOfWeek("TUE")
	RotationDayOfWeekWed = RotationDayOfWeek("WED")
	RotationDayOfWeekThu = RotationDayOfWeek("THU")
	RotationDayOfWeekFri = RotationDayOfWeek("FRI")
	RotationDayOfWeekSat = RotationDayOfWeek("SAT")
	RotationDayOfWeekSun = RotationDayOfWeek("SUN")
)

func (RotationDayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationDayOfWeek)(nil)).Elem()
}

func (e RotationDayOfWeek) ToRotationDayOfWeekOutput() RotationDayOfWeekOutput {
	return pulumi.ToOutput(e).(RotationDayOfWeekOutput)
}

func (e RotationDayOfWeek) ToRotationDayOfWeekOutputWithContext(ctx context.Context) RotationDayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RotationDayOfWeekOutput)
}

func (e RotationDayOfWeek) ToRotationDayOfWeekPtrOutput() RotationDayOfWeekPtrOutput {
	return e.ToRotationDayOfWeekPtrOutputWithContext(context.Background())
}

func (e RotationDayOfWeek) ToRotationDayOfWeekPtrOutputWithContext(ctx context.Context) RotationDayOfWeekPtrOutput {
	return RotationDayOfWeek(e).ToRotationDayOfWeekOutputWithContext(ctx).ToRotationDayOfWeekPtrOutputWithContext(ctx)
}

func (e RotationDayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RotationDayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RotationDayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RotationDayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RotationDayOfWeekOutput struct{ *pulumi.OutputState }

func (RotationDayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationDayOfWeek)(nil)).Elem()
}

func (o RotationDayOfWeekOutput) ToRotationDayOfWeekOutput() RotationDayOfWeekOutput {
	return o
}

func (o RotationDayOfWeekOutput) ToRotationDayOfWeekOutputWithContext(ctx context.Context) RotationDayOfWeekOutput {
	return o
}

func (o RotationDayOfWeekOutput) ToRotationDayOfWeekPtrOutput() RotationDayOfWeekPtrOutput {
	return o.ToRotationDayOfWeekPtrOutputWithContext(context.Background())
}

func (o RotationDayOfWeekOutput) ToRotationDayOfWeekPtrOutputWithContext(ctx context.Context) RotationDayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RotationDayOfWeek) *RotationDayOfWeek {
		return &v
	}).(RotationDayOfWeekPtrOutput)
}

func (o RotationDayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RotationDayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RotationDayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RotationDayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RotationDayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RotationDayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RotationDayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (RotationDayOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationDayOfWeek)(nil)).Elem()
}

func (o RotationDayOfWeekPtrOutput) ToRotationDayOfWeekPtrOutput() RotationDayOfWeekPtrOutput {
	return o
}

func (o RotationDayOfWeekPtrOutput) ToRotationDayOfWeekPtrOutputWithContext(ctx context.Context) RotationDayOfWeekPtrOutput {
	return o
}

func (o RotationDayOfWeekPtrOutput) Elem() RotationDayOfWeekOutput {
	return o.ApplyT(func(v *RotationDayOfWeek) RotationDayOfWeek {
		if v != nil {
			return *v
		}
		var ret RotationDayOfWeek
		return ret
	}).(RotationDayOfWeekOutput)
}

func (o RotationDayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RotationDayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RotationDayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RotationDayOfWeekInput is an input type that accepts values of the RotationDayOfWeek enum
// A concrete instance of `RotationDayOfWeekInput` can be one of the following:
//
//	RotationDayOfWeekMon
//	RotationDayOfWeekTue
//	RotationDayOfWeekWed
//	RotationDayOfWeekThu
//	RotationDayOfWeekFri
//	RotationDayOfWeekSat
//	RotationDayOfWeekSun
type RotationDayOfWeekInput interface {
	pulumi.Input

	ToRotationDayOfWeekOutput() RotationDayOfWeekOutput
	ToRotationDayOfWeekOutputWithContext(context.Context) RotationDayOfWeekOutput
}

var rotationDayOfWeekPtrType = reflect.TypeOf((**RotationDayOfWeek)(nil)).Elem()

type RotationDayOfWeekPtrInput interface {
	pulumi.Input

	ToRotationDayOfWeekPtrOutput() RotationDayOfWeekPtrOutput
	ToRotationDayOfWeekPtrOutputWithContext(context.Context) RotationDayOfWeekPtrOutput
}

type rotationDayOfWeekPtr string

func RotationDayOfWeekPtr(v string) RotationDayOfWeekPtrInput {
	return (*rotationDayOfWeekPtr)(&v)
}

func (*rotationDayOfWeekPtr) ElementType() reflect.Type {
	return rotationDayOfWeekPtrType
}

func (in *rotationDayOfWeekPtr) ToRotationDayOfWeekPtrOutput() RotationDayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(RotationDayOfWeekPtrOutput)
}

func (in *rotationDayOfWeekPtr) ToRotationDayOfWeekPtrOutputWithContext(ctx context.Context) RotationDayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RotationDayOfWeekPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactChannelChannelTypeInput)(nil)).Elem(), ContactChannelChannelType("SMS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContactChannelChannelTypePtrInput)(nil)).Elem(), ContactChannelChannelType("SMS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContactTypeInput)(nil)).Elem(), ContactType("PERSONAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContactTypePtrInput)(nil)).Elem(), ContactType("PERSONAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*RotationDayOfWeekInput)(nil)).Elem(), RotationDayOfWeek("MON"))
	pulumi.RegisterInputType(reflect.TypeOf((*RotationDayOfWeekPtrInput)(nil)).Elem(), RotationDayOfWeek("MON"))
	pulumi.RegisterOutputType(ContactChannelChannelTypeOutput{})
	pulumi.RegisterOutputType(ContactChannelChannelTypePtrOutput{})
	pulumi.RegisterOutputType(ContactTypeOutput{})
	pulumi.RegisterOutputType(ContactTypePtrOutput{})
	pulumi.RegisterOutputType(RotationDayOfWeekOutput{})
	pulumi.RegisterOutputType(RotationDayOfWeekPtrOutput{})
}
