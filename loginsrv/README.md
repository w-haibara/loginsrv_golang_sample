docker build -t loginsrv .

docker run -itd --rm \
	--network proxy-loginsrv \
	--name loginsrv loginsrv \
	/bin/sh -c '/loginsrv \
	--redirect-check-referer=false --redirect-host-file=allowed-hosts \
	-jwt-expiry 30m \
	-jwt-algo HS512 \
	-jwt-secret-file secret.key \
	-cookie-secure=false \
	-simple bob=secret'
