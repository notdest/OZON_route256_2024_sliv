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

    for set := 1; set <= setsCount; set++ {         // перебираем наборы данных
        var n,val int

        fmt.Fscan(in, &n)

        nums := make([]int, n)
        for i := 0; i < n; i++ {                    // Вводим все цифры
            fmt.Fscan(in, &val)
            nums[i] =  val
        }

        count := 0
        lefts := make(map[int]int)
        for i := 1; i < n-2; i++ {
            left := nums[i]-nums[i-1]
            right := nums[i+2]-nums[i+1]

            val, ok := lefts[left]                  // пишем количество левых разностей
            if ok {
                lefts[left] = lefts[left] + 1
            }else{
                lefts[left] = 1
            }

            val, ok = lefts[right]                 // смотрим количество пар к правой разности
            if ok {
                count += val
            }
        }

        fmt.Fprintln(out, count)
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
