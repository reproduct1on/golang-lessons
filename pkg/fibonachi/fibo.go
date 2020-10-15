package fibonachi

import(

"fmt"

)



func calcFibonachi(n int) int{
	if (n == 0){
		return 0
	} else if (n == 1)||(n == 2){
		return 1
	}else {
		return calcFibonachi(n - 1) + calcFibonachi(n - 2)
	}
}



func Fibonachi(n int) {
	fmt.Println("fibonachi",calcFibonachi(n))
}

func FibonachiMulti(nums ...int){
	for _, n := range nums {
        fmt.Println("fibonachi myltiple",calcFibonachi(n))
    }
}
