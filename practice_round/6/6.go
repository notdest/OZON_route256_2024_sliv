package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

const boxSizes = 30
type BoxList [boxSizes]float64

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscanln(in, &setsCount)
    for set := 1; set <= setsCount; set++ {     // перебираем наборы данных
        var carsCount int
        var carsCapacity float64
        fmt.Fscan(in, &carsCount, &carsCapacity)

        boxes := inputBoxes(in)

        min, max := bounds(boxes)

        trips := 0

        for{
            for car := 0; car < carsCount; car++ {      // грузим все машины в рейсе, поглядывая осталось ли что грузить
                newBounds := loadCar(&boxes,carsCapacity, min, max)

                if newBounds {
                    min, max = bounds(boxes)

                    if(min >= boxSizes){
                        break
                    }
                }
            }

            trips ++

            if(min >= boxSizes){
                break
            }
        }

        fmt.Fprintln(out, trips)
    }
}

func inputBoxes(in *bufio.Reader) BoxList{
    var boxes BoxList
    for k := 0; k < boxSizes; k++ {
        boxes[k] = 0
    }

    var boxCount, val int
    fmt.Fscan(in, &boxCount)
    for j := 0; j < boxCount; j++ {
        fmt.Fscan(in, &val)

        boxes[val] ++
    }

    return boxes
}


func loadCar(boxes *BoxList, carsCapacity float64, min, max int) bool{
    newBounds := false

    remainder := carsCapacity
    minBox := math.Pow(2,float64(min))


    maxBox := math.Floor( math.Log2(remainder) )
    realMax := math.Min(maxBox, float64(max))

    for i := int(realMax); i >= min; i-- {
        if (boxes[i] == 0){
            continue
        }

        boxSize := math.Pow(2,float64(i))
        maxBoxes := math.Floor(remainder/boxSize)

        if(maxBoxes<boxes[i]){
            boxes[i] -= maxBoxes
            remainder -= maxBoxes*boxSize
        }else{
            newBounds = true
            remainder -= boxSize*boxes[i]
            boxes[i] = 0
        }


        if(remainder<minBox){
            break
        }

    }

    return newBounds
}

func bounds(boxes BoxList) (int, int) {
    min := boxSizes
    for i := 0; i < boxSizes; i++ {
        if(boxes[i] != 0){
            min = i
            break
        }
    }

    max := -1
    for i := boxSizes-1; i >=0 ; i-- {
        if(boxes[i] != 0){
            max = i
            break
        }
    }
    return min, max
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
