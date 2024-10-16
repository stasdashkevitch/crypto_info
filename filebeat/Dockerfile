FROM docker.elastic.co/beats/filebeat-oss:7.15.0

COPY filebeat.yaml /usr/share/filebeat/filebeat.yaml
USER root
RUN chown -R root /usr/share/filebeat/
RUN chmod -R go-w /usr/share/filebeat/
