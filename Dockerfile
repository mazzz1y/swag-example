FROM golang

RUN curl https://glide.sh/get | sh && \
    go get -u -v github.com/swaggo/swag/cmd/swag

