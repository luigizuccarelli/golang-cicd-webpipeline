#!/bin/bash

LOG_LEVEL=info
SERVER_PORT=9002
VERSION=1.0.1
NAME=golang-cicd-webpipeline
TEMPLATE_FILE=html-templates/template.html
CICD_CONSOLE_DIR=../console

export LOG_LEVEL NAME SERVER_PORT VERSION TEMPLATE_FILE CICD_CONSOLE_DIR

./build/microservice
