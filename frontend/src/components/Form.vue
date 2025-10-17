<script setup>
    import Input from "./Input.vue"
    import Button from  "./Button.vue"
    import { ref } from "vue"

    const states = {
        name: Symbol("name"),
        password: Symbol("password"),
        email: Symbol("email"),
    };

    const state = ref(states.name);

    function next() {
        switch (state.value) {
            case states.name:
                state.value = states.password;
                break;
            case states.password:
                state.value = states.email;
                break;
            case states.email:
                state.value = states.name;
                break;
        }
    }
</script>

<template>
    <transition name="fade" mode="out-in">
        <template v-if="state == states.name">
            <div class="form-fields">
                <Input label="Your name" placeholder="Your first name..."/>
                <Input label="Your last name" placeholder="Your last name..."/>
            </div>
        </template>

        <template v-else-if="state == states.password">
            <div class="form-fields">
                <p>Choose a good password UwU</p>
                <Input type="password" label="Your email" placeholder="Your secure email..."/>
            </div>
        </template>

        <template v-else-if="state == states.email">
            <div class="form-fields">
                <p>Please give us your email</p>
                <Input type="password" label="Your password" placeholder="Your secure password..."/>
            </div>
        </template>
    </transition>
    <Button :rows="3" :cols="12" width="200px" @click="next">Next</Button>
</template>

<style>
.form-fields {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 400px;
    gap: 1em;
}

.fade-enter-active, .fade-leave-active {
  transition: all 0.5s ease;
}

.fade-leave-to {
    transform: translateX(-200%);
}

.fade-enter-from {
    transform: translateX(200%); 
}
</style>