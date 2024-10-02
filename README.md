# Prueba Golang
<!-- 
La prueba está dividida en varias secciones con distintos enfoques. Aquí te guiaré en cómo abordar cada una de ellas utilizando Golang y buenas prácticas de desarrollo. -->

## 1. Problemas de diseño – 1 punto

**Pregunta:** Solución a un cuello de botella en la base de datos bajo alta demanda de transacciones.

**Solución propuesta:**

- **Colas de Mensajes (Message Queue):** Implementar un sistema de colas como RabbitMQ o AWS SQS para encolar las transacciones y procesarlas de manera asíncrona, desacoplando el sistema del cuello de botella que es la base de datos.
- **Bulk Inserts:** En lugar de procesar una transacción a la vez, agrupar múltiples transacciones en un solo insert masivo (batch insert). Esto reduce el número de operaciones en la base de datos.
- **Optimización de índices:** Revisa los índices existentes en la base de datos, asegurando que las consultas críticas estén adecuadamente indexadas para mejorar el rendimiento.
- **Caching:** Implementar un caché (como Redis) para almacenar temporalmente transacciones frecuentes, reduciendo la carga en la base de datos.

---

## 2. Encripta tu mensaje – 3.5 puntos

**Enunciado:** Desarrollar una función que encripte las vocales del mensaje precedidas por una clave.

**Código en Golang:**

```go
package main

import (
	"fmt"
	"strings"
)

func x(key, message string) string {
	if message == "" {
		return ""
	}
	if key == "" {
		key = "DCJ"
	}

	vowels := "aeiouAEIOU"
	var result strings.Builder

	for _, char := range message {
		if strings.ContainsRune(vowels, char) {
			result.WriteString(key)
		}
		result.WriteRune(char)
	}
	return result.String()
}

func main() {
	key := "dcj"
	message := "I love prOgrAmming!"
	encrypted := x(key, message)
	fmt.Println(encrypted)
}
```

---

## 3. Suma a cero – 3.5 puntos

**Enunciado:** Dado un arreglo de enteros, eliminar todos los nodos consecutivos que la suma sea cero.

**Código en Golang:**

```go
package main

import (
	"fmt"
)

func fn(val []int) []int {
	i := 0

	for i < len(val) {
		acc := 0

		for j := i; j < len(val); j++ {
			acc += val[j]

			if acc == 0 {
				val = append(val[:i], val[j+1:]...)
				break
			}
		}

		if acc != 0 {
			i++
		}
	}

	return val
}

func main() {
	val := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := fn(val)
	fmt.Println(result)
}
```


## 4. Aprendizaje – 1 punto

***La arquitectura mostrada utiliza un API Gateway que dirige las solicitudes de clientes hacia microservicios independientes (Catalog, Shopping Cart, Discount, Ordering), cada uno con su propia base de datos. El Message Broker facilita la comunicación asíncrona entre estos microservicios, permitiendo que trabajen de forma desacoplada y escalable.***


### Le agregaria CQRS

##### ¿Por qué implementar CQRS?
Implementar CQRS permite separar las operaciones de lectura y escritura, optimizando el rendimiento al evitar la competencia por recursos entre ambas. Usar una base de datos de lectura, como Elasticsearch o una base de datos NoSQL, facilita la creación de vistas especializadas para consultas rápidas y complejas, uniendo múltiples entidades en un solo registro. Esto mejora la escalabilidad, ya que puedes ajustar las capacidades de lectura y escritura de manera independiente. Las bases de datos de lectura se pueden escalar horizontalmente mediante la adición de nodos, permitiendo gestionar grandes volúmenes de datos y consultas sin afectar la base de datos de escritura, que permanece optimizada para transacciones.



## 5. Demostrando tus hallazgos – 1 punto

***Pregunta:*** Explicar microservicios en términos no técnicos.

***Explicación***:

Imaginemos que tenemos una fábrica con diferentes secciones: cada sección hace una tarea específica, como ensamblar piezas, pintar o empaquetar. Si una sección se daña, no afecta a las demás, y además, podemos ampliar o mejorar una sección sin cambiar las demás. Los microservicios funcionan de manera similar: dividen una aplicación grande en pequeñas piezas independientes. Cada pieza es responsable de una tarea específica. Esto nos permite manejar más pedidos sin que todo se colapse, haciendo que el sistema sea más rápido y fácil de escalar.