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

    for set := 1; set <= setsCount; set++ {       // перебираем наборы данных
        var val string
        fmt.Fscan(in, &val)

        opt := optimalNumber(val)
        fmt.Fprintln(out,opt)
    }
}
// wrong answer 1st lines differ - expected: '98724', found: '98474'


func optimalNumber(str string) string{

    if(len(str)<2){
        return "0"
    }

    if (str[0]<str[1]){
        return str[1:]
    }

    mini := 1
    for i := 1; i < len(str); i++ {

        if(str[i]<str[i-1]){
            mini = i
        }

        if(str[i]>str[i-1]){
            break
        }
    }

    newStr := str[:mini] + str[mini+1:]


    return newStr
}


func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
