<script setup>
import Card from "@/components/product/Card.vue"
import { onBeforeMount, ref } from "vue";
import { get_products_api } from "@/services/product";
import CreateProduct from "@/components/product/CreateProduct.vue";

const products = ref([])

onBeforeMount(async () => {
  await getProducts()
})

const getProducts = async () => {
  await get_products_api().then(res => {
    products.value = res
    console.log(products.value)
  })
}

const showPopupCreate = ref(false)

const openPopupCreate = () => {
    showPopupCreate.value = true
}

const closePopupCreate = () => {
    showPopupCreate.value = false
}

</script>

<template>
  <div class="main-ctn">
    <div class="main-product">
      <button @click="openPopupCreate">Create post</button>

     

      <div class="ctn-products">
        <Card v-for="(product, index) of products" :product="product"/>
      </div>
    </div>

    <CreateProduct 
      v-if="showPopupCreate"
      @close="closePopupCreate"
    />
      

    <div class="footer">
      <div class="block">
        <h2>Thông tin liên hệ</h2>
        <span>46 Nguyễn Văn Tố, Quận Hoàn Kiếm, Hà Nội</span>
        <span>0829721097</span>
        <span>lostboyshp68@gmail.com</span>
      </div>

      <div class="block">
        <h2>Mua hàng</h2>
        <span>Trang chủ</span>
        <span>0829721097</span>
        <span>lostboyshp68@gmail.com</span>
      </div>

      <div class="block">
        <h2>Dịch vụ khách hàng</h2>
        <span>46 Nguyễn Văn Tố, Quận Hoàn Kiếm, Hà Nội</span>
        <span>0829721097</span>
        <span>lostboyshp68@gmail.com</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main-ctn {
  display: flex;
  flex-flow: column;
  width: 100%;
  height: 100%;
  overflow-y: auto;
  flex: 1;
}

.main-product {
  width: 100%;
  height: min-content;
  padding: 0 200px;
}

.desc-ctn {
  display: flex;
  height: 500px;
  align-items: center;
  justify-content: space-between;
}

.desc-text {
  width: 50%;
}

.ctn-products {
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  justify-content: space-between;
}

.footer {
  display: flex;
  min-height: 200px;
  background-color: #161616;;
  width: 100%;
  padding: 24px 200px;
}

.block {
  display: flex;
  flex-flow: column;
  gap: 10px;
  width: 30%;
}
</style>
