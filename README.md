# Resistance API



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
