FROM golang:latest

WORKDIR /app

# copy all your .go source files
# (or use a .dockerignore and COPY . .)
COPY *.go .

# remove any previously initialized go.mod and go.sum files
# (this is in case the container data wasn't destroyed)
RUN rm -f go.mod rm go.sum

# initialize Go modules
RUN go mod init app

# fetch dependencies
RUN go mod tidy

# build (switch to 1 to use the CGO SQLite)
RUN CGO_ENABLED=0 go build -o /build

# run
CMD ["/build"]