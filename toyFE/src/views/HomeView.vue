<script setup>
import Card from "@/components/product/Card.vue"
import { ref } from "vue";
// import { products } from "@/stores";

</script>

<template>
  <div class="main-ctn">
    <div class="main-product">
      <div class="desc-ctn">
        <div class="desc-text">
          <h2>Cartoons And Anime Series</h2>
          <p>Stream Anytime On Animado</p>
          <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugit amet est
          aliquam, debitis atque odio neque doloribus ea commodi illum quam, cupiditate
          ex, dolorem fuga.</p>
          <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugit amet est
          aliquam, debitis atque odio neque doloribus ea commodi illum quam, cupiditate
          ex, dolorem fuga.</p>
        </div>

        <img src="/public/anh1.png" alt="">
      </div>

      <div class="ctn-products">
        <Card v-for="(product, index) of products" :product="product"/>
      </div>
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
