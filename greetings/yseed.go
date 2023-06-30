package greetings
import (
	"fmt"
	"math/rand"
	"time"
)


// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomYSeed(t int, b int, y int) string {
	imageName := []string { "天","泽","火","雷","风","水","山","地" }
	
	top := mod(t,8)		
	bom := mod(b,8)		
	yao := mod(y,6)		
	topname :=  imageName[top]
	bomname :=  imageName[bom]
	
	//fmt.Println(`current Hour=`, dt.Hour())
	// => 2021-01-22 08:28:55.300846 +0900 JST m=+0.000052154
	//re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} (\d{2}):\d{2}:\d{2}`)
	
	return fmt.Sprintf("origin(上：%d %s, 下：%d %s, 爻：%d).", top+1,topname, bom+1,bomname, yao+1)
}
func mod(x int, m int) int {
	if(x<=0){
		x = rand.Intn(1000) 
	}
	for {
		if (x > m){
			x %= m
		}else{
			return x
		}
	}
}