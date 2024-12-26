<!-- <script setup>
import { computed, onBeforeMount, ref } from 'vue';
// import { products } from '@/stores';
import { useRoute } from 'vue-router';

const route = useRoute()

const nameProduct = computed(() => {
    return route.params.nameProduct
})

const product = computed(() => {
    return products.find(product => product.name == nameProduct.value)
})
</script>

<template>
    <div class="container">
        <div class="sample-ctn">
            <img :src="product.img" alt="">
        </div>
        <div class="pay-ctn">
            <h1>{{ product.name }}</h1>
            <span>{{ product.desc }}</span>
            <span>{{ product.price }}</span>

            <form class="form">
                <input type="email" placeholder="Họ và tên">

                <input type="password" placeholder="Số điện thoại">

                <input type="password" placeholder="Số lượng">
            
                <input type="password" placeholder="Địa chỉ">
            </form>

            <span>Gọi đặt mua: 0829721097 (miễn phí 8:30 - 21:30).</span>

            <button>Mua ngay</button>
        </div>
    </div>
</template>

<style scoped>
.form {
    display: flex;
    flex-flow: column;
    gap: 10px;
}
.form input {
    width: 100%;
}

.form button {
    width: 100%;
}

.form label {
    font-size: 17px;
    font-weight: 500;
    display: block;
    margin-bottom: 4px;
}

button {
    background-color: white;
    margin-top: auto;
}

.container {
    display: flex;
    gap: 10px;
    padding: 24px 200px;
}

.sample-ctn {
    width: 50%;
    height: 80%;
}

img {
    width: 100%;
    height: 100%;
}

.pay-ctn {
    width: 50%;
    display: flex;
    flex-flow: column;
    gap: 5px;
}

</style> -->