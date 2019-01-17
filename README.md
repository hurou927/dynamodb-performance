# DynamoTest

## DB Setup

```
$ sls deploy
$ cd seed
$ yarn
$ node seed
```

## go

```                                   
$ go get github.com/jmespath/go-jmespath //require???   
$ go get github.com/aws/aws-sdk-go-v2                       
$ go get github.com/aws/aws-sdk-go
```

```
# aws-sdk-go
$ cd go
$ go build dynamo.go
$ ./dynamo

# aws-sdk-go2
$ cd go
$ go build dynamo-v2.go
$ ./dynamo-v2
```

## node

```
$ cd node
$ yarn
$ node get
```


## result

macbook pro
```
$ system_profiler SPHardwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro13,1
      Processor Name: Intel Core i5
      Processor Speed: 2 GHz
      Number of Processors: 1
      Total Number of Cores: 2
      L2 Cache (per Core): 256 KB
      L3 Cache: 4 MB
      Memory: 16 GB
```


|  Lang. |  Time |
| ------ | ---- |
|  go    |  294 ms  |
|  go(v2)|  295 ms  |
| node   |  191 ms| 