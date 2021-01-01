FROM golang:1.15

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🥧 View the status of your repo checks from the RPi sense hat"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing make
RUN apt-get update && apt-get install make=4.2.1-1.2 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

CMD ["make", "local-test"]
