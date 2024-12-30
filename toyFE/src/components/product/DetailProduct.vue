<script setup>
import { useRouter } from 'vue-router';
const props = defineProps(["product"])
const emits = defineEmits(["close"])

const close = () => {
    emits("close")
}

const router = useRouter()
const goPay = () => {
    router.push(`/product/` + props.product.name)
}
</script>

<template>
    <div class="overlay" @click="close">
        <div class="container-popup" @click.stop="">
            <div class="img-ctn" v-if="product.image_urls">
                <div v-for="(file, index) of product.image_urls" :key="file.link">
                    <img :src="$loadFile(file.link)">
                </div>
            </div>
            <div class="detail-product">
                <h1>{{ product.name }}</h1>
                <h3>{{ product.description }}</h3>
                <div class="ctn-price">
                    <span>{{ product.price }}</span>
                    <button @click="goPay">Buy now</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
* {
    color: black;
}

.overlay {
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background-color: #00000080;
    z-index: 20;
    display: flex;
    align-items: center;
    justify-content: center;
}

.container-popup {
    display: flex;
    flex-flow: column;
    background-color: white;
    min-width: 400px;
    border-radius: 8px;
    overflow-y: auto;
}

.detail-product {
    width: 100%;
    height: 100%;
    padding: 12px;
}

.ctn-price {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

img {
    width: 100%;
}

button {
    color: white;
}
</style>
