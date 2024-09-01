package main


//go run task_3_with_profiler.go -cpuprofile=task_3.prof
//go tool pprof task_3.prof

// https://go.dev/blog/pprof

import (
    "fmt"
    "bufio"
    "os"

    "flag"              //!!!!!!!!!!!!!!!  Это внёс для профилирования
    "runtime/pprof"
    "log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")         //!!!!!!!!!!!!!!

func main() {

    flag.Parse()                                                                    //!!!!!!!!!!!!!!
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }




    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    for set := 1; set <= setsCount; set++ {       // перебираем наборы данных

        fmt.Fprintln(out, getRoot(in))
    }
}



// специфичный формат ввода графа, вводим потомков
func getRoot(in *bufio.Reader)int{
    var numCount, vnikuda int

    sum := 0
    count := 0

    fmt.Fscan(in, &numCount)            // для каждого дерева
    i := 1
    for i <= numCount {                 // ограничиваем количество вводимых символов
        fmt.Fscan(in, &vnikuda)         // номер вершины, не интересен
        i++

        var childCount int              // Количество потомков
        fmt.Fscan(in, &childCount)
        i++

        for child := 1; child <= childCount; child++ { // записываем потомков
            var val int
            fmt.Fscan(in, &val)
            i++

            sum += val
            count ++
        }
    }

    count ++ // одно число было пропущено
    return ((1+count)*count/2 - sum) // сумма членов арифметической прогрессии минус сума потомков
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
