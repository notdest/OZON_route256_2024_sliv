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

    var setsCount int
    fmt.Fscanln(in, &setsCount)

    sets:
    for set := 1; set <= setsCount; set++ {             // перебираем наборы данных
        var n, width, lines int

        fmt.Fscanln(in, &n)
        fmt.Fscanln(in, &lines, &width)

        fields := make([]string,lines)

        for i := 0; i < lines; i++ {
            var line string
            line, _ = in.ReadString('\n')
            line = strings.TrimRight(line, "\n")

            fields[i] = line
        }


        //-------- далее смотрим выигрышные ситуации, и возможность их возникновения

        XHorCount := make([]int,lines)                      // Смотрим все горизонтальные строки
        xCount := 0
        oCount := 0

        for line := 0; line < lines; line++ {
            xCount = 0
            oCount = 0
            for col := 0; col < width; col++ {
                if (fields[line][col] == 'X'){
                    XHorCount[line]++
                    xCount++
                    oCount = 0

                    if (xCount >= n){
                        fmt.Fprintln(out, "NO")
                        continue sets
                    }
                }else if(fields[line][col] == 'O'){
                    xCount = 0
                    oCount++

                    if (oCount >= n){
                        fmt.Fprintln(out, "NO")
                        continue sets
                    }
                }else{
                    xCount = 0
                    oCount = 0
                }
            }
        }


        XVerCount := make([]int,width)                      // Вертикальные строки

        for col := 0; col < width; col++ {
            xCount = 0
            oCount = 0
            for line := 0; line < lines; line++ {
                if (fields[line][col] == 'X'){
                    XVerCount[col]++
                    xCount++
                    oCount = 0

                    if (xCount >= n){
                        fmt.Fprintln(out, "NO")
                        continue sets
                    }
                }else if(fields[line][col] == 'O'){
                    xCount = 0
                    oCount++

                    if (oCount >= n){
                        fmt.Fprintln(out, "NO")
                        continue sets
                    }
                }else{
                    xCount = 0
                    oCount = 0
                }
            }
        }

        hasDiagonals := (lines>=n)&&(width>=n)
        var XMainCount, xSecondaryCount []int
        diagonals := lines+width+1-2*n // количество диагоналей
        if hasDiagonals {
            XMainCount = make([]int,diagonals)
            lineMax := lines-n
            colMax := width-n

            for lineStart := 0; lineStart <= lineMax; lineStart++ {  // от главной диагонали вниз
                xCount = 0
                oCount = 0
                diagonal := lineMax-lineStart
                for i := 0; (lineStart+i < lines)&&(i<width); i++ {
                    val := fields[lineStart+i][i]

                    if (val == 'X'){
                        XMainCount[diagonal]++
                        xCount++
                        oCount = 0

                        if (xCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else if(val == 'O'){
                        xCount = 0
                        oCount++

                        if (oCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else{
                        xCount = 0
                        oCount = 0
                    }
                }
            }

            for colStart := 1; colStart <= colMax; colStart++ {  // от главной диагонали вправо
                xCount = 0
                oCount = 0
                diagonal := colStart+lineMax
                for i := 1; (i < lines)&&(i+colStart <width); i++ {
                    val := fields[i][colStart+i]

                    if (val == 'X'){
                        XMainCount[diagonal]++
                        xCount++
                        oCount = 0

                        if (xCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else if(val == 'O'){
                        xCount = 0
                        oCount++

                        if (oCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else{
                        xCount = 0
                        oCount = 0
                    }
                }
            }



            xSecondaryCount = make([]int,diagonals)                        // Диагонали в направлении побочной

            for lineStart := 0; lineStart <= lineMax; lineStart++ { // ниже побочной
                xCount = 0
                oCount = 0
                diagonal := width-n+lineStart

                for i := 0; (lineStart+i < lines)&&(i<width); i++ {
                    val := fields[lineStart+i][width-i-1]

                    if (val == 'X'){
                        xSecondaryCount[diagonal]++
                        xCount++
                        oCount = 0

                        if (xCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else if(val == 'O'){
                        xCount = 0
                        oCount++

                        if (oCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else{
                        xCount = 0
                        oCount = 0
                    }
                }
            }

            for colStart := width-2; colStart >= n-1; colStart-- {  // левее побочной
                xCount = 0
                oCount = 0
                diagonal := colStart - n +1

                for i := 0; (i < lines)&&(colStart-i >= 0); i++ {
                    val := fields[i][colStart-i]

                    if (val == 'X'){
                        xSecondaryCount[diagonal]++
                        xCount++
                        oCount = 0

                        if (xCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else if(val == 'O'){
                        xCount = 0
                        oCount++

                        if (oCount >= n){
                            fmt.Fprintln(out, "NO")
                            continue sets
                        }
                    }else{
                        xCount = 0
                        oCount = 0
                    }
                }


            }
        }

        for line := 0; line < lines; line++ {
            for col := 0; col < width; col++ {
                if (fields[line][col] == '.'){

                    if (XHorCount[line]>=n-1){
                        if (sumHor(fields,line,col,width)>=n){
                            fmt.Fprintln(out, "YES")
                            continue sets
                        }
                    }

                    if (XVerCount[col]>=n-1){
                        if (sumVert(fields,line,col,lines)>=n){
                            fmt.Fprintln(out, "YES")
                            continue sets
                        }
                    }


                    if hasDiagonals {
                        diagonal := col-line+(lines-n)
                        if( (diagonal>=0) && (diagonal < diagonals) && (XMainCount[diagonal]>=n-1)){
                            if(sumMainDiag(fields,line,col,lines,width)>=n){
                                fmt.Fprintln(out, "YES")
                                continue sets
                            }
                        }

                        diagonal = col + line - n + 1
                        if( (diagonal>=0) && (diagonal < diagonals) && (xSecondaryCount[diagonal]>=n-1)){
                            if(sumSecDiag(fields,line,col,lines,width)>=n){
                                fmt.Fprintln(out, "YES")
                                continue sets
                            }
                        }
                    }
                }       
            }                
        }

        fmt.Fprintln(out, "NO")
    }
}

func sumSecDiag(fields []string,startLine,startCol,lines,width int) int {
    sum := 1 // один Х ещё ставим

    for i := 1; (startLine+i < lines)&&(startCol-i >=0) ; i++ {    // влево вниз
        if (fields[startLine+i][startCol-i] == 'X'){
            sum++
        }else{
            break
        }
    }

    for i := 1; (startLine-i >= 0)&&(startCol+i < width) ; i++ {    // вправо вверх
        if (fields[startLine-i][startCol+i] == 'X'){
            sum++
        }else{
            break
        }
    }

    return sum
}

func sumMainDiag(fields []string,startLine,startCol,lines,width int) int {
    sum := 1 // один Х ещё ставим

    for i := 1; (startLine-i >= 0)&&(startCol-i >=0) ; i++ {    // влево вверх
        if (fields[startLine-i][startCol-i] == 'X'){
            sum++
        }else{
            break
        }
    }

    for i := 1; (startLine+i < lines)&&(startCol+i < width) ; i++ {    // влево вверх
        if (fields[startLine+i][startCol+i] == 'X'){
            sum++
        }else{
            break
        }
    }

    return sum
}

func sumVert(fields []string,startLine,col,lines int) int {
    sum := 1 // один Х ещё ставим

    for line := startLine-1; line >= 0; line-- {   // верхняя сумма
        if (fields[line][col] == 'X'){
            sum++
        }else{
            break
        }
    }

    for line := startLine+1; line < lines; line++ {   // нижняя сумма
        if (fields[line][col] == 'X'){
            sum++
        }else{
            break
        }
    }

    return sum
}

func sumHor(fields []string,line,StartCol,width int) int {
    sum := 1 // один Х ещё ставим

    for col := StartCol-1; col >= 0; col-- { // левая сумма
        if (fields[line][col] == 'X'){
            sum++
        }else{
            break
        }
    }

    for col := StartCol+1; col < width; col++ { //правая сумма
        if (fields[line][col] == 'X'){
            sum++
        }else{
            break
        }
    }

    return sum
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}
