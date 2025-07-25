// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.
type App struct {
	pulumi.CustomResourceState

	// The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// Unique Id for the Amplify App.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Name for the Amplify App.
	AppName pulumi.StringOutput `pulumi:"appName"`
	// ARN for the Amplify App.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig AppAutoBranchCreationConfigPtrOutput `pulumi:"autoBranchCreationConfig"`
	// The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig AppBasicAuthConfigPtrOutput `pulumi:"basicAuthConfig"`
	// The build specification (build spec) for an Amplify app.
	BuildSpec pulumi.StringPtrOutput `pulumi:"buildSpec"`
	// The cache configuration for the Amplify app. If you don't specify the cache configuration `type` , Amplify uses the default `AMPLIFY_MANAGED` setting.
	CacheConfig AppCacheConfigPtrOutput `pulumi:"cacheConfig"`
	// The Amazon Resource Name (ARN) of the IAM role for an SSR app. The Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
	ComputeRoleArn pulumi.StringPtrOutput `pulumi:"computeRoleArn"`
	// The custom HTTP headers for an Amplify app.
	CustomHeaders pulumi.StringPtrOutput `pulumi:"customHeaders"`
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules AppCustomRuleArrayOutput `pulumi:"customRules"`
	// Default domain for the Amplify App.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// The description of the Amplify app.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Automatically disconnect a branch in Amplify Hosting when you delete a branch from your Git repository.
	EnableBranchAutoDeletion pulumi.BoolPtrOutput `pulumi:"enableBranchAutoDeletion"`
	// The environment variables for the Amplify app.
	//
	// For a list of the environment variables that are accessible to Amplify by default, see [Amplify Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the *Amplify Hosting User Guide* .
	EnvironmentVariables AppEnvironmentVariableArrayOutput `pulumi:"environmentVariables"`
	// AWS Identity and Access Management ( IAM ) service role for the Amazon Resource Name (ARN) of the Amplify app.
	IamServiceRole pulumi.StringPtrOutput `pulumi:"iamServiceRole"`
	// The configuration details that apply to the jobs for an Amplify app.
	JobConfig AppJobConfigPtrOutput `pulumi:"jobConfig"`
	// The name of the Amplify app.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	OauthToken pulumi.StringPtrOutput `pulumi:"oauthToken"`
	// The platform for the Amplify app. For a static app, set the platform type to `WEB` . For a dynamic server-side rendered (SSR) app, set the platform type to `WEB_COMPUTE` . For an app requiring Amplify Hosting's original SSR support only, set the platform type to `WEB_DYNAMIC` .
	//
	// If you are deploying an SSG only app with Next.js version 14 or later, you must set the platform type to `WEB_COMPUTE` and set the artifacts `baseDirectory` to `.next` in the application's build settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the *Amplify Hosting User Guide* .
	Platform AppPlatformPtrOutput `pulumi:"platform"`
	// The Git repository for the Amplify app.
	Repository pulumi.StringPtrOutput `pulumi:"repository"`
	// The tag for an Amplify app.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		args = &AppArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("aws-native:amplify:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("aws-native:amplify:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	AccessToken *string `pulumi:"accessToken"`
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig *AppAutoBranchCreationConfig `pulumi:"autoBranchCreationConfig"`
	// The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig *AppBasicAuthConfig `pulumi:"basicAuthConfig"`
	// The build specification (build spec) for an Amplify app.
	BuildSpec *string `pulumi:"buildSpec"`
	// The cache configuration for the Amplify app. If you don't specify the cache configuration `type` , Amplify uses the default `AMPLIFY_MANAGED` setting.
	CacheConfig *AppCacheConfig `pulumi:"cacheConfig"`
	// The Amazon Resource Name (ARN) of the IAM role for an SSR app. The Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
	ComputeRoleArn *string `pulumi:"computeRoleArn"`
	// The custom HTTP headers for an Amplify app.
	CustomHeaders *string `pulumi:"customHeaders"`
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules []AppCustomRule `pulumi:"customRules"`
	// The description of the Amplify app.
	Description *string `pulumi:"description"`
	// Automatically disconnect a branch in Amplify Hosting when you delete a branch from your Git repository.
	EnableBranchAutoDeletion *bool `pulumi:"enableBranchAutoDeletion"`
	// The environment variables for the Amplify app.
	//
	// For a list of the environment variables that are accessible to Amplify by default, see [Amplify Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the *Amplify Hosting User Guide* .
	EnvironmentVariables []AppEnvironmentVariable `pulumi:"environmentVariables"`
	// AWS Identity and Access Management ( IAM ) service role for the Amazon Resource Name (ARN) of the Amplify app.
	IamServiceRole *string `pulumi:"iamServiceRole"`
	// The configuration details that apply to the jobs for an Amplify app.
	JobConfig *AppJobConfig `pulumi:"jobConfig"`
	// The name of the Amplify app.
	Name *string `pulumi:"name"`
	// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	OauthToken *string `pulumi:"oauthToken"`
	// The platform for the Amplify app. For a static app, set the platform type to `WEB` . For a dynamic server-side rendered (SSR) app, set the platform type to `WEB_COMPUTE` . For an app requiring Amplify Hosting's original SSR support only, set the platform type to `WEB_DYNAMIC` .
	//
	// If you are deploying an SSG only app with Next.js version 14 or later, you must set the platform type to `WEB_COMPUTE` and set the artifacts `baseDirectory` to `.next` in the application's build settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the *Amplify Hosting User Guide* .
	Platform *AppPlatform `pulumi:"platform"`
	// The Git repository for the Amplify app.
	Repository *string `pulumi:"repository"`
	// The tag for an Amplify app.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	AccessToken pulumi.StringPtrInput
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig AppAutoBranchCreationConfigPtrInput
	// The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig AppBasicAuthConfigPtrInput
	// The build specification (build spec) for an Amplify app.
	BuildSpec pulumi.StringPtrInput
	// The cache configuration for the Amplify app. If you don't specify the cache configuration `type` , Amplify uses the default `AMPLIFY_MANAGED` setting.
	CacheConfig AppCacheConfigPtrInput
	// The Amazon Resource Name (ARN) of the IAM role for an SSR app. The Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
	ComputeRoleArn pulumi.StringPtrInput
	// The custom HTTP headers for an Amplify app.
	CustomHeaders pulumi.StringPtrInput
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules AppCustomRuleArrayInput
	// The description of the Amplify app.
	Description pulumi.StringPtrInput
	// Automatically disconnect a branch in Amplify Hosting when you delete a branch from your Git repository.
	EnableBranchAutoDeletion pulumi.BoolPtrInput
	// The environment variables for the Amplify app.
	//
	// For a list of the environment variables that are accessible to Amplify by default, see [Amplify Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the *Amplify Hosting User Guide* .
	EnvironmentVariables AppEnvironmentVariableArrayInput
	// AWS Identity and Access Management ( IAM ) service role for the Amazon Resource Name (ARN) of the Amplify app.
	IamServiceRole pulumi.StringPtrInput
	// The configuration details that apply to the jobs for an Amplify app.
	JobConfig AppJobConfigPtrInput
	// The name of the Amplify app.
	Name pulumi.StringPtrInput
	// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	OauthToken pulumi.StringPtrInput
	// The platform for the Amplify app. For a static app, set the platform type to `WEB` . For a dynamic server-side rendered (SSR) app, set the platform type to `WEB_COMPUTE` . For an app requiring Amplify Hosting's original SSR support only, set the platform type to `WEB_DYNAMIC` .
	//
	// If you are deploying an SSG only app with Next.js version 14 or later, you must set the platform type to `WEB_COMPUTE` and set the artifacts `baseDirectory` to `.next` in the application's build settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the *Amplify Hosting User Guide* .
	Platform AppPlatformPtrInput
	// The Git repository for the Amplify app.
	Repository pulumi.StringPtrInput
	// The tag for an Amplify app.
	Tags aws.TagArrayInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

// The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
//
// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
//
// You must specify either `AccessToken` or `OauthToken` when you create a new app.
//
// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
func (o AppOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.AccessToken }).(pulumi.StringPtrOutput)
}

