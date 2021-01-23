package common

import "encoding/json"

func Swap(resquest,category interface{})(err error){
	dataByte,err:=json.Marshal(resquest)
	if err!=nil{
		return
	}
	err=json.Unmarshal(dataByte,category)

	return
}



