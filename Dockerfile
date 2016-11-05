FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-gateway /kyp-gateway
ENV KYP_SECRET_KEY="c91267c27a8599ca0480ea505487d052e3b63a1dd39819db853225a518200399"
ENTRYPOINT ["/kyp-gateway"]