// Unique Id for the Amplify App.
func (o AppOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// Name for the Amplify App.
func (o AppOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.AppName }).(pulumi.StringOutput)
}

// ARN for the Amplify App.
func (o AppOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Sets the configuration for your automatic branch creation.
func (o AppOutput) AutoBranchCreationConfig() AppAutoBranchCreationConfigPtrOutput {
	return o.ApplyT(func(v *App) AppAutoBranchCreationConfigPtrOutput { return v.AutoBranchCreationConfig }).(AppAutoBranchCreationConfigPtrOutput)
}

// The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format `user:password` .
func (o AppOutput) BasicAuthConfig() AppBasicAuthConfigPtrOutput {
	return o.ApplyT(func(v *App) AppBasicAuthConfigPtrOutput { return v.BasicAuthConfig }).(AppBasicAuthConfigPtrOutput)
}

// The build specification (build spec) for an Amplify app.
func (o AppOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.BuildSpec }).(pulumi.StringPtrOutput)
}

// The cache configuration for the Amplify app. If you don't specify the cache configuration `type` , Amplify uses the default `AMPLIFY_MANAGED` setting.
func (o AppOutput) CacheConfig() AppCacheConfigPtrOutput {
	return o.ApplyT(func(v *App) AppCacheConfigPtrOutput { return v.CacheConfig }).(AppCacheConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM role for an SSR app. The Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
func (o AppOutput) ComputeRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.ComputeRoleArn }).(pulumi.StringPtrOutput)
}

