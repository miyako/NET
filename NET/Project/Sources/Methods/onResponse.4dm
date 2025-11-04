//%attributes = {"invisible":true,"preemptive":"capable"}
#DECLARE($worker : 4D:C1709.SystemWorker; $params : Object)

TRACE:C157

var $result : Object

If ($worker.response="{@")
	$result:=JSON Parse:C1218($worker.response; Is object:K8:27)
Else 
	$result:=Null:C1517
End if 
