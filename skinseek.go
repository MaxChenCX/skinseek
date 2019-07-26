package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	name    string
	uuid    string
	outfile string
)

func init() {
	flag.StringVar(&name, "n", "", "get from player name, example: MaxChen_CX")
	flag.StringVar(&uuid, "u", "", "get from player uuid, example: c343a67ff3fd49059e8ff58385a0bbfe")
	flag.StringVar(&outfile, "o", "", "output filename or path, example: myskin.png; D:\\myskin.png")
}

func main() {
	fmt.Println(" Minecraft")
	fmt.Println(" ╔═╗┬┌─┬┌┐┌  ╔═╗┌─┐┌─┐┬┌─ ")
	fmt.Println(" ╚═╗├┴┐││││  ╚═╗├┤ ├┤ ├┴┐ ")
	fmt.Println(" ╚═╝┴ ┴┴┘└┘  ╚═╝└─┘└─┘┴ ┴ v1.0")
	fmt.Println("")

	flag.Parse()
	if name == "" && uuid == "" {
		flag.Usage()
		return
	}

	if uuid == "" {
		UUIDp := getUUID(name)
		uuid = UUIDp.UUID
	}
	Sessp := getSessionPro(uuid)
	Skinp := getSkinPro(Sessp.Properties[0].Value)
	fmt.Println("Player Name: " + Skinp.Name)
	fmt.Println("Player UUID: " + Skinp.ID)
	fmt.Println("Skin URL: " + Skinp.Skin.SkinInfo.SkinURL)
	if Skinp.Skin.SkinInfo.MetaData.Model != "" {
		fmt.Println("Skin Type: " + Skinp.Skin.SkinInfo.MetaData.Model)
	}

	if outfile == "" {
		outfile = Skinp.Name + ".png"
	}
	skindata := httpGet(Skinp.Skin.SkinInfo.SkinURL)
	err := ioutil.WriteFile(outfile, skindata, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("Size: " + strconv.Itoa(len(skindata)) + " bytes")
	fmt.Println("SHA1: " + sha1String(skindata))
	fmt.Println("File Saved to: " + outfile)
}

func getSkinPro(base64str string) textures {
	decodeBytes, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		panic(err)
	}
	txu := &textures{}
	err = json.Unmarshal(decodeBytes, txu)
	if err != nil {
		panic(err)
	}
	return *txu
}

func getSessionPro(uuid string) proSession {
	rep := httpGet("https://sessionserver.mojang.com/session/minecraft/profile/" + uuid)
	ps := &proSession{}
	err := json.Unmarshal(rep, ps)
	if err != nil {
		panic(err)
	}
	return *ps
}

func getUUID(name string) proUser {
	rep := httpGet("https://api.mojang.com/users/profiles/minecraft/" + name)
	p := &proUser{}
	err := json.Unmarshal(rep, p)
	if err != nil {
		panic(err)
	}
	return *p
}

func httpGet(url string) []byte {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14")
	if err != nil {
		panic(err)
	}
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func sha1String(data []byte) string {
	SHA1 := sha1.New()
	SHA1.Write(data)
	SHA1Data := SHA1.Sum([]byte(nil))
	return hex.EncodeToString(SHA1Data)
}

type textures struct {
	TimeStamp int64  `json:"timestamp"`
	ID        string `json:"profileId"`
	Name      string `json:"profileName"`
	Skin      skin   `json:"textures"`
}

type skin struct {
	SkinInfo skininfo `json:"SKIN"`
}

type skininfo struct {
	SkinURL  string   `json:"url"`
	MetaData metadata `json:"metadata"`
}

type metadata struct {
	Model string `json:"model"`
}

type proUser struct {
	UUID string `json:"id"`
	Name string `json:"name"`
}

type properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type proSession struct {
	UUID       string       `json:"id"`
	Name       string       `json:"name"`
	Properties []properties `json:"properties"`
}
