<script></script>

<template>
  <h1>
    <slot></slot>
  </h1>
</template>

<style>
h1 {
  background-clip: text;
  color: transparent;
  background-image: var(--title);
  background-position: top left;
  background-size: 300% 100%;
  transition: background-position 1s ease-in-out;
  font-size: var(--font-huge);
  margin: 0;
}

h1:hover {
  background-position: top left 100%;
}
</style>
