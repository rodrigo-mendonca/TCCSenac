package main

import (
    "testing"
    "errors"
    //"fmt"
    somf "github.com/rodrigo-mendonca/TCCSenac/somfunctions"
)

func TestExecute(t *testing.T) {
    PngBefore = "Before.png"
    PngAfter = "After.png"
    Filename = "Food.txt"
    ValidFilename := "Food.txt"
    Train = true

    somf.Gridsize = 25
    somf.Dimensions = 3
    somf.Interactions = 5000
    somf.TxVar = 0.5

    patterns,labels :=somf.LoadFile(ValidFilename)

    Execute()

    for i := 0; i < len(patterns); i++ {
        _,label := somf.Koh.Test(patterns[i])

        if(labels[i] != label){
            err := errors.New("Puts!Deu Erro!")
            t.Fatal(err)
        }
        /*
        fmt.Printf("\nO:%s   R:%s\n",labels[i],label)

        fmt.Printf(" [")
        for j := 0; j < len(Weights); j++ {
            fmt.Printf(" %d ",int(Weights[j]*100))
        }
        fmt.Printf("]")
        */
    }
}