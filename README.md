# Operator Library

[![Godoc Reference](https://godoc.org/github.com/ankitrgadiya/operatorlib?status.svg)](https://godoc.org/github.com/ankitrgadiya/operatorlib)
[![Go Report Card](https://goreportcard.com/badge/github.com/ankitrgadiya/operatorlib)](https://goreportcard.com/report/github.com/ankitrgadiya/operatorlib)
[![Build Status](https://travis-ci.com/ankitrgadiya/operatorlib.svg?branch=master)](https://travis-ci.com/ankitrgadiya/operatorlib)
[![codecov](https://codecov.io/gh/ankitrgadiya/operatorlib/branch/master/graph/badge.svg)](https://codecov.io/gh/ankitrgadiya/operatorlib)
[![Maintainability](https://api.codeclimate.com/v1/badges/2c6f0689231cab164aad/maintainability)](https://codeclimate.com/github/ankitrgadiya/operatorlib/maintainability)
[![BSD License](https://img.shields.io/github/license/ankitrgadiya/operatorlib)](LICENSE)

## About

While working on many Kubernetes operators, I realised that a lot of
code is repetative across operators. Generating, creating, updating,
deleting objects is a common thing and yet the code is being repeated
in every operator. I decided to work on project which removes the the
repeatative code from all Kubernetes Operators. The benefits of this
approach are that common well tested and stable functions can be used
by all operators. Also, this kind of tries to reduce the complexity of
dealing with Kubernetes objects which (hopefully!) will lead to more
and more vendors building there own operators. This project also
attempts to reduce the overall work required to build the operator
which means operators for small projects can be build quickly.

## Usage

The Kubernetes interactions are abstracted away so you do not have to
care about the Client for the simple usage. For instance, if you want
to create a `ConfigMap` object using Operatorlib, import `configmap`
package and add the following lines of code whereever it seems fit.

```go
import "github.com/ankitrgadiya/operatorlib/pkg/configmap

...
	result, err := configmap.CreateOrUpdate(configmap.Conf{
		Name: "test",
		Namespace: "default",
	})
	if err != nil {
		return result, err
	}
...
```

This however will create an empty `Configmap`. To see all the options
check the [`Conf`
section](https://godoc.org/github.com/ankitrgadiya/operatorlib/pkg/configmap#Conf)
in the documentation. Check the list of supported objects
below. Operatorlib also provides low-level package `operation` which
can be used to create all non-supported objects (including custom
objects) as long as they implement
[`Object`](https://godoc.org/github.com/ankitrgadiya/operatorlib/pkg/interfaces#Object)
interface defines in `interface` package.

## Supported Objects

* [x] [`Configmap`](https://godoc.org/github.com/ankitrgadiya/operatorlib/pkg/configmap)
* [x] [`Secret`](https://godoc.org/github.com/ankitrgadiya/operatorlib/pkg/secret)
* [ ] `Service`
* [ ] `Pod`
* [ ] `Deployment`
* [ ] `StatefulSet`
* [ ] `Job`
* [ ] `CronJob`
* [ ] `Volume`
* [ ] `PersistentVolumeClaim`

## License

[BSD 3-Clause License](LICENSE)
