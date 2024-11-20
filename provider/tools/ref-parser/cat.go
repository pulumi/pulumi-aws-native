package main

import (
	"regexp"
	"strings"
)

type Category interface {
	isCategory()
	Name() string
}

type simpleCategory string

func (s simpleCategory) isCategory() {}

func (s simpleCategory) Name() string {
	return string(s)
}

var Uncategorized Category = simpleCategory("Uncategorized")
var Undocumented Category = simpleCategory("Undocumented")
var RefReturnsArn Category = simpleCategory("RefReturnsArn")
var RefReturnsName Category = simpleCategory("RefReturnsName")
var RefReturnsID Category = simpleCategory("RefReturnsID")
var ShouldNotBeUsed Category = simpleCategory("ShouldNotBeUsed")

type categorizationRule struct {
	Category Category
	Pattern  *regexp.Regexp
}

var categorizationRules = []categorizationRule{
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile(`^The Amazon Resource Name ..ARN.. of the certificate authority.?.?$`),
	},
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile("When you pass the logical ID of this resource to the intrinsic .Ref..?function, .Ref.?.?returns the ARN"),
	},
	{
		Category: ShouldNotBeUsed,
		Pattern:  regexp.MustCompile("^This reference should not be used in CloudFormation templates"),
	},
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the Arn of the .* that is created by the"),
	},
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the .* ARN"),
	},
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the resource ARN"),
	},
	{
		Category: RefReturnsArn,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the Amazon Resource Name"),
	},
	{
		Category: RefReturnsName,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the value of [`][^`]*Name[`]"),
	},
	{
		Category: RefReturnsName,
		Pattern:  regexp.MustCompile("When the logical ID of this resource is provided to the Ref intrinsic function, .Ref. returns the .* name"),
	},
	{
		Category: RefReturnsName,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the name"),
	},
	{
		Category: RefReturnsName,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the .* name"),
	},
	{
		Category: RefReturnsID,
		Pattern:  regexp.MustCompile("^When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns the.*ID"),
	},
	{
		Category: RefReturnsID,
		Pattern:  regexp.MustCompile("When you pass the logical ID of this resource to the intrinsic .Ref.?.?function, .Ref.?.?returns a .* ID"),
	},
	// {
	// 	Name: "RefReturnsArn",
	// 	Pattern: regexp.MustCompile("When you pass the logical ID of an [^ ]+ resource to the intrinsic " +
	// 		"[`]Ref[`] function, the function returns the (ARN)"),
	// },
	// {
	// 	Name: "RefReturnsArn",
	// 	Pattern: regexp.MustCompile("When the logical ID of this resource is provided to the `Ref` " +
	// 		"intrinsic function, `Ref` returns the (ARN)"),
	// },
	// {
	// 	Name: "RefReturnsSomething",
	// 	Pattern: regexp.MustCompile("When you pass the logical ID of this resource to the intrinsic" +
	// 		" [`]Ref[`]function, `Ref`returns ([^,.]+)[,.]"),
	// },
}

func Categorize(rawRefSection string) Category {
	clean := strings.TrimSpace(rawRefSection)
	for _, rule := range categorizationRules {
		if mm := rule.Pattern.FindStringSubmatch(clean); mm != nil {
			return rule.Category
		}
	}
	return Uncategorized
}
