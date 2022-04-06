# adv-avito-task

## Выполненное Тестовое задание для backend-стажёра в команду Advertising Avito

Задача: Необходимо создать сервис для хранения и указать объявления. Объявления должны храниться в базе данных. Сервис должен использовать API, работающее поверх HTTP в формате JSON.

Стек приложения: Go, Fiber, Gorm, PostgreSQL. 

## Эндпоинты


### Методы получения списка объявлений(GET)
Количество возвращаемых объявлений = 10. 

__Пример самого простого запроса списка объявлений__: ___http://localhost:8080/get/groupAdv___

Пример возвращаемого JSON:
 ```json
[
    {
        "Name": "bike",
        "PhotoURL": "someurl",
        "Price": 32323322
    },
    {
        "Name": "bike",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqweasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqweasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqsdke",
        "PhotoURL": "som233euqwel",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqe",
        "PhotoURL": "som233euqwel",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqe",
        "PhotoURL": "som233euqwel",
        "Price": 6767767
    }
]
```

Также можно получить отсортированный список объявлений по дате или цене передав параметр.





Сначала дорогие = ___/get/groupAdv?type=ExpensiveToCheapest___

Сначала дешёвые = ___/get/groupAdv?type=CheapestToExpensive___

Сначала новые   = ___/get/groupAdv?type=findNewest___

Сначала дорогие = ___/get/groupAdv?type=findEldest___





Пример получения отсортированного списка(сначала дорогие): ___http://localhost:8080/get/groupAdv___

Пример полученного JSON:
```json
[
    {
        "Name": "bid78eweqe",
        "PhotoURL": "som233euqwel",
        "Price": 6767989897678
    },
    {
        "Name": "bid78eweqe",
        "PhotoURL": "som233euqwel",
        "Price": 6767989897678
    },
    {
        "Name": "bike",
        "PhotoURL": "someurl",
        "Price": 32323322
    },
    {
        "Name": "bike",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233eurl",
        "Price": 6767767
    },
    {
        "Name": "bidasdasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqweasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqweasdke",
        "PhotoURL": "som233euqweqwerl",
        "Price": 6767767
    },
    {
        "Name": "bidasdqweqsdke",
        "PhotoURL": "som233euqwel",
        "Price": 6767767
    }
]
```
### Методы получения объявления(GET)

__Пример самого простого запроса списка объявлений__: ___http://localhost:8080/get/adv/20___

Пример полученного JSON:
```json
{
    "name": "bidasdqweqsdke",
    "photoURL": "som233euqwel",
    "price": 6767767
}
```

Также запрос можно модифицировать и получить дополнительную информацию.


Описание объявления = ___/get/adv/10?fields=desc___

Дополнительные изображения объявления = ___/get/adv/10?fields=photos___


Параметры можно использовать совместно.


Пример полученного JSON: ___http://localhost:8080/get/adv/13?fields=photos___

```json
{
    "name": "bike",
    "photoURL": [
        "someurl",
        "dasd",
        "adsad"
    ],
    "price": 32323322
}
```
### Метод добавления объявления(POST)

Путь запроса: ___/add/adv/___

Запрос должен содержать поля названия, описания, ссылок на изображения, цену объявления.

Требования к полям: описание не более 1000 символов, название не более 200 символов, не более 3 изображений.

Пример запроса JSON:

```json
{
    "name":"bidytyweqe",
    "description":"ttrtweqwew",
    "price": 67673237678,
    "photoURL":[
        "som2ytyuqwel",
        "d2as909yuy0090dd32", 
        "ads32342uyuy1qwed"
        ]
}
```

Пример ответа на успешный запрос JSON:

```json
{
    "Id": 26,
    "message": "success"
}
```






