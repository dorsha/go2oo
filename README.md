# Go to OO
Go CLI for HP OO actions

[![Build Status](https://travis-ci.org/dorsha/go2oo.svg)](https://travis-ci.org/dorsha/go2oo)

## Prerequisite
* HP OO 10.x Central installed and running
* Go set up

## Install
* go get -u -t -v github.com/dorsha/go2oo
* Navigate to github.com/dorsha/go2oo/go2oo
* go install
* go build

## Usage
```go2oo.exe -help ```  
```go2oo.exe --url https://localhost:8443/oo --action show-content-packs ```  
With Authentication enabled:  
```go2oo.exe --url https://localhost:8443/oo --action show-content-packs --user <user> --password <password> ```  
*CSRF supported*
