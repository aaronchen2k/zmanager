VERSION=1.0
PROJECT=zmanager
QINIU_DIR=/Users/aaron/work/zentao/qiniu/
QINIU_DIST_DIR=${QINIU_DIR}${PROJECT}/${VERSION}/
PACKAGE=${PROJECT}-${VERSION}
BINARY=zmanager
BIN_DIR=bin
BIN_ZIP_DIR=${BIN_DIR}/zip/${PROJECT}/${VERSION}/
BIN_OUT=${BIN_DIR}/${PROJECT}/${VERSION}/
BIN_WIN64=${BIN_OUT}win64/${BINARY}/
BIN_WIN32=${BIN_OUT}win32/${BINARY}/
BIN_LINUX=${BIN_OUT}linux/${BINARY}/
BIN_MAC=${BIN_OUT}mac/${BINARY}/

default: prepare_res compile_all copy_files package

win64: prepare_res compile_win64 copy_files package
win32: prepare_res compile_win32 copy_files package
linux: prepare_res compile_linux copy_files package
mac: prepare_res compile_mac copy_files package
upload: upload_to

prepare_res:
	@echo 'start prepare res'
	@go-bindata -o=res/res.go -pkg=res res/...
	@rm -rf ${BIN_DIR} && mkdir -p ${BIN_DIR}

compile_all: compile_win64 compile_win32 compile_linux compile_mac

compile_win64:
	@echo 'start compile win64'
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BIN_WIN64}${BINARY}.exe cmd/main.go

compile_win32:
	@echo 'start compile win32'
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${BIN_WIN32}${BINARY}.exe cmd/main.go

compile_linux:
	@echo 'start compile linux'
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN_LINUX}${BINARY} cmd/main.go

compile_mac:
	@echo 'start compile mac'
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BIN_MAC}${BINARY} cmd/main.go

copy_files:
	@echo 'start copy files'

package:
	@echo 'start package'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@for platform in `ls ${BIN_OUT}`; do mkdir -p ${QINIU_DIST_DIR}$${platform}; done

	@cd ${BIN_OUT} && \
		for platform in `ls ./`; \
		   do cd $${platform} && \
		   zip -r ${QINIU_DIST_DIR}$${platform}/${BINARY}.zip ${BINARY} && \
		   md5sum ${QINIU_DIST_DIR}$${platform}/${BINARY}.zip | awk '{print $$1}' | \
		          xargs echo > ${QINIU_DIST_DIR}$${platform}/${BINARY}.zip.md5 && \
           cd ..; \
		done

upload_to:
	@echo 'upload...'
	@find ${QINIU_DIR} -name ".DS_Store" -type f -delete
	@qshell qupload2 --src-dir=${QINIU_DIR} --bucket=download --thread-count=10 --log-file=qshell.log \
					 --skip-path-prefixes=zendata --rescan-local --overwrite --check-hash