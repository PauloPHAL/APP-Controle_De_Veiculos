FROM node:22.8.0-slim

RUN apt update && \
    apt install openssh-server procps -y &&\
    npm install -g @nestjs/cli@10.4.8

USER root

WORKDIR /home/node/app

CMD ["tail", "-f", "/dev/null"]