FROM golang
COPY output output
WORKDIR output
RUN chmod -R +x bin
EXPOSE 8888
ENTRYPOINT ["sh", "bootstrap.sh"]