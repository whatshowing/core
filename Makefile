
install:
	go install \
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
