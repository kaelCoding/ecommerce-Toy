<script setup>
import ItemProduct from "@/components/product/ItemProduct.vue";
import { onBeforeMount, ref} from "vue";
import { get_products_api } from "@/services/product";
import CreateProduct from "@/components/product/CreateProduct.vue";
import { get_auth_user } from "@/stores/auth";

const products = ref([]);

onBeforeMount(async () => {
  await getProducts()
  console.log(get_auth_user.value)
})

const getProducts = async () => {
  await get_products_api().then((res) => {
    products.value = res
  })
}

const showPopupCreate = ref(false);

const openPopupCreate = () => {
  showPopupCreate.value = true
}

const closePopupCreate = () => {
  showPopupCreate.value = false
}

const createProduct = async (product) => {
    products.value.unshift(product)
    await getProducts()
    closePopupCreate()
}
</script>

<template>
  <div class="main">
    <div class="main-product">
      <div class="btn-container">
        <button v-if="get_auth_user.admin" style="margin-left: auto;" @click="openPopupCreate">Create post</button>
      </div>

      <div class="ctn-products">
        <ItemProduct
          v-for="(product, index) of products" 
          :key="product.ID"
          :product="product" 
        </ItemProduct>
      </div>

      <CreateProduct
        v-if="showPopupCreate"
        @close="closePopupCreate"
        @createProduct="createProduct"
      />
    </div>

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
.main {
  flex-wrap: wrap;
}

.main-product {
  height: min-content;
  width: 100%;
  height: min-content;
  padding: 0 200px;
}

.btn-container {
  margin-top: 24px;
  display: flex;
  margin-bottom: 24px;
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
  gap: 24px;
  width: 100%;
  flex-wrap: wrap;
  margin-bottom: 24px;
  justify-content: space-between;
}

.footer {
  display: flex;
  min-height: 250px;
  align-items: center;
  justify-content: center;
  background-color: var(--c-black);
  width: 100%;
  color: var(--c-white);
}

.block {
  display: flex;
  flex-flow: column;
  gap: 10px;
  width: 30%;
}
</style>
