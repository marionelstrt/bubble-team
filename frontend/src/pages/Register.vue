<script setup>
import Title from "../components/Title.vue";
import Form from "../components/Form.vue";
import icon from "../assets/boba.svg";
</script>

<template>
  <main>
    <div class="title">
      <img :src="icon" />
      <Title>Bubble</Title>
    </div>
    <Form />
  </main>
</template>

<style scoped>
main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  gap: 3em;

  padding: 40px;
  border-radius: 10%;
  overflow: hidden;
  border: 5px dashed var(--mauve-5);
  background-color: var(--mauve-2);
}

.title {
  display: flex;
  gap: 3em;
  align-items: center;
  flex-direction: row;
}

.title img {
  width: 80px;
}
</style>
