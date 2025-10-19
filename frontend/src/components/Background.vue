<script setup>
import { HeartIcon } from "@heroicons/vue/24/solid";
import { ref, computed, onMounted, onBeforeUnmount } from "vue";

const windowWidth = ref(window.innerWidth);
const windowHeight = ref(window.innerHeight);

const start = 15;
const spacing = 60;
const xCoords = computed(() =>
  Array.from(
    {
      length: windowWidth.value,
    },
    (_, i) => i,
  ).filter((n) => n % spacing === start),
);
const yCoords = computed(() =>
  Array.from(
    {
      length: windowHeight.value,
    },
    (_, i) => i,
  ).filter((n) => n % spacing === start),
);

function handleResize() {
  windowWidth.value = window.innerWidth;
  windowHeight.value = window.innerHeight;
}

function handleMouseMove(ev) {
  const mouse = { x: ev.clientX, y: ev.clientY };
  const points = xCoords.value.flatMap(x => yCoords.value.map(y => ({ x, y })));

  const maxDist = 80;
  const movementRatio = 0.3;

  points.forEach((point, i) => {
    const dist = Math.hypot(point.x-mouse.x, point.y-mouse.y);
    if (dist < maxDist) {
      const xDist = mouse.x-point.x;
      const yDist = mouse.y-point.y;

      document.querySelectorAll(".background-icon")[i].style.transform =
        `translateX(${xDist*movementRatio}px) translateY(${yDist*movementRatio}px) rotate(-20deg)`;
    } else {
      document.querySelectorAll(".background-icon")[i].style.transform = "";
    }
  });
}

onMounted(() => {
  window.addEventListener("resize", handleResize);
  window.addEventListener("mousemove", handleMouseMove);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", handleResize);
  window.removeEventListener("mousemove", handleMouseMove);
});
</script>

<template>
  <div class="background">
    <template v-for="x in xCoords" :key="x">
      <template v-for="y in yCoords" :key="y">
        <div
          class="background-icon"
          :style="{
            marginLeft: x + 'px',
            marginTop: y + 'px',
          }"
        >
          <template v-if="Math.random() > 0.5">
            <HeartIcon />
          </template>
          <template v-else>
            <div class="circle" />
          </template>
        </div>
      </template>
    </template>
  </div>
</template>

<style scoped>
.background {
  top: 0;
  left: 0;
  z-index: -1;
  position: absolute;
}

.background-icon {
  color: #00000010;
  width: 32px;
  height: 32px;
  position: absolute;
  transform: rotate(-20deg);
  transition: transform 0.2s linear;
}

.circle {
  background-color: #00000010;
  width: 24px;
  height: 24px;
  position: absolute;
  border-radius: 50%;
}
</style>
