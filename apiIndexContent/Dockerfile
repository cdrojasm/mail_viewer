FROM ubuntu:20.04
RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt search golang-go
RUN apt search gcc-go

RUN echo "tzdata tzdata/Areas select America" > /tmp/preseed.txt; \
    debconf-set-selections /tmp/preseed.txt && \
    apt-get update && \
    apt-get install -y tzdata

RUN apt install golang-go -y
RUN apt-get install git -y
#RUN go mod download github.com/go-chi/chi/v5

COPY . /app
RUN go version
WORKDIR app

CMD ["/app/startServices.sh"]