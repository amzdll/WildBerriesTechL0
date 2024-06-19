#!/bin/bash

# Проверка количества аргументов
if [ $# -ne 1 ]; then
    echo "Usage: $0 <dev|local|prod>"
    exit 1
fi

# Переменная для аргумента
ENV=$1

# Путь к исполняемому файлу
EXECUTABLE="./cmd/publisher/main.go"

# Проверка существования исполняемого файла
if [ ! -f "$EXECUTABLE" ]; then
    echo "Executable file '$EXECUTABLE' not found"
    exit 1
fi

# Выполнение и вывод результатов
go run $EXECUTABLE --stage=$ENV
