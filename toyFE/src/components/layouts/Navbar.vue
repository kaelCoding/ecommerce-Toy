<script setup>
import { RouterLink } from 'vue-router';
import { get_auth_user } from '@/stores/auth';
import { logout } from '@/stores/auth';

</script>

<template>
  <div class="nav-ctn">
    <RouterLink to="/"><h1>Toni Toku</h1></RouterLink>
    <div class="text-white info-ctn" v-if="get_auth_user">
      <h2>{{ get_auth_user.username }}</h2>
      <b @click="logout">Logout</b>
    </div>
    <div v-else style="display: flex; gap: 10px;">
      <RouterLink to="login">Login</RouterLink>
      <RouterLink to="register">Register</RouterLink>
    </div>
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

.info-ctn {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 12px;
}
</style>
