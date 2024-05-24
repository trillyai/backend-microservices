FROM golang:latest as build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./core ./core
COPY ./services/profile ./services/profile

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o profile-service ./services/profile/main.go

FROM scratch
WORKDIR /app
COPY --from=build /build/profile-service /app/

# Set environment variables
ENV HTTP_PORT=:8081
ENV JWT_SECRET_KEY=topsecret
ENV DB_HOST=192.168.215.2
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=cimri
ENV DB_DBNAME=grad-auth

EXPOSE 8081

ENTRYPOINT ["/app/profile-service"]

#  docker build -f ./Dockerfile.profile -t mehmetali10/profile-service .
