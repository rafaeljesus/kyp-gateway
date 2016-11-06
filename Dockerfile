FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-gateway /kyp-gateway

ENV KYP_GATEWAY_PORT=3004
ENV KYP_AUTH_ENDPOINT=http://localhost:3000/v1/token
ENV KYP_TODO_ENDPOINT=http://localhost:3001/v1/todos
ENV KYP_USERS_ENDPOINT=http://localhost:3003/v1/users
ENV KYP_SECRET_KEY="c91267c27a8599ca0480ea505487d052e3b63a1dd39819db853225a518200399"

ENTRYPOINT ["/kyp-gateway"]
