FROM centos:centos7
RUN yum clean all
RUN yum update -y
RUN curl -fsSL https://rpm.nodesource.com/setup_16.x | bash -
RUN yum install nodejs -y