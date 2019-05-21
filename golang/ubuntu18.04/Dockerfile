FROM ubuntu:18.04

LABEL maintainer="novemberde.github.io"

RUN apt-get update && \
    apt-get install -y \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    curl \
    lxc \
    iptables \
    git

# Install python & pip3
RUN apt-get -y install python3-pip

# Install nodejs 10.x & npm
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt install -y nodejs && \
    apt install build-essential

# Install golang
RUN add-apt-repository ppa:gophers/archive && \
    apt-get update && \
    apt-get install -y golang-go

# Install AWSCLI
RUN pip3 install awscli --upgrade

# Install Docker from Docker Inc. repositories.
RUN curl -sSL https://get.docker.com/ | sh

# Copy sripts
COPY modprobe.sh /usr/local/bin/modprobe
COPY docker-entrypoint.sh /usr/local/bin/

# Define additional metadata for our image.
VOLUME /var/lib/docker

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["go"]