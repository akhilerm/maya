#
# This Dockerfile builds a recent maya agent using the latest binary from
# maya-agent  releases.
#

FROM ubuntu:16.04

# TODO: The following env variables should be auto detected.
ENV MAYA_AGENT_NETWORK="eth0"

RUN apt-get update && apt-get install -y \
    iproute2 \
    curl \
    net-tools \   
    procps \
    open-iscsi

COPY maya-agent /usr/local/bin/
COPY entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

ARG ARCH
ARG DBUILD_DATE
ARG DBUILD_REPO_URL
ARG DBUILD_SITE_URL
LABEL org.label-schema.name="m-agent"
LABEL org.label-schema.description="OpenEBS Agent"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$DBUILD_DATE
LABEL org.label-schema.vcs-url=$DBUILD_REPO_URL
LABEL org.label-schema.url=$DBUILD_SITE_URL

ENTRYPOINT entrypoint.sh "${MAYA_AGENT_NETWORK}"
EXPOSE 7676
