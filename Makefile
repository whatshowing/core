
install:
	go install \
		github.com/nats-io/stan.go \
		github.com/golang/protobuf/proto \
		github.com/joho/godotenv \
		github.com/dgrijalva/jwt-go \
		github.com/google/uuid \
		github.com/stretchr/testify


#generate:
#	buf generate
#	statik -m -f -src third_party/OpenAPI/
#
#
#buf-update:
#	buf beta mod update
