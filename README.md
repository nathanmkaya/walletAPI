## Simple Wallet API
[![Go Report Card](https://goreportcard.com/badge/github.com/nathanmkaya/walletAPI)](https://goreportcard.com/report/github.com/nathanmkaya/walletAPI)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ebfaf44753654f8d8158aa0571f093fc)](https://www.codacy.com/manual/nathanmkaya/walletAPI?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=nathanmkaya/walletAPI&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://api.codacy.com/project/badge/Coverage/ebfaf44753654f8d8158aa0571f093fc)](https://www.codacy.com/manual/nathanmkaya/walletAPI?utm_source=github.com&utm_medium=referral&utm_content=nathanmkaya/walletAPI&utm_campaign=Badge_Coverage)
[![Coverage Status](https://coveralls.io/repos/github/nathanmkaya/walletAPI/badge.svg?branch=master)](https://coveralls.io/github/nathanmkaya/walletAPI?branch=master)
[![Build Status](https://travis-ci.com/nathanmkaya/walletAPI.svg?branch=master)](https://travis-ci.com/nathanmkaya/walletAPI)

Write a simple wallet REST API.

We expect the following:
1. account deposits
2. account withdraws
3. balance enquiry
4. Mini statement

Added advantage:
1. Using Go Modules
2. Using Postgres.
3. Unit tests
4. Code coverage on the tests

Clearly document how to set up and run your application.

To run the app
```shell script
go run ./cmd/server
```

To setup the app
```shell script
go get ./...
```

To run tests
```shell script
go test ./...
```

To run tests with coverage
```shell script
go test ./... -cover
```
