#!/bin/bash

# 실행 파일 경로 설정
APP_PATH="/var/www/todo-app/todo-app"

# 실행 파일 백그라운드에서 실행
nohup $APP_PATH &

# 백그라운드에서 실행 중인 프로세스의 PID 저장
echo $! > /tmp/todo-app.pid

echo "Application started."