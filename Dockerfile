FROM debian:latest

RUN apt update
RUN apt install curl -y
RUN curl -O https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN tar xvf go1.12.7.linux-amd64.tar.gz
RUN chown -R root:root ./go
RUN mv go /usr/local
RUN useradd -ms /bin/bash bazyli

WORKDIR /home/bazyli/go
RUN chown -R bazyli:bazyli /home/bazyli/go
RUN chmod 755 /home/bazyli/go
ENV PATH="/usr/local/go/bin:$HOME/go:${PATH}"
COPY firstfile.go ${HOME}/go/source.go
USER bazyli
# RUN go get -t github.com/surullabs/lint
