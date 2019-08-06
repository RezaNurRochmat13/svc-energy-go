# Service Energy Using Golang

This project is example manage meter resource data using Golang.

## Table of Contents

List of all readme contents:

- [Key Features](#key-features)
- [Built With](#built-with)
- [Deployment](#deployment)
- [API Reference](#api-reference)
- [Release History](#release-history)
- [Authors](#authors)

## Key Features

List of all features/functionalities handled by this app/service:

 - Get All Meter Read With Pagination
 - Get Detail Meter Read by Meter Read ID
 - Create New Meter Read Data
 - Update Meter Read Data
 - Delete Meter Read Data

## Built With
List of all core technologies used to build this app/service along with their functions, versions (if any), and link to their online pages, grouped based on development stack:

### Backend
- [Golang](https://golang.org/) - Programming Language [v1.12.4]
- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP Routing Framework for Golang [v1.3.0]
- [Dep Package Management](https://github.com/golang/dep) - Dep Package Management For Golang [v0.5.1]
- [Viper Configuration Tools](https://github.com/spf13/viper) - Viper Configuration Tools [v1.3.2]
- [Go SQL Driver](https://github.com/go-sql-driver/mysql) - SQL Driver For MySQL DB [v1.4.1]

### Frontend
- None

### Automation Test
- For Running All Test, run command go test

### Database
- [MySQL](https://www.percona.com/) - Core Database [v5.7]

## Deployment

List all required steps to deploy this app/service in server, like environment variables, server requirements, amount of compute resources (CPU, RAM), and dependency services that will communicate with this app/service.

### Running Locally
- Pull this repo using SSH or HTTPS
- Import DB in your local DB (DB using MySQL)
- Install dep package in `https://github.com/golang/dep`
- After install dep, run command `dep ensure` to download dependencies
- After finish install dep, you can running apps using command shell script aka `sh run.sh`

### Running using Docker
- Pull this repo using SSH or HTTPS
- Import DB in your local DB (DB using MySQL)
- Install dep package in `https://github.com/golang/dep`
- After install dep, run command `dep ensure` to download dependencies
- Compile your apps to binary using command shell aka `sh compile.sh`
- After successfully compiled, build docker images using command `docker build -t $NAME_IMAGES:TAGS .`
- After build docker images successful, you can run your apps in docker using command `docker run --rm -it -p $PORT:$PORT $NAME_IMAGES`

### Runing in Kubernetes
- On Going

### Environment Variables

**MySQL Config:**

- `MYSQL_DATABASE` Name of the database using MySQL.
- `MYSQL_USER` Username for the database using MySQL.
- `MYSQL_PASSWORD` Password for the database using MySQL.
- `MYSQL_HOST` Hostname of the database server using MySQL.

## API Reference

Depending on the size of this app/service APIs, if it is small and simple enough the reference docs can be added to this README. For medium to larger app/service please use Gitlab wiki or provide a link to where the API reference docs live.

### API Docs

## Release History

List of all released versions of this app/service along with their version logs, sorted from newest to oldest:

- 0.2.1
  - CHANGE: Update docs (module code remains unchanged)
- 0.2.0
  - CHANGE: Remove `setDefaultXYZ()`
  - ADD: Add `init()`
- 0.1.1
  - FIX: Crash when calling `baz()`
- 0.1.0
  - The first proper release
  - CHANGE: rename `foo()` to `bar()`
- 0.0.1
  - Work in progress

## Authors
List of all people working on this app/service along with their respective roles & contact emails:

- [Reja Nur Rochmat](mailto:hal.rezanur@uii.ac.id) - Software Engineer
