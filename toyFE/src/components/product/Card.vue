<script setup>

</script>

<template>
  <div>
    <form class="card form" @submit.prevent="login">
      <label>Name</label>
      <input type="text" v-model="dataLogin.username" placeholder="Name">

      <label>Password</label>
      <input type="password" v-model="dataLogin.password" placeholder="Password">

      <button type="submit" class="btn btn-primary">Register</button>
    </form>
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
