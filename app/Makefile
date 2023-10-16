# Makefile 示例

# Go 编译器和参数
GO = go
GOCMD = $(GO) run

# Go 源文件
APP1_SOURCE_FILES = gameway/main.go
APP2_SOURCE_FILES = slot/main.go
APP3_SOURCE_FILES = user/main.go

# 默认目标，启动所有三个应用程序
all: start-app1 start-app2 start-app3


# 启动应用程序 1
start-app1: $(APP1_SOURCE_FILES)
	$(GOCMD) $(APP1_SOURCE_FILES)

# 启动应用程序 2
start-app2: $(APP2_SOURCE_FILES)
	$(GOCMD) $(APP2_SOURCE_FILES)

# 启动应用程序 3
start-app3: $(APP3_SOURCE_FILES)
	$(GOCMD) $(APP3_SOURCE_FILES)
