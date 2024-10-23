package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    sets:
    for set := 1; set <= setsCount; set++ {         // перебираем наборы данных
        var n,val int
        fmt.Fscan(in, &n)
        fmt.Fscanln(in,&val)

        var tops string
        tops, _ = in.ReadString('\n')
        tops = strings.TrimRight(tops, "\n")

        var bottoms string
        bottoms, _ = in.ReadString('\n')
        bottoms = strings.TrimRight(bottoms, "\n")


        if(len(tops) != len(bottoms)){          // не совпали по длине
            fmt.Fprintln(out, "no")
            continue
        }


        strTops := strings.Split(tops," ")
        mapTops := make(map[string] int)
        for i := 0; i < n; i++ {
            num := strTops[i]

            var cnt int
            cnt, ok := mapTops[num]
            if ok {
                mapTops[num] = cnt+1
            }else{
                mapTops[num] = 1
            }
        }


        strBottoms := strings.Split(bottoms," ")
        if(len(strBottoms) != n){
            fmt.Fprintln(out, "no")
            continue
        }

        str := strBottoms[0]
        previous, _ := strconv.Atoi(str)
        cnt, ok := mapTops[str]
        if (!ok || (cnt<1)){
            fmt.Fprintln(out, "no")
            continue
        }else{
            mapTops[str] = cnt-1
        }

        for i := 1; i < n; i++ {
            str = strBottoms[i]
            current, _ :=strconv.Atoi(str)

            cnt, ok = mapTops[str]
            if (!ok || (cnt<1) || (previous>current)){
                fmt.Fprintln(out, "no")
                continue sets
            }else{
                mapTops[str] = cnt-1
            }

            previous = current
        }
        fmt.Fprintln(out, "yes")
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
