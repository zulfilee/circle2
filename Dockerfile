FROM scratch
EXPOSE 8080
ENTRYPOINT ["/circle2"]
COPY ./bin/ /