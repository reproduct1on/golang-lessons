package main
import(

	"lesson1/pkg/fibonachi"
	"flag"
)

func main(){
	
	var nFlag = flag.Int("n", 10, "help message for flag n")
	flag.Parse()

	fibonachi.Fibonachi(*nFlag)
	fibonachi.FibonachiMulti(4,3,4,2,10)


}


