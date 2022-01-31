package util

import (
	"math/rand"
	"time"
)

func RandomName(i int) string {
	var str=[]byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKZXCVBN1234567890")
	result:=make([]byte,i)
	rand.Seed(time.Now().Unix())
	for j:=range result{
		result[j]=str[rand.Intn(len(str))]
	}
	return string(result)
}
