# Project Repo Title

A little introduction here...

<hr />

## **Requirements**
- Golang 1.20 or higher
- Openssl3

<hr />

## **Dependencies**
- [PymigrateDB](https://pypi.org/project/pymigratedb/)
- [Gcloud CLI](https://cloud.google.com/sdk/docs/install)

<hr />

## Setup
- To install project run ```make```.
    - It will create two new files in the project root, called: ".env" and "config.yaml"
    - Fill it with correct values to procceed with development

<hr />

## Run
- To run project execute ```make run``` into the terminal. It will start the API and serve the requests connected to the resources filled in the config files.

<hr />

## Tests
- To run and perform test cases, run the following command: ```make test```. It will begin the tests execution.
