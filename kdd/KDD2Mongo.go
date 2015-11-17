package main

import (
	"fmt"
	"os"
	"strings"
	"reflect"
	"strconv"
	"bufio"
    "gopkg.in/mgo.v2"
)

const ( 
   MaxDuration float64 = 58329.0
   MaxSrc_bytes float64 = 1379963888.0
   MaxDst_bytes float64 = 1309937401.0
   MaxWrong_fragment float64 = 3.0
   MaxUrgent float64 = 14.0
   MaxHot float64 = 77.0
   MaxNum_failed_logins float64 = 5.0
   MaxNum_compromised float64 = 7479.0
   MaxRoot_shell float64 = 1.0
   MaxSu_attempted float64 = 2.0
   MaxNum_root float64 = 7468.0
   MaxNum_file_creations float64 = 43.0
   MaxNum_shells float64 = 2.0
   MaxNum_access_files float64 = 9.0
   MaxCount float64 = 511.0
   MaxSrv_count float64 = 511.0
   MaxDst_host_count float64 = 255.0
   MaxDst_host_srv_count float64 = 255.0
)

var Protocol_type []string = []string { 
	"tcp",
	"udp",
	"icmp"}

var Service []string = []string { 
	"http",
	"smtp",
	"domain_u",
	"auth",
	"finger",
	"telnet",
	"eco_i",
	"ftp",
	"ntp_u",
	"ecr_i",
	"other",
	"urp_i",
	"private",
	"pop_3",
	"ftp_data",
	"netstat",
	"daytime",
	"ssh",
	"echo",
	"time",
	"name",
	"whois",
	"domain",
	"mtp",
	"gopher",
	"remote_job",
	"rje",
	"ctf",
	"supdup",
	"link",
	"systat",
	"discard",
	"X11",
	"shell",
	"login",
	"imap4",
	"nntp",
	"uucp",
	"pm_dump",
	"IRC",
	"Z39_50",
	"netbios_dgm",
	"ldap",
	"sunrpc",
	"courier",
	"exec",
	"bgp",
	"csnet_ns",
	"http_443",
	"klogin",
	"printer",
	"netbios_ssn",
	"pop_2",
	"nnsp",
	"efs",
	"hostnames",
	"uucp_path",
	"sql_net",
	"vmnet",
	"iso_tsap",
	"netbios_ns",
	"kshell",
	"urh_i",
	"http_2784",
	"harvest",
	"aol",
	"tftp_u",
	"http_8001",
	"tim_i",
	"red_i"}

var Flag []string = []string { 
	"SF",
    "S2",
    "S1",
    "S3",
    "OTH",
    "REJ",
    "RSTO",
	"S0",
	"RSTR",
	"RSTOS0",
	"SH"}

var Land []string = []string { 
	"0",
	"1"}

var Logged_in []string = []string { 
	"0",
	"1"}

var Is_host_login []string = []string { 
	"0",
	"1"}

var Is_guest_login []string = []string { 
	"0",
	"1"}

var Attack map[string]string = map[string]string{
	"ack" : "dos",
	"buffer_overflow" : "u2r",
	"guess_passwd" : "r2l",
	"imap" : "r2l",
	"land" : "dos",
	"ipsweep" : "probe",
	"loadmodule" : "u2r",
	"multihop" : "r2l",
	"neptune" : "dos",
	"nmap" : "probe",
	"perl" : "u2r",
	"phf" : "r2l",
	"pod" : "dos",
	"portsweep" : "probe",
	"rootkit" : "u2r",
	"satan" : "probe",
	"smurf" : "dos",
	"spy" : "r2l",
	"teardrop" : "dos",
	"warezclient" : "r2l",
	"warezmaster" : "r2l",
	"normal" : "normal"}

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
	Is_host_login float64
	Is_guest_login float64
	Count float64
	Srv_count float64
	Dst_host_count float64
	Dst_host_srv_count float64
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
		if totalreg % 1000 == 0 {
			fmt.Printf("%d Registros...\n",totalreg)
		}
	}
	fmt.Printf("Total de registros: %n",totalreg)
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

	Colletion := Db.DB("TCC").C("10KDDCup")

	err =Colletion.Insert(newreg)
	if err != nil {
		fmt.Printf("Erro Linha: %d\n",totalreg)
		checkerro(err)
    }

    Colletion = Db.DB("TCC").C("10KDDNormal")

	err =Colletion.Insert(newregnormal)
	if err != nil {
		fmt.Printf("Erro Linha: %d\n",totalreg)
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

func flen(l []string) float64{
	return float64(len(l))
}

func Indexof(arr []string,fin string) float64 {
	for i := 0; i < len(arr); i++ {
		if(arr[i] == fin){
			return float64(i + 1)
		}
	}
	return float64(0)
}

func Normalize(kdd *KDDCup) *KDDNormal {
	newregnormal := new(KDDNormal)

	// nome do ataque
	newregnormal.Attack 			= Attack[kdd.Attack]

	// valores variaveis
	newregnormal.Duration 			= float64(kdd.Duration) / MaxDuration
	newregnormal.Src_bytes			= float64(kdd.Src_bytes) / MaxSrc_bytes
   	newregnormal.Dst_bytes			= float64(kdd.Dst_bytes) / MaxDst_bytes
   	newregnormal.Wrong_fragment		= float64(kdd.Wrong_fragment) / MaxWrong_fragment
   	newregnormal.Urgent				= float64(kdd.Urgent) / MaxUrgent
   	newregnormal.Hot				= float64(kdd.Hot) / MaxHot
   	newregnormal.Num_failed_logins	= float64(kdd.Num_failed_logins) / MaxNum_failed_logins
   	newregnormal.Num_compromised	= float64(kdd.Num_compromised) / MaxNum_compromised
   	newregnormal.Root_shell			= float64(kdd.Root_shell) / MaxRoot_shell
   	newregnormal.Su_attempted		= float64(kdd.Su_attempted) / MaxSu_attempted
   	newregnormal.Num_root			= float64(kdd.Num_root) / MaxNum_root
   	newregnormal.Num_file_creations	= float64(kdd.Num_file_creations) / MaxNum_file_creations
   	newregnormal.Num_shells			= float64(kdd.Num_shells) / MaxNum_shells
   	newregnormal.Num_access_files	= float64(kdd.Num_access_files) / MaxNum_access_files
   	newregnormal.Count				= float64(kdd.Count) / MaxCount
   	newregnormal.Srv_count			= float64(kdd.Srv_count) / MaxSrv_count
   	newregnormal.Dst_host_count		= float64(kdd.Dst_host_count) / MaxDst_host_count
   	newregnormal.Dst_host_srv_count	= float64(kdd.Dst_host_srv_count) / MaxDst_host_srv_count

   	// indices
   	newregnormal.Protocol_type		= Indexof(Protocol_type,kdd.Protocol_type) / flen(Protocol_type)
   	newregnormal.Service			= Indexof(Service,kdd.Service) / flen(Service)
   	newregnormal.Flag				= Indexof(Flag,kdd.Flag) / flen(Flag)
   	newregnormal.Land				= Indexof(Land,kdd.Land) / flen(Land)
   	newregnormal.Logged_in			= Indexof(Logged_in,kdd.Logged_in) / flen(Logged_in)
   	newregnormal.Is_host_login		= Indexof(Is_host_login,kdd.Is_host_login) / flen(Is_host_login)
   	newregnormal.Is_guest_login		= Indexof(Is_guest_login,kdd.Is_guest_login) / flen(Is_guest_login)

	return newregnormal
}
