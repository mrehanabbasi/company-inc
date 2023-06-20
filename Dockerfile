##
## STEP 1 - BUILD
##

# the base image to  be used for the application
FROM golang:1.18-alpine AS build

RUN adduser -u 1001 -D gouser

RUN apk add --no-cache git ca-certificates

ENV GO111MODULE=on GOOS=linux

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy all files and folders
COPY ./ ./

# compile application
RUN CGO_ENABLED=0 go build -o /company-service -ldflags="-w -s"

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /company-service /company-service

COPY --from=build /etc/passwd /etc/passwd

USER 1001

EXPOSE 8080

ENTRYPOINT ["/company-service"]
