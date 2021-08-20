# **Notification Service**

This repository for notification service – sms, email (comming soon), or push notification (comming soon). Built with ❤️ and Go

## Requirements

 - Golang version 1.12+
 - Basic knowledge about go mod https://github.com/golang/go/wiki/Modules
 - With new feature from Golang `1.11+` you can work on this project outside `GOPATH`

## Build
 - Build
    ```console
    $ go build
    ```

## Building and running tests

The software is designed to be able to run both on the machine on top of docker. You can see `Makefile` to see what
kind of targets supported.

### Integration Test

This test should be run with environment variable set.

```
make test
```

### Integration Test With Coverage

This test is to produce covarage report. We use `gocovmerge` to merge coverage reports from the packages. To run, use

```
make cover
```

## How to Use/Run


- run command `make docker` if you want to run with docker or just `make run`
- import postman collection from `docs/notif.postman_collection.json` into [postman app](https://www.postman.com/downloads/) and specify the `{{uri}}` variable with your running app host.

## What you can do

### Send SMS
request:
```zsh
curl --location --request POST '{{your/app/host}}/sms/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to": 628199999999,
    "message": "Hi, folks!"
}'
```
response:
```json
// success
{
    "success": true,
    "code": 200,
    "message": "Success",
    "error": null,
    "data": {
        "to": "628199999999",
        "message": "sms is 0",
        "status": "SENT"
    }
}
// failed
{
    "success": false,
    "code": 400,
    "message": "the error message",
    "error": {
        "Offset": 1
    },
    "data": null
}
```

### Enable/Disable SMS Vendor
request:
```zsh
curl --location --request POST '{{your/app/host}}/sms/vendor/toggle' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "ngehe",
    "status": true
}'
```
response:
```zsh
{
    "success": true,
    "code": 200,
    "message": "Success",
    "error": null,
    "data": {
        "name": "ngehe",
        "status": true
    }
}
```

### Get SMS Vendor
request:
```zsh
curl --location --request GET '{{your/app/host}}/sms'
```
response:
```json
{
    "success": true,
    "code": 200,
    "message": "Success",
    "error": null,
    "data": [
        {
            "active": true,
            "name": "bacrit"
        },
        {
            "active": false,
            "name": "ngehe"
        }
    ]
}
```
