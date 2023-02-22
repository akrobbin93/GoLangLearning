package main

import(
    "fmt"
    "sort"
)

func main(){
    anInt := 42
    var p = &anInt
    fmt.Println("Pointers")
    fmt.Println("Value of p: ", *p)

    value1 := 42.13
    pointer1 := &value1
    fmt.Println("Value1: ", *pointer1)
    
    *pointer1 = *pointer1/31
    fmt.Println("Poionter1: ", *pointer1)
    fmt.Println("Value1: ", value1)

    fmt.Println("Arrays")
    var colors [3]string
    colors[0] = "Red"
    colors[1] = "Green"
    colors[2] = "Blue"
    fmt.Println(colors)
    fmt.Println(colors[0])

    var numbers = [5]int{5, 3,  1,  2,  4}
    fmt.Println(numbers)

    fmt.Println("Number of colors: ", len(colors))

    fmt.Println("Slices")
    var colors2 = []string{"Red", "Green", "Blue"}
    fmt.Println(colors2)
    colors2 = append(colors2, "Purple")
    fmt.Println(colors2)

    colors2 = append(colors2[1:len(colors)])
    fmt.Println(colors2)

    numbers2 := make([]int, 5, 5)
    numbers2[0] = 134
    numbers2[1] = 72
    numbers2[2] = 254
    numbers2[3] = 33
    numbers2[4] = 44
    fmt.Println(numbers2)

    sort.Ints(numbers2)
    fmt.Println(numbers2)

    fmt.Println("Maps")
    states := make(map[string]string)
    fmt.Println(states)
    states["WA"] = "Washington"
    states["OR"] = "Oregon"
    states["NC"] = "North Carolina"
    fmt.Println(states)

    oregon := states["OR"]
    fmt.Println(oregon)
    
    delete(states, "OR")
    fmt.Println(states)

    states["NY"] = "New York"
    fmt.Println(states)

    for k,v := range states {
        fmt.Printf("%v : %v\n", k, v) 
    }

    keys := make([]string, len(states))
    i := 0
    for k:= range states {
        keys[i] = k
        i++
    }
    fmt.Println(keys)
    sort.Strings(keys)
    fmt.Println(keys)

    for i := range keys {
        fmt.Println(states[keys[i]])
    }

    fmt.Println("Structs")
    poodle := Dog{"Poodle", 10}
    fmt.Println(poodle)
    fmt.Printf("%+v\n",poodle)
    fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
    poodle.Weight = 9 
    fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

// Dog is a struct
type Dog struct {
    Breed string
    Weight int
}
