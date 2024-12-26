<script setup>
import { ref } from "vue"
import { useRouter } from "vue-router"
import { auth_register_api } from "@/services/auth";

const router = useRouter()

const dataRegister = ref({
    username: "",
    email: "",
    password: "",
})

const register = async () => {
    try {
        await auth_register_api(dataRegister.value)
        
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
    <div class="">
        <div class="ctn">
            <label>Name</label>
            <input type="text" v-model="dataRegister.username" placeholder="Name">

            <label>Email</label>
            <input type="email" v-model="dataRegister.email" placeholder="Email address">

            <label>Password</label>
            <input type="password" v-model="dataRegister.password" placeholder="Password">

            <button @click="register">Register</button>

            <div class="help-block">
                Already have an account ?
                <RouterLink to="/login">Login</RouterLink>
            </div>
        </div>
    </div>
</template>

<style scoped>

.help-block {
    margin-top: 24px;
}
</style>