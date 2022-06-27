# Laboratorio 3 - Hash 🔐

Autor: Cristian Villavicencio  😎

## Dependencias

- [Golang](https://go.dev)
- [Docker](https://www.docker.com)

## ¿Como ejecutar?

Podemos ejecutar directamente con Golang, para ello debemos utilizar:

```sh
    go run ./src/hash.go
```

Para ejecutar con docker debemos utilizar:

```sh
    docker run --rm -it $(docker build -q .)
```

Este ultimo comando hace un build y ejecuta inmediatamente el contenedor de forma interactiva, por lo que al terminar el contenedor desaparecera, mas no la imagen del mismo.

## Comparativa de hash

Tabla comparativa de algoritmos, cabe destacar que Time(i) se refiere al tiempo total por cantidad de palabras segun algoritmo, ademas este tiempo de la tabla fue probado con las mismas palabras para los distintos algoritmos, por lo que al probar de forma aleatoria podrian ir variando estos tiempos. La implementacion de los algortimos SHA1, SHA256 y MD5 provienen de [crypto](https://pkg.go.dev/crypto) de la STL de Golang.

| Metrica | Hash69 | SHA1 | SHA256 | MD5 |
| :-:| :-: | :-:| :-: | :-: |
| Hash(12345) | MTIzNDUHNUJ0OEVHNUJ0OEVHNUJ0OEVHNUJ0OEVHNUJ0OEVHNUJ0OEVHNUJ0OEVHNUJ0O | 8cb2237d0679ca88db6464eac60da96345513964 | 5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5 | 827ccb0eea8a706c4c34a16891f84e7b |
| Entropia | 414 | 240 | 384 | 192 |
| Time(1) [s] | 0.0000210830 [s] | 0.0000030410 [s] | 0.0000017500 [s] | 0.0000016670 [s] |
| Time(10) [s] | 0.0002765000 [s] | 0.0000065000 [s] | 0.0000041250 [s] | 0.0000042920 [s] |
| Time(20) [s] | 0.0012737500 [s] | 0.0000121250 [s] | 0.0000051660 [s] |0.0000052500 [s] |
| Time(50) [s] | 0.0015496660 [s] | 0.0000235420 [s] | 0.0000090830 [s] | 0.0000091670 [s] |

Notamos que el algoritmo de Hash69 es notoriamente mas lento que el resto de algoritmos, en definitiva esto se debe a la programacion del mismo, pues puede ser optimizado aun de forma de que los tiempos mejoren, la mayoria de los algoritmos de hash, involucran operacion a bajo nivel, lo cual para el algoritmo Hash69 no aplica. Un ejemplo podria ser operaciones a nivel matematica mas compleja o nivel de bits, con lo cual se podria obtener una clara mejora en los tiempos de ejecucion.

Otro detalle a destacar es el nivel de base y entropia, si comparamos nuestro algoritmo tiene un mayor nivel de entropia, el mas similar es SHA256, que comparte un largo similar, pero un detalle es la base del algoritmo, en este caso al usar base64 tenemos un contraste a las bases del resto de algoritmos que poseen base16, esto puede reflejar tambien en tiempos de espera mayor al calculo, pues se podria optimizar teniendo menos candidatos posibles a la hora de elegir caracteres.

Se concluye que al tener un algoritmo que involucre una base mas grande y una cadena de texto mas grande podria ser mas complejo en terminos de tiempo y calculos, aunque se podria llegar a un nivel de eficiencia con ayuda de operaciones matematicas y a nivel binario con el fin de optimizar y lograr un tiempo de espera mas cercano respecto a los algoritmos comparados.