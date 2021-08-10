// Copyright 2021, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.AwsNative.Logs;
using Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var name = new RandomString("name", new RandomStringArgs
        {
            Length = 8,
            Special = false,
            Upper = false
        });

        var logGroup = new LogGroup("log-test", new LogGroupArgs
        {
            LogGroupName = name.Result,
            RetentionInDays = 90
        });

        this.Arn = logGroup.Arn;
    }

    [Output]
    public Output<string> Arn { get; set; }
}
