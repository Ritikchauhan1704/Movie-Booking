FROM golang:1.22.1

WORKDIR /app

COPY . .


# download and install the dependencies:
RUN go get -d -v ./...


# Build the go app
RUN go build -o movie-booking .

EXPOSE 8080

CMD [ "./movie-booking" ]