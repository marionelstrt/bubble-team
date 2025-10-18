<script setup>
  import { HeartIcon } from "@heroicons/vue/24/solid";
  import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

  const windowWidth = ref(window.innerWidth);
  const windowHeight = ref(window.innerHeight);

  const start = 15;
  const spacing = 60;
  const xCoords = computed(() => Array.from({
    length: windowWidth.value
  }, (_, i) => i)
    .filter(n => n % spacing === start));
  const yCoords = computed(() => Array.from({
    length: windowHeight.value
  }, (_, i) => i)
    .filter(n => n % spacing === start));

  function handleResize() {
    windowWidth.value = window.innerWidth;
    windowHeight.value = window.innerHeight;
  }

  onMounted(() => {
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
  })
</script>

<template>
  <div class="background">
    <template v-for="x in xCoords" :key="x">
      <template v-for="y in yCoords" :key="y">
        <div class="background-icon" :style="{
            marginLeft: x + 'px',
            marginTop: y + 'px',
          }">
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
}

.circle {
  background-color: #00000010;
  width: 24px;
  height: 24px;
  position: absolute;
  border-radius: 50%;
}
</style>
