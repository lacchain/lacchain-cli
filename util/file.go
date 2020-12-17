package util

import (
	"io/ioutil"
    "log"
    "bytes"
	"strings"
    "regexp"
    "encoding/hex"
    "fmt"
    "strconv"
    "github.com/lacchain/lacchain-cli/model" 
)

//ReplaceConfigFile ...
func ReplaceConfigFile(addBootList, removeBootList []model.Event){
	input, err := ioutil.ReadFile("/root/lacchain/config.toml")
    if err != nil {
            log.Fatalln(err)
    }

	r, err := regexp.Compile("^bootnodes")
    if err != nil {
        log.Fatal(err)
	}
	
	lines := strings.Split(string(input), "\n")
	
    for i, line := range lines {
        if r.MatchString(line) {
            
                var buffer bytes.Buffer
                for _,event := range addBootList{
                    buffer.WriteString(",\"")
                    _newBootnode := generateEnode(event.EnodeHigh[:],event.EnodeLow[:],event.EnodeIp[:],event.EnodePort)
                    buffer.WriteString(_newBootnode+"\"")
                }
                line = addBootnode(line,buffer.String()+"]")

                for _,event := range removeBootList{
                    _oldBootnode := generateEnode(event.EnodeHigh[:],event.EnodeLow[:],event.EnodeIp[:],event.EnodePort)
                    line = removeBootnode(line, _oldBootnode)
                }
                lines[i] = line
        }
    }
	
	output := strings.Join(lines, "\n")
    err = ioutil.WriteFile("/root/lacchain/config2.toml", []byte(output), 0644)
    if err != nil {
            log.Fatalln(err)
    }
}

func addBootnode(s, newBootNode string) string {
	result := strings.Replace(s, "]", newBootNode, -1)
	return result
}

func removeBootnode(s, oldBootNode string) string {
    result := strings.Replace(s, ",\""+oldBootNode+"\"", "", -1)
    return result
}

func generateEnode(enodeHigh,enodeLow,enodeIp []byte, enodePort uint16) string {
    high := fmt.Sprintf("0x%s",hex.EncodeToString(enodeHigh[:]))
    low := hex.EncodeToString(enodeLow[:])
    ipHex := hex.EncodeToString(enodeIp[:])
    ip1, _ := strconv.ParseInt(ipHex[24:26], 16, 64)
    ip2, _ := strconv.ParseInt(ipHex[26:28], 16, 64)
    ip3, _ := strconv.ParseInt(ipHex[28:30], 16, 64)
    ip4, _ := strconv.ParseInt(ipHex[30:32], 16, 64)

    ip := strconv.FormatInt(ip1, 10)+"."+strconv.FormatInt(ip2, 10)+"."+strconv.FormatInt(ip3, 10)+"."+strconv.FormatInt(ip4, 10)

    return "enode://"+high+low+"@"+ip+":"+strconv.Itoa(int(enodePort))
}