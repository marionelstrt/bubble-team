<script setup>
defineProps({
    rows: Number,
    cols: Number,
    width: String,
});
</script>

<template>
    <div class="button" :style="{ width: width }">
        <div class="bubbles">
            <template v-for="n in rows" :key="n">
                <div class="bubble-row" :data-n="n">
                    <template v-for="m in cols" :key="m">
                        <div class="bubble"></div>
                    </template>
                </div>
            </template>
        </div>
        <button>
            <p class="button-label">
                <slot></slot>
            </p>
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
    justify-content: center;
    align-items: center;

    box-shadow: inset 6px 4px 20px 0px #b04bc187, inset -4px -8px 20px 0px #ca93e475;
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

.button:hover {
    transform: scale(1.1);
}

.button:active {
    transform: scale(0.9);
}

.button-label {
    position: fixed;
    z-index: 999;
}

.bubble-row {
    position: absolute;
    display: flex;
    gap: 5px;
    --rotate: rotate(-20deg);

    --bubble-x-sep: 10px;
    --bubble-y-sep: 20px;
    --bubble-y-margin: 70px;
}

.bubble-row[data-n="1"] {
    transform: var(--rotate) translateY(var(--bubble-y-margin));
}
.bubble-row[data-n="2"] {
    transform:
        var(--rotate)
        translateX(calc(var(--bubble-x-sep) * 1))
        translateY(calc(var(--bubble-y-margin) + var(--bubble-y-sep) * 1));
}
.bubble-row[data-n="3"] {
    transform:
        var(--rotate)
        translateX(calc(var(--bubble-x-sep) * 2))
        translateY(calc(var(--bubble-y-margin) + var(--bubble-y-sep) * 2));
}

.bubble {
    --bubble-size: 16px;
    --bubble-color: rgba(88, 50, 50, 0.103);

    border-radius: 50%;
    width: var(--bubble-size);
    height: var(--bubble-size);
    background-color: var(--bubble-color);
}
</style>