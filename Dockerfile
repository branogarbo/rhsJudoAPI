FROM golang:1.16.3-alpine AS builder
WORKDIR /go/src/rhsJudoAPI
COPY . .
RUN apk add --no-cache git
RUN go get -v ./...
RUN go install -v ./...

FROM alpine
COPY --from=builder /go/src/rhsJudoAPI/data /data
COPY --from=builder /go/bin/rhsJudoAPI /
EXPOSE 3000
CMD [ "/rhsJudoAPI" ]