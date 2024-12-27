<script setup>
import { useRouter } from 'vue-router';
import { ref } from 'vue';
// const props = defineProps(["product"])
const emits = defineEmits(["close"])

const close = () => {
    emits("close")
}

const router = useRouter()

const product = ref({
    name: "",
    description: "",
    price: "",
    category: "",
    image_urls: "",
})

const filesInp = ref(null)


const handleUploadFiles = async () => {
    const files = filesInp.value.files

    const formData = new FormData()

    for (const file of files) {
        formData.append("files", file)
    }

    try {
        await files_upload_api(formData).then(res => {
            post.value.files = res.map(item => {
                return {
                    link: item.id
                }
            })
            console.log(post.value.files)
        })
    } catch (error) {
        console.log(error)
    }
}

const handleClickUploadFile = () => {
    filesInp.value.click()
}
</script>

<template>
    <div class="overlay" @click="close">
        <div class="container-popup" @click.stop="">
            <form class="product-ctn">
                <label>Name</label>
                <input type="text" v-model="product.name" placeholder="Enter name">
                <label>Description</label>
                <input type="text" v-model="product.description" placeholder="Enter description">
                <label>Price</label>
                <input type="text" v-model="product.price" placeholder="Enter price">
                <label>Category</label>
                <input type="text" v-model="product.category" placeholder="Enter category">

                <hr>
                <input style="display: none;" type="file" ref="filesInp" multiple @input="handleUploadFiles">
                <button style="width: max-content" type="button" @click="handleClickUploadFile">
                    <span class="font-content">Upload files</span>
                </button>
            </form>
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
    background-color: white;
    min-width: 400px;
    max-width: 700px;
    border-radius: 8px;
    overflow-y: auto;
}

.product-ctn {
    display: flex;
    flex-flow: column;
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
</style>
