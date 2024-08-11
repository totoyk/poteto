FROM golang:1.22-bookworm

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
  && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN go install github.com/air-verse/air@latest

COPY . ./

EXPOSE 1323
