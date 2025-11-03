![version](https://img.shields.io/badge/version-20%2B-E23089)
![platform](https://img.shields.io/static/v1?label=platform&message=mac-intel%20|%20mac-arm%20|%20win-64&color=blue)
[![license](https://img.shields.io/github/license/miyako/ping)](LICENSE)
![downloads](https://img.shields.io/github/downloads/miyako/ping/total)

# ping
NET_Ping replacement

## Usage

```4d
#DECLARE($params : Object)

If (Count parameters=0)
	
	CALL WORKER(1; Current method name; {})
	
Else 
	
	var $ping : cs.ping
	$ping:=cs.ping.new()
	
	//atomic
	$result:=$ping.ping({host: "us.4d.com"; timeout: 3; text: "Hello from 4D"})
	
	//async
	$ping.ping({host: "us.4d.com"; timeout: 3; text: "Hello from 4D"}; Formula(onResponse))
	
End if
```
