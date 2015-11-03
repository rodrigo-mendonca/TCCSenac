package main

import (
    //"./somstructs"
    somf "./somfunctions"
    "fmt"
    "flag"
    "testing"
    "os"
    "bufio"
)
var Test bool
var Loadtype int
var Filename, Savename, ValidFilename string
var Saved, Train bool
var PngBefore,PngAfter string

func main() {
    LoadParams()
    PngBefore = "Before.png"
    PngAfter = "After.png"
    ShowParams()

    if Test {
        result := testing.Benchmark(ExecuteTest)
        
        patterns,labels :=somf.LoadFile(ValidFilename)

        seconds := float64(result.T.Seconds()) / float64(result.N)
        fmt.Printf("%13.2f s\n\n", seconds)

        for i := 0; i < len(patterns); i++ {
            Weights,label := somf.Koh.Test(patterns[i])
            fmt.Printf("\nO:%s   R:%s\n",labels[i],label)

            fmt.Printf(" [")
            for j := 0; j < len(Weights); j++ {
                fmt.Printf(" %d ",int(Weights[j]*100))
            }
            fmt.Printf("]")
        }
    }
    
    if !Test  {
        Execute()
    }
}
func ExecuteTest(b *testing.B) {
    Execute()
}

func Execute(){
    var patterns [][]float64
    var labels []string

    // faz a leitura dos dados de treinamento
    if Loadtype == 0 {
        patterns,labels =somf.LoadFile(Filename)
    }
    if Loadtype == 1 {
        patterns,labels =somf.LoadKDDCup("KDDCup")
    }
    if Loadtype == 2 {
        somf.Koh = somf.LoadJson(Filename)
    }else{
        somf.Koh.Patterns = patterns
        somf.Koh.NumReg   = len(patterns)
        somf.Koh.DimensionsOut = len(labels)
        somf.Koh.Labels   = labels

        somf.Koh = somf.Koh.Create(somf.Gridsize,somf.Dimensions,somf.Interactions,somf.TxVar)
    }

    if somf.Koh.Empty() {
        return
    }

    // Desenha o estado atual da grade antes do treino
    somf.Koh.Draw(PngBefore)

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
    flag.IntVar(&somf.Gridsize,"grid", 10, "Grid Size")
    flag.IntVar(&somf.Dimensions,"dim", 3, "Dimensions Weigths")
    flag.IntVar(&somf.Interactions,"ite", 5000, "Iteractions")
    flag.Float64Var(&somf.TxVar,"var", 0.5, "Taxa Variation")

    flag.BoolVar(&Saved,"s", false, "Save?")
    flag.BoolVar(&Train,"t", false, "Train?")
    flag.StringVar(&Savename,"sname", "", "Save file name")
    flag.IntVar(&Loadtype,"type", 0, "0-Load file, 1-Load KddCup, 2-Json File")
    flag.StringVar(&Filename,"f", "", "File name")
    flag.StringVar(&ValidFilename,"fv", Filename, "Valid File name")
    flag.BoolVar(&Test,"test", false, "Test time")

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
        fmt.Println("-DataBase:", somf.Server)
    }

    if Loadtype == 2{
        fmt.Println("-Json:", Filename)
    }

    fmt.Println("-Grid Size:", somf.Koh.Gridsize)
    fmt.Println("-Interactions:", somf.Koh.Interactions)
    fmt.Println("-Variation:", somf.Koh.TxVar)
    fmt.Println("-Save?:", Saved)

    if Saved{
        fmt.Println("  -File name:", Savename)
    }

}