<script setup>
import Detailproduct from './Detailproduct.vue';
import { ref } from 'vue';
const props = defineProps(["product"])

const showPopup = ref(false)

const openPopup = () => {
  showPopup.value = true
}

const closePopup = () => {
  showPopup.value = false
}

</script>

<template>
  <Detailproduct :product="product" v-if="showPopup" @close="closePopup"/>
  <div class="ctn" @click="openPopup">
    <img :src="product.img">
    <h2>{{ product.name }}</h2>
  </div>
</template>

<style scoped>
.ctn {
  display: flex;
  flex-flow: column;
  height: 300px;
  width: 290px;
  cursor: pointer;
}

img {
  width: 100%;
  height: 80%;
}

img:hover {
  transform: scale(1.1); 
  transition: transform 0.3s ease;
  border: 2px solid #fff;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);
}
</style>
