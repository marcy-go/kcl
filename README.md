# Amazon Kinesis Client Library for Go

This library provides an interface to the KCL MultiLangDaemon, which is part of the
[Amazon Kinesis Client Library](https://github.com/awslabs/amazon-kinesis-client). This interface manages the
interaction with the MultiLangDaemon so that developers can focus on
implementing their record processor executable. A record processor executable
typically looks something like:

```go
package main

import(
	"github.com/marcy-go/kcl"
)

type Sample struct {}

func (s *Sample) Init(str string) error {
	return nil
}

func (s *Sample) ProcessRecords(rs []*kcl.Record, cp *kcl.CheckPointer) error {
	return nil
}

func (s *Sample) Shutdown(cp *kcl.CheckPointer, reason string) error {
	return nil
}

func main() {
	var rp kcl.RecordProcessor
	rp = &Sample{}
	p := kcl.NewProcess(rp)
	p.Run()
}
```

Note, the initial implementation of this library is largely based on the reference [Amazon Kinesis Client Library for Python](https://github.com/awslabs/amazon-kinesis-client-python) provided by Amazon.

## Usage
#### 1. Write and build your `RecordProcessor`.  
If your want to run a sample implementation, you get the sample binary from here:  
https://drone.io/github.com/marcy-go/kclsample/files
#### 2. Write `properties` file like this:  
```properties
executableName          = /path/to/kclsample
streamName              = kclsample
applicationName         = GoKCLSample
AWSCredentialsProvider  = DefaultAWSCredentialsProviderChain
processingLanguage      = go/1.3.1
initialPositionInStream = TRIM_HORIZON
regionName              = ap-northeast-1
```
#### 3. Get `kclhelper` binary for your environment from here:  
https://drone.io/github.com/marcy-go/kclhelper/files
#### 4. Exec `kclhelper` command.
```sh
# Get Jar files from Maven projects
kclhelper setup
# Run KCL
kclhelper run -p <path-to-properties> [-j <path-to-java>]
```

# Changelog

See [CHANGELOG](https://github.com/marcy-go/kcl/blob/master/CHANGELOG.md)

## Contributing

1. Fork it ( https://github.com/marcy-go/kcl/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

# License

[MIT License](https://github.com/marcy-go/kcl/blob/master/LICENSE.txt)
