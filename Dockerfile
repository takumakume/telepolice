FROM debian:stable-slim

RUN apt-get update && apt-get -uy upgrade
RUN apt-get -y install ca-certificates && update-ca-certificates

ADD dist/telepolice_linux_amd64/telepolice /telepolice
