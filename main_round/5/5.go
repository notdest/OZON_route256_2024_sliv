package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount,linesCount int
    fmt.Fscanln(in, &setsCount)
    for set := 1; set <= setsCount; set++ {     // перебираем наборы данных

        fmt.Fscanln(in, &linesCount)

        lastIsVariable := false
        variableLevel := 0
        first := true

        dictionary := make([]string , 50)
        for lineNum := 1; lineNum <= linesCount; lineNum++ {
            line, _ := in.ReadString('\n')
            line = strings.TrimRight(line, "\n")

            if(strings.HasSuffix(line,":")){
                trimmed := strings.TrimLeft(line, " ")
                level := (len(line)-len(trimmed))/4
                dictName := strings.TrimRight(trimmed, ":")

                if (len(dictionary) < level+1){
                    dictionary = append(dictionary,dictName)
                }else{
                    dictionary[level] = dictName
                }
                lastIsVariable = false
            }else{
                trimmed := strings.TrimLeft(line, " ")
                level := (len(line)-len(trimmed))/4
                variable, value, _ := strings.Cut(trimmed, ": ")

                if !(lastIsVariable && (variableLevel == level)){
                    if(!first){
                        fmt.Fprintln(out, "")
                    }

                    if (level>0){
                        fmt.Fprint(out,"[")
                        fmt.Fprint(out,strings.Join(dictionary[0:level],"."))
                        fmt.Fprintln(out,"]")
                    }
                }
                fmt.Fprintln(out, variable, "=",value)

                variableLevel = level
                lastIsVariable = true
                first = false
            }

        }


        fmt.Fprintln(out, "")
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}