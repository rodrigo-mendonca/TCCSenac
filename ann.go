package main

import (
	"bufio"
	"flag"
	"fmt"
	somf "github.com/rodrigo-mendonca/TCCSenac/somfunctions"
	ann "github.com/tadvi/ann"
	"os"
)

var Loadtype int
var Filename, Savename string
var Saved, Train, Normalize bool
var PngBefore, PngAfter string

func main() {
	LoadParams()

	Execute()
}

func Execute() {
	ShowParams()
	var patterns [][]float64
	var out [][]float64
	//var labels []string

	// faz a leitura dos dados de treinamento
	if Loadtype == 0 {
		patterns, out, _ = somf.LoadFile(Filename)
	}
	if Loadtype == 1 {
		patterns, out, _ = somf.LoadKDDCup()
	}

	nn := ann.NewBackprop(10, 19, 1)

	fmt.Println("Preparando dados.")
	tr := []*ann.TrainingData{}
	const inputSize = 26

	for i := 0; i < len(patterns); i++ {
		data := &ann.TrainingData{
			Input:  make([]float64, inputSize, inputSize),
			Output: make([]float64, 1, 1),
		}

		for j := 0; j < inputSize; j++ {
			data.Input[j] = patterns[i][j]
		}

		if out[i][0] == 1 {
			data.Output[0] = 0
		} else {
			data.Output[0] = 1
		}

	}

	fmt.Println("training")
	nn.Train(5000, tr)
}

func LoadParams() {
	flag.StringVar(&somf.Server, "server", "localhost", "Server name")
	flag.StringVar(&somf.Dbname, "base", "TCC", "Data base name")
	flag.StringVar(&somf.Colname, "colletion", "10KDDNormal", "Data base name")
	flag.IntVar(&somf.Gridsize, "grid", 10, "Grid Size")
	flag.IntVar(&somf.Dimensions, "dim", 3, "Dimensions Weigths")
	flag.IntVar(&somf.Interactions, "ite", 5000, "Iteractions")
	flag.Float64Var(&somf.TxVar, "var", 0.5, "Taxa Variation")

	flag.BoolVar(&Saved, "s", false, "Save?")
	flag.BoolVar(&Train, "t", false, "Train?")
	flag.BoolVar(&Normalize, "n", true, "Normalize?")
	flag.StringVar(&Savename, "sname", "Train.json", "Save file name")
	flag.IntVar(&Loadtype, "type", 0, "0-Load file, 1-Load KddCup, 2-Json File")
	flag.StringVar(&Filename, "f", "", "File name")

	config := flag.String("config", "", "Config file")

	flag.Parse()

	// usando arquivo de configuracao
	if *config != "" {
		fmt.Println("-Config:", *config)
		file, err := os.Open(*config)
		somf.Checkerro(err)

		reader := bufio.NewReader(file)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("--" + line)
		}
	}

	if Train {
		fmt.Println("Trainning...")
	} else {
		fmt.Println("Loading...")
	}

}

func ShowParams() {
	fmt.Println("Params")
	fmt.Println("-Type:", Loadtype)

	if Loadtype == 0 {
		fmt.Println("-File:", Filename)
	}

	if Loadtype == 1 {
		fmt.Println("-Server:", somf.Server)
		fmt.Println("-DataBase:", somf.Dbname)
	}

	if Loadtype == 2 {
		fmt.Println("-Json:", Filename)
	}

	fmt.Println("-Grid Size:", somf.Gridsize)
	fmt.Println("-Interactions:", somf.Interactions)
	fmt.Println("-Variation:", somf.TxVar)
	fmt.Println("-Save?:", Saved)
	fmt.Println("-Train?:", Train)
	fmt.Println("-Normalize?:", Normalize)

	if Saved {
		fmt.Println("  -File name:", Savename)
	}

}
