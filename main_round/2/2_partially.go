package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    for set := 1; set <= setsCount; set++ {       // перебираем наборы данных
        var n int
        var val float64
        fmt.Fscan(in, &n)

        l := make([]float64, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &val)
            l[i] = val
        }

        r := make([]float64, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &val)
            r[i] = val
        }

        count := float64(1)

        for i := 0; i < n; i++ {
            minx := math.Ceil(  l[i]/float64(i+1)  )
            maxx := math.Floor(  r[i]/float64(i+1))

            mult := (maxx-minx + 1)
            if(mult > 1000000007){
                mult =  math.Mod(mult, 1000000007)
            }
            
            count = count * mult

            if(count > 1000000007){
                count =  math.Mod(count, 1000000007)
            }
        }

        

        fmt.Fprintf(out, "%.0f\n", count)
    }
}

//wrong answer 1st numbers differ - expected: '195789499', found: '997846926'

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
