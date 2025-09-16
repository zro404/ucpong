const p1Score = document.getElementById("p1-score");
const p2Score = document.getElementById("p2-score");

const modal = document.getElementById("modal");
const modalMessage = document.getElementById("modal-text");
const modalButton = document.getElementById("modal-button");

const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

const height = canvas.height;
const width = canvas.width;

var syncEnabled = true;

let data = {
  p1: { x: 0, y: height / 2, width: 25, height: 150, score: 0 },
  p2: { x: width - 25, y: height / 2, width: 25, height: 150, score: 0 },
  ball: { x: (width - 15) / 2, y: (height - 15) / 2, radius: 15 },
};

const animationLoop = () => {
  ctx.clearRect(0, 0, width, height);

  ctx.fillStyle = "#0000FF";
  ctx.fillRect(data.p1.x, data.p1.y - data.p1.height / 2, data.p1.width, data.p1.height);

  ctx.fillStyle = "#FF0000";
  ctx.fillRect(data.p2.x, data.p2.y - data.p2.height / 2, data.p2.width, data.p2.height);

  ctx.fillStyle = "#FFFFFF";
  ctx.beginPath();
  ctx.arc(data.ball.x, data.ball.y, data.ball.radius, 0, Math.PI * 2);
  ctx.fill();

  setTimeout(() => requestAnimationFrame(animationLoop), 1000 / 10);
};

requestAnimationFrame(animationLoop);

let ws = new WebSocket("ws://" + window.location.host + "/ws");
ws.onerror = (error) => {
  console.error("WebSocket error:", error);
};

ws.onclose = () => {
  console.log("WebSocket connection closed");
};

window.onbeforeunload = () => {
  ws.close();
};

ws.onopen = () => {
  console.log("WebSocket connection established");
  const pathParts = window.location.pathname.split("/");
  const roomCode =
    pathParts[pathParts.length - 1] || pathParts[pathParts.length - 2];
  ws.send(JSON.stringify({ roomCode }));

  document.addEventListener("keydown", (event) => {
      let action;
      switch (event.key) {
        case "ArrowUp":
          action = 2;
          break;
        case "ArrowDown":
          action = 1;
          break;
      }

    if (syncEnabled) {
      ws.send(JSON.stringify({ action }));
      syncEnabled = false;
    }
  });
};

ws.onmessage = (event) => {
  const parsedData = JSON.parse(event.data);
  const {type} = parsedData;

  console.log(type)

  switch(type) {
    case "disconnect":
      showModal("Opponent disconnected!", "NEW GAME", () => {
        ws.close();
        window.location.href = "/";
      });
      return;
    case "gameOver":
      const winner = parsedData.playerScore1 > parsedData.playerScore2 ? "Player 1" : "Player 2";
      showModal(`${winner} wins!`, "PLAY AGAIN", () => {
        ws.send(JSON.stringify({ action: 3 }));
      });
      break;
    case "inProgress":
      break;
    default:
      console.error("Unknown message type:", type);
  }

  const { playerPos1, playerPos2, playerScore1, playerScore2, ballX, ballY } = parsedData;
  data = {
    p1: { ...data.p1, y: playerPos1, score: playerScore1 },
    p2: { ...data.p2, y: playerPos2, score: playerScore2 },
    ball: {
      ...data.ball,
      x: ballX,
      y: ballY
    },
  };
  syncEnabled = true;

  p1Score.textContent = playerScore1;
  p2Score.textContent = playerScore2;
};

const showModal = (message, buttonText, handler) => {
  modalMessage.textContent = message;
  modalButton.textContent = buttonText;
  modal.classList.remove("hidden");

  modalButton.onclick = () => {
    modal.classList.add("hidden");
    handler();
  };
};