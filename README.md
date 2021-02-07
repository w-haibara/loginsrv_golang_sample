# loginsrv_golang_sample

openssl rand -base64 172 | tr -d '\n' > secret.key
cp secret.key ./loginsrv/
cp secret.key ./jwt-go-sample/
