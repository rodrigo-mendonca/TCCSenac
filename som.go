package main

import (
    //"./somstructs"
    //somf "github.com/rodrigo-mendonca/TCCSenac/somfunctions"
    somf "./somfunctions"
    "fmt"
    "flag"
    "os"
    "bufio"
)

var Loadtype int
var Filename, Savename string
var Saved, Train, Normalize bool
var PngBefore,PngAfter string

func main() {
    LoadParams()
    
    Execute()
}

func Execute(){
    ShowParams()
    var patterns [][]float64
    var out [][]float64
    var labels []string

    // faz a leitura dos dados de treinamento
    if Loadtype == 0 {
        patterns,out,labels =somf.LoadFile(Filename)
    }
    if Loadtype == 1 {
        patterns, out, labels =somf.LoadKDDCup()
        Normalize = false
    }
    if Loadtype == 2 {
        somf.Koh = somf.LoadJson(Filename)
    }else{
        somf.Koh.Patterns = patterns
        somf.Koh.NumReg   = len(patterns)
        somf.Koh.DimensionsOut = len(labels)
        somf.Koh.Labels   = labels
        somf.Koh.Result = out

        somf.Koh = somf.Koh.Create(somf.Gridsize,somf.Dimensions,somf.Interactions,somf.TxVar)
    }

    if somf.Koh.Empty() {
        return
    }

    // Desenha o estado atual da grade antes do treino
    somf.Koh.Draw(PngBefore)

    if Normalize {
        somf.Koh = somf.Koh.NormalisePatterns()
    }

    if Train {
        // faz o treinamento da base de dados
        somf.Koh = somf.Koh.Train()
    }
    // Desenha o estado atual da grade depois do treino
    somf.Koh.Draw(PngAfter)

    // verifica se deve salvar o treinamento
    if Saved{
        somf.SaveJson(Savename)
    }
}

func LoadParams(){
    flag.StringVar(&somf.Server,"server", "localhost", "Server name")
    flag.StringVar(&somf.Dbname,"base", "TCC", "Data base name")
    flag.StringVar(&somf.Colname,"colletion", "10KDDNormal", "Data base name")
    flag.IntVar(&somf.Gridsize,"grid", 10, "Grid Size")
    flag.IntVar(&somf.Dimensions,"dim", 3, "Dimensions Weigths")
    flag.IntVar(&somf.Interactions,"ite", 5000, "Iteractions")
    flag.Float64Var(&somf.TxVar,"var", 0.5, "Taxa Variation")

    flag.BoolVar(&Saved,"s", false, "Save?")
    flag.BoolVar(&Train,"t", false, "Train?")
    flag.BoolVar(&Normalize,"n", true, "Normalize?")
    flag.StringVar(&Savename,"sname", "Train.json", "Save file name")
    flag.IntVar(&Loadtype,"type", 0, "0-Load file, 1-Load KddCup, 2-Json File")
    flag.StringVar(&Filename,"f", "", "File name")

    config:= flag.String("config", "", "Config file")

    flag.Parse()
    
    // usando arquivo de configuracao
    if *config!="" {
        fmt.Println("-Config:", *config)
        file,err := os.Open(*config)
        somf.Checkerro(err)

        reader := bufio.NewReader(file)
        scanner := bufio.NewScanner(reader)
        for scanner.Scan() {
            line:=scanner.Text()
            fmt.Println("--"+line)
        }
    }

    if  Train {
        fmt.Println("Trainning...")
    } else {
        fmt.Println("Loading...")
    }
    
}

func ShowParams() {
    fmt.Println("Params")
    fmt.Println("-Type:", Loadtype)

    if Loadtype == 0{
        fmt.Println("-File:", Filename)
    }

    if Loadtype == 1{
        fmt.Println("-Server:", somf.Server)
        fmt.Println("-DataBase:", somf.Dbname)
    }

    if Loadtype == 2{
        fmt.Println("-Json:", Filename)
    }

    fmt.Println("-Grid Size:", somf.Gridsize)
    fmt.Println("-Interactions:", somf.Interactions)
    fmt.Println("-Variation:", somf.TxVar)
    fmt.Println("-Save?:", Saved)
    fmt.Println("-Train?:", Train)
    fmt.Println("-Normalize?:", Normalize)

    if Saved{
        fmt.Println("  -File name:", Savename)
    }

}