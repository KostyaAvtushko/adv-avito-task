# adv-avito-task

## Выполненное Тестовое задание для backend-стажёра в команду Advertising Avito

Стек приложения: Go, Fiber, Gorm, PostgreSQL. 
## Эндпоинты
____
### Методы получения списка объявлений 
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
### Методы получения объявления 

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


Описание объявления = ___/get/adv/10?filter=desc___

Дополнительные изображения объявления = ___/get/adv/10?filter=photos___


Параметры можно использовать совместно.


Пример полученного JSON: ___http://localhost:8080/get/adv/13?fields=photos___

'''json
{
    "name": "bike",
    "photoURL": [
        "someurl",
        "dasd",
        "adsad"
    ],
    "price": 32323322
}
'''





