package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "math"
    "time"
)

const aConst int = 64

func main(){
    fmt.Println("Variable basics")
    var aString string = "This is Go"
    fmt.Println(aString)
    fmt.Printf("The var's type is %T\n", aString)

    var anInteger int = 42
    fmt.Println(anInteger)
    fmt.Printf("The var's type is %T\n", anInteger)

    var defaultInt int
    fmt.Println(defaultInt)

    var anotherString = "This is another string"
    fmt.Println(anotherString)
    fmt.Printf("The var's type is %T\n", anotherString)

    var anotherInt = 53
    fmt.Println(anotherInt)
    fmt.Printf("The var's type is %T\n", anotherInt)

    myString := "This is also a string"
    fmt.Println(myString)
    fmt.Printf("The var's type is %T\n", myString)

    fmt.Println(aConst)
    fmt.Printf("The var's type is %T\n", aConst)
    //---------------------------------------------------------------------
    fmt.Println("Console Input")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Text: ")
    input, _ := reader.ReadString('\n')
    fmt.Println("You entered: ", input)

    fmt.Print("Enter a number: ")
    numInput, _ := reader.ReadString('\n')
    aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
    if err != nil{
        fmt.Println(err)
    } else{
        fmt.Println("A Value of number: ", aFloat)
    }
    //---------------------------------------------------------------------
    fmt.Println("Math package")

    i1, i2, i3 := 12, 45, 68
    intSum := i1 + i2 + i3
    fmt.Println("Int sum is: ", intSum)

    f1, f2, f3 := 23.5, 65.1, 76.3
    floatSum := f1 + f2 + f3
    fmt.Println("Float sum is: ", floatSum)

    floatSum = math.Round(floatSum*100) / 100
    fmt.Println("The Float sum is now: ", floatSum)

    circleRadius := 15.5
    circumference := circleRadius*2 * math.Pi
    fmt.Printf("Circumference %.2f\n", circumference)
    //---------------------------------------------------------------------
    fmt.Println("Date and Time")
    n := time.Now()
    fmt.Println("This is taking place at ", n)
    
    t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    fmt.Println("Go launched at ", t)
    fmt.Println(t.Format(time.ANSIC))

}
