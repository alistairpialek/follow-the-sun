CDK_DEPLOY_ACCOUNT := $(shell aws sts get-caller-identity --profile follow_the_sun --output text --query 'Account')
CDK_DEPLOY_REGION := $(shell aws configure get region)

CMD := cdk
CDK = docker run -it \
	-e CDK_DEPLOY_ACCOUNT=$(CDK_DEPLOY_ACCOUNT) \
	-e CDK_DEPLOY_REGION=$(CDK_DEPLOY_REGION) \
	-e AWS_PROFILE=follow_the_sun \
	-w /src \
	-v ~/.aws:/root/.aws \
	-v $(PWD):/src \
	follow-the-sun $(CMD)

clean:
	rm -rf cdk.out
#
# Build the CDK docker container used for running CDK code.
#
build:
	docker build -t follow-the-sun .

#
# Used to jump into the container for debugging.
#
shell: CMD=/bin/bash
shell:
	$(CDK)

#
# This only needs to be run the once for an AWS account.
#
bootstrap:
	$(CDK) bootstrap aws://$(CDK_DEPLOY_ACCOUNT)/$(CDK_DEPLOY_REGION)
	$(CDK) bootstrap aws://$(CDK_DEPLOY_ACCOUNT)/us-west-1

#
# Good for testing CloudFormation output.
#
synth: clean build
	$(CDK) synth

#
# Deploy CloudFormation scripts.
#
deploy: clean build worker-build
	$(CDK) deploy --require-approval never --all

#
# Build go binary used by the lambda deployed by CDK code.
#
worker-build: CMD=go build -o function/bootstrap function/main.go
worker-build:
	$(CDK)

#
# Can be used to test the work by invoking it via browser URL.
# http://localhost:9000/2015-03-31/functions/function/invocations
#
worker-run: worker-build
	docker run -it \
	-e AWS_REGION=us-west-1 \
	-v $(PWD)/function:/var/task \
	-p 9000:8080 \
	public.ecr.aws/lambda/go:1.2024.07.10.12 bootstrap
