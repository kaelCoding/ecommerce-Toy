<script setup>
import Navbar from './components/layouts/Navbar.vue';
</script>

<template>
  <div class="app">
    <nav>
      <Navbar/>
    </nav>

    <main>
      <RouterView/>
    </main>

    <a href="https://www.facebook.com/huynh.o.thanh.tung" class="chat-ctn">
      <i class="bi bi-messenger"></i>
      <span class="messLink">Chat</span>
    </a>
  </div>
</template>

<style scoped>
.chat-ctn {
  position: fixed;
  display: flex;
  align-items: center;
  background-color: rgb(10, 124, 255);
  width: max-content;
  height: 46px;
  border-radius: 24px;
  padding: 12px;
  bottom: 50px;
  right: 5px;
  z-index: 3;
}

.bi {
  color: white;
  margin-right: 8px;
  font-size: 20px;
}

.messLink {
  font-size: 17px;
  font-weight: 600;
}
</style>
