# Go to OO
Go CLI for HP OO actions

[![Build Status](https://travis-ci.org/dorsha/go2oo.svg)](https://travis-ci.org/dorsha/go2oo)

## Prerequisite
* HP OO 10.x Central installed and running
* Go set up

## Install
* go get -u -t -v github.com/dorsha/go2oo  
  (or download as .zip / .tar.gz)
* Navigate to github.com/dorsha/go2oo/go2oo
* go build

## Usage
```go2oo.exe --help ```  

**Show deployed content packs example**  
```go2oo.exe --url https://localhost:8443/oo --action show-content-packs ```  

With authentication enabled in Central:  
```go2oo.exe --url https://localhost:8443/oo --action show-content-packs --user <user> --password <password> ```  
  
*CSRF supported*
