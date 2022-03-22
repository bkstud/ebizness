FROM debian:latest

RUN apt-get update
RUN apt install -y curl git build-essential
RUN curl -O https://dl.google.com/go/go1.18.linux-amd64.tar.gz
RUN tar xvf go1.18.linux-amd64.tar.gz
RUN chown -R root:root ./go
RUN mv go /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"
# Download linter will be available in /root/go/bin
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

RUN apt install -y default-jdk
RUN curl -sSLO https://github.com/pinterest/ktlint/releases/download/0.45.1/ktlint 
RUN chmod a+x ktlint
RUN mv ktlint /usr/local/bin