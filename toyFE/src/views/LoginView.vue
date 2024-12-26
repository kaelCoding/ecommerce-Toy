<script setup>
import { ref } from "vue"
import { save_token_local } from "@/stores/auth";
import { useRouter } from "vue-router"
import { auth_login_api } from "@/services/auth";

const router = useRouter()

const dataLogin = ref({
    username: "",
    password: "",
})

const login = async () => {
    try {
        const data = await auth_login_api(dataLogin.value)
        console.log(data)
        save_token_local(data.token)

        router.push("/")
    } catch (error) {
        console.log('on login error ', error)
    }
}
</script>

<template>
    <div class="main">
        <div class="main-center ctn">
            <form class="card form" @submit.prevent="login">
                <label>Name</label>
                <input type="text" v-model="dataLogin.username" placeholder="Name">

                <label>Password</label>
                <input type="password" v-model="dataLogin.password" placeholder="Password">

                <button type="submit" class="btn btn-primary">Register</button>
            </form>

            <div class="help-block">
                Don't have account ?
                <RouterLink to="/register">Register</RouterLink>
            </div>
        </div>
    </div>
</template>

<style scoped>
.help-block {
    margin-top: 24px;
}
</style>