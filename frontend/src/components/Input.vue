<script setup>
    defineProps({
        label: String,
        placeholder: String,
        type: {
            type: String,
            default: "text",
        },
    });
</script>

<template>
    <div class="input-field">
        <label for="input">{{ label }}</label>
        <input name="input" :type="type" :placeholder="placeholder" />
    </div>
</template>

<style scoped>
.input-field {
    gap: 1em;
    transition: all 0.1s ease-in-out;
    display: flex;
    flex-direction: column;
    display: flex;
}

.input-field:active {
    transform: scale(0.9);
}

input {
    font-size: var(--font-big);
    background-color: white;
    border-radius: var(--input-border-radius);
    border-color: var(--input-border);
    box-shadow: 0px 2px 10px 1px #00000030;
    border-style: solid;
    border-width: 5px;
    color: var(--input-text-color);
    padding: 20px;
    font-weight: 400;
}

input:focus {
    outline: thick double var(--input-outline); 
}

label {
    font-size: 2em;
    font-weight: 800;
    align-self: flex-start;
    color :white; 
}
</style>