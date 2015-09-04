SHELL := /bin/bash

.PHONY : build
build : frontend/node_modules
	@echo "Building frontend"
	@(cd frontend; npm run build)
	@echo "Building backend"
	@(go build)

.PHONY : watch
watch : frontend/node_modules
	@(cd frontend; npm start)

.PHONY : clean
clean :
	rm -rf frontend/node_modules static/css static/js

frontend/node_modules :
	@(cd frontend; npm install)
