<script setup>
import { HeartIcon } from "@heroicons/vue/24/solid";

defineProps({
  rows: Number,
  cols: Number,
  width: String,
  disabled: Boolean,
});
</script>

<template>
  <div
    :class="['button', disabled ? 'button-disabled' : null]"
    :style="{ width: width }"
  >
    <div class="bubbles">
      <template v-for="n in rows" :key="n">
        <div class="bubble-row" :data-n="n">
          <template v-for="m in cols" :key="m">
            <div class="bubble">
              <template v-if="Math.random() > 0.5">
                <div class="is-heart">
                  <HeartIcon />
                </div>
              </template>
              <template v-else>
                <div class="is-bubble"></div>
              </template>
            </div>
          </template>
        </div>
      </template>
    </div>
    <button :disabled="disabled">
      <p class="button-label">
        <slot></slot>
      </p>
      <div class="button-icon">
        <slot name="icon"></slot>
      </div>
    </button>
  </div>
</template>

<style scoped>
button {
  font-size: var(--font-medium);
  background-color: var(--button-background);
  color: white;
  padding: 20px;
  border: none;
  font-weight: 800;

  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;

  box-shadow:
    inset 6px 4px 20px 0px #b04bc187,
    inset -4px -8px 20px 0px #ca93e475;
}

.button {
  display: block;
  transition: all 0.1s ease-in-out;
  box-shadow: 0px 2px 10px 1px #00000030;
  border: 5px solid var(--button-border);
  border-radius: var(--button-border-radius);
  height: 100px;
  overflow: hidden;
  transform: scale(1);
}

.button-disabled {
  border: 5px solid rgb(142 122 149 / 67%);
  filter: brightness(0.5);
}

.button:hover:not(.button-disabled) {
  transform: scale(1.1);
}

.button:active:not(.button-disabled) {
  transform: scale(0.9);
}

.button-label {
  z-index: 999;
  flex-basis: 70%;
}

.button-icon {
  width: 40px;
  height: 40px;
  flex-basis: 20%;
  margin-top: -6px;
}

.bubble-row {
  position: absolute;
  display: flex;
  gap: 5px;

  --rotate: rotate(-20deg);
  --bubble-x-sep: 10px;
  --bubble-y-sep: 20px;
  --bubble-y-margin: 70px;

  transform: var(--rotate) translateX(calc(var(--bubble-x-sep) * var(--n)))
    translateY(calc(var(--bubble-y-margin) + var(--bubble-y-sep) * var(--n)));
}

.bubble-row[data-n="1"] {
  --n: 0;
}
.bubble-row[data-n="2"] {
  --n: 1;
}
.bubble-row[data-n="3"] {
  --n: 2;
}

.bubble {
  --bubble-color: rgba(88, 50, 50, 0.103);
}

.is-bubble {
  --bubble-size: 12px;
  border-radius: 50%;
  width: var(--bubble-size);
  height: var(--bubble-size);
  background-color: var(--bubble-color);
}

.is-heart {
  --bubble-size: 16px;
  width: var(--bubble-size);
  height: var(--bubble-size);
  color: var(--bubble-color);
}
</style>
