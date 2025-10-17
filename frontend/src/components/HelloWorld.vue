<script setup>
import Button from './Button.vue'
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <Button :rows="3" :cols="12" width="200px">Next</Button>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
