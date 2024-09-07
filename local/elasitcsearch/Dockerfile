FROM elasticsearch:8.4.1

COPY . /elastic

RUN bin/elasticsearch-plugin install analysis-nori

ENTRYPOINT ["/elastic/init.sh"]