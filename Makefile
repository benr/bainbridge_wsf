all: static

standard:
	go build 

static:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s'  

docker: static
	docker build -t cuddletech/wsf .  

docker-clean:
	echo "Deleting cuddletech/wsf Docker Images"
	docker rmi -f cuddletech/wsf

clean:
	rm -f bainbridge_wsf*
	rm -f wsf

clean-all: clean docker-clean


