
# Go Arrays - MCQ Questions with Answers

## ✅ Round 1

**1. What is the index of the first element in a Go array?**  
a) -1  
b) 0  
c) 1  
d) Depends on initialization  
**Answer:** b) 0

---

**2. Which of the following is the correct way to declare and initialize a fixed-size array of 5 integers?**  
a) arr := int[5]{1,2,3,4,5}  
b) arr := [5]int{1,2,3,4,5}  
c) arr := []int{1,2,3,4,5}  
d) arr := array[5]int{1,2,3,4,5}  
**Answer:** b) arr := [5]int{1,2,3,4,5}

---

**3. What is the default value of elements in an integer array in Go?**  
a) nil  
b) undefined  
c) 0  
d) -1  
**Answer:** c) 0

---

**4. What is the length of the array declared as `arr := [...]int{1, 2, 3, 4}`**  
a) Compilation error  
b) 4  
c) 3  
d) Cannot determine  
**Answer:** b) 4

---

**5. What happens if you access an index beyond the length of an array?**  
a) Zero is returned  
b) It wraps around  
c) Compile-time error  
d) Run-time panic  
**Answer:** d) Run-time panic

---

## ✅ Round 2

**1. What will be the output of the following code?**

```go
arr := [3]int{1, 2, 3}
brr := arr
brr[0] = 100
fmt.Println(arr[0])
```

a) 1  
b) 100  
c) Compilation Error  
d) Runtime Error  
**Answer:** a) 1

---

**2. Which of the following statements is true about arrays in Go?**  
a) Arrays are passed by reference to functions  
b) Arrays are passed by value  
c) Arrays do not support indexing  
d) Arrays can be resized dynamically  
**Answer:** b) Arrays are passed by value

---

**3. How do you access the last element of a Go array named `arr`?**  
a) arr[-1]  
b) arr[len(arr)-1]  
c) arr[last]  
d) arr[len(arr)]  
**Answer:** b) arr[len(arr)-1]

---

**4. What is the output of the following?**

```go
arr := [4]int{10, 20, 30, 40}
fmt.Println(len(arr))
```

a) 5  
b) 4  
c) Compilation Error  
d) 0  
**Answer:** b) 4

---

**5. Which of the following initializations is invalid?**  
a) arr := [3]int{1, 2, 3}  
b) arr := [5]int{}  
c) arr := [...]int{4, 5, 6}  
d) arr := [2]int{1, 2, 3}  
**Answer:** d) arr := [2]int{1, 2, 3}

