package main
import(
	"fmt"
)
var count int
func permutation(arr []int,pos int,n int){
	var i int
	if pos==n{
		print_array(arr,n)
		return
	}
	for i=pos;i<n;i++{
		swap(arr,i,pos)
		permutation(arr,pos+1,n)
		swap(arr,i,pos)
	}
}
func swap(arr []int,i int,j int){
	if i!=j{
		var temp int
		temp=arr[i]
		arr[i]=arr[j]
		arr[j]=temp
	}
}
func print_array(arr[]int,len int){
	var i int
	count++
	fmt.Printf("\n第%d个全排列为：",count)
	for i=0;i<len;i++{
		fmt.Printf("%d",arr[i])
	}
}
func main() {
	var b int
	fmt.Print("请输入全排列个数：")
	fmt.Scan(&b)

	if b >= 13 {
		fmt.Print("请输入12或以下的数")
		return
	}

	var a[12] int
	for i:=0;i<12;i++{
		a[i]=i+1
	}
	c:=a[0:b]
	permutation(c,0,b)
}