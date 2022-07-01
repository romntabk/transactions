FROM golang:1.18

RUN mkdir /app

RUN cd app

COPY . .

RUN go build main_program.go db.go

CMD ["go run /app/main_program"]
