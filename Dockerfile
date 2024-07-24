ARG AWS_CDK_VERSION=2.149.0
FROM node:22-bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates

RUN npm install -g \
    aws-cdk@${AWS_CDK_VERSION}

COPY --from=golang:1.22-bookworm /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
