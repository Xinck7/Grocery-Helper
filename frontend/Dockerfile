#rockylinux:8 and had to tag as a separate image because it wouldn't pull??
FROM vm-image 
RUN yum clean all
RUN yum update -y
RUN curl -fsSL https://rpm.nodesource.com/setup_18.x | bash -
RUN yum install nodejs -y
RUN npm install -g @angular/cli
# RUN npm install 
