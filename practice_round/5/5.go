package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "regexp"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    fmt.Fprint(out, "[")

    var setsCount int
    fmt.Fscanln(in, &setsCount)
    for set := 1; set <= setsCount; set++ {     // перебираем наборы данных

        output := inputTrimmedJson(in)

        for{
            oldlen := len(output)
            output = pretty(output)
            if(oldlen == len(output)){
                break
            }
        }

        fmt.Fprint(out, output)

        if (set!=setsCount){
            fmt.Fprint(out, ",")
        }else{
            fmt.Fprintln(out, "]")
        }
    }
}

func inputTrimmedJson(in *bufio.Reader)string{
    output := ""

    var strCount int
    fmt.Fscanln(in, &strCount)
    for j := 1; j <= strCount; j++ {
        var str string
        str, _ = in.ReadString('\n')
        str = strings.ReplaceAll(str, " ", "")
        str = strings.ReplaceAll(str, "\t", "")
        output += strings.TrimSpace( str )
    }

    return output
}

func pretty(inp string)string {
    var out  string

    out = inp

    re1 := regexp.MustCompile(`[^,\[\]\{\}]*\[\],*`)
    re2 := regexp.MustCompile(`[^,\[\]\{\}]*\{\},*`)
    out  = re1.ReplaceAllString(out, "")
    out  = re2.ReplaceAllString(out, "")
    out = strings.ReplaceAll(out, ",}", "}")
    out = strings.ReplaceAll(out, ",]", "]")

    return out
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
