FROM scratch

MAINTAINER Ben Rockwood <benr@cuddletech.com>
ADD 	   bainbridge_wsf wsf

ENTRYPOINT ["/wsf"]
