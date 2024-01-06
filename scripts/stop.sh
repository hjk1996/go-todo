#!/bin/bash

# PID 파일 경로
PID_FILE="/tmp/todo-app.pid"

# PID 파일 존재 확인
if [ ! -f $PID_FILE ]; then
    echo "PID file not found: $PID_FILE"
    exit 1
fi

# PID 파일에서 PID 읽기
PID=$(cat $PID_FILE)

# 프로세스 종료
kill $PID

# 종료 성공 여부 확인
if [ $? -eq 0 ]; then
    echo "Application stopped."
    rm -f $PID_FILE
else
    echo "Failed to stop application."
fi
