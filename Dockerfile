# мультистейджинговая сборка
# билдим бинарь пока ещё не в контейнере
# в контейнер положим уже готовый бинарник
# 1
# для сборки можно ипользовать golang:latest пусть он и большой, но нужен только для сборки
FROM golang:latest AS build

WORKDIR /app

COPY . .

# три строчки ниже позволят при выполнении go build не качать все гошные зависимости по новой, которые кэшуруются
COPY go.mod .
COPY go.sum .
RUN go mod download

# RUn выполняется в момент сборки 
# -a скачать все пакеты и зависимости
# -netgo перекомипилировать сетевые библиотеки, ослабляет зависимость от стетевых системных библиотек 
# -static  компилируем статически связанное приложение, в которое встраиваются все необходимые системные зависимости
# CGO_ENABLED=0 отключить CGO чтобы отвязаться от glibc и стандартных библиотек, только в момент сборки
# 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./main ./cmd/urlshortener

# создание образа из пустого контейнера
# предпочтительный способо, т.к. в нем нет никаких зависимостей, которые не используются нашим приложением
# однако необходимые зависимости нужно подтянуть вручную
#2

FROM scratch

WORKDIR /app

# подтянуть образы из предыдущего стейджа
COPY --from=build /app/main /app/main
# в scrathc нет сертификатов и таймзон, подтягиваем их
# важно их подтянуть для меток времени в бд
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Moscow

EXPOSE 8000

ENV URLSHORTENER_STORE="pgst"
# вместо CMD можно использовать ENTRYPOINT
CMD ["./main"]