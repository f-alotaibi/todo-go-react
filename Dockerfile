FROM node:20-alpine AS wwwbuild

WORKDIR /build

COPY ./web .

RUN npm install
RUN npm run build

FROM golang:1.22.0-alpine AS webbuild

WORKDIR /build

COPY . /build

WORKDIR /build/cmd

RUN go build -o main -ldflags '-s -w'

FROM alpine:latest

WORKDIR /app
RUN mkdir /app/dist

COPY --from=wwwbuild /build/dist /app/dist
COPY --from=webbuild /build/cmd/main /app/

EXPOSE 80

ENTRYPOINT ["./main"]