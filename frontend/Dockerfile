FROM ubuntu:22.04
COPY . /opt/groceryhelper
WORKDIR /opt/groceryhelper

#Update
RUN apt update
RUN apt install -y curl

#install node
RUN curl -s https://deb.nodesource.com/setup_18.x | sh
RUN apt install -y nodejs
RUN npm install -g npm@9.8.1
