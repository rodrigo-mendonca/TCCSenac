package somf

import (
	somk "../kohonen"
	"os"
    "os/exec"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "io/ioutil"
)

var Koh somk.Kohonen
var Server, Dbname string
var Gridsize,Dimensions int
var Error float64
var Interactions int
var TxVar float64


func ShowPng(name string) {
    command := "open"
    arg1 := "-a"
    arg2 := "/Applications/Preview.app"
    cmd := exec.Command(command, arg1, arg2, name)
    err := cmd.Run()

    Checkerro(err)
}

func Checkerro(e error) {
    if e != nil {
        panic(e)
    }
}

func LoadColletion(name string) *mgo.Collection{
    // faz a conexao com a base de dados
    session, err := mgo.Dial(Server)
    if err != nil {
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)

    return session.DB(Dbname).C(name)
}

func LoadFile(f string) ([][]float64,[]string) {
    // faz a leitura do arquivo
    file,err := os.Open(f)
    Checkerro(err)

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    var patterns [][]float64
    var labels []string

    for scanner.Scan() {
        line:=scanner.Text()

        params:=strings.Split(line,",")

        // primeiro parametro deve ser a label do registro
        // verifica se a label ja existe
        find:=false
        for i := 0; i < len(labels); i++ {
            find = labels[i] == params[0]
        }
        if !find{
            labels = append(labels, params[0])
        }
        inputs := make([]float64,Dimensions)

        for i := 1; i <= Dimensions; i++ {
            p:=params[i]

            num,err:=strconv.ParseFloat(p, 64)
            
            inputs[i - 1] = num
            Checkerro(err)
        }
        patterns = append(patterns, inputs)
    }

    return patterns,labels
}

func LoadKDDCup(col string) ([][]float64,[]string){
    //var patterns [][]float64
    var labels []string

    Colletion := LoadColletion(col)
    
    var patterns [][]float64

    err := Colletion.Find(bson.M{}).All(&patterns)
    Checkerro(err)
    //numlines:=0
    //for _,reg:= range patterns{
        //labels = append(labels, string(reg[0]))
        
        //fmt.Printf(reg[0]+"\n")

        //numlines++
    //}

    return patterns,labels
}

func SaveDB(col string){
    Colletion := LoadColletion(col)
    Colletion.RemoveAll(nil)

    Colletion.Insert(Koh)
    fmt.Printf("Treinamento Salvo\n")
    /*
    ind:=0
    for _, newline:= range Koh {
        for _, newreg:= range newline {
            err := Colletion.Insert(newreg)
            Checkerro(err)

            ind++
        }
    }
    */
}

func LoadDB(col string) somk.Kohonen{
    Colletion := LoadColletion(col)

    err := Colletion.Find(bson.M{}).All(&Koh)
    Checkerro(err)

    return Koh
}

func SaveJson(f string){
    o, err := os.Create(f)
    if err != nil {
        panic(err)
    }

    b, err := json.Marshal(Koh)
    if err != nil {
        fmt.Println(err)
        return
    }
    o.WriteString(string(b))

    fmt.Printf("Treinamento Salvo\n")
}

func LoadJson(f string) somk.Kohonen{
    file, e := ioutil.ReadFile(f)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    json.Unmarshal(file, &Koh)

    return Koh
}