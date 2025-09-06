const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

const height = canvas.height;
const width = canvas.width;

const animationLoop = () => {
    setTimeout(() => requestAnimationFrame(animationLoop), 1000/10);
}

requestAnimationFrame(animationLoop)