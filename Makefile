
upgradeLibs:
	go get -u ./...
	go mod tidy

ifndef version
	@echo 'version is not defined'
	exit 1
endif
release:
	git fetch
	git tag ${version}
	git push --tags
