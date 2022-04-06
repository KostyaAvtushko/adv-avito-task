# adv-avito-task

## Выполненное Тестовое задание для backend-стажёра в команду Advertising Avito
____
Стек приложения: Go, Fiber, Gorm, PostgreSQL. 
### Эндпоинты
____
####__Методы получения списка объявлений (bold)__
Количество возвращаемых объявлений = 10. 
__Пример самого простого запроса списка объявлений(bold)__: ___ http://localhost:8080/get/groupAdv (bold italic)___
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
