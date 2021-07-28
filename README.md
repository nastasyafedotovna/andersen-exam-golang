[![ci/cd](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/lint-build-test-push-deploy.yml/badge.svg?branch=main)](https://github.com/nastasyafedotovna/andersen-exam-golang/actions/workflows/lint-build-test-push-deploy.yml)

# Golang(gin framework) hello world web app
+ CI/CD: [**Github Actions**](https://docs.github.com/en/actions):
    + Docker images as artifacts
    + Export images to AWS ECR
    + Deploy to AWS ECS-EC2
+ SCM/ControlVersion: [**Git**](https://git-scm.com/)
+ Registry: [**AWS ECR**](https://aws.amazon.com/en/ecr/)
+ Infrastructure: [**Terraform**](https://www.terraform.io) + [**AWS**](https://aws.amazon.com/) 
+ IaC: [**Infrastructure repo**](https://github.com/nastasyafedotovna/andersen-exam-golang-infrastructure)
