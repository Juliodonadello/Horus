FROM quay.io/influxdb/influxdb:v2.0.3
WORKDIR /
COPY influx-setup.sh /
RUN apt install -y procps
RUN chmod +x influx-setup.sh
RUN bash influx-setup.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["influxd"]