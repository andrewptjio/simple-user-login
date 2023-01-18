FROM golang:1.17

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o main .

EXPOSE 8000

##jalankan executeable
CMD ["/app/main"]