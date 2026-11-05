// 1. WebAssembly Initialization
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    console.log("WebAssembly engine successfully connected.");
});

// 2. DOM Elements
const mainMenu = document.getElementById("mainMenu");
const visualizerWorkspace = document.getElementById("visualizerWorkspace");
const workspaceTitle = document.getElementById("workspaceTitle");
const btnBackToMenu = document.getElementById("btnBackToMenu");

const btnMenuQueue = document.getElementById("btnMenuQueue");
const controlsQueue = document.getElementById("controlsQueue");
const controlsAVL = document.getElementById("controlsAVL");

const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

// Queue Elements
const inputQueue = document.getElementById("inputQueue");
const btnEnqueue = document.getElementById("btnEnqueue");
const btnDequeue = document.getElementById("btnDequeue");

let currentMode = "";

// 3. Navigation Events
btnMenuQueue.addEventListener("click", () => {
    currentMode = "queue";
    workspaceTitle.innerText = "Visualizer - List Queue";
    mainMenu.classList.add("hidden");
    visualizerWorkspace.classList.remove("hidden");
    controlsQueue.classList.remove("hidden");
    drawQueue();
});

btnBackToMenu.addEventListener("click", () => {
    currentMode = "";
    mainMenu.classList.remove("hidden");
    visualizerWorkspace.classList.add("hidden");
    controlsQueue.classList.add("hidden");
    controlsAVL.classList.add("hidden");
    ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear canvas on exit
});

// 4. Queue Events
btnEnqueue.addEventListener("click", () => {
    const val = inputQueue.value.trim();
    if (val !== "") {
        window.wasmQueueEnqueue(val);
        inputQueue.value = "";
        drawQueue();
    }
});

btnDequeue.addEventListener("click", () => {
    window.wasmQueueDequeue();
    drawQueue();
});

// 5. Rendering Logic
function drawQueue() {
    if (currentMode !== "queue") return; // Prevent drawing if not in queue mode
    
    ctx.clearRect(0, 0, canvas.width, canvas.height); 
    const values = window.wasmQueueGetValues(); 
    
    if (!values || values.length === 0) {
        ctx.fillStyle = "#888";
        ctx.font = "20px Arial";
        ctx.textAlign = "center";
        ctx.fillText("Queue is empty (Front = nil)", canvas.width / 2, canvas.height / 2);
        return;
    }

    const boxWidth = 60;
    const boxHeight = 60;
    const startX = 50;
    const startY = 70;
    const gap = 30; 

    ctx.font = "20px Arial";
    ctx.textAlign = "center";
    ctx.textBaseline = "middle";

    values.forEach((val, index) => {
        const x = startX + index * (boxWidth + gap);
        const y = startY;

        ctx.fillStyle = "#4CAF50";
        ctx.fillRect(x, y, boxWidth, boxHeight);
        ctx.strokeStyle = "#fff";
        ctx.lineWidth = 2;
        ctx.strokeRect(x, y, boxWidth, boxHeight);

        ctx.fillStyle = "#fff";
        ctx.fillText(val, x + boxWidth / 2, y + boxHeight / 2);

        if (index < values.length - 1) {
            ctx.beginPath();
            ctx.moveTo(x + boxWidth, y + boxHeight / 2);
            ctx.lineTo(x + boxWidth + gap, y + boxHeight / 2);
            ctx.strokeStyle = "#aaa";
            ctx.stroke();
            
            ctx.beginPath();
            ctx.moveTo(x + boxWidth + gap, y + boxHeight / 2);
            ctx.lineTo(x + boxWidth + gap - 8, y + boxHeight / 2 - 5);
            ctx.lineTo(x + boxWidth + gap - 8, y + boxHeight / 2 + 5);
            ctx.fillStyle = "#aaa";
            ctx.fill();
        }
    });
}