package main

import "net/http"

func main() {
	//server := &http.Server{
	//	//  自定义的请求核心处理函数
	//	Handler: framework.NewCore(),
	//	//  请求监听地址
	//	Addr: ":8080",
	//}
	//server.ListenAndServe()
	//resp, err := http.Get("http://dcim-ito.hit.edu.cn")
	//fmt.Println(err, resp)
	//http.ListenAndServe("", nil)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("coredemo/cert.pem", "coredemo/key.pem")
	//max := new(big.Int).Lsh(big.NewInt(1), 128)
	//serialNumber, _ := rand.Int(rand.Reader, max)
	//subject := pkix.Name{
	//	Organization:       []string{"Manning Publications Co."},
	//	OrganizationalUnit: []string{"Books"},
	//	CommonName:         "Go Web Programming",
	//}
	//
	//template := x509.Certificate{
	//	SerialNumber: serialNumber,
	//	Subject:      subject,
	//	NotBefore:    time.Now(),
	//	NotAfter:     time.Now().Add(365 * 24 * time.Hour),
	//	KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	//	ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	//	IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	//}
	//
	//pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	//
	//derBytes, _ := x509.CreateCertificate(rand.Reader, &template,
	//	&template, &pk.PublicKey, pk)
	//certOut, _ := os.Create("/Users/xieth/go/src/GoNote/coredemo/cert.pem")
	//pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	//certOut.Close()
	//
	//keyOut, _ := os.Create("/Users/xieth/go/src/GoNote/coredemo/key.pem")
	//pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	//keyOut.Close()
}
