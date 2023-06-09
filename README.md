# Поиск наибольшей группы смежных блоков в матрице

## Дано
- Матрица m*n, состоящая из блоков, размер которой заранее неизвестен.
- Каждый блок имеет цвет.

## Требуется
Найти наибольшую группу смежных блоков одного цвета.

## Вход
- Высота и ширина матрицы, количество цветов.

## Выход
- Сгенерированная матрица
- Найденная группа, в виде по которой ее можно найти в сгенеренной матрице.


## Использование
Запуск программы: _make_ или _make run_ 

Билд прогрыммы: _make build_

Запуск бенчей: _make bench_
## Пример

### Найдена одна группа
input:

    5 5 10

output:

    Generated matrix:
    1 1 2 2 1 
    1 2 2 2 2 
    2 1 2 1 2 
    2 1 1 1 2 
    2 2 2 1 1 
    Max group size: 9
    Number of groups: 1
    Group 1:
    _ _ 2 2 _ 
    _ 2 2 2 2 
    _ _ 2 _ 2 
    _ _ _ _ 2 
    _ _ _ _ _ 

### Найдено больше одной группы

input:

    2 2 2

output:

    Generated matrix:
    2 1 
    2 1 
    Max group size: 2
    Number of groups: 2
    Group 1:
    2 _ 
    2 _ 
    Group 2:
    _ 1 
    _ 1 