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
    for set := 1; set <= setsCount; set++ {
        var packCount int
        var commision float64
        fmt.Fscan(in, &packCount, &commision)

        var sum float64 = 0
        for pack := 1; pack <= packCount; pack++ {
            var val float64
            fmt.Fscan(in, &val)

            goodCommision := math.Floor(val*commision) / 100
            badCommision := math.Floor(val*commision / 100)
            err := goodCommision - badCommision

            sum += err
        }

        fmt.Fprintf(out, "%.2f\n", sum)
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
