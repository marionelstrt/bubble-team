<script setup>
    defineProps({
        placeholder: String,
    });
</script>

<template>
    <input type="text" :placeholder="placeholder" />
</template>

<style scoped>
input {
    font-size: var(--font-big);
    background-color: var(--input-background);
    border-radius: var(--input-border-radius);
    border-color: var(--input-border);
    transition: all 0.1s ease-in-out;
    box-shadow: 0px 2px 10px 1px #00000030;
    border-style: solid;
    border-width: 5px;
    color: white;
    padding: 20px;
    font-weight: 600;
}

input:hover {
    transform: scale(1.1);
}
</style>