package main

import (
    "fmt"
    "strconv"
    //"math"
)


func main() {
    var param0 float64 = -1000000000000000
    var param1 float64 = 1000000000000001
    var param2 float64 = 50000000000000
    fmt.Println(fmt.Sprint(param0))
    //var space float64 = math.Max(len(fmt.Sprint(param0)), len(fmt.Sprint(param1)))
    var space int = 0

    for i:=param0; i<param1; i+=param2 {

        var iString string = ""
        var iLen int = len(strconv.Itoa(i))

        switch {
            case i<0: 
                iString = strconv.Itoa(i) // + " : " + strconv.Itoa(iLen)
            case i==0:
                iString = strconv.Itoa(i) // + " : " + strconv.Itoa(iLen)
            case i>0:
                iLen++
                iString = "+" + strconv.Itoa(i) // + " : " + strconv.Itoa(iLen)
        }


        
        var spaceBefore int = space - iLen

        var spaceString string

        for counter:=0; counter<spaceBefore+1; counter++ {
            spaceString += " "
        }




        iString = spaceString + iString


        fmt.Println(iString)
    }
}



