# Kubernetes Scheduled Queries

Experiment to run SQL queries against Redshift using Kubernetes CronJobs.

## Requirements

To do this you'll need a Kubernetes cluster with the CronJob API enabled as stated in their docs. Make sure you have access to the Postgress/Redshift database you want to run queries against!

## Usage

Create a `redshift` secrets resource in Kubernetes with the environment variables showed in [`cronjob.yaml`](kubernetes/cronjob.yaml). Then you can simply deploy the CronJob resource `kubectl create -f cronjob.yaml`.
