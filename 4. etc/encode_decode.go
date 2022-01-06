package main

import "fmt"
import "encoding/base64"

func main(){

	/*----------------------------- CARA 1 ------------------------------- */
	var data = "john wick"

	//encode
	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded : ", encodeString)

	//decode
	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("Decoded : ", decodeString)
	/*----------------------------- CARA 1 ------------------------------- */

	/*----------------------------- CARA 2 ------------------------------- */

	//encode
	var data2 = "Hisbikal Haqqi"
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data2)))
	base64.StdEncoding.Encode(encoded, []byte(data2))
	var encodedString = string(encoded)
	fmt.Println(encodedString)


	//decode
	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err  != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)
	/*----------------------------- CARA 2 ------------------------------- */


	/*----------------------------- URL ENCODING ------------------------------- */

	//encode
	var data3 = "http://google.com"
	var encodeUrl = base64.URLEncoding.EncodeToString([]byte(data3))
	fmt.Println(encodeUrl)

	//decode
	var decodeBytes, _  = base64.URLEncoding.DecodeString(encodeUrl)
	var decodedStrings = string(decodeBytes)
	fmt.Println(decodedStrings)

}