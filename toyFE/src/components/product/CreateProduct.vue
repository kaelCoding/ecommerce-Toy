<script setup>
import { useRouter } from 'vue-router';
import { ref } from 'vue';
import { add_product_api } from '@/services/product';
import { files_upload_api } from '@/services/file';

const emits = defineEmits(["close", "createProduct"])
const close = () => {
    emits("close")
}

const router = useRouter()

const product = ref({
    name: "",
    description: "",
    price: "",
    category_name: "",
    image_urls: [],
})

const filesInp = ref(null)

const handleUploadFiles = async () => {
    const files = filesInp.value.files
    const formData = new FormData()

    for (const file of files) {
        formData.append("image", file)
    }

    try {
        await files_upload_api(formData).then(res => {
            product.value.image_urls = res.map(item => {
                return {
                    link: item.Link
                }
            })
        })
    } catch (error) {
        console.log(error)
    }
}

const handleClickUploadFile = () => {
    filesInp.value.click()
}

const createProduct = async () => {
    try {
        await add_product_api(product.value).then(res => {
        emits("createProduct", res)
    })
    } catch (error) {
      console.log(error)  
    }
}
</script>

<template>
    <div class="overlay" @click="close">
        <div class="container-popup" @click.stop="">
            <form class="form card" @submit.prevent="createProduct">
                <h1 style="text-align: center;">CREATE PRODUCT</h1>
                <label>Name</label>
                <input type="text" v-model="product.name" placeholder="Enter name">
                <label>Description</label>
                <input type="text" v-model="product.description" placeholder="Enter description">
                <label>Price</label>
                <input type="text" v-model="product.price" placeholder="Enter price">
                <label>Category</label>
                <input type="text" v-model="product.category_name" placeholder="Enter category">

                <hr>
                <input style="display: none;" type="file" ref="filesInp" multiple @input="handleUploadFiles">
                <button style="width: max-content" type="button" @click="handleClickUploadFile">
                    <span class="font-content">Upload files</span>
                </button>

                <div class="img-ctn" v-if="product.image_urls.length > 0">
                    <div v-for="(file, index) of product.image_urls" :key="file.link">
                        <!-- <i class="bi bi-x-circle-fill" @click="removeFile(index)"></i> -->
                        <img class="product-img" :src="$loadFile(file.link)">
                    </div>
                </div>
                <hr>

                <button type="submit">Create</button>
            </form>
        </div>
    </div>
</template>

<style scoped>
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
    background-color: var(--c-white);
    width: 400px;
    border-radius: 8px;
}

.form {
    margin: 0;
}

.card {
    width: 100%;
}

.product-ctn {
    display: flex;
    flex-flow: column;
    padding: 12px;
    width: 100%;
    height: 100%;
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

.img-ctn {
    display: flex;
    gap: 5px;
    max-width: 400px;
    overflow-x: auto;
}

.product-img {
    width: 216px;
    height: 216px;
}
</style>
