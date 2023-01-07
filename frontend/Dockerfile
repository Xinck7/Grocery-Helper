#rockylinux:8 and had to tag as a separate image because it wouldn't pull??
FROM vm-image
COPY . /opt
WORKDIR /opt/GroceryHelper

#update container
RUN yum clean all
RUN yum update -y

#Install node
RUN curl -fsSL https://rpm.nodesource.com/setup_18.x | bash -
RUN dnf install nodejs -y
RUN npm install npm@latest -g
RUN npm install -g @angular/cli

#Install packages
RUN npm install 
