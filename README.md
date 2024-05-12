# Comparison of Go Web App and TensorFlow Serving API in Kubernetes

* [General info](#general-info)
* [Technologies](#technologies)
* [Project structure](#project-structure)
* [Setup](#setup)
* [App view](#app-view)

## General info
This project was created to conduct performance tests and comparisons between two different applications that convert temperatures from Fahrenheit to Celsius in a Kubernetes environment.

## Technologies 
* Go
* Python
* Docker
* Helm Charts
* Ansible
* Kubernetes cluster (1 master node, 3 worker nodes)
* NFS server (not required but the CronJobs will be happy to use it)
* Github actions

## Project structure
* **.github** - here are all the tasks that run in the pipeline. They are invoked every time there is a new push to main or pull request. Workflow consists of 4 stages:
  - lint - the linters analyse the code (Go and Python) and look for potential errors, also the tests are ran for Go project
  - build_and_push - docker images are built and pushed to docker hub
  - helm_lint - the linter verifies, whether the charts will be created correctly
  - release - helm charts are packaged and pushed to github pages on the same repository (available to be downloaded from https://filbary.github.io/test)
* **src** - in this directory there are 3 folders, each representing different app:
  - webapp - The Go backend server has a single endpoint /convert that accepts POST requests with the body: {"fahrenheit": float}, returning the Celsius value along with a special random identifier, which is unique for each instance of this app
  - serving - The TensorFlow Serving system includes a computational graph that converts Fahrenheit temperatures to Celsius. It offers an endpoint, /v1/models/model:predict, which accepts POST requests with the body: {"inputs": float}.
  - locust - a client app which performs load tests and gathers statistics
* **charts** - 3 helm charts (for each app)
  - webapp, serving - Deployment, Service, Horizontal Pod Autoscaler (scales the number of pod based on CPU and Memory usage)
  - locust - Deployment, Service, ConfigMap (with locust.conf), 2 CronJobs (each CronJob can create a job, which starts the locust pod, performs the tests and then saves the statistics to persistent volumes)
  Each helm chart is configurable by values that can modify for instance: number of min and max replicas, app configs, cronjob configs, exposed ports...
* **configuration** - this folder contains ansible tasks: namespace, storage class, pvc's creation and deployments of helm charts (in order for nfs-provisioner and CronJobs to work, the NFS server must be configured!). There are 2 scripts: run.sh (deployment of everything) and run_results.sh (it fetches the created statistics from pv's and copies them to dir: "configuration/results/")
* **results** - here are statistics and reports downloaded from locust web UI

## Setup
To build and push Docker images and Helm charts, simply commit to the main branch or create a pull request. To deploy the applications, start the run.sh script located in the configuration directory. To initiate tests, either navigate to the Locust web UI exposed by the node port, port forward to Locust, or trigger the CronJobs :)
 
## App view
![image](https://github.com/filbary/test/assets/62804065/c37fb375-5a39-43d7-8cea-6051daebee19)
