package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
    "math/big"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    for set := 1; set <= setsCount; set++ {             // перебираем наборы данных
        var n int
        var val float64
        fmt.Fscan(in, &n)

        l := make([]float64, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &val)
            l[i] =  math.Ceil( val/float64(i+1)  )      // минимальный множитель для делителя в этой ячейке
        }

        r := make([]float64, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &val)
            r[i] = math.Floor(  val/float64(i+1))       // максимальный множитель для делителя в этой ячейке
        }

        count := big.NewInt(1)
        divisor := big.NewInt(1000000007)

        for i := 0; i < n; i++ {
            mult := big.NewInt(int64(r[i]-l[i] + 1))    // количество вариантов в этой ячейке
            count.Mul(count,mult)
            count.Mod(count,divisor)
        }

        fmt.Fprintln(out, count.String())
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
