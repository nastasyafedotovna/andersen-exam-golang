[![build-test](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/go.yml) [![golangci-lint](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/golint.yml/badge.svg?branch=main)](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/golint.yml) [![publish-image-to-ecr-and-deploy-to-ecs](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/aws.yml/badge.svg?branch=main)](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/aws.yml)

# Golang(gin framework) hello world web app
+ CI/CD: [**Github Actions**](https://docs.github.com/en/actions):
    + Docker images as artifacts
    + Export images to AWS ECR
    + Deploy to AWS ECS-EC2
+ SCM/ControlVersion: [**GitHub.com**](https://github.com)
+ Registry: **AWS ECR**
+ Infrastructure: [**Terraform**](https://www.terraform.io) + **AWS** 
+ IaC: [**Infrastructure**](https://github.com/nastasyafedotovna/andersen-exam-golang-infrastructure)
