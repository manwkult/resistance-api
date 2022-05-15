# Resistance API

API para identificar las coordenadas de una nave enemiga que esta enviando mensajes encriptados segun la distancia obtenida por 2 satelites de la resistencia.


## Algoritmo de Trilateración

### Trilateración

La trilateración es un método matemático para determinar las posiciones relativas de objetos usando la geometría de triángulos de forma análoga a la triangulación.

![](https://reglasyrelojes.files.wordpress.com/2015/04/trilateracion.png)


## CI / CD - Github Actions

### Github Actions para el CI / CD

![](https://miro.medium.com/max/3404/1*k99_arb0x9B7LI4I5hhCPw.png)

En el aplicativo se creo una pipeline de github actions que corre cuando se realiza un merge con la rama main.

El pipeline genera la imagen de docker y la sube al container registry de Dockerhub. Una vez realizado esto se reemplazan los datos de ambiente en los manifiestos .yaml de Kubernetes para posteriormente ser aplicados y desplegados en el cluster de Digital Ocean.

[Resistance API](http://143.244.202.236/swagger/index.html)

## Infraestructura sobre Kubernetes

[![Kubernetes](https://cdn.filestackcontent.com/RlUuJIVESsOwxSF6qcD9?auto=compress,format)](http://143.244.202.236/swagger/index.html)

### Cluster de Kubernetes en Digital Ocean

Se crea un cluster de kubernetes para el despliegue de los diferentes componentes de la API y se accede a la API a travez de un balancedor de carga que la misma nube provee

## Estructura de Directorio

```bash
.
├── .dockerignore
├── .github
│   └── workflows
│       └── build-and-deploy.yaml
├── .gitignore
├── .vscode
│   └── launch.json
├── Dockerfile
├── README.md
├── core
│   ├── api.go
│   ├── models.go
│   └── trilateration.go
├── deployment
│   └── deployment.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go
```

## Instrucciones para correr el proyecto localmente

Primeramente debemos de descargar el repositorio con git

```bash
git clone https://github.com/manwkult/resistance-api.git
```

_Golang 1.18 debe de estar instalado en tu maquina_

[Install Golang](https://go.dev/dl/)

Una vez descargado el repositorio debes de descargar las librerias del proyecto

```bash
go mod download
```

Y luego lo ejecutas

```bash
go run .
```

Estando el proyecto ejecutado ingresa al link

[Rhttp://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)


## Arquitectura

 - gin gonic
 - swagger

## Link del proyecto publicado

[Resistance API](http://143.244.202.236/swagger/index.html)
