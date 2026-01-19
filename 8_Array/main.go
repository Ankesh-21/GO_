package main
import "fmt"

func main(){
	var arr[10] int
	fmt.Println(arr)
	var str string
	str = "hello"
	for i:= 0;i < len(str);i++{
		if i % 2 == 1{
			continue
		}
		fmt.Printf("%c\n",str[i])
	}
	nums := [3]int{1,2,3}
	fmt.Println(nums)
	v := [2][2]int{{1,2},{3,4}}
	fmt.Println(v)
}