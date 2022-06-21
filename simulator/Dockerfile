FROM golang:1.16
#pasta do container
WORKDIR /go/src
#variavel go
ENV PATH="/go/bin:${PATH}"
#instalando lib do kafka
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

#ativando container 
CMD ["tail", "-f", "/dev/null"]