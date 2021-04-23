# golang
Golang




#### String Manipualtion

##### FormatInt returns the string representation of i in the given base, 
```Go

//FormatInt returns the string representation of i in the given base, 
strconv.FormatInt(8 , 10)

```

##### Int to String, 

```Go

strconv.Itoa(count)

```

##### Int64 to String, 

```Go

strconv.FormatInt(value, 10)

```

##### String to Int, 

```Go

strconv.Atoi(intvalue)

```


### String to int 64

```Go
var s string = "9223372036854775807"
i, _ := strconv.ParseInt(s, 10, 64)
fmt.Printf("val: %v ; type: %[1]T\n", i)

```

##### Get Env variable
```Go

os.Getenv("Envvalue")

```




##### Append to array 

		var storeList []Stores

		for _, store := range anotherStore {

			storeList = append(storeList, InsideStores{
			})

		}










