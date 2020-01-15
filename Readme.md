
1) go get -u github.com/gorilla/mux
2) go get github.com/joho/godotenv
3) go get github.com/jmoiron/sqlx
4) go get -u github.com/kshvakov/clickhouse

_____________________________________________________________________________________________________
Параметры пагинации(pagination) : 
1) page=? - номер страницы которую хотим посмотреть
2) offset=? - количество записей на одной странице

$domain=blabla.com

Methods("GET"):
1) $domain/api/clicks - получить все клики за все время

2) $domain/api/clicks/pagination - получить все клики за все время с пагинацией
пример:
$domain/api/clicks/pagination?page=2&offset=30 - означает покажи все записи на 2-ой странице, 
если всего записей на странице 30

3) $domain/api/click/offer/{offerId}/pagination
пример:
$domain/api/click/offer/{offerId}/pagination?page=2&offset=30 - означает покажи все записи на 2-ой странице, 
если всего записей на странице 30 для данного offerId

4) $domain/api/click/thread/{thread}/pagination
пример:
$domain/api/click/thread/{thread}/pagination?page=2&offset=30 - означает покажи все записи на 2-ой странице, 
если всего записей на странице 30 для данного потока(thread)

5) $domain/api/click/pid/{pid}/pagination
пример:
$domain/api/click/pid/{pid}/pagination?page=2&offset=30 - означает покажи все записи на 2-ой странице, 
если всего записей на странице 30 для данного веба(pid)

6) $domain/api/click/id/{id} - показать конкретный клик по его clickId

7) $domain/api/click/offer/{offerId} - показать все клики по данному offerId

8) $domain/api/click/thread/{thread} - показать все клики по данному потоку(thread)

9) $domain/api/click/pid/{pid} - показать все клики по данному вебу(pid)

10) $domain/api/click/date/{time}/pid/{pid} - показать все клики за последние {time} часов для данного веба {pid}
time - целое положительное число, то есть вычитаем количество часов начиная с данной секунды

11) $domain/api/sub/{pid}/{sub} - показать все клики по данному вебу и данному sub_id_1 за все время

12) $domain/api/clicksByPidAndTime/pid/{pid}/start/{start}/end/{end} - 
показать все клики для данного веба({pid}) и интервала дат {start} - {end} в формате 
(YYYY-mm-dd-hh-mm-ss):2019-11-06-12-33-22, что означает 2019год-ноябрь-6число-12часов-33минуты-22секунды, пример:
$domain/api/clicksByPidAndTime/pid/2278/start/2019-11-06-12-33-22/end/2019-11-06-15-21-35

13) $domain/api/clicksByOfferIdAndTime/offer/{offerId}/start/{start}/end/{end} - 
показать все клики для данного офера({offerId}) и интервала дат {start} - {end} в формате 
(YYYY-mm-dd-hh-mm-ss):2019-11-06-12-33-22, что означает 2019год-ноябрь-6число-12часов-33минуты-22секунды, пример:
$domain/api/clicksByPidAndTime/pid/2278/start/2019-11-06-12-33-22/end/2019-11-06-15-21-35

____________________________________________________________________________________________________________________

1) Через файлик .env мы можем задать список ip адресов допущенных к запросам как по отдельным адресам так и задав диапазон
#array of trusted ips with "," separator and without space, example: ips_trusted = 127.12.41.77,128.13.45.77,128.12.49.71,
ips_trusted = 128.12.41.77,128.13.45.77,128.12.49.71,128.22.15.99,198.162.1.102,192.168.88.205,192.168.88.208,

ip_range_start =
ip_range_end = 128.12.45.77

2) Можем задать список конечных точек для которых не нужен токен для получения данных тоже в .env,
как и значение самого токена указано в .env
#array of endpoints that dont need jwt token "," separator and without space, example: endpointsNoAuth = /api/some/path,/api/some/path/1,/api/some/path/with/value,
endpointsNoAuth = /api/clicks,/api/click/id/,

3) Можем поменять дефолтное значение пагинации для количества записей на одной странице в .env
items_per_page = 10
