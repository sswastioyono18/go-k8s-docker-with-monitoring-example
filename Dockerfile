FROM golang:alpine as build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM kitabisa/debian-base-image:latest
RUN mkdir /app
COPY --from=build /app/main /app/main
CMD ["/app/main"]