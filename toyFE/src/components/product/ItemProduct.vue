<script setup>
import { useRouter } from "vue-router";
import { ref } from "vue";
const props = defineProps(["product"]);

const router = useRouter();
const goPay = () => {
  router.push(`/product/` + props.product.name + `/` + props.product.ID);
};
</script>

<template>
  <div class="ctn card" @click="goPay">
    <div class="img-ctn" v-if="product.image_urls">
      <img class="product-img" :src="$loadFile(product.image_urls[0].link)" />
    </div>

    <div class="info-product">
      <h2><b>{{ product.name }}</b></h2>
      <span>
        <b>{{ product.price }}</b>
        <i class="bi bi-currency-dollar"></i>
      </span>
    </div>
    <button>Buy now</button>
  </div>
</template>

<style scoped>
.ctn {
  display: flex;
  flex-flow: column;
  min-height: 350px;
  width: 300px;
  cursor: pointer;
  margin-bottom: 50px;
}

.card {
  padding: 0;
}

.img-ctn {
  width: 100%;
  height: 90%;
}

.product-img {
  width: 100%;
  height: 100%;
  border-radius: 8px;
}

img:hover {
  transform: scale(1.1);
  transition: transform 0.3s ease;
  border: 2px solid #fff;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);
}

.info-product {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 5px;
}
</style>
