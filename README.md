# 🛡️ Hackademy: Docker y K8s

![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)
![Security](https://img.shields.io/badge/Security-Hardened-green?style=for-the-badge)

**Proyecto Final - Arquitectura de Contenedores y Orquestación.**
Este repositorio contiene la implementación de una arquitectura (Frontend, Backend, Database) segura, migrada de un entorno local (Docker Compose) a producción (Kubernetes).

---

## 📋 Prerrequisitos
Para ejecutar este proyecto correctamente, asegúrate de tener instalado:

1. **Docker Engine & Docker Compose:** Incluido en Docker Desktop.
2. **Kubernetes Cluster:** Habilitar la opción "Enable Kubernetes" en Docker Desktop (o usar Minikube/Kind).
3. **kubectl:** Herramienta de línea de comandos para Kubernetes (necesaria para el Entregable 4).
4. **Git:** Para clonar el repositorio.

---

## 🛠️ Instalación
Abre tu terminal y clona este repositorio:

```bash
git clone [https://github.com/AradiaEtreshka/hackademy-DockerK8s.git]
cd hackademy-DockerK8s

---

## 🚀 Opción A: Ejecución Local (Docker Compose)
Opción para levantar el entorno en PC (Entorno de Desarrollo).

1. **Iniciar los servicios:**
docker-compose up -d

2. **Verificar funcionamiento**
docker-compose ps

3. **Acceder a la aplicación**
    Frontend (Web): http://localhost:8080
    Backend (API): http://localhost:8081

4. **Detener Entorno**
docker-compose down

---

## ☸️ Opción B: Despliegue en Kubernetes (Producción)
Despliegue utilizando Manifiestos (`/k8s`) .

1. **Aplicar configuración al cluster:**
   ```bash
   kubectl apply -f k8s/

2. **Verificar estado de los Pods**
kubectl get all -n hackademy-k8s

3. **Acceder a la aplicación**

    Frontend (Público): http://localhost:8080 (Servicio LoadBalancer expuesto).

    Backend (Privado): Configurado como ClusterIP para máxima seguridad. No es accesible directamente desde internet.
    
Dado que la API es privada se debe establecer un tunel administrativo temporal:

    kubectl port-forward -n hackademy-k8s svc/backend 8081:80
    Mantener la terminal abierta. Acceder a: http://localhost:8081 para ver el JSON de respuesta.

4. Terminar el despliegue.
    kubectl delete -f k8s/
