package main

var pivot bool

var divNum func(i int) int = div10

func div10(i int) int {
    return i / 10
}

func div100(i int) int {
    return i / 100
}

func ifelse(i int) int {
    if pivot {
        return div10(i)
    } else {
        return div100(i)
    }
}

func funcpoint(i int) int {
    return divNum(i)
}

