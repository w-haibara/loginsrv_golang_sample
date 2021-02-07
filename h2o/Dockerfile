From lkwg82/h2o-http2-server:latest

USER root
RUN chown h2o /home/h2o/* 

USER h2o
WORKDIR /home/h2o
COPY h2o.conf .
CMD h2o --conf h2o.conf
