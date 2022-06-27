package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"math/rand"
	"encoding/base64"
	"strings"
	"time"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/md5"
)

const hash_size = 69 // largo fijo del hash
const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/" // caracteres base 64
var hashes = []string{"hash69", "md5", "sha1", "sha256"}

func timeTrack(start time.Time, name string) { // funcion de calculo de tiempo de ejecucion
    elapsed := time.Since(start)
    fmt.Printf("%s tardo %s", name, elapsed)
}

func hash69 (input string) string {
	output := ""
	text64 := base64.StdEncoding.EncodeToString([]byte(input)) // encodeamos el texto en base64
	if len(text64) == 0 { // si no existe nada, devolvemos un string "vacio" de meme
		return ""
	} else if len(text64) > hash_size { // si existe y es mayor al largo de 69 aplicaremos operaciones para hashear
		output_nomod := "" // string auxiliar para output
		mod := len(text64) % hash_size // primero obtenemos el modulo
		for i := 0; i < (len(text64) - mod); i++ { // iteramos sobre el string
			if i < mod { // los primeros i terminos que correspondan a la cantidad del modulo obtenido se operaran para reducir el largo a una potencia de 69
				// esta operacion hace un calculo de la siguiente forma:
				// primero obtiene la posicion dentro del arreglo de base 64 del caracter en la i-esima posicion
				// luego obtiene la posicion del caracter en el arreglo de base 64 de la posicion largo de cadena - modulo + i
				// se suman las posiciones y se saca el modulo 64 en absoluto
				output_nomod += string(b64[int(math.Abs(float64((strings.Index(b64, string(text64[i])) + strings.Index(b64, string(text64[len(text64) - mod + i])))%64)))])
			} else { // el resto de caracteres simplemente se agregaran
				output_nomod += string(text64[i])
			}
		}
		// en este punto tenemos una cadena de largo 69 a la N, siendo N >= 1
		size := len(output_nomod) / hash_size // obtenemos la potencia de 69 de la cadena
		for i := 0; i < hash_size; i++ { // iteramos sobre el largo final es decir 69
			index := strings.Index(b64, string(output_nomod[i])) // obtenemos la posicion del caracter del ouput auxiliar, respecto al arreglo de base 64
			for j := 1; j < size; j++ { // sumaremos todas las posiciones que correspondan a la potencia de la posicion i*69 + j, con el fin de obtener todas las posiciones cada 69 posiciones
				index += strings.Index(b64, string(output_nomod[j*hash_size + i]))  // se van almacenando en la variable index
			}
			index %= 64 // sacamos modulo finalmente para obtener caracter
			output += string(b64[index]) // se agregar el caracter al mod correspondiente del arreglo base64
		}
		// finalmente se obtiene la cadena de largo 69
	} else { // en caso de que la cadena sea inferior a 69 y mayor a 0
		// 2 variables auxiliares para iterar sobre la cadena y ver las vueltas que llevamos
		i := 0
		vueltas := 0
		for len(output) < hash_size { // iteramos hasta que la cadena tenga el largo de 69
			if i == len(text64) { // si llegamos al final de la cadena, volvemos a empezar y aumenta la vuelta
				i = 0
				vueltas++
			}
			if vueltas == 0 { // si estamos en la vuelta 0
				if text64[i] == '=' { // un padding de b64 indica que simplemente se sacara el caracter de la posicion i mod 69 del arreglo de b64
					output += string(b64[i%64])
				} else { // en otro caso se agrega el caracter de forma normal
					output += string(text64[i])
				}
			} else { // si no estamos en la vuelta 0
				if text64[i] == '=' { // esta regla se mantiente
					output += string(b64[i%64])
				} else { // esta regla cambia, ahora se agrega el caracter que deberia seguir segun la cadena b64 respecto a la posicion del caracter en el arreglo de base64
					index := strings.Index(b64, string(text64[i]))
					output += string(b64[(index+1)%64])
				}
			}
			i++
		}
	}
	return output
}

