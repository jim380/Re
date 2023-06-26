# syntax=docker/dockerfile:1

FROM golang:1.20-bullseye
RUN apt-get update && apt-get install -y jq
EXPOSE 26656 26657 1317 9090
COPY --from=app . /opt/re
RUN cd /opt/re && make install-test-binary
WORKDIR /opt/re

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD \
  curl -f http://127.0.0.1:1317/blocks/1 >/dev/null 2>&1 || exit 1

CMD bash /opt/re/network/init.sh && \
  bash /opt/re/network/start.sh