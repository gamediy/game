# Makefile 示例

# Go 编译器和参数
GO = go
GOCMD = $(GO) run

# Go 源文件
APP1_SOURCE_FILES = gateway/main.go
APP2_SOURCE_FILES = slot/main.go
APP3_SOURCE_FILES = user/main.go

# 启动参数
APP_ARGS = --gf.gcfg.file=config-dev.yaml

# 默认目标，启动所有三个应用程序
all: start-app1 start-app2 start-app3

# 用于检查应用程序是否已经在运行
APP1_PIDFILE = app1.pid
APP2_PIDFILE = app2.pid
APP3_PIDFILE = app3.pid

# 启动应用程序 1
start-app1: $(APP1_SOURCE_FILES)
	@if [ ! -e $(APP1_PIDFILE) ]; then \
		$(GOCMD) $(APP1_SOURCE_FILES) $(APP_ARGS) > app1.log 2>&1 & \
		echo $$! > $(APP1_PIDFILE); \
	fi

# 启动应用程序 2
start-app2: $(APP2_SOURCE_FILES)
	@if [ ! -e $(APP2_PIDFILE) ]; then \
		$(GOCMD) $(APP2_SOURCE_FILES) $(APP_ARGS) > app2.log 2>&1 & \
		echo $$! > $(APP2_PIDFILE); \
	fi

# 启动应用程序 3
start-app3: $(APP3_SOURCE_FILES)
	@if [ ! -e $(APP3_PIDFILE) ]; then \
		$(GOCMD) $(APP3_SOURCE_FILES) $(APP_ARGS) > app3.log 2>&1 & \
		echo $$! > $(APP3_PIDFILE); \
	fi
# 端口列表
PORTS = 5000 5001 5002
# 停止应用程序
stop:
	@for PORT in $(PORTS); do \
		PID=$(shell lsof -ti :$$PORT); \
		if [ -n "$$PID" ]; then \
			kill -9 "$$PID"; \
			echo "Process with PID $$PID on port $$PORT has been stopped."; \
		else \
			echo "No process found listening on port $$PORT."; \
		fi; \
	done
