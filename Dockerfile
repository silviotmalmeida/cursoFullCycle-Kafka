# definindo a imagem base
FROM golang:1.16

# definindo a pasta de trabalho a ser criada e focada no acesso
WORKDIR /go/src

# inserindo o go no path
ENV PATH="/go/bin:${PATH}"

# comandos necessários
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

# comando para manter o container funcionando
CMD ["tail", "-f", "/dev/null"]