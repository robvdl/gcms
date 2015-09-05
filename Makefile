SHELL := /bin/bash

.PHONY : build
build : build-frontend build-backend

.PHONY : build-frontend
build-frontend : frontend/node_modules
	@echo "Building frontend"
	@(cd frontend; npm run build)

.PHONY : build-backend
build-backend :
	@echo "Building backend"
	@(go build)

.PHONY : watch
watch : frontend/node_modules
	@(cd frontend; npm start)

.PHONY : clean-assets
clean-assets :
	rm -rf static/css static/js

.PHONY : clean-npm
clean-npm :
	rm -rf frontend/node_modules

.PHONY : clean
clean : clean-assets clean-npm

frontend/node_modules :
	@(cd frontend; npm install)
