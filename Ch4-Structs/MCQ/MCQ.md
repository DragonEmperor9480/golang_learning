# Go Structs - MCQs with Answers

---

### 1. What is a struct in Go?

a) A built-in type  
b) A composite type that groups fields together ✅  
c) A type for string handling  
d) A type of array  

---

### 2. How do you create a struct instance?

a) `new Struct{}`  
b) `Struct.create()`  
c) `Struct{}` ✅  
d) `make(Struct)`

---

### 3. What is the correct way to associate a method with a struct?

a) `func structName.methodName()`  
b) `func (structName) methodName()`  
c) `func (s Struct) methodName()` ✅  
d) `func Struct::methodName()`

---

### 4. Which of these is true about struct embedding?

a) Embedding copies fields into another struct ✅  
b) Only methods are embedded  
c) Go supports multiple inheritance via embedding  
d) Fields can't be accessed directly  

---

### 5. What happens if both embedded and outer struct have a field with the same name?

a) Outer struct's field shadows the embedded one ✅  
b) Embedded one always wins  
c) Compilation error  
d) Fields are merged  
