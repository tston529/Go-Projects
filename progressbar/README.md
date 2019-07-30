Yet Another Go Progress Bar
===========================

[=============_______]

Usage:
======
With a number for size:  
  
```go
pb, err := NewPbar(500000)
if err != nil {
    t.Errorf("How did this happen")
}
for i := 0; i < 500000; i++ {
    pb.IncreaseBar()
}
```
    
With a slice for size:  
```go
arr := make([]int, 100000)
pb, err := NewPbar(arr)
if err != nil {
    t.Errorf("How did this happen")
}
for i := 0; i < len(arr); i++ {
    pb.IncreaseBar()
}
```
  
Change foreground/background text rendering!  
```go
pb.SetGraphics("\u001b[1mX", " ")
```
[XXXXXXXXXXXXXX      ]