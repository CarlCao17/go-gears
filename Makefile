BIN_DIR := $(abspath ./bin)
SRC_DIR := $(abspath ./cmd)
GOOS := linux
GOARCH := amd64
CGO_DISABLED := 0


# 获取所有源码目录
DIRS := $(shell find $(SRC_DIR) -mindepth 1 -maxdepth 1 -type d)
# 获取所有目标文件名
EXES := $(foreach dir, $(DIRS), $(BIN_DIR)/$(notdir $(dir)))

all: $(EXES)

# 编译指定程序
$(BIN_DIR)/%: $(SRC_DIR)/%
	cd $(SRC_DIR)/$* && CGO_ENABLED=$(CGO_DISABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$*

# 清理目标
clean:
	rm -rf $(EXES)

# 允许我们使用 make depready
%:
	$(MAKE) $(BIN_DIR)/$@

.PHONY: all clean %


