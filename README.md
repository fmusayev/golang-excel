[![CircleCI](https://circleci.com/gh/fmusayev/golang-excel.svg?style=svg)](https://circleci.com/gh/fmusayev/golang-excel)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fmusayev_golang-excel&metric=alert_status)](https://sonarcloud.io/dashboard?id=fmusayev_golang-excel)
[![Go Report Card](https://goreportcard.com/badge/github.com/fmusayev/golang-excel)](https://goreportcard.com/report/github.com/fmusayev/golang-excel)
# Generate excel file for input data using custom tags and reflection

Name: *golang-excel* \
Created By: @fmusayev \

### Run

* Run \
`go run main.go`
* Run using custom port \
`go run main.go -port :5000`
* Run in disturbuted env. (set log format to json for log parsing tools, like Datadog) \
`go run main.go -port :80 -log json`
