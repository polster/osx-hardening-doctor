OSX Hardening Doctor
====================

This tool runs a collection of pre-defined security checks against your system printing a small report at the end.

Although the OSX Hardening Doctor may require sudo, the verified system will not be touched.

## Ops

### Usage

* Basic usage:
```
./ohd
```
* For further details on the Doctor's usage, simply run:
```
./ohd --help
```

### Hardening Checks

See [checks](checks.yml)

## Dev

### How to build

* Checkout the project and cd into the project folder
* Set the GO path:
```
export GOPATH=$HOME/Src/Git/osx-hardening-doctor 
```
* Build and install:
```
go install github.com/polster/ohd
```