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
        Weights,label := somf.Koh.Test(patterns[i])

        if(labels[i] != label){
            err := errors.New("Erro: Resposta errada!")
            t.Fatal(err)
        }

        if int(Weights[i]*100) < 99 {
            err := errors.New("Porcentagem menor que 99%")
            t.Fatal(err)
        }
    }
}