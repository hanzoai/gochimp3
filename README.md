# gochimp3
[![GoDoc][godoc-img]][godoc-url] [![Build Status][travis-img]][travis-url] [![Gitter chat][gitter-img]][gitter-url]

## Introduction
Golang client for [MailChimp API 3.0](http://developer.mailchimp.com/documentation/mailchimp/).

## Install
Install with `go get`:

```bash
$ go get github.com/hanzoai/gochimp3
```

## Usage
```go
package main

import (
    "fmt"
    "os"

    "github.com/hanzoai/gochimp3"
)

const (
    apiKey = "f6f6eb412g2b9677b00550d14d86db5e-us4"
)

func main() {
    client := gochimp3.New(apiKey)

    // Fetch list
	list, err := client.GetList("28a3d7a5", nil)
	if err != nil {
		fmt.Println("Failed to get list '%s'", listId)
		os.Exit(1)
	}

    // Add subscriber
    req := &gochimp3.MemberRequest{
        EmailAddress: "spam@hanzo.ai",
	Status: "subscribed",
    }

	if _, err := list.CreateMember(req); err != nil {
		fmt.Println("Failed to subscribe '%s'", req.EmailAddress)
		os.Exit(1)
	}
}
```

[godoc-img]:      https://godoc.org/github.com/hanzoai/gochimp3?status.svg
[godoc-url]:      https://godoc.org/github.com/hanzoai/gochimp3
[travis-img]:     https://img.shields.io/travis/hanzoai/gochimp3.svg
[travis-url]:     https://travis-ci.org/hanzoai/gochimp3
[gitter-img]:     https://badges.gitter.im/join-chat.svg
[gitter-url]:     https://gitter.im/hanzoai/chat

<!-- not used -->
[coveralls-img]:    https://coveralls.io/repos/hanzoai/gochimp3/badge.svg?branch=master&service=github
[coveralls-url]:    https://coveralls.io/github/hanzoai/gochimp3?branch=master
