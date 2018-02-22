package security

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func SecurityHivridApp(user, token, hash string) (answer bool, err error) {
	var hashGenerated string
	// Paso 1 generar Hash
	if hashGenerated, err = GeneratedHash(user, token); err == nil {
		//println(hashGenerated, hash)
		// Paso 2 Validar que el hash sea el esperado
		if hashGenerated == hash {
			// Consultar sesion # falta
			println("IGUAl")
			answer = true
		} else {
			err = errors.New("the hash does not match")
		}
	}
	return
}

func GeneratedHash(user, token string) (outputHash string, err error) {
	if secret := beego.AppConfig.String("Secret"); secret != "" {
		var buffer bytes.Buffer
		buffer.WriteString(user)
		buffer.WriteString(token)
		buffer.WriteString(secret)
		allString := buffer.String()
		//hash
		h := sha1.New()
		h.Write([]byte(allString))
		outputHash = hex.EncodeToString(h.Sum(nil))
		println(allString, outputHash)
	} else {
		err = errors.New("secret not defined")
	}
	return
}

func GetSesionAcademica(token string) (sesion bool) {
	var dataSeion []interface{}
	if err := request.GetJson("http://jbpm.udistritaloas.edu.co:8280/services/uranoPruebasProxy/get_usuario_session/"+token, &dataSeion); err == nil && dataSeion != nil {
		beego.Info(dataSeion)
	}

}
