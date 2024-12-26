<script setup>
</script>

<template>
  <div class="nav-ctn">
    <RouterLink to="/"><h1>Toni Toku</h1></RouterLink>
    <ul>
      <li>Home</li>
      <li>About Us</li>
      <li>Gallery</li>
      <li>Products</li>
      <li>Blog</li>
      <li>Contact</li>
    </ul>
  </div>
</template>

<style scoped>
.nav-ctn {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 200px;
}

ul {
  display: flex;
  list-style: none;
  gap: 15px;
}

.tab-text {
  width: 25%;
  text-align: center;
  font-size: 18px;
}
</style>
