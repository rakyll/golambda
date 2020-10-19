# golambda

AWS Lambda Go functions made easy...

golambda allows you to build and deploy Lambda function in Go easily.

## Installation

```
$ go get github.com/rakyll/golambda
```

## Usage

The following command will build a zip file (main.zip) with the Go binary.

```
$ golambda build
```

Create a Lambda function from a zip file:

```
# by default, it uploads ./main.zip.
$ golambda create -name helloworld -role arn:aws:iam::951969755123:role/lamda
```

Update and publish a Lambda function:

```
# by default, it uploads ./main.zip.
$ golambda update -name helloworld
```

Note: This is a personal project and is not officially supported.
