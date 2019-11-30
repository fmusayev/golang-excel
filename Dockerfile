FROM debian

RUN apt update 
RUN apt install -y ca-certificates 
RUN update-ca-certificates

WORKDIR /app
COPY golang-excel ./
COPY resources/main.xlsx ./resources/main.xlsx

EXPOSE 80

ENTRYPOINT [ "./golang-excel" ]
