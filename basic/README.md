#### Belajar



### Variabel
A variable is the name of a memory location. That memory location may store a value of any type.

```go 
 var aaa int
 fmt.Println(aaa)


  var aaa, bbb int
  fmt.Println(aaa)
  fmt.Println(bbb)

 var (
    aaa int
    bbb int    = 8
    ccc string = "a"
  )

   

```

### Data Types

```go 
    var t = 123      //Type Inferred will be int
    var u = "circle" //Type Inferred will be string
    var v = 5.6      //Type Inferred will be float64
    var w = true     //Type Inferred will be bool
    var x = 'a'      //Type Inferred will be rune
    var y = 3 + 5i   //Type Inferred will be complex128
    var z = sample{name: "test"}  //Type Inferred will be main.Sample
     var r byte = 'a'
    
    //Print Size
    fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
    
    //Print Type
    fmt.Printf("Type: %s\n", reflect.TypeOf(r))

    fmt.Printf("Type: %T Value: %v\n", t, t)
    fmt.Printf("Type: %T Value: %v\n", u, u)
    fmt.Printf("Type: %T Value: %v\n", v, v)
    fmt.Printf("Type: %T Value: %v\n", w, w)
    fmt.Printf("Type: %T Value: %v\n", x, x)
    fmt.Printf("Type: %T Value: %v\n", y, y)
    fmt.Printf("Type: %T Value: %v\n", z, z)

```

### Constant

A constant is anything that doesn’t change its value. In Go const can be either of type string, numeric, boolean, and characters.

```go 
    const name = "test"
    func main() {
        const a = getValue()
    }
    func getValue() int {
        return 1
    }


    type myString string

    //Typed String constant
    const aa string = "abc"
    var uu = aa
    fmt.Println("Untyped named string constant")
    fmt.Printf("uu: Type: %T Value: %v\n\nn", uu, uu)

    //Below line will raise a compilation error
    //var v myString = aa

    //Untyped named string constant
    const bb = "abc"
    var ww myString = bb
    var xx = bb
    fmt.Println("Untyped named string constant")
    fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
    fmt.Printf("xx: Type: %T Value: %v\n\n", xx, xx)

    //Untyped unnamed string constant
    var yy myString = "abc"
    var zz = "abc"
    fmt.Println("Untyped unnamed string constant")
    fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
    fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)

```

### For loop

```go
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    i := 0
    for {
        fmt.Println(i)
        i++
        time.Sleep(time.Second * 1)
    }

    i := 0
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break
        }
    }

    for i := 1; i < 10; i++ {
        if i%3 == 0 {
            continue
        }
        fmt.Println(i)
    }


    // Nested
    for i := 0; i < 3; i++ {
        fmt.Printf("Outer loop iteration %d\n", i)
        for j := 0; j < 2; j++ {
            fmt.Printf("i= %d j=%d\n", i, j)
        }
    }


    func main() {
        i := 1
        //Function call in the init part in for loop
        for test(); i < 3; i++ {
            fmt.Println(i)
        }

        //Assignment in the init part in for loop
        for i = 2; i < 3; i++ {
            fmt.Println(i)
        }
    }
    func test() {
        fmt.Println("In test function")
    }

```

### For range loop
```go
    letters := []string{"a", "b", "c"}

    //With index and value
    fmt.Println("Both Index and Value")
    for i, letter := range letters {
        fmt.Printf("Index: %d Value:%s\n", i, letter)
    }

    //Only value
    fmt.Println("\nOnly value")
    for _, letter := range letters {
        fmt.Printf("Value: %s\n", letter)
    }

    //Only index
    fmt.Println("\nOnly Index")
    for i := range letters {
        fmt.Printf("Index: %d\n", i)
    }

    //Without index and value. Just print array values
    fmt.Println("\nWithout Index and Value")
    i := 0
    for range letters {
        fmt.Printf("Index: %d Value: %s\n", i, letters[i])
        i++
    }


    for i := 0; i < len(sample); i++ {
        fmt.Printf("%c\n", sample[i])
    }
```
### If


```go
    a := 6
    if a > 5 {
        fmt.Println("a is greater than 5")
    }

    a := 4
    if a > 3 && a < 6 {
        fmt.Println("a is within range")
    }

    a := 1
    b := 2

    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("b is greater than a")
    }

    a := 1
    b := 2
    c := 3
    if a > b {
        if a > c {
            fmt.Println("Biggest is a")
        } else if b > c {
            fmt.Println("Biggest is b")
        }
    } else if b > c {
        fmt.Println("Biggest is b")
    } else {
        fmt.Println("Biggest is c")
    }

```

