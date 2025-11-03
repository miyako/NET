![version](https://img.shields.io/badge/version-20%2B-E23089)
![platform](https://img.shields.io/static/v1?label=platform&message=mac-intel%20|%20mac-arm%20|%20win-64&color=blue)
[![license](https://img.shields.io/github/license/miyako/NET)](LICENSE)
![downloads](https://img.shields.io/github/downloads/miyako/NET/total)

# NET
NET_Ping replacement

## Usage

```4d
#DECLARE($params : Object)

If (Count parameters=0)
	
	CALL WORKER(1; Current method name; {})
	
Else 
	
	var $ping : cs.NET.ping
	$ping:=cs.NET.ping.new()
	
	//atomic
	$result:=$ping.ping({host: "us.4d.com"; timeout: 3; text: "Hello from 4D"})
	
	//async
	$ping.ping({host: "us.4d.com"; timeout: 3; text: "Hello from 4D"}; Formula(onResponse))
	
End if
```

## Callback

```4d
#DECLARE($worker : 4D.SystemWorker; $params : Object)

var $result : Object
$result:=JSON Parse($worker.response; Is object)
```

<img width="500" height="auto" alt="" src="https://github.com/user-attachments/assets/f379c8ac-e971-4c61-921d-f62182c2ec33" />

<img width="500" height="auto" alt="" src="https://github.com/user-attachments/assets/0b01dcea-bd5d-431e-9026-dfe489d45e6a" />
