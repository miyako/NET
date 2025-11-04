//%attributes = {}
#DECLARE($params : Object)

If (Count parameters:C259=0)
	
	CALL WORKER:C1389(1; Current method name:C684; {})
	
Else 
	
	var $ping : cs:C1710.ping
	$ping:=cs:C1710.ping.new()
	
	//atomic
	$result:=$ping.ping({host: "us.4d.com"; timeout: 1; text: "Hello from 4D"})
	
	//async
	$ping.ping({host: "us.4d.com"; timeout: 1; text: "Hello from 4D"}; Formula:C1597(onResponse))
	
End if 