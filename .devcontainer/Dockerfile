FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm as go

RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends g++ \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app/

RUN curl -O https://bootstrap.pypa.io/get-pip.py \
      && python3 get-pip.py --break-system-packages \
      && rm get-pip.py

COPY ./requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt --break-system-packages
