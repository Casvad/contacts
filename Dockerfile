FROM golang:1.22 as build
ARG FOLDER_PATH="contacts"
RUN mkdir -p /go/src/$FOLDER_PATH
ADD . /go/src/$FOLDER_PATH
WORKDIR /go/src/$FOLDER_PATH
RUN GOOS=linux GOARCH=amd64 eval `ssh-agent -s` && \
    go mod download && \
    go get -u && \
    go build -o app .
RUN rm -rf /root/.ssh
RUN cp /go/src/$FOLDER_PATH/app /app
ENTRYPOINT [ "/app" ]