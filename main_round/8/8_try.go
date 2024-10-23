package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

type Rectangle struct{
    w int
    h int
    used bool
}

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscanln(in, &setsCount)

    for set := 1; set <= setsCount; set++ {             // перебираем наборы данных
        var count,val1,val2 int
    

        fmt.Fscanln(in, &count)
        boxes := make([]Rectangle, count)

        for num := 0; num < count; num++ {
            fmt.Fscanln(in, &val1, &val2)
            if(val1>val2){
                boxes[num].w = val1
                boxes[num].h = val2
            }else{
                boxes[num].w = val2
                boxes[num].h = val1
            }
            boxes[num].used = false
        }

        sort.Slice(boxes, func(a, b int) bool {
            if(boxes[a].w != boxes[b].w){
                return boxes[a].w < boxes[b].w
            }else{
                return boxes[a].h < boxes[b].h
            }
        })


        maxHeights := make([]int, count)
        maxHeight := 0
        for i := count-1; i >= 0; i-- {
            if(boxes[i].h >  maxHeight){
                maxHeight = boxes[i].h
            }

            maxHeights[i] = maxHeight
        }

        //-------------------------------
        
        var inp Rectangle

        bad := false
        needCount := 0
        fmt.Fscanln(in, &count)
        for num := 0; num < count; num++ {
            fmt.Fscanln(in, &val1, &val2)

            if (bad){           // Нужно вхолостую собрать остальные строки
                continue
            }


            if(val1>val2){
                inp.w = val1
                inp.h = val2
            }else{
                inp.w = val2
                inp.h = val1
            }

            minIndex, found := sort.Find(len(boxes), func(i int) int {
                if inp.w <= boxes[i].w {
                    return 0
                }else{
                    return 1
                }
            })

            if !found {
                needCount = -1
                bad = true
                continue
            }


            maxIndex, found2 := sort.Find(len(maxHeights), func(i int) int {
                if inp.h > maxHeights[i] {
                    return 0
                }else{
                    return 1
                }
            })

            maxIndex --

            if !found2 {
                needCount = -1
                bad = true
                continue
            }


            if (minIndex>maxIndex){
                needCount = -1
                bad = true
                continue
            }else{
                if (!boxes[minIndex].used){
                    boxes[minIndex].used = true
                    needCount ++
                }
            }

                
        }

        fmt.Fprintln(out, needCount)
    }
}




func contain(inner,outer Rectangle) bool {
    return (inner.w <= outer.w)&&(inner.h <= outer.h)
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
