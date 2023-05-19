FROM centos:7

RUN yum -y install nginx

RUN mkdir -p /opt/app

