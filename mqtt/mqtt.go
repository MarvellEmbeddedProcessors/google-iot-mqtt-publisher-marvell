// Package mqtt is used to publish message to MQTT broker.
package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	paho "github.com/eclipse/paho.mqtt.golang"
)

// generateJWT from the Project ID and the private key.
func generateJWT(projectId string) (*string, error) {
	// Generate JWT.
	privateKeyBytes, err := ioutil.ReadFile("certs/rsa_private.pem")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	jwtToken := jwt.New(jwt.SigningMethodRS256)
	claims := jwtToken.Claims.(jwt.MapClaims)
	if err != nil {
		return nil, err
	}
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["aud"] = projectId
	jwtTokenString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &jwtTokenString, nil
}

// PublishMessage publishes a message to MQTT broker.
func PublishMessage(projectId, registryId, deviceId, topic, message string) {
	// Generate JWT.
	jwtToken, err := generateJWT(projectId)
	if err != nil {
		log.Println("error generating JWT: ", err)
		os.Exit(1)
	}

	opts := paho.NewClientOptions()
	opts.AddBroker("tls://mqtt.googleapis.com:8883")
	opts.SetClientID("projects/" + projectId + "/locations/us-central1/registries/" + registryId + "/devices/" + deviceId)
	// Username is not used by the Google Cloud Platform but we need to set it here since
	// the MQTT client library will not set the password if the username is also not set.
	opts.SetUsername("unused")
	opts.SetPassword(*jwtToken)

	// Apply roots certificate.
	rootPEM, err := ioutil.ReadFile("certs/roots.pem")
	if err != nil {
		log.Println("error opening root certificate: ", err)
		os.Exit(1)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		log.Println("failed to parse root certificate")
		os.Exit(1)
	}
	tlsConf := &tls.Config{RootCAs: roots}
	opts.SetTLSConfig(tlsConf)

	// Create and start a client using the above ClientOptions.
	c := paho.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		// Unable to connect to the MQTT broker.
		log.Println(token.Error())
		os.Exit(1)
	}

	// Publish message to topic at QOS 1 and wait for the receipt from the server after sending each message.
	token := c.Publish("/devices/"+deviceId+"/"+topic, 0, false, message)
	token.Wait()

	// Disconnect client from broker.
	c.Disconnect(250)
}
