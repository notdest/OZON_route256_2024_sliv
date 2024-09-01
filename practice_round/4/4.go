package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)
    for set := 1; set <= setsCount; set++ {     // перебираем наборы данных
        var len1, maxlen, len2, first, second int

        var numCount int
        fmt.Fscan(in, &numCount)                // для каждой последовательности
        for j := 0; j < numCount; j++ {
            var val int
            fmt.Fscan(in, &val)

            if (j==0){
                len1 = 0
                len2 = 0
                maxlen = 0
                first = val
                second = val
            }

            if (first==val){
                len1++
                len2++
            }else{
                if (second==val){
                    len2++
                }else{
                    len2 = len1 + 1
                }
                len1 = 1
                second = first
                first = val
            }


            if (len2>maxlen){
                maxlen = len2
            }
        }


        fmt.Fprintln(out, maxlen)
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
