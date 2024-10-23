package main

import (
    "fmt"
    "bufio"
    "os"
)

type Rectangle struct{
    w int
    h int
}

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscanln(in, &setsCount)

    sets:
    for set := 1; set <= setsCount; set++ {             // перебираем наборы данных
        var boxCount, paintingsCount,val1, val2 int

        fmt.Fscanln(in, &boxCount)                                              // Цикл 1, коробки, не устраним
        boxes := make([]Rectangle, boxCount)
        prevBoxCount := make([]int, boxCount)
        for num := 0; num < boxCount; num++ {
            fmt.Fscanln(in, &val1, &val2)
            boxes[num].w = val1
            boxes[num].h = val2

            prevBoxCount[num] = 9999999  // предыщее кол-во картин, поместившихся в коробке, для пропуска тактов в цикле 4
        }

        fmt.Fscanln(in, &paintingsCount)
        paintings := make([]Rectangle, paintingsCount)
        for num := 0; num < paintingsCount; num++ {                             // Цикл 2, картины, не устраним
            fmt.Fscanln(in, &val1, &val2)
            paintings[num].w = val1
            paintings[num].h = val2
        }

        paintingsByBox := make([][]int,boxCount)
        flags := make(map[int] []bool)
        for pnt := 0; pnt < paintingsCount; pnt++ {
            count := 0
            flags[pnt] = make([]bool,boxCount)
            for box := 0; box < boxCount; box++ {                               // Цикл 3, "коробки" Х "картины"
                if contain(paintings[pnt],boxes[box]){
                    paintingsByBox[box] = append(paintingsByBox[box],pnt)       // Можно упихнуть в предыдущий
                    count++                                                     // Может можно как-то пропускать коробки
                    flags[pnt][box] = true
                }

            }

            if (count == 0){
                fmt.Fprintln(out, "-1")
                continue sets 
            }
        }


        needBoxCount := 0
        for true {

            betterBox := 0
            betterCount := 0
            for box := 0; box < boxCount; box++ {

                if (prevBoxCount[box]<=betterCount){
                    continue
                }

                currentCount := 0
                for pnt := range flags {                                        // Цикли 4, "коробки" Х "картины" Х "необходимые коробки"
                    if flags[pnt][box]{                                         // Хранить предыдущие коробки, и пропускать слишком мелкие
                        currentCount ++
                    }
                }

                prevBoxCount[box] = currentCount

                if currentCount > betterCount {
                    betterCount = currentCount
                    betterBox = box
                }
            }

            if(betterCount == 0){
                break
            }

            needBoxCount ++

            for _,delpnt := range paintingsByBox[betterBox] {                   // Цикл 5, картины
                delete(flags, delpnt)
            }
        }

        fmt.Fprintln(out, needBoxCount)
    }
}

func contain(inner,outer Rectangle) bool {
    return ((inner.w <= outer.w)&&(inner.h <= outer.h))||((inner.h <= outer.w)&&(inner.w <= outer.h))
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
