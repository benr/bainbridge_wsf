```
      ______            _            __                 _        __                       
     |_   _ \          (_)          [  |               (_)      |  ]                      
       | |_) |  ,--.   __   _ .--.   | |.--.   _ .--.  __   .--.| |  .--./) .---.         
       |  __'. `'_\ : [  | [ `.-. |  | '/'`\ \[ `/'`\][  |/ /'`\' | / /'`\;/ /__\\        
      _| |__) |// | |, | |  | | | |  |  \__/ | | |     | || \__/  | \ \._//| \__.,        
     |_______/ \'-;__/[___][___||__][__;.__.' [___]   [___]'.__.;__].',__`  '.__.'        
 ____      ____   ______   ________    ______     ______  ____  ___( (___))___  ______    
|_  _|    |_  _|.' ____ \ |_   __  | .' ____ \  .' ___  ||_   ||   _||_   __  ||_   _ `.  
  \ \  /\  / /  | (___ \_|  | |_ \_| | (___ \_|/ .'   \_|  | |__| |    | |_ \_|  | | `. \ 
   \ \/  \/ /    _.____`.   |  _|     _.____`. | |         |  __  |    |  _| _   | |  | | 
    \  /\  /    | \____) | _| |_     | \____) |\ `.___.'\ _| |  | |_  _| |__/ | _| |_.' / 
     \/  \/      \______.'|_____|     \______.' `.____ .'|____||____||________||______.'  
                                                                                          
    A Seattle-Bainbridge Washington State Ferry (WSF) Schedule Display Tool for the CLI
```


* __Author:__ Ben Rockwood <benr@cuddletech>
* __Created:__  10/27/15
* __License:__  BSD

---

# Purpose

This tool was created for 2 reasons: 

1. I like Go and look for reasons to better explore it.
2. I commute on the Bainbridge-Settle Ferry every day and wanted a CLI alternative to schedule web page.

It is _not_ intended to be a multi-purpose tool.

# Usage

This tool queries the Washington State Ferry (WSF) API for Ferry Schedules.  To use this 
API you must obtain an Access Key, this is quick and easy, just supply your email address
and you'll immediately recieve a valid code.  Get it here: http://www.wsdot.wa.gov/traffic/api/

Store that Access Key in the "WSFAPI" environment variable (_the key shown here is not valid, it is for example purposes only_):

```bash
$ export WSFAPI="bb2ea927-43e7-4b24-9f10-a152e33579a5"
```

Then execute the tool either using _go run wsf.go_ or build and install via _go install_.

No arguments are required for use. 

For static compilation, use:

```bash
$ CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' 
```

# Dockerization

To create a docker container of this tool, run _make docker_.  The resulting image will be named **cuddletech/wsf**.

Run the container like so:

```bash
$ docker run -it -e WSFAPI=fb2cc927-43e7-XXXXXX-XXXXX-XXXXXXXXXX cuddletech/wsf
Spring 2016
Seattle --> Bainbridge Island 
  12:30   --  Wenatchee (1)
  13:10   --  Tacoma (2)
  14:05   --  Wenatchee (1)
  14:55   --  Tacoma (2)
  15:45   --  Wenatchee (1)
  16:35   --  Tacoma (2)
  17:40   --  Wenatchee (1)
  18:25   --  Tacoma (2)
 ....
```

# See Also

To learn more about the Washington State Department of Transpertation (WSDOT) and Washington State Ferry (WSF) API's, visit http://www.wsdot.wa.gov/traffic/api/
