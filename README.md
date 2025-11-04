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

If ($worker.response="{@")
	$result:=JSON Parse($worker.response; Is object)
Else 
	$result:=Null
End if 
```

<img width="500" height="612" alt="" src="https://github.com/user-attachments/assets/36a7e864-adfe-4fbb-ad7e-dedb375634bc" />  
<img width="500" height="629" alt="" src="https://github.com/user-attachments/assets/5354f3f0-12b7-4a50-bc34-af1bfd7ea992" />

