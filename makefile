# Name of output
BINARY=Z_Builder

# Sources 
SOURCE_DIR=package_assembler
OUT_DIR=bin
ALL_GO_FILES := $(SOURCE_DIR)/packager.go

# Setup build vars
GO_BUILD=go build

# Builds the project
build:
	@echo Building ${BINARY} Binary
	${GO_BUILD} -o ${OUT_DIR}/${BINARY} ${ALL_GO_FILES}

	@echo
	@echo All Binaries are built to the ${OUT_DIR} Folder

mac_win_linux:
	@echo Building ${BINARY} Binary for Mac Windows and Linux

	GOOS=darwin GOARCH=amd64 go build -o ${OUT_DIR}/darwin_amd64_${BINARY} ${ALL_GO_FILES}
	GOOS=linux GOARCH=amd64 go build -o ${OUT_DIR}/linux_amd64_${BINARY} ${ALL_GO_FILES}
	GOOS=windows GOARCH=amd64 go build -o ${OUT_DIR}/windows_amd64_${BINARY} ${ALL_GO_FILES}


clean:
	rm -r $(OUT_DIR)