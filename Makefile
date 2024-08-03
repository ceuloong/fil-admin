PROJECT:=fil-admin

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o fil-admin .

# make build-linux
build-linux:
	@docker build -t fil-admin:latest .
	@echo "build successful"

build-sqlite:
	go build -tags sqlite3 -ldflags="-w -s" -a -installsuffix -o fil-admin .

# make run
run:
    # delete fil-admin-api container
	@if [ $(shell docker ps -aq --filter name=fil-admin --filter publish=8000) ]; then docker rm -f fil-admin; fi

    # 启动方法一 run fil-admin-api container  docker-compose 启动方式
    # 进入到项目根目录 执行 make run 命令
	@docker-compose up -d

	# 启动方式二 docker run  这里注意-v挂载的宿主机的地址改为部署时的实际决对路径
    #@docker run --name=fil-admin -p 8000:8000 -v /home/code/go/src/fil-admin/fil-admin/config:/fil-admin-api/config  -v /home/code/go/src/fil-admin/fil-admin-api/static:/fil-admin/static -v /home/code/go/src/fil-admin/fil-admin/temp:/fil-admin-api/temp -d --restart=always fil-admin:latest

	@echo "fil-admin service is running..."

	# delete Tag=<none> 的镜像
	@docker image prune -f
	@docker ps -a | grep "fil-admin"

stop:
    # delete fil-admin-api container
	@if [ $(shell docker ps -aq --filter name=fil-admin --filter publish=8000) ]; then docker-compose down; fi
	#@if [ $(shell docker ps -aq --filter name=fil-admin --filter publish=8000) ]; then docker rm -f fil-admin; fi
	#@echo "fil-admin stop success"


#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t fil-admin:latest

# make deploy
deploy:

	#@git checkout master
	#@git pull origin master
	make build-linux
	make run
