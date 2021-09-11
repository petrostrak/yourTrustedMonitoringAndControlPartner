# syntax=docker/dockerfile:1
FROM golang as builder

ENV GO111MODULE=on

# Create a directory inside the image that we are building
WORKDIR /app

# Copy our source code into the image
COPY . .

# Download necessary Go modules
COPY go.mod ./

# Compile application
RUN go build -o /app/

FROM golang:1.17-bullseye


COPY --from=builder /app/app .

CMD [ "/docker-inaccess" ]
