# Follow The Sun

Renewables aware AWS infrastructure example.

## Components

1. Go lambda that runs every hour to determine the cleanest country to run compute workloads by updating Route53
   weighted routing records to desired country endpoint.
2. Go CDK deployment code which:
   1. Deploys follow-the-sun lambda and run schedule.
   1. Deploys workload lambdas that run a simple Go webserver.

## Getting started

### 1. Run bootstrap step

[What is CDK bootstrapping?][1]<br>
[When do I need to bootstrap an AWS account?][2]

```
make bootstrap
```

### 2. Check for code errors

Prepare for deployment by synthesizing a CloudFormation template.

```
make synth
```

[1]: https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html
[2]: https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping-env.html#bootstrapping-env-when

### 3. Perform deployment

Apply CloudFormation scripts to your AWS account.

```
make deploy
```
