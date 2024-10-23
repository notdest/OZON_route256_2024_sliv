package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type Change struct{
    step int
    name string
}

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    for set := 1; set <= setsCount; set++ {             // перебираем наборы данных
        var n,val int
        fmt.Fscan(in, &n)
        fmt.Fscanln(in,&val)    //Вникуда

        
        changes := make(map[int] []Change , n+1)
        idByName := make(map[string] int)

        steps:
        for i := 1; i <= n; i++ {
            var line string
            line, _ = in.ReadString('\n')
            line = strings.TrimRight(line, "\n")

            command := strings.Split(line," ")
            if(command[0] == "CHANGE"){
                var ch Change

                id, _ := strconv.Atoi(command[2])
                name := command[1]
                ch.step = i
                ch.name = name

                oldId, ok := idByName[name]                 // Если товар с этим именем уже есть -- удаляем старый id
                if ok {
                    var ch2 Change
                    ch2.step = i
                    ch2.name = "404"
                    changes[oldId] = append(changes[oldId],ch2)
                }

                _, ok2 := changes[id]                       // Из вспомогательной карты удаляем товары с этим id (мы его перезаписали)
                if ok2 {
                    delete(idByName, changes[id][len(changes[id])-1].name)
                }

                changes[id] = append(changes[id],ch)
                idByName[name] = id
            }else{
                id, _ := strconv.Atoi(command[1])
                step, _ := strconv.Atoi(command[2])

                curChanges, ok := changes[id]

                if ok {

                    for j := len(curChanges)-1; j >= 0; j-- {
                        if (curChanges[j].step<=step){
                            fmt.Fprintln(out, curChanges[j].name)
                            continue steps
                        }
                    }
                }
                fmt.Fprintln(out, "404")
            }
        }

    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
