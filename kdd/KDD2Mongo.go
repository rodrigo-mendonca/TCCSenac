package kdd

import (
	"fmt"
	"os"
	"strings"
	"reflect"
	"strconv"
	"bufio"
    "gopkg.in/mgo.v2"
)

/* ataques
back,
buffer_overflow,
ftp_write,
guess_passwd,
imap,ipsweep,
land,
loadmodule,
multihop,
neptune,
nmap,
normal,
perl,
phf,
pod,
portsweep,
rootkit,
satan,
smurf,
spy,
teardrop,
warezclient,
warez
*/
type KDDCup struct{
	Duration int
	Protocol_type string
	Service string
	Flag string
	Src_bytes int
	Dst_bytes int
	Land string
	Wrong_fragment int
	Urgent int
	Hot int
	Num_failed_logins int
	Logged_in string
	Num_compromised int
	Root_shell int
	Su_attempted int 
	Num_root int
	Num_file_creations int
	Num_shells int
	Num_access_files int
	Num_outbound_cmds int
	Is_host_login string
	Is_guest_login string
	Count int
	Srv_count int
	Serror_rate int
	Srv_serror_rate int
	Rerror_rate int
	Srv_rerror_rate int
	Same_srv_rate int
	Diff_srv_rate int
	Srv_diff_host_rate int
	Dst_host_count int
	Dst_host_srv_count int
	Dst_host_same_srv_rate int
	Dst_host_diff_srv_rate int
	Dst_host_same_src_port_rate int
	Dst_host_srv_diff_host_rate int
	Dst_host_serror_rate int
	Dst_host_srv_serror_rate int
	Dst_host_rerror_rate int
	Dst_host_srv_rerror_rate int
	Attack string
}

type KDDNormal struct{
	Attack string
	Duration float64
	Protocol_type float64
	Service float64
	Flag float64
	Src_bytes float64
	Dst_bytes float64
	Land float64
	Wrong_fragment float64
	Urgent float64
	Hot float64
	Num_failed_logins float64
	Logged_in float64
	Num_compromised float64
	Root_shell float64
	Su_attempted float64 
	Num_root float64
	Num_file_creations float64
	Num_shells float64
	Num_access_files float64
	Num_outbound_cmds float64
	Is_host_login float64
	Is_guest_login float64
	Count float64
	Srv_count float64
	Serror_rate float64
	Srv_serror_rate float64
	Rerror_rate float64
	Srv_rerror_rate float64
	Same_srv_rate float64
	Diff_srv_rate float64
	Srv_diff_host_rate float64
	Dst_host_count float64
	Dst_host_srv_count float64
	Dst_host_same_srv_rate float64
	Dst_host_diff_srv_rate float64
	Dst_host_same_src_port_rate float64
	Dst_host_srv_diff_host_rate float64
	Dst_host_serror_rate float64
	Dst_host_srv_serror_rate float64
	Dst_host_rerror_rate float64
	Dst_host_srv_rerror_rate float64
}
type t struct {
        Duration int
    }

var totalreg int
var server string
var Db *mgo.Session
var err error

func main() { 
	server ="localhost"
	filename := "kddcup"
	

	// faz a conexao com a base de dados
	Db, err = mgo.Dial(server)
    if err != nil {
        panic(err)
    }

    // Optional. Switch the session to a monotonic behavior.
    Db.SetMode(mgo.Monotonic, true)

	// faz a leitura do arquivo
	file,err := os.Open(filename)
	checkerro(err)
	totalreg = 0

	reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        line:=scanner.Text()

		scanline(line)
	}
	fmt.Printf("Total de registros: %i",totalreg)
	Db.Close()
}

func scanline(l string){
	newreg := new(KDDCup)
	
	
	// carrega todos os registros no padrao da KDDCup
	val := reflect.ValueOf(newreg).Elem()
	
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

        f := val.FieldByName(typeField.Name)        

        if f.IsValid() {
            if f.Kind() == reflect.Int {
                x,_ := strconv.Atoi(readreg(&l))
                f.SetInt(int64(x))
            }

            if f.Kind() == reflect.String {
                x := readreg(&l)
                 f.SetString(x)
            }
        }
	}
	
	newregnormal := Normalize(newreg)

	Colletion := Db.DB("TCC").C("KDDCup")

	err =Colletion.Insert(newreg)
	if err != nil {
		fmt.Printf("Erro Linha: %n\n",totalreg)
		checkerro(err)
    }

    Colletion = Db.DB("TCC").C("KDDNormal")

	err =Colletion.Insert(newregnormal)
	if err != nil {
		fmt.Printf("Erro Linha: %n\n",totalreg)
		checkerro(err)
    }
	totalreg = totalreg + 1
}

func readreg(l *string) string{
	i 	:= strings.Index(*l, ",")

	// valida final de arquivo
	if i < 0 {
		i=len(*l)-1
	}

	reg := (*l)[:i]
	*l 	= (*l)[i+1:]

	return reg
}

func checkerro(e error) {
    if e != nil {
        panic(e)
    }
}

func Normalize(kdd KDDCup) KDDNormal {
	newregnormal := new(KDDNormal)

	newregnormal.Attack 	= kdd.Attack
	newregnormal.Duration 	= float64(kdd.Duration)
}
