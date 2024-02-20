FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm as go

RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends g++

RUN curl -O https://bootstrap.pypa.io/get-pip.py \
      && python3 get-pip.py --break-system-packages \
      && rm get-pip.py

COPY ./requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt --break-system-packages
RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
ENV PATH $PATH:$(go env GOPATH)/bin