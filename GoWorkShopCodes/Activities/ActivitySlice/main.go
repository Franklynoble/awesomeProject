package main

import "fmt"

func main(){
	//sliceDemo()
	arrangeDays()

}

func sliceDemo() {
	scliseA := []string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
  //  scliseA[1] += "Sunday"

  //  scliseA[0],scliseA[1],scliseA[2],scliseA[3],scliseA[4],scliseA[5],scliseA[6] =	 scliseA[6],scliseA[0],scliseA[1],scliseA[2],scliseA[3],scliseA[4],scliseA[5]
     //  slb := scliseA[0:5]
       // scliseA = append(slb,scliseA...)
        fmt.Println("From Monday: ",scliseA)
        scliseA =  append([]string{},scliseA[6],scliseA[0],scliseA[1],scliseA[2],scliseA[3],scliseA[4],scliseA[5])
        fmt.Println("From Sunday : ",scliseA)

}


func arrangeDays(){

	sclise1 :=  []string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}

      sclise1[0],sclise1[1],sclise1[2],sclise1[3],sclise1[4],sclise1[5],sclise1[6] = "Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"

       n1 := append(sclise1)
       fmt.Println(n1)
}