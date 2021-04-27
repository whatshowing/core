
install:
	go install \
		github.com/nats-io/stan.go \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golang/protobuf/proto \
		github.com/joho/godotenv \
		github.com/dgrijalva/jwt-go \
		github.com/google/uuid \
		github.com/stretchr/testify

gen:
	buf generate

buf-update:
	buf beta mod update
