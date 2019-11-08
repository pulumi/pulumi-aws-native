import * as pulumi from "@pulumi/pulumi";
import * as cloudformation from "@pulumi/cloudformation";

const region = cloudformation.region!;

const region2S3WebsiteSuffix: any = {
    "us-east-1":      ".s3-website-us-east-1.amazonaws.com",
    "us-west-1":      ".s3-website-us-west-1.amazonaws.com",
    "us-west-2":      ".s3-website-us-west-2.amazonaws.com",
    "eu-west-1":      ".s3-website-eu-west-1.amazonaws.com",
    "eu-west-2":      ".s3-website-eu-west-2.amazonaws.com",
    "eu-west-3":      ".s3-website-eu-west-3.amazonaws.com",
    "ap-northeast-1": ".s3-website-ap-northeast-1.amazonaws.com",
    "ap-northeast-2": ".s3-website-ap-northeast-2.amazonaws.com",
    "ap-northeast-3": ".s3-website-ap-northeast-3.amazonaws.com",
    "ap-southeast-1": ".s3-website-ap-southeast-1.amazonaws.com",
    "ap-southeast-2": ".s3-website-ap-southeast-2.amazonaws.com",
    "ap-south-1":     ".s3-website-ap-south-1.amazonaws.com",
    "us-east-2":      ".s3-website-us-east-2.amazonaws.com",
    "ca-central-1":   ".s3-website-ca-central-1.amazonaws.com",
    "sa-east-1":      ".s3-website-sa-east-1.amazonaws.com",
    "cn-north-1":     ".s3-website.cn-north-1.amazonaws.com.cn",
    "cn-northwest-1": ".s3-website.cn-northwest-1.amazonaws.com.cn",
    "eu-central-1":   ".s3-website-eu-central-1.amazonaws.com",
    "eu-north-1":     ".s3-website-eu-north-1.amazonaws.com"
};

const s3BucketForWebsiteContent = new cloudformation.s3.Bucket("s3BucketForWebsiteContent", {
    AccessControl: "PublicRead",
    WebsiteConfiguration: {
        IndexDocument: "index.html",
        ErrorDocument: "error.html",
    },
});

const websiteCdn = new cloudformation.cloudfront.Distribution("websiteCdn", {
    DistributionConfig: {
        Comment: "CDN for S3-backed website",
        Enabled: true,
        DefaultCacheBehavior: {
            ForwardedValues: { QueryString: true },
            TargetOriginId: "only-origin",
            ViewerProtocolPolicy: "allow-all",
        },
        DefaultRootObject: "index.html",
        Origins: [{
            CustomOriginConfig: {
                HTTPPort: 80,
                HTTPSPort: 443,
                OriginProtocolPolicy: "http-only",
            },
            DomainName: pulumi.interpolate`${s3BucketForWebsiteContent}.${region2S3WebsiteSuffix[region]}`,
            Id: "only-origin",
        }],
    },
});
