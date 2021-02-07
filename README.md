# loginsrv_golang_sample

## Attention
This sample use the flag `-cookie-secure=false`. If you use this for public network, setup HTTPS and remove the flag.

## Usage

```
# genelate JWT secret key
$ openssl rand -base64 172 | tr -d '\n' > secret.key
$ cp secret.key ./loginsrv/
$ cp secret.key ./jwt-go-sample/

# create docker network
$ docker network create proxy-loginsrv

# run loginsrv container
$ docker build -t loginsrv .
$ docker run -itd --rm \
        --network proxy-loginsrv \
        --name loginsrv loginsrv \
        /bin/sh -c '/loginsrv \
        --redirect-check-referer=false --redirect-host-file=allowed-hosts \
        -jwt-expiry 30m \
        -jwt-algo HS512 \
        -jwt-secret-file secret.key \
        -cookie-secure=false \
        -simple bob=secret'

# run jwt-go-sample container
$ docker build -t jwt-go-sample .
$ docker run -itd --rm \
        --network proxy-loginsrv \
        --name jwt-go-sample jwt-go-sample

# run h20 container
$ docker build ./ -t h2o-loginsrv
$ docker run -itd \
  --network proxy-loginsrv \
  -p 8080:80 \
  --name h2o-loginsrv h2o-loginsrv
```

Web app: http://\<your domain\>/sample
