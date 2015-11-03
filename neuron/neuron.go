package kohonen

import (
    "math"
)

type Neuron struct {
    Weights []float64
    WeightsOut []float64
    RGB []int
    X,Y,Length int
    Nf float64
    MaxInteraction float64
    Txvar float64
}

func (r Neuron) Create(x int, y int, l int,maxi int, v float64) Neuron{
    r.X = x
    r.Y = y
    r.Length = l
    r.MaxInteraction = float64(maxi)
    r.Txvar = v

    dl:=float64(l)

    log:=math.Log(dl)
    r.Nf = float64(maxi) / log
    
    return r
}

func (r Neuron) Gauss(it int,le float64,dist float64) float64 {
    dit:=float64(it)

    return math.Exp(-0.1*math.Pow(dist, 2) / (2 * le * dit))
}

func (r Neuron) Strength(it int) float64 {
    dit:=float64(it)
    dl:=float64(r.Length / 2)

    result:= math.Exp(-dit / r.Nf) * dl

    return result
}

func (r Neuron) LearningRate(it int) float64 {
    dit:=float64(it)

    return math.Exp(-1.0*dit / r.Nf) * r.Txvar
}

func (r Neuron) UpdateWeigths(pattern,patternout []float64,winner Neuron,it int) (Neuron) {
    
    le:=r.Strength(it)
    dx:=float64(winner.X - r.X)
    dy:=float64(winner.Y - r.Y)

    dist:=math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

    if(dist < le){
        Gau:=r.Gauss(it,le,dist)
        Lea:=r.LearningRate(it)

        for i := 0; i < len(r.Weights); i++ {
            delta:= Lea * Gau * (pattern[i] - r.Weights[i])

            r.Weights[i]+=delta

            r.RGB[i] = int((r.Weights[i] * 255))

            if r.RGB[i] > 255 {
                r.RGB[i] = 255
            }
        }

        for i := 0; i < len(r.WeightsOut); i++ {
            delta:= Lea * Gau * (patternout[i] - r.WeightsOut[i])

            r.WeightsOut[i]+=delta
        }
    }
    return r
}