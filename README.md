# Redsfhit Kubernetes Scheduled Queries 

Experiment to run SQL queries against Redshift using Kubernetes CronJobs.

## Requirements

To do this you'll need a Kubernetes cluster with the CronJob API enabled as stated in their docs. Make sure you have access to the Postgress/Redshift database you want to run queries against!

## Usage

First, create a `redshift` secrets resource in Kubernetes with the environment variables showed in [`cronjob.yaml`](kubernetes/cronjob.yaml). 

You'll need to tweak a bit this example to meet your needs:

- Update the file `query.sql` with your desired query.
- Change the image name in the Makefile and run `make build` to create the container. 
- Push your container to a Docker registry.
- You can now deploy the CronJob resource `kubectl create -f cronjob.yaml`.
