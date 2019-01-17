package timeStamp
import (
    "time"
    "fmt"
    // "strings"
)
type TimeStamp struct{
    Length int
    Times [] time.Time
    Tag [] string
    Index int
}

func NewTimeStamp(l int) (*TimeStamp){
    var ts TimeStamp
    ts.Length = l
    ts.Times = make([]time.Time, l, l)
    ts.Index = 0
    ts.Tag = make([]string, l, l)
    return &ts
}

func (this *TimeStamp) Stamp () (){
    this.Times[this.Index] = time.Now()
    this.Index = this.Index + 1 ;
}
func (this *TimeStamp) StampWithTag (tag string) (){
    this.Times[this.Index] = time.Now()
    this.Tag[this.Index] = tag
    this.Index = this.Index + 1 ;
}

func (this *TimeStamp) Interval (p int, n int) (float64){
    return (this.Times[n].Sub(this.Times[p])).Seconds()*1000
}

func (this *TimeStamp) Print () (){
    for i:= 0; i < this.Index-1; i++ {
        fmt.Printf("%f,ms,%s\n",this.Interval(i,i+1),this.Tag[i]);
    }
}


//
// func()
