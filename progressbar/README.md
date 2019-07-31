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
  
Change nearly every aspect of the progress bar! 
```go
pb.SetWidth(15) // Progress bar will print out 15 ascii chars wide (not including endcaps)
pb.SetGraphics("\u001b[1mX", " ") // Here I set a bolded 'X' as fg, whitespace as bg
pb.ToggleColor([]string{"cyan", "magenta", "green",}) // handles a single color string or an array of up to four strings
pb.SetPrintNumbers("percent") // handles keywords for displaying percent and fraction
```
[**XXXXXXXXX**&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;]  
60%
