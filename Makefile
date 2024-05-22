push:
	git add . && git commit -m "update" && git push origin main
get:
	git pull origin main
all: push get


	
.PHONY: all