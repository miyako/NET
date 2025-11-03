Class extends _NET

Class constructor($controller : 4D:C1709.Class)
	
	Super:C1705("ping"; $controller)
	
Function ping($option : Variant; $formula : 4D:C1709.Function) : Collection
	
	var $stdOut; $isStream; $isAsync : Boolean
	var $options : Collection
	var $results : Collection
	$results:=[]
	
	Case of 
		: (Value type:C1509($option)=Is object:K8:27)
			$options:=[$option]
		: (Value type:C1509($option)=Is collection:K8:32)
			$options:=$option
		Else 
			$options:=[]
	End case 
	
	var $commands : Collection
	$commands:=[]
	
	If (OB Instance of:C1731($formula; 4D:C1709.Function))
		$isAsync:=True:C214
		This:C1470.controller.onResponse:=$formula
	End if 
	
	For each ($option; $options)
		
		If ($option=Null:C1517) || (Value type:C1509($option)#Is object:K8:27)
			continue
		End if 
		
		$stdOut:=Not:C34(OB Instance of:C1731($option.output; 4D:C1709.File))
		
		$command:=This:C1470.escape(This:C1470.executablePath)
		
		If (Value type:C1509($option.text)=Is text:K8:3) && (Length:C16($option.text)#0)
			$command+=" --text "
			$command+=This:C1470.escape($option.text)
		End if 
		
		$isStream:=True:C214
		
		If (Value type:C1509($option.text)=Is text:K8:3) && (Length:C16($option.text)#0)
			$command+=" --text "
			$command+=This:C1470.escape($option.text)
		End if 
		
		If (Value type:C1509($option.host)=Is text:K8:3) && (Length:C16($option.host)#0)
			$command+=" --host "
			$command+=This:C1470.escape($option.host)
		End if 
		
		If (Value type:C1509($option.timeout)=Is real:K8:4) && ($option.timeout>0)
			$command+=" --timeout "
			$command+=String:C10($option.timeout*1000)
		End if 
		
		This:C1470.controller.variables.PATH:="/sbin/"
		
		var $worker : 4D:C1709.SystemWorker
		$worker:=This:C1470.controller.execute($command; $isStream ? $option.file : Null:C1517; $option.data).worker
		
		If (Not:C34($isAsync))
			$worker.wait()
		End if 
		
		If ($stdOut) && (Not:C34($isAsync))
			//%W-550.26
			//%W-550.2
			
			If (This:C1470.controller.stdOut#"")
				$results.push(JSON Parse:C1218($worker.response; Is object:K8:27))
			End if 
			
			This:C1470.controller.clear()
			//%W+550.2
			//%W+550.26
		End if 
		
	End for each 
	
	If ($stdOut) && (Not:C34($isAsync))
		return $results
	End if 