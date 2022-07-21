package helpers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/parnurzeal/gorequest"
)

type Auth struct {
	AppID       string `json:"app_id"`
	KeyID       string `json:"key_id"`
	RequestTS   string `json:"request_ts"`
	SignedToken string `json:"signed_token"`
}
type JwtStruct struct {
	Expiry_ts string
	Jwt_token string
}
type ConfigComponent struct {
	KeyIdFrmConfFile   string
	KeyPathFrmConfFile string
	AppIdFrmConfFile   string
	BaseUrlFrmConfFile string
	SecureStatus       string
}

func Autherisation() JwtStruct {
	confData := ReadConfFile()
	var jwtResponseArray JwtStruct
	if confData.SecureStatus == "true" {
		path, _ := filepath.Abs(confData.KeyPathFrmConfFile)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			//return nil, err
		}
		password := []byte("sACaCKCh]Ac|")
		block, _ := pem.Decode(data)
		if block == nil {
			//return nil, errors.New("ssh: no key found")
		}
		der, err := x509.DecryptPEMBlock(block, password)
		if err != nil {
			log.Fatalf("Decrypt failed: %v", err)
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(der)
		if err != nil {
			//return nil, err
			log.Println("err")
		}
		request_ts := time.Now().UTC().Format(time.RFC3339)
		log.Println("request_ts", request_ts)
		toSign := []byte(confData.AppIdFrmConfFile + ":" + request_ts)
		h := sha256.New()
		h.Write(toSign)
		d := h.Sum(nil)
		signed, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, d)
		stringSignKey := hex.EncodeToString(signed)
		//stringSignKey := string(signed)
		var auths Auth
		auths.AppID = confData.AppIdFrmConfFile
		auths.KeyID = confData.KeyIdFrmConfFile
		auths.RequestTS = request_ts
		auths.SignedToken = stringSignKey
		request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		request.SetDebug(true)
		_, body, errs := request.Post("https:/" + confData.BaseUrlFrmConfFile + "/kochi-bus/partner/auth").
			Send(auths).
			End()
		if errs != nil {
			log.Println("error", errs)
		}

		json.Unmarshal([]byte(body), &jwtResponseArray)

	} else {
		log.Println("insecure")
	}
	return jwtResponseArray

}

func ReadConfFile() ConfigComponent {
	confData := ConfigComponent{}
	confData.AppIdFrmConfFile = beego.AppConfig.String("AppId")
	confData.KeyIdFrmConfFile = beego.AppConfig.String("KeyID")
	confData.BaseUrlFrmConfFile = beego.AppConfig.String("BaseUrl")
	confData.KeyPathFrmConfFile = beego.AppConfig.String("PrivateKeyPath")
	confData.SecureStatus = beego.AppConfig.String("HttpSeureStatus")
	return confData

}
