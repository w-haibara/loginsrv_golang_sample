docker build -t jwt-go-sample .

docker run -itd --rm \
	--network proxy-loginsrv \
	--name jwt-go-sample jwt-go-sample
