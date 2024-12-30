<script setup>
import { get_productID_api } from '@/services/product';
import { computed, onBeforeMount, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()

const product = ref({})

const idProduct = computed(() => {
    return route.params.idProduct
})

const getProduct = async () => {
    await get_productID_api(idProduct.value).then(res => {
        console.log(res)
        product.value = res
    })
} 

onBeforeMount(async () => {
    await getProduct()
})
</script>

<template>
    <div class="main">
        <div class="product-ctn">
            <div class="img-ctn" v-if="product.image_urls.length > 0" v-for="(file, index) of product.image_urls" :key="file.link">
                <img class="product-img" :src="$loadFile(file.link)">
            </div>
        </div>

        <div class="pay-ctn">
            <div>
                <h1>{{ product.name }}</h1>
                <span>Mô tả: {{ product.description }}</span>
            </div>

            <div>
                <h2>Giá:{{ product.price }}đ</h2>
                <span>Tình trạng: <span style="color: red;">còn hàng</span></span>
            </div>
            <form class="form card">
                <h1 style="text-align: center;">Thông tin cá nhân</h1>
                <label>Nhập tên</label>
                <input type="email" placeholder="Họ và tên">

                <label>Nhập số điện thoại</label>
                <input type="password" placeholder="Số điện thoại">

                <label>Địa chỉ</label>
                <input type="password" placeholder="Địa chỉ">
            </form>

            <span>Gọi đặt mua: 0829721097 (miễn phí 8:30 - 21:30).</span>

            <button style="margin-top: auto;">Mua ngay</button>
        </div>
    </div>
</template>

<style scoped>
.main {
    padding: 12px;
    gap: 12px;
}

.product-ctn {
    display: flex;
    flex-flow: column;
    overflow-y: auto;
    width: 600px;
    height: 500px;
}

.img-ctn {
    display: flex;
    flex-flow: column;
    gap: 5px;
}

.product-img {
    width: 600px;
    height: 500px;
    border-radius: 12px;
}

.pay-ctn {
    width: 500px;
    height: 500px;
    display: flex;
    flex-flow: column;
    gap: 5px;
}

</style>