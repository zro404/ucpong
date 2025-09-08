const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

console.log("Canvas initialized");

const height = canvas.height;
const width = canvas.width;

let data = {
  p1: { x: 5, y: (height - 100) / 2, width: 25, height: 100 },
  p2: { x: width - 30, y: (height - 100) / 2, width: 25, height: 100 },
  ball: { x: (width - 15) / 2, y: (height - 15) / 2, radius: 15 },
};

const animationLoop = () => {
  ctx.clearRect(0, 0, width, height);

  ctx.fillStyle = "#0000FF";
  ctx.fillRect(data.p1.x, data.p1.y, data.p1.width, data.p1.height);
  console.log(data.p1.x, data.p1.y, data.p1.width, data.p1.height);

  ctx.fillStyle = "#FF0000";
  ctx.fillRect(data.p2.x, data.p2.y, data.p2.width, data.p2.height);

  ctx.fillStyle = "#FFFFFF";
  ctx.beginPath();
  ctx.arc(data.ball.x, data.ball.y, data.ball.radius, 0, Math.PI * 2);
  ctx.fill();

  setTimeout(() => requestAnimationFrame(animationLoop), 1000 / 10);
};

requestAnimationFrame(animationLoop);
