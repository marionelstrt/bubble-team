<script setup>
import Button from './Button.vue'
import Input from './Input.vue'
import Title from './Title.vue'
import icon from '../assets/boba.svg'
</script>

<template>
  <div>
    <div class="title">
      <img :src="icon" />
      <Title>Bubble</Title>
    </div>
    <Input label="Your name" placeholder="Your first name..."/>
    <Input label="Your last name" placeholder="Your last name..."/>
    <Button :rows="3" :cols="12" width="200px">Next</Button>
  </div>
</template>

<style scoped>
div {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  gap: 3em;
}

.title {
  flex-direction: row;
}

.title img {
  width: 80px;
}
</style>
