# Micro [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/micro?status.svg)](https://godoc.org/github.com/micro/micro) [![Travis CI](https://travis-ci.org/micro/micro.svg?branch=master)](https://travis-ci.org/micro/micro) [![Go Report Card](https://goreportcard.com/badge/micro/micro)](https://goreportcard.com/report/github.com/micro/micro)

Micro is a framework for distributed systems development in the Cloud and beyond.

## Overview

Micro addresses the key requirements for building distributed systems. It leverages the microservices
architecture pattern and provides a set of services which act as the building blocks of a platform. Micro deals
with the complexity of distributed systems and provides simpler programmable abstractions to build on.

## Features

The framework is composed of the following features:

- **Server:** A distributed systems runtime server composed of building block services which abstract away the underlying infrastructure 
and provide a programmable abstraction layer. Authentication, configuration, messaging, storage and more built in.

- **Client:** Multiple entrypoints through which you can access your services. Write services once and access them through every means 
you've already come to know. An API Gateway, gRPC proxy, slack bot, web dashboard and commmand line interface.

- **Library:** A Go library which makes it drop dead simple to write your services without having to piece together lines and lines of 
boilerplate. Auto configured and initialised by default, just import and get started quickly.

- **Plugins:** Micro is runtime and infrastructure agnostic. Each underlying building block service uses the Go Micro standard library 
to provide a pluggable foundation. We make it simple to use by pre-initialising of local use and the cloud.

## Install

Install from source

```
go get github.com/micro/micro/v2
```

Using a docker image

```
docker pull micro/micro
```

Latest release binaries

```
# MacOS
curl -fsSL https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh | /bin/bash

# Linux
wget -q  https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh -O - | /bin/bash

# Windows
powershell -Command "iwr -useb https://raw.githubusercontent.com/micro/micro/master/scripts/install.ps1 | iex"
```

## Getting Started

Run the server locally

```
micro server
```

Create a service

```
# generate a service (follow instructions in output)
micro new helloworld

# set to use server
micro env set server

# run the service
micro run helloworld

# list services
micro services

# call a service
micro call helloworld Helloworld.Call '{"name": "John"}'
```

## Usage

See all the options

```
micro --help
```

See the [docs](https://github.com/micro/docs) for detailed information on the architecture, installation and use of the platform.

## License

Micro is Apache 2.0 Licensed

