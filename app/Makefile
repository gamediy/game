# Makefile 示例

# Go 编译器和参数
GO = go
GOCMD = $(GO) run

# Go 源文件
APP1_SOURCE_FILES = gateway/main.go
APP2_SOURCE_FILES = slot/main.go
APP3_SOURCE_FILES = user/main.go
APP4_SOURCE_FILES = game/main.go

# 启动参数
APP_ARGS = --gf.gcfg.file=config-dev.yaml

# 默认目标，启动所有三个应用程序
all: start-app1 start-app2 start-app3 start-app4

# 启动应用程序 1
start-app1: $(APP1_SOURCE_FILES)
	@$(GOCMD) $(APP1_SOURCE_FILES) $(APP_ARGS) >app1.log 2>&1 &

# 启动应用程序 2
start-app2: $(APP2_SOURCE_FILES)
	@$(GOCMD) $(APP2_SOURCE_FILES) $(APP_ARGS) > app2.log 2>&1 &

# 启动应用程序 3
start-app3: $(APP3_SOURCE_FILES)
	@$(GOCMD) $(APP3_SOURCE_FILES) $(APP_ARGS) >app3.log 2>&1 &

# 启动应用程序 4
start-app4: $(APP4_SOURCE_FILES)
	@$(GOCMD) $(APP4_SOURCE_FILES) $(APP_ARGS) >app4.log 2>&1 &
# 端口列表
PORTS := 5000 5001 5002 5003

stop:
	@echo "Stopping programs on ports $(PORTS)..."
	@for port in $(PORTS); do \
        pid=$$(lsof -t -i :$$port); \
        if [ -n "$$pid" ]; then \
            echo "Stopping program on port $$port (Process ID: $$pid)..."; \
            kill $$pid; \
            echo "Program on port $$port has been stopped."; \
        else \
            echo "No running program found on port $$port."; \
        fi \
    done