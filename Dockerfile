FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/app/
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]