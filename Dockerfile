# Etapa 1: Construcción de la aplicación Go
FROM golang:1.23-alpine AS build

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos de configuración de Go para instalar las dependencias
COPY go.mod go.sum ./

# Instalar las dependencias de Go
RUN go mod download

# Copiar el código fuente del backend
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Etapa 2: Contenedor final para ejecutar la aplicación
FROM alpine:latest

# Crear un directorio de trabajo en el contenedor
WORKDIR /app

# Copiar el ejecutable desde la etapa de construcción
COPY --from=build /app/main .

# Exponer el puerto en el que el servidor escuchará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
