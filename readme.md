## 1.At the command line, change to the directory that contains hello/hello.go.
## 2.Discover the Go install path, where the go command will install the current package
go list -f '{{.Target}}'  

## 3.Add the Go install directory to your system's shell path
  On Linux or Mac, run the following command:  
    export PATH=$PATH:/path/to/your/install/directory  
  On Windows, run the following command:  
    set PATH=%PATH%;C:\path\to\your\install\directory  

  go env -w GOBIN=/path/to/your/bin  
or  
  go env -w GOBIN=C:\path\to\your\bin  

## 4.Once you've updated the shell path, run the go install command to compile and install the package.
$ go install

## 5.Run your application by simply typing its name.
$ hello  
map[Darrin:Hail, Darrin! Well met! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met!]  