[![CircleCI](https://circleci.com/gh/fmusayev/golang-excel.svg?style=svg)](https://circleci.com/gh/fmusayev/golang-excel)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fmusayev_golang-excel&metric=alert_status)](https://sonarcloud.io/dashboard?id=fmusayev_golang-excel)
[![Go Report Card](https://goreportcard.com/badge/github.com/fmusayev/golang-excel)](https://goreportcard.com/report/github.com/fmusayev/golang-excel)
# Excel generation

Name: *golang-excel* \
Created By: @fmusayev \
\
Sample web api type project for generating excel, and using custom tags in struct to specify the excel column in which the field will be inserted. Also added reflection when writing struct to excel for preventing fields listing. 


### Run

* Run \
`go run main.go`
* Run using custom port \
`go run main.go -port :5000`
* Run in disturbuted env. (set log format to json for log parsing tools, like Datadog) \
`go run main.go -port :80 -log json`
