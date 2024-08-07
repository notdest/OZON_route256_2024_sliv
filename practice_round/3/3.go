package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func main() {
    in,out := getBuffers()
    defer out.Flush()

    var setsCount int
    fmt.Fscan(in, &setsCount)

    for set := 1; set <= setsCount; set++ {       // перебираем наборы данных

        nums := inputChilds(in)
        root := missingElement(nums);

        fmt.Fprintln(out, root)
    }
}



// специфичный формат ввода графа, вводим потомков
func inputChilds(in *bufio.Reader)[]int{
    var numCount, vnikuda int

    nums := make([]int, 0)
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

            nums = append(nums, val)
        }
    }

    return nums
}


func missingElement(nums []int) int {
    sort.Slice(nums, func(a, b int) bool {
            return nums[a] < nums[b]
    })

    var cnt int
    cnt = len(nums)

    if ((cnt == 0) || nums[0] != 1){
        return 1;
    }

    for i := 0; i < cnt-1; i++ {
        if (nums[i+1] - nums[i] > 1){
            return nums[i] +1
        }
    }

    return cnt+1;
}


func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
