<script setup>
import Input from "./Input.vue";
import Button from "./Button.vue";
import { ChevronRightIcon } from "@heroicons/vue/24/solid";
import { useRoute, useRouter } from "vue-router";
import { ref, watch } from "vue";

const states = {
  bubble: "bubble", 
  name: "name",
  password: "password",
  email: "email",
};

const route = useRoute();
const router = useRouter();
const state = ref(route.query.state || states.name);

watch(route, newRoute => {
  state.value = newRoute.query.state || states.bubble;
});

function next() {
  switch (state.value) {
    case states.bubble:
      state.value = states.name; 
      break;
    case states.name:
      state.value = states.password;
      break;
    case states.password:
      state.value = states.email;
      break;
    case states.email:
      state.value = states.bubble;
      break;
  }

  router.push({
    query: { state: state.value },
  });
}

function handleBobaClick(target)
{
  state.value = states.name;
  router.push({
    query: { state: states.name, flavor: target},
  });
}
</script>

<template>
  <transition name="fade" mode="out-in">
    <template v-if="state == states.bubble">
      <div class="form-fields"> 
        <div class="image-grid">
            <img src="/boba.svg" alt="Boba 4" @click="handleBobaClick('monster')" />
            <img src="/boba.svg" alt="Boba 4" @click="handleBobaClick('taro')" />
            <img src="/boba.svg" alt="Boba 4" @click="handleBobaClick('mango')" />
            <img src="/boba.svg" alt="Boba 4" @click="handleBobaClick('yay')" />
        </div>
      </div>
    </template>
  
    <template v-else-if="state == states.name">
      <div class="form-fields">
        <Input label="Your name" placeholder="Your first name..." />
        <Input label="Your last name" placeholder="Your last name..." />
      </div>
    </template>

    <template v-else-if="state == states.password">
      <div class="form-fields">
        <p class="info">In order to easily recover your account in case of loss and to prevent spam, we need to your email address. An email will be sent to it to verify its authenticity.</p>
        <Input
          type="password"
          label="Your email address"
          placeholder="Your email address..."
        />
      </div>
    </template>

    <template v-else-if="state == states.email">
      <div class="form-fields">
        <p class="info">Please choose a password for your account. You may also register a passkey instead. Common passwords are insecure and thus not allowed.</p>
        <Input
          type="password"
          label="Your password"
          placeholder="Your secure password..."
        />
      </div>
    </template>
  </transition>
  <Button :rows="3" :cols="12" width="200px" @click="next">
    Next
    <template #icon>
      <ChevronRightIcon />
    </template>
  </Button>
</template>

<style>
.form-fields {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 400px;
  gap: 1em;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr); 
  grid-template-rows: repeat(2, 1fr); 
  gap: 30px; 
  width: 200px; 
  height: 200px;
  margin: 0 auto; 
}
.info {
  text-wrap: wrap;
  color: white;
  max-width: 100%;
  align-self: center;
  font-size: var(--font-small);
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.5s ease;
}

.fade-leave-to {
  transform: translateX(-200%);
}

.fade-enter-from {
  transform: translateX(200%);
}
</style>