### Switch
Switch statement are a perfect way to prevent a if-else ladder. Here is the format for switch statement

```go
    switch ch := "b"; ch {
    case "a":
        fmt.Println("a")
    case "b", "c":
        fmt.Println("b or c")    
    default:
        fmt.Println("No matching character")    
    }
    

```

### Defer
Defer as the name suggests is used to defer the cleanup activities in a function. These cleanup activities will be performed at the end of the function. 

```go
func main() {
    err := writeToTempFile("Some text")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
    file, err := os.Open("temp.txt")
    if err != nil {
        return err
    }
    n, err := file.WriteString("Some text")
    if err != nil {
        return err
    }
    fmt.Printf("Number of bytes written: %d", n)
    file.Close()
    return nil
}

```

### Pointer

```go
func main() {
    //Declare
    var b *int
    a := 2
    b = &a
    
    //Will print a address. Output will be different everytime.
    fmt.Println(b)
    fmt.Println(*b)
    b = new(int)
    *b = 10
    fmt.Println(*b) 
}
```

### Struct
GO struct is named collection of data fields which can be of different types. Struct acts as a container that has different heterogeneous data types which together represents an entity.

```go
type employee struct {
    name   string
    age    int
    salary int
}

```


### Array

```go
func main() {
    //Both number of elements and actual elements
    sample1 := [2]int{1, 2}
    fmt.Printf("Sample1: Len: %d, %v\n", len(sample1), sample1)

    //Only actual elements
    sample2 := [...]int{2, 3}
    fmt.Printf("Sample2: Len: %d, %v\n", len(sample2), sample2)

    //Only number of elements
    sample3 := [2]int{}
    fmt.Printf("Sample3: Len: %d, %v\n", len(sample3), sample3)

    //Without both number of elements and actual elements
    sample4 := [...]int{}
    fmt.Printf("Sample4: Len: %d, %v\n", len(sample4), sample4)
}

```

### Slice

```go
func main() {
    numbers := []int{1, 2, 3, 4, 5}

    //Both start and end
    num1 := numbers[2:4]
    fmt.Println("Both start and end")
    fmt.Printf("num1=%v\n", num1)
    fmt.Printf("length=%d\n", len(num1))
    fmt.Printf("capacity=%d\n", cap(num1))

    //Only start
    num2 := numbers[2:]
    fmt.Println("\nOnly start")
    fmt.Printf("num1=%v\n", num2)
    fmt.Printf("length=%d\n", len(num2))
    fmt.Printf("capacity=%d\n", cap(num2))

    //Only end
    num3 := numbers[:3]
    fmt.Println("\nOnly end")
    fmt.Printf("num1=%v\n", num3)
    fmt.Printf("length=%d\n", len(num3))
    fmt.Printf("capacity=%d\n", cap(num3))

    //None
    num4 := numbers[:]
    fmt.Println("\nOnly end")
    fmt.Printf("num1=%v\n", num4)
    fmt.Printf("length=%d\n", len(num4))
    fmt.Printf("capacity=%d\n", cap(num4))
}

```

### Method

```go
type employee struct {
    name   string
    age    int
    salary int
}

func (e employee) details() {
    fmt.Printf("Name: %s\n", e.name)
    fmt.Printf("Age: %d\n", e.age)
}

func (e employee) getSalary() int {
    return e.salary
}

func main() {
    emp := employee{name: "Sam", age: 31, salary: 2000}
    emp.details()
    fmt.Printf("Salary %d\n", emp.getSalary())
}
```

### Interface

```go
type animal interface {
    breathe()
    walk()
}

type lion struct {
    age int
}

func (l lion) breathe() {
    fmt.Println("Lion breathes")
}

func (l lion) walk() {
    fmt.Println("Lion walk")
}

func main() {
    var a animal
    a = lion{age: 10}
    a.breathe()
    a.walk()
}
```

### Iota

```go
const (
    a = iota
    b
    c
)
func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}

```

### Goroutines


```go
func main() {
    go func() {
        fmt.Println("In Goroutine")
    }()

    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")
}


```

### Channel

```go
func main() {
  ch := make(chan int)
  go send(ch)

  go receive(ch)
  time.Sleep(time.Second * 2)
}

func send(ch chan int) {
  time.Sleep(time.Second * 1)
  fmt.Println("Timeout finished")
  ch <- 1
}

func receive(ch chan int) {
  val := <-ch
  fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
}
```

### Select

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go goOne(ch1)
    go goTwo(ch2)

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}

func goOne(ch chan string) {
    ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
    ch <- "From goTwo goroutine"
}
```

### Error

```go
func main() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}
}
```

### Panic and recover

```go
func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
```