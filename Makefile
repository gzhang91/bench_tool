
PROJECT_PATH=$(shell cd "$(dirname "$0" )" &&pwd)
PROJECT_NAME=$(shell basename "$(PWD)")
DESTDIR=${PROJECT_PATH}/build

.PHONY: all clean

export

all : ${PROJECT_NAME}

${PROJECT_NAME}:
	@echo "创建 ${PROJECT_NAME}目录"
	@mkdir -p ${DESTDIR}/bin

	@echo "编译 ${PROJECT_NAME}"
	@env GOARCH=amd64 go build -o ${DESTDIR}/bin/${PROJECT_NAME} cmd/main.go


clean:
	rm -rf ${DESTDIR}
