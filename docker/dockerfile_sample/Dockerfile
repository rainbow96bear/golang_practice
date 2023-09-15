# 기반 이미지를 설정합니다.
FROM golang:1.16

# 작업 디렉토리를 설정합니다.
WORKDIR /app

# 현재 디렉토리의 파일들을 컨테이너 내부의 /app 디렉토리로 복사합니다.
COPY . /app

RUN go mod init myapp

# Go 애플리케이션을 빌드합니다.
RUN go build -o myapp

# 컨테이너가 실행될 때 실행할 명령어를 설정합니다.
CMD ["./myapp"]