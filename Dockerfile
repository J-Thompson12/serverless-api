FROM hashicorp/terraform

LABEL maintainer="J-Thompson12"

WORKDIR /srv

ADD providers.tf .

RUN terraform init