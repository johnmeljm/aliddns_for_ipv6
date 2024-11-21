# aliddns_for_ipv6

## Introduction

If you have an ipv6 global unicast address, you can run 
this application on your own machine, which is used to 
publish your ipv6 address to the alidns service, and 
then you can access your local machine from anywhere.


## Compile

```
git clone https://github.com/johnmeljm/aliddns_for_ipv6.git
cd aliddns_for_ipv6
make build
```

## Run

Download the `config/example.yaml` file to the path as same as 
the binary file, rename `example. yaml` to `aliddns.yaml`, then 
modify your config information, run the binary.
