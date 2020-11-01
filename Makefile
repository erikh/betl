react-build:
	cd frontend && npm run build

statik: react-build
	statik -src=frontend/build -f

.PHONY: statik
