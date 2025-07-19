FROM golang:1.24-alpine

# Instala ferramentas básicas
RUN apk add --no-cache curl wget git unzip

# Baixa e instala o Air manualmente
RUN wget https://github.com/cosmtrek/air/releases/download/v1.51.0/air_1.51.0_linux_amd64.tar.gz && \
  tar -xzf air_1.51.0_linux_amd64.tar.gz && \
  mv air /usr/local/bin/air && \
  chmod +x /usr/local/bin/air && \
  rm air_1.51.0_linux_amd64.tar.gz

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]