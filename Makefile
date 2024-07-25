ROOT_DIR := .
FRONTEND_DIR := blockpost-frontend

# Default target
all:
	$(MAKE) chain &
	$(MAKE) frontend &

# Target to run ignite chain serve in the root directory
chain:
	konsole --hold -e "bash -c 'cd $(ROOT_DIR) && ignite chain serve; exec bash'"

# Target to run node index.js and npm start in the frontend directory
frontend:
	konsole --hold -e "bash -c 'cd $(FRONTEND_DIR) && node index.js; exec bash'" &
	konsole --hold -e "bash -c 'cd $(FRONTEND_DIR) && npm start; exec bash'" &

npm_build:
	cd $(FRONTEND_DIR) && npm install

.PHONY: all chain frontend

TEST_DIR := x/blockpost/tests
TEST_FILES := keeper_test.go query_server_test.go msg_server_test.go

# Targets for each test file
keeper_test:
	go test $(TEST_DIR)/keeper_test.go

query_server_test:
	go test $(TEST_DIR)/query_server_test.go

msg_server_test:
	go test $(TEST_DIR)/msg_server_test.go

# Target to run all tests
runAllTests: keeper_test query_server_test msg_server_test

.PHONY: keeper_test query_server_test msg_server_test runAllTests
