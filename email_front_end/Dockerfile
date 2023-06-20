FROM ubuntu:20.04
RUN apt update -y
RUN apt upgrade -y
RUN apt-get install curl -y
RUN echo "tzdata tzdata/Areas select America" > /tmp/preseed.txt; \
    debconf-set-selections /tmp/preseed.txt && \
    apt-get update && \
    apt-get install -y tzdata
#RUN apt install npm -y
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install nodejs -y
RUN npm install -g npm@latest
RUN npm cache clean --force
RUN rm -rf node_modules
#RUN node -v
RUN npm install -g vue
RUN npm install -g http-server

RUN mkdir /app
WORKDIR /app

COPY . /app

CMD ["/app/startServices.sh"]
# Understand how to build image with multistage mode