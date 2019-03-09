# Google Assistant chat written in go

This repository is a simple example of how to use Golang, Docker, Terraform and actions on google.

The app is deployed to Google Assistant and you can try it here: [stranger chat](https://assistant.google.com/services/a/uid/000000ee76914812?hl=en)

![Stranger chat](https://i.imgur.com/zkdo7jn.png)

# Prerequisites
- Installed [Terraform](https://www.terraform.io/)
- Installed [Actions on google command line interface](https://developers.google.com/actions/tools/gactions-cli)
- Created [Google cloud](https://cloud.google.com) project
- Created [Actions on google](https://developers.google.com/actions) project

# Getting started
The app consists of a service in golang, an infrastructure provided by terraform and actions on google configuration files.

## Service configuration

Set the service domain in `chat-service/service/Config.go`. The HTTPS is provided by [cert magic](https://github.com/mholt/certmagic) library.

## Terraform configuration

The terraform configuration file is designed to work with google cloud. Follow [this tutorial](https://cloud.google.com/video-intelligence/docs/common/auth) to get the JSON file with credentials and
copy it to `terraform/config/google_cloud_credentials.json`

Check out `terraform/config/outputs.tf` to set the domain and your google cloud project id.

## Deploying service

There are two helpful scripts available:

`/scripts/depoly.sh` - creates infrastructure on Google Cloud

`/scripts/update.sh` - builds the service, uploads binaries to virtual machine and spins up a Docker instance

## Deploying actions on google

The last step is to configure actions on google. The `google-actions/` contains configuration files for every supported language.
You need change the service domain.

To apply the changes `gactions update --action_package action.en.json --action_package action.pl.json --project {googleActionsProjectId}`
