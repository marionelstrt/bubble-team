<script setup>
import Input from "./Input.vue";
import Button from "./Button.vue";
import MonsterBoba from "../assets/monster.svg";
import TaroBoba from "../assets/boba.svg";
import BrownSugarBoba from "../assets/sugar.svg";
import MatchaBoba from "../assets/matcha.svg";
import { ChevronRightIcon } from "@heroicons/vue/24/solid";
import { useRoute, useRouter } from "vue-router";
import { ref, onMounted, watch } from "vue";
import rockyou from "../assets/rockyou.txt?raw";

const states = {
  bubble: "bubble",
  name: "name",
  password: "password",
  email: "email",
};
const defaultState = states.bubble;

const route = useRoute();
const router = useRouter();
const state = ref(route.query.state || defaultState);

watch(route, (newRoute) => {
  state.value = newRoute.query.state || defaultState;
});

const bobas = [
  { name: "monster", desc: "soirée énergétique", icon: MonsterBoba },
  { name: "matcha", desc: "selon la vibe", icon: MatchaBoba },
  { name: "taro", desc: "claquage régulier", icon: TaroBoba },
  { name: "sugar", desc: "mariage", icon: BrownSugarBoba },
];
const selectedBoba = ref(route.query.boba || null);

function handleBobaClick(boba) {
  selectedBoba.value = boba;
}

const errors = ref({});

const name = defineModel("name");
const lastName = defineModel("lastName");
const password = defineModel("password");
const email = defineModel("email");

onMounted(() => {
  name.value = route.query.name;
  lastName.value = route.query.lastName;
  password.value = route.query.password;
  email.value = route.query.email;
});

function commonPassword() {
  const badpasswords = rockyou.split("\n");
  return badpasswords.includes(password.value);
}

function badPasswordLength() {
  return password.value?.length <= 6;
}

function validEmail() {
  return email.value.match(
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
  );
}

function checkErrors() {
  errors.value = {};

  switch (state.value) {
    case states.name:
      if (!name.value?.trim()) errors.value.name = "This must be filled.";
      if (!lastName.value?.trim())
        errors.value.lastName = "This must be filled.";
      break;
    case states.password:
      if (!password.value?.trim()) {
        errors.value.password = "This must be filled.";
      } else if (commonPassword()) {
        errors.value.password = "This password is too common.";
      } else if (badPasswordLength()) {
        errors.value.password = "This must be at least 6 characters.";
      }
      break;
    case states.email:
      if (!email.value?.trim()) {
        errors.value.email = "This must be filled.";
      } else if (!validEmail()) {
        errors.value.email = "This email is invalid";
      }
      break;
  }

  return !!Object.keys(errors.value).length;
}
function next() {
  if (checkErrors() || selectedBoba.value == null) return;

  let query = {};

  switch (state.value) {
    case states.bubble:
      state.value = states.name;
      query = { boba: selectedBoba.value };
      break;
    case states.name:
      state.value = states.password;
      query = { name: name.value, lastName: lastName.value };
      break;
    case states.password:
      state.value = states.email;
      query = { password: password.value };
      break;
    case states.email:
      state.value = states.bubble;
      query = { email: email.value };
      break;
  }

  fetch("http://localhost:8081/account/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      boba: selectedBoba.value,
      name: name.value,
      lastName: lastName.value,
      password: password.value,
      email: email.value,
    }),
  });
  router.push({
    query: { ...route.query, state: state.value, ...query },
  });
}
</script>

<template>
  <transition name="fade" mode="out-in">
    <template v-if="state == states.bubble">
      <div class="form-fields">
        <div class="image-grid">
          <template v-for="boba in bobas">
            <div
              @click="handleBobaClick(boba.name)"
              :class="[
                'boba-tile',
                selectedBoba == boba.name ? 'boba-tile-selected' : null,
              ]"
            >
              <img :src="boba.icon" :alt="'Boba ' + boba.name" />
              <p class="boba-label">{{ boba.name }}</p>
              <p class="boba-desc">{{ boba.desc }}</p>
            </div>
          </template>
        </div>
      </div>
    </template>

    <template v-else-if="state == states.name">
      <div class="form-fields">
        <Input
          v-model="name"
          label="Your name"
          placeholder="Your first name..."
          :error="errors.name"
        />
        <Input
          v-model="lastName"
          label="Your last name"
          placeholder="Your last name..."
          :error="errors.lastName"
        />
      </div>
    </template>

    <template v-else-if="state == states.password">
      <div class="form-fields">
        <p class="info">
          Please choose a password for your account. You may also register a
          passkey instead. Common passwords are insecure and thus not allowed.
        </p>
        <Input
          v-model="password"
          type="password"
          label="Your password"
          placeholder="Your secure password..."
          :error="errors.password"
        />
      </div>
    </template>

    <template v-else-if="state == states.email">
      <div class="form-fields">
        <p class="info">
          In order to easily recover your account in case of loss and to prevent
          spam, we need to your email address. An email will be sent to it to
          verify its authenticity.
        </p>
        <Input
          v-model="email"
          type="email"
          label="Your email address"
          placeholder="Your email address..."
          :error="errors.email"
        />
      </div>
    </template>
  </transition>
  <Button
    :rows="3"
    :cols="12"
    width="200px"
    :disabled="state == states.bubble && selectedBoba == null"
    @click="next"
  >
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
  width: 100%;
  height: 100%;
}

.boba-tile {
  background-color: var(--pink-lavender);
  border-radius: 15%;
  height: 180px;
  aspect-ratio: 1/1;
  display: flex;
  transition:
    border-color 0.2s ease-in,
    transform 0.1s ease-in-out;
  justify-content: center;
  border: 3px solid transparent;
  user-select: none;
}

.boba-tile:hover {
  transform: scale(1.1);
}

.boba-tile:active {
  transform: scale(0.9);
}

.boba-tile-selected {
  border: 3px solid white;
}

.boba-tile-selected:after {
  position: absolute;
  width: 20px;
  height: 20px;
  left: 0;
  background-color: var(--mauve-5);
  content: url("data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2024%2024%22%20stroke-width%3D%221.5%22%20stroke%3D%22white%22%20class%3D%22size-6%22%3E%3Cpath%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20d%3D%22m4.5%2012.75%206%206%209-13.5%22%20%2F%3E%3Cpath%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20d%3D%22m4.5%2012.75%206%206%209-13.5%22%20%2F%3E%3C%2Fsvg%3E");
  border-radius: 50%;
  margin: 10px;
  padding: 4px;
}

.boba-tile img {
  width: 40%;
  margin: auto;
}

.boba-tile {
  position: relative;
}

.boba-label {
  position: absolute;
  justify-self: anchor-center;
  align-self: flex-start;
  background: var(--title);
  background-clip: text;
  color: transparent;
  font-weight: 500;
  font-size: var(--font-small);
  margin: 7px;
}

.boba-desc {
  position: absolute;
  align-self: flex-end;
  color: #00000080;
  font-weight: 400;
  font-size: var(--font-very-small);
  margin: 5px;
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
