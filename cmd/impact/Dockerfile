FROM golang:1.9.3-alpine3.7 AS build
COPY . /go/src/github.com/hinshun/libpackage
RUN go install github.com/hinshun/libpackage/cmd/impact

FROM golang:1.9.3-alpine3.7
COPY --from=build /go/bin/impact bin/
ENTRYPOINT ["impact"]
CMD ["--help"]
