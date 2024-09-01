package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    inputs := make(map[uint16] []string)

    var count int
    fmt.Fscanln(in, &count)

    var inp string
    var key uint16
    for i := 1; i <= count; i++ {           // Пишем входные строки
        fmt.Fscanln(in, &inp)

        key = getHash(inp)

        inputs[key] = append(inputs[key],inp)
        
    }


    fmt.Fscanln(in, &count)
    for i := 1; i <= count; i++ {           // Проверяем входные значения
        fmt.Fscanln(in, &inp)

        key = getHash(inp)

        val, ok := inputs[key]

        if ok {
            fmt.Fprintln(out, similarLogin(val,inp))            
        }else{
            fmt.Fprintln(out, 0)
        }
        
    }
}

func similarLogin(variants []string, login string) int{
    count := len(variants)

    strlen := len(login)

    variants:
    for i := 0; i < count; i++ {

        current := variants[i]
        if (strlen != len(current)){
            continue
        }

        if ((strlen<2)&&(current==login)){
            return 1
        }

        mismatches := 0
        for j := 0; j < strlen; j++ {

            if(login[j]!=current[j]){
                next := j+1

                if(next>=strlen){
                    continue variants
                }

                if( ( login[j]==current[next] )&&( login[next]==current[j] ) ){
                    if (mismatches>0){
                        continue variants
                    }else{
                        mismatches++
                        j++
                    }
                }else{
                    continue variants
                }
            }        
        }


        return 1
    }

    return 0
}

func getHash(inp string)uint16{
    var sum uint64 = 0

    strlen := len(inp)

    for i := 0; i < strlen; i++ {
        sum += uint64(inp[i])
    }

    return uint16(sum%65536)
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
