FROM golang:1.24.2-alpine

WORKDIR /app

# copy all your .go source files
# (or use a .dockerignore and COPY . .)
COPY ./ /app

# fetch dependencies
RUN go mod tidy

# build (switch to 1 to use the CGO SQLite)
RUN CGO_ENABLED=0 go build -o /app/build

# run
CMD ["/app/build"]