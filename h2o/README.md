docker build ./ -t h2o-loginsrv
docker run -itd \
  --network proxy-loginsrv \
  -p 8080:80 \
  --name h2o-loginsrv h2o-loginsrv
