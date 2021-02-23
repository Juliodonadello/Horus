FROM quay.io/influxdb/influxdb:v2.0.3

WORKDIR /
COPY influx-setup.sh /
RUN apt install -y procps dos2unix
RUN dos2unix influx-setup.sh
RUN chmod +x influx-setup.sh
ARG INFLUX_HOST=http://localhost:8086
ARG INFLUXD_HTTP_BIND_ADDRESS=:8086
RUN bash influx-setup.sh
ENV INFLUX_HOST=http://metrics-db:8086
ENTRYPOINT ["/entrypoint.sh"]
CMD ["influxd"]