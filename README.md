# Golang HTTP/2 Push example

[Based on this article](https://blog.golang.org/h2push) and also [folowing code here](https://gist.github.com/denji/12b3a568f092ab951456).


## Requisites

  - **Golang 1.8**
  - **OpenSSL**, if you want to generate your own certificates

## Build

    $ go build


### Generate keys

#### Generate private key (.key)

    # Key considerations for algorithm "RSA" ≥ 2048-bit 
    openssl genrsa -out server.key 2048

    # Key considerations for algorithm "ECDSA" ≥ secp384r1
    # List ECDSA the supported curves (openssl ecparam -list_curves)
    openssl ecparam -genkey -name secp384r1 -out server.key   


#### Generate public key (.crt) based on the private

    openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650


## Run

    $ ./http2-push
    $ firefox https://localhost:8443
