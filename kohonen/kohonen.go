package kohonen

import (
    neu "../neuron"
    "os"
    "math"
    "math/rand"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "time"
)

var (
    white color.Color = color.RGBA{255, 255, 255, 255}
    black color.Color = color.RGBA{0, 0, 0, 255}

    red  color.Color = color.RGBA{255, 0, 0, 255}
    green  color.Color = color.RGBA{0, 255, 0, 255}
    blue  color.Color = color.RGBA{0, 0, 255, 255}
)

type Kohonen struct {
    Grid [][]neu.Neuron

    Interactions int
    Gridsize, Dimensions, DimensionsOut, NumReg int
    TxVar float64

    Patterns, Result [][]float64
    Labels []string
    Normal []float64
}

func (r Kohonen) Create(l int, d int, i int, v float64) Kohonen{
    r.Gridsize = l
    r.Dimensions = d
    r.Interactions = i
    r.TxVar = v

    // seguindo o numero de registros diferentes, cria a grid de retorno sendo uma matriz identidade
    r.Result = make([][]float64,len(r.Labels))
    for i := 0; i < len(r.Labels); i++ {
        r.Result[i] = make([]float64,len(r.Labels))
        r.Result[i][i] = 1
    }

    rand.Seed(time.Now().Unix())
    r = r.Initialise()

    return r
}

func (r Kohonen) Empty() bool{
    return len(r.Grid) == 0
}

func (r Kohonen) Initialise() Kohonen{

    r.Grid=make([][]neu.Neuron,r.Gridsize)

    for i := 0; i < r.Gridsize; i++ {
        r.Grid[i] = make([]neu.Neuron,r.Gridsize)
        for j := 0; j < r.Gridsize; j++ {

            neu:=r.Grid[i][j]
            neu = neu.Create(i,j,r.Gridsize,r.Interactions,r.TxVar)

            neu.Weights = make([]float64,r.Dimensions)
            neu.WeightsOut = make([]float64, r.DimensionsOut)
            neu.RGB = make([]int,r.Dimensions)


            for k := 0; k < r.Dimensions; k++ {
                neu.Weights[k] = rand.Float64()
                neu.RGB[k] = int((neu.Weights[k] * 255))
            }

            for k := 0; k < r.DimensionsOut; k++ {
                neu.WeightsOut[k] = rand.Float64();
            }
            r.Grid[i][j] = neu
        }
    }

    return r
}

func (r Kohonen) Draw(f string){
    Screen := image.NewRGBA(image.Rect(0, 0, r.Gridsize, r.Gridsize))
    draw.Draw(Screen, Screen.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

    for i := 0; i < r.Gridsize; i++ {
        for j := 0; j < r.Gridsize; j++ {
            red:=uint8(r.Grid[i][j].RGB[0])
            green:=uint8(r.Grid[i][j].RGB[1])
            blue:=uint8(r.Grid[i][j].RGB[2])

            Screen.Set(i, j, color.RGBA{red, green, blue, 255})
        }
    }

    After, _ := os.Create(f)
    png.Encode(After, Screen)
    After.Close()
}

func (r Kohonen) NormalisePatterns() Kohonen{
    r.Normal = make([]float64,r.Dimensions)

    for j := 0; j < r.Dimensions; j++ {
        max:=float64(0)
        for _, num := range r.Patterns {
            if(max < num[j]){
                max = num[j]
            }
        }
        r.Normal[j] = max
        for i := 0; i < r.NumReg; i++ {
            r.Patterns[i][j] = r.Patterns[i][j] / max
        }
    }
    return r
}

func (r Kohonen) Train() Kohonen{
    r = r.NormalisePatterns()

    for inter := 1; inter <= r.Interactions; inter++ {
        for i := 0; i < r.NumReg; i++ {
            r = r.TrainPattern(inter,r.Patterns[i],r.Result[i])
        }
    }
    return r
}

func (r Kohonen) TrainPattern(inter int, pattern []float64,out []float64) (Kohonen){
    var winner neu.Neuron
    winner = r.Winner(pattern)
    
    aux:=r.Grid

    for i := 0; i < r.Gridsize; i++ {
        for j := 0; j < r.Gridsize; j++ {
            aux[i][j] = r.Grid[i][j].UpdateWeigths(pattern, out, winner, inter)
        }   
    }
    r.Grid = aux
    
    return r
}

func (r Kohonen) Test(pattern []float64) ([]float64,string) {
    
    // normaliza a entrada
    for i := 0; i < r.Dimensions; i++ {
        pattern[i] = pattern[i] / r.Normal[i]
    }

    neu := r.Winner(pattern)
    max:=0
    //fmt.Printf("[")
    for i := 0; i < len(neu.WeightsOut); i++ {
        if(neu.WeightsOut[max] <= neu.WeightsOut[i]){
            max = i
        }
        //fmt.Printf(" %d ",int(neu.WeightsOut[i]*100))
    }

    return neu.WeightsOut,r.Labels[max]
}

func (r Kohonen) Winner(pattern []float64) neu.Neuron{
    var winner neu.Neuron

    min:= math.Sqrt(float64(len(pattern)))

    for i := 0; i < r.Gridsize; i++ {
        for j := 0; j < r.Gridsize; j++ {
            
            dist:=r.Distance(r.Grid[i][j].Weights,pattern)

            if(dist< min){
                min = dist
                
                winner = r.Grid[i][j]
            }
        }
    }
    return winner
}

func (r Kohonen) Distance(v1 []float64,v2 []float64)  float64{
    v:=float64(0)

    for i := 0; i < len(v1); i++ {
        v+=math.Pow(v1[i] -v2[i],2)
    }
    return math.Sqrt(v)
}