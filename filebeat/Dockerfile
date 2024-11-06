FROM docker.elastic.co/beats/filebeat:8.15.3

COPY filebeat.yml /usr/share/filebeat/filebeat.yml
COPY certs/logstash.crt /usr/share/filebeat/certs/logstash.crt
USER root
RUN chown -R root /usr/share/filebeat/
RUN chmod -R go-w /usr/share/filebeat/