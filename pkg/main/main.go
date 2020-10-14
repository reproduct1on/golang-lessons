package main


import(
	// "fibonachi"
	"fmt"
)

func main(){
	Fibonachi(4)
	FibonachiMulti(4,3,4,2,10)
}


func calcFibonachi(n int) int{

	// switch (n) {
	// 	case 0:
	// 		return 0
	// 	break	
	// 	case 1: 
	// 		return 1
	// 	break
	// 	default:
	// 		return Fibonachi(n - 1) + Fibonachi(n - 2)
	// }
	if (n == 1){
		return 0
	} else if (n == 2)||(n == 3){
		return 1
	}else {
		return calcFibonachi(n - 1) + calcFibonachi(n - 2)
	}
	return n
}
func Fibonachi(n int) {
	fmt.Println("fibonachi",calcFibonachi(n))
}

func FibonachiMulti(nums ...int){
	for _, n := range nums {
        fmt.Println("fibonachi myltiple",calcFibonachi(n))
    }
}