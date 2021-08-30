FROM golang:1.16.6-buster

# cp

COPY ./scripts /app/scripts
COPY ./src /app/src

WORKDIR /app/src

EXPOSE 3000

CMD [ "go", "run", "." ]