// The custom HTTP headers for an Amplify app.
func (o AppOutput) CustomHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.CustomHeaders }).(pulumi.StringPtrOutput)
}

// The custom rewrite and redirect rules for an Amplify app.
func (o AppOutput) CustomRules() AppCustomRuleArrayOutput {
	return o.ApplyT(func(v *App) AppCustomRuleArrayOutput { return v.CustomRules }).(AppCustomRuleArrayOutput)
}

// Default domain for the Amplify App.
func (o AppOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// The description of the Amplify app.
func (o AppOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Automatically disconnect a branch in Amplify Hosting when you delete a branch from your Git repository.
func (o AppOutput) EnableBranchAutoDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *App) pulumi.BoolPtrOutput { return v.EnableBranchAutoDeletion }).(pulumi.BoolPtrOutput)
}

// The environment variables for the Amplify app.
//
// For a list of the environment variables that are accessible to Amplify by default, see [Amplify Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the *Amplify Hosting User Guide* .
func (o AppOutput) EnvironmentVariables() AppEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *App) AppEnvironmentVariableArrayOutput { return v.EnvironmentVariables }).(AppEnvironmentVariableArrayOutput)
}

// AWS Identity and Access Management ( IAM ) service role for the Amazon Resource Name (ARN) of the Amplify app.
func (o AppOutput) IamServiceRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.IamServiceRole }).(pulumi.StringPtrOutput)
}

// The configuration details that apply to the jobs for an Amplify app.
func (o AppOutput) JobConfig() AppJobConfigPtrOutput {
	return o.ApplyT(func(v *App) AppJobConfigPtrOutput { return v.JobConfig }).(AppJobConfigPtrOutput)
}

// The name of the Amplify app.
func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
//
// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
//
// You must specify either `OauthToken` or `AccessToken` when you create a new app.
//
// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
func (o AppOutput) OauthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.OauthToken }).(pulumi.StringPtrOutput)
}

// The platform for the Amplify app. For a static app, set the platform type to `WEB` . For a dynamic server-side rendered (SSR) app, set the platform type to `WEB_COMPUTE` . For an app requiring Amplify Hosting's original SSR support only, set the platform type to `WEB_DYNAMIC` .
//
// If you are deploying an SSG only app with Next.js version 14 or later, you must set the platform type to `WEB_COMPUTE` and set the artifacts `baseDirectory` to `.next` in the application's build settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the *Amplify Hosting User Guide* .
func (o AppOutput) Platform() AppPlatformPtrOutput {
	return o.ApplyT(func(v *App) AppPlatformPtrOutput { return v.Platform }).(AppPlatformPtrOutput)
}

// The Git repository for the Amplify app.
func (o AppOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Repository }).(pulumi.StringPtrOutput)
}

// The tag for an Amplify app.
func (o AppOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *App) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInput)(nil)).Elem(), &App{})
	pulumi.RegisterOutputType(AppOutput{})
}
