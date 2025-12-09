# crea el builder para levantar el backend en Go
FROM golang:alpine AS builder

#carpeta interna de alpine para la aplicacion
WORKDIR /app

# copiar backend
COPY go.mod ./
COPY main.go ./

# compilar
RUN CGO_ENABLED=0 go build -o backend-app main.go

# imagen para producción
FROM alpine:latest

# aradia, el mismo usuario del front que sera el administrador
RUN adduser -D -u 1000 aradia

# entorno de producción
WORKDIR /home/aradia

# traer el ejecutable de la app
COPY --from=builder /app/backend-app .

# permisos
RUN chown aradia:aradia backend-app

# aradia rules
USER aradia

# puerto
EXPOSE 8080

# ejecuución
CMD ["./backend-app"]