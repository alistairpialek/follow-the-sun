package main

import (
	"follow-the-sun/config"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaCdkStackProps struct {
	awscdk.StackProps
}

func MainCdkStack(scope constructs.Construct, id string, props *LambdaCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	worker := awslambda.NewFunction(stack, jsii.String(config.FunctionName), &awslambda.FunctionProps{
		FunctionName: jsii.String(*stack.StackName() + "-" + config.FunctionName),
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		MemorySize:   jsii.Number(config.MemorySize),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(config.MaxDuration)),
		Code:         awslambda.AssetCode_FromAsset(jsii.String(config.CodePath), nil),
		Handler:      jsii.String(config.Handler),
	})

	workerUrl := worker.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	awscdk.NewCfnOutput(stack, jsii.String("WorkerURL"), &awscdk.CfnOutputProps{
		Value:       workerUrl.Url(),
		Description: jsii.String("Worker function public URL"),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	MainCdkStack(app, config.StackName, &LambdaCdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(os.Getenv("CDK_DEPLOY_ACCOUNT")),
				Region:  jsii.String("ap-southeast-2"),
			},
		},
	})

	MainCdkStack(app, config.StackName+"-us-west-1", &LambdaCdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(os.Getenv("CDK_DEPLOY_ACCOUNT")),
				Region:  jsii.String("us-west-1"),
			},
		},
	})

	app.Synth(nil)
}
