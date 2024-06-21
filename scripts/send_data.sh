#!/bin/bash

EXECUTABLE="./cmd/publisher/main.go"

if [ ! -f "$EXECUTABLE" ]; then
    echo "Executable file '$EXECUTABLE' not found"
    exit 1
fi

go run $EXECUTABLE
