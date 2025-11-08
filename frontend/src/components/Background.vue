<script setup>
import { HeartIcon } from "@heroicons/vue/24/solid";
import { ref, computed, onMounted, onBeforeUnmount } from "vue";

const windowWidth = ref(window.innerWidth);
const windowHeight = ref(window.innerHeight);
const mouse = ref({ x: 0, y: 0 });

const points = ref([]);
const spacing = 50;

function initializePoints() {
  points.value = [];
  const width = window.innerWidth;
  const height = window.innerHeight;

  for (
    let y = 0;
    y < height - 32 /* icon width, so we don't overflow */;
    y += spacing
  ) {
    for (let x = 0; x < width - 32; x += spacing) {
      points.value.push({
        ox: 15+x,
        oy: 15+y,
        x: 15+x,
        y: 15+y,
        vx: 0,
        vy: 0,
        shape: Math.random() > 0.5 ? "heart" : "circle",
      });
    }
  }
}

function handleResize() {
  windowWidth.value = window.innerWidth;
  windowHeight.value = window.innerHeight;
  initializePoints();
}

function handleMouseMove(e) {
  mouse.value.x = e.clientX;
  mouse.value.y = e.clientY;
}

let animationFrameId = null;
function animate() {
  for (const p of points.value) {
    const dx = mouse.value.x - p.x;
    const dy = mouse.value.y - p.y;
    const dist = Math.sqrt(dx * dx + dy * dy);
    const force = Math.max(100 - dist, 0) / 250; // attraction strength
    const angle = Math.atan2(dy, dx);

    // mouse attraction
    p.vx += Math.cos(angle) * force * 0.5;
    p.vy += Math.sin(angle) * force * 0.5;

    // elastic pull back to origin
    p.vx += (p.ox - p.x) * 0.02;
    p.vy += (p.oy - p.y) * 0.02;

    // damping
    p.vx *= 0.9;
    p.vy *= 0.9;

    p.x += p.vx;
    p.y += p.vy;
  }
  animationFrameId = requestAnimationFrame(animate);
}

onMounted(() => {
  initializePoints();
  window.addEventListener("resize", handleResize);
  window.addEventListener("mousemove", handleMouseMove);
  animate();
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", handleResize);
  window.removeEventListener("mousemove", handleMouseMove);
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId);
  }
});
</script>

<template>
  <div class="background">
    <div
      v-for="(p, index) in points"
      :key="index"
      class="background-icon"
      :style="{
        transform: `translateX(${p.x - p.ox}px) translateY(${p.y - p.oy}px) rotate(-20deg)`,
        left: p.ox + 'px',
        top: p.oy + 'px',
      }"
    >
      <template v-if="p.shape === 'heart'">
        <HeartIcon />
      </template>
      <template v-else>
        <div class="circle" />
      </template>
    </div>
  </div>
</template>

<style scoped>
.background {
  top: 0;
  left: 0;
  z-index: -1;
  position: absolute;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

.background-icon {
  color: #00000010;
  width: 32px;
  height: 32px;
  position: absolute;
  transition: none;
}

.circle {
  background-color: #00000010;
  width: 24px;
  height: 24px;
  position: absolute;
  border-radius: 50%;
}
</style>
