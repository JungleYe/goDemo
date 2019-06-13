package bigmap

import (
	"testing"
	"fmt"
	"time"
)

func TestItem(t *testing.T) {
	for j:=0;j<=100;j++{
		for i:=0;i<=100000;i++{
			AddOneItem(fmt.Sprintf("%d",i))
		}
		time.Sleep(10 * time.Second)
	}
}