func calculateTime(size int, words []string){ // funcion que calcula el tiempo de ejecucion de cada algoritmo
	index := []int{} // arreglo de indices de palabras a utilizar
	for i := 0; i < size; i++ {
		index = append(index, rand.Intn(len(words)))
	}
	for i := 0; i < 4; i++ { // iteramos sobre los 4 algoritmos
		start := time.Now() // iniciamos el tiempo de ejecucion del algoritmo i-esimo
		for j := 0; j < len(index); j++ { // iteramos sobre las palabras
			if i == 0 { // si estamos en i = 0, se aplica el algoritmo de hash69
				hash69(words[index[j]])
			} else if i == 1 { // si estamos en i = 1, se aplica el algoritmo de md5
				md5.Sum([]byte(words[index[j]]))
			} else if i == 2 { // si estamos en i = 2, se aplica el algoritmo de sha1
				sha1.Sum([]byte(words[index[j]]))
			} else if i == 3 { // si estamos en i = 3, se aplica el algoritmo de sha256
				sha256.Sum256([]byte(words[index[j]]))
			}
		}
		elapsed := time.Since(start) // obtenemos el tiempo de ejecucion
		sec, _ := time.ParseDuration(elapsed.String()) // se obtiene el tiempo en segundos
		fmt.Printf("Tiempo de %s: %.10f s\n", hashes[i], sec.Seconds()) // se imprime el tiempo de ejecucion
	}
}

func speedtest () { // funcion que realiza una prueba de velocidad
	fmt.Println("Ingrese la ruta del archivo:")
	ruta := ""
	fmt.Scanln(&ruta)
	lectura, err := os.Open(ruta)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(lectura)
	fileScanner.Split(bufio.ScanLines) // leemos linea por linea
	words := []string{} // arreglo de palabras
	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())  // agregamos la palabra a la lista de palabras
	}
	lectura.Close()
	comando := ""
	for comando != "5" { // menu de pruebas
		fmt.Println("(1) Para hashear con 1 string.")
		fmt.Println("(2) Para hashear con 10 strings.")
		fmt.Println("(3) Para hashear con 20 strings.")
		fmt.Println("(4) Para hashear con 50 strings.")
		fmt.Println("(5) Para salir.")
		fmt.Scanln(&comando)
		if comando == "1" { // 1 palabra
			calculateTime(1, words)
		} else if comando == "2" { // 10 palabras
			calculateTime(10, words)
		} else if comando == "3" { // 20 palabras
			calculateTime(20, words)
		} else if comando == "4" { // 50 palabras
			calculateTime(50, words)
		} else if comando != "5" {
			fmt.Println("Comando no reconocido, intente de nuevo.")
		}
	}
}

func entropia (input string) float64 { // Entropia de un string largo*log2(64)
	return float64(len(input)) * math.Log2(64)
}

func menu() {
	fmt.Println("Bienvenido al sistema de hash.")
	fmt.Println("(1) Para hashear un string.")
	fmt.Println("(2) Para hashear desde archivo.")
	fmt.Println("(3) Para obtener entropia.")
	fmt.Println("(4) Para comparar algoritmos.")
	fmt.Println("(5) Para salir.")
	comando := ""
	fmt.Scanln(&comando)
	if comando == "1" {
		fmt.Println("Ingrese el string a hashear:")
		str := ""
		fmt.Scanln(&str)
		fmt.Println(hash69(str))
	} else if comando == "2" {
		fmt.Println("Ingrese la ruta del archivo:")
		ruta := ""
		fmt.Scanln(&ruta)
		lectura, err := os.Open(ruta)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileScanner := bufio.NewScanner(lectura)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			fmt.Println(hash69(fileScanner.Text()))
		}
		lectura.Close()
	} else if comando == "3" {
		fmt.Println("Ingrese el string a calcular entropia:")
		str := ""
		fmt.Scanln(&str)
		hstr := hash69(str)
		fmt.Println(entropia(hstr))
	} else if comando == "4" {
		speedtest()
	} else if comando == "5" {
		return
	} else {
		fmt.Println("Comando no reconocido, intente de nuevo.")
	}
	menu()
}

func main() {
	menu()
}