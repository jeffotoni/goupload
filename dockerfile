FROM ubuntu:16.04
MAINTAINER Jefferson Otoni <jeff.otoni@gmail.com>

RUN apt update 

RUN atp upgrade

RUN apt install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core

RUN curl -s https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz | tar -v -C /usr/local -xz


ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/jeffotoni/goupload

# Build the jeffotoni command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/jeffotoni/goupload

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/goupload

# Document that the service listens on port 8080.
EXPOSE 8080

CMD ["bash"]