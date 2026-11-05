# Data Structures Visualizer

An interactive data structures visualizer built in Go and compiled to WebAssembly (Wasm). 

This project was created as a tool for the Data Structures and Algorithms II course (Algoritmos y Programación 2).
The goal is to implement dynamic structures from scratch, without relying on third-party libraries or heavy frontend frameworks, delegating all memory management and processing logic directly to Go.

*(Leer versión en español más abajo).*

### Architecture

- **Core (Go):** All the hard logic, node management, and mathematical calculations of the data structures live in the `internal/` directory.
- **Bridge (WebAssembly):** Through the `syscall/js` package in `cmd/wasm/`, Go exposes its functions directly to the browser's global environment.
- **Frontend (Vanilla JS + HTML5 Canvas):** JavaScript acts solely as a lightweight rendering layer. It listens to DOM events, calls the in-memory `.wasm` binary, receives a data array, and manually draws the elements using geometric coordinates on the Canvas API. Zero dependencies.

### Project Status

- [x] Queue (Enqueue / Dequeue / Front).
- [ ] AVL Tree (In development).
- [ ] Stack (Pending).
- [ ] Binary Search Tree (Pending).
- [ ] Linked List (Pending).
- [ ] Double Linked List (Pending).
- [ ] Double Circular Linked List - Ring Buffer (Pending).
- [ ] Sentinel Linked List (Pending).
- [ ] Hash Table (Pending).
- [ ] Binary Heap (Pending).

### Live Demo

The project is compiled and statically hosted on GitHub Pages:
**[View Visualizer](https://pedroantinolfi.github.io/visualizer/)**

### Compilation & Local Usage (Windows)

If you want to clone the repository and run it locally, you need to have Go installed.

1. Compile the Go code into a WebAssembly binary:
```powershell
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o web/main.wasm ./cmd/wasm

2. Start the HTTP server (configured to disable caching for easier development):
go run server.go

3. Open http://localhost:8080 in your web browser.

---

# Data Structures Visualizer (Español)

Un visualizador interactivo de estructuras de datos construido en Go y compilado a WebAssembly (Wasm). 

Este proyecto nace como una herramienta para ayudar a estudiantes de programación a comprender más facilmente algunas estructuras de datos.
El objetivo es implementar estructuras dinámicas desde cero, sin depender de librerías de terceros o frameworks pesados en el frontend, delegando toda la lógica de memoria y procesamiento a Go.

### Arquitectura

- **Core (Go):** Toda la lógica dura, manejo de nodos y cálculos matemáticos de las estructuras vive en la carpeta `internal/`.
- **Puente (WebAssembly):** A través del paquete `syscall/js` en `cmd/wasm/`, Go expone sus funciones directamente al entorno global del navegador.
- **Frontend (Vanilla JS + HTML5 Canvas):** JavaScript funciona únicamente como una capa tonta de dibujado. Escucha los clicks, llama al binario `.wasm` en memoria, recibe un array de datos y los dibuja manualmente usando coordenadas geométricas en el Canvas. Cero dependencias.

### Estado del Proyecto

- [x] Cola mediante Lista Enlazada (Enqueue / Dequeue / Frente).
- [ ] Árbol AVL (En desarrollo: rotaciones, factor de balance y renderizado jerárquico).
- [ ] Pila (Pendiente).
- [ ] Árbol ABB (Pendiente).
- [ ] Lista enlazada (Pendiente).
- [ ] Lista doblemente enlazada (Pendiente).
- [ ] Lista circular doblemente enlazada (Pendiente).
- [ ] Lista enlazada Sentinela (Pendiente).
- [ ] Tabla de Hash (Pendiente).
- [ ] Montículos binarios (Pendiente).

### Live Demo

El proyecto está compilado y hosteado de forma estática en GitHub Pages:
**[Ver Visualizador](https://pedroantinolfi.github.io/visualizer/)**

### Compilación y Uso Local (Windows)

Si querés clonar el repositorio y probarlo localmente, necesitás tener Go instalado. 

1. Compilar el código de Go a binario WebAssembly:
```powershell
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o web/main.wasm ./cmd/wasm
```

2. Levantar el servidor HTTP (que desactiva la caché para facilitar el desarrollo):
```bash
go run server.go
```

3. Abrir `http://localhost:8080` en el navegador.
