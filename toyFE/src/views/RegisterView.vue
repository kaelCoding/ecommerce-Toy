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
        await auth_register_api(dataRegister.value).then(res => {
            router.push("/login")
        })
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
    <div class="main">
        <div class="ctn">
            <form class="form card">
                <h1 style="text-align: center;">REGISTER</h1>
                <label>Name</label>
                <input type="text" v-model="dataRegister.username" placeholder="Name">

                <label>Email</label>
                <input type="email" v-model="dataRegister.email" placeholder="Email address">

                <label>Password</label>
                <input type="password" v-model="dataRegister.password" placeholder="Password">

                <button type="button" @click="register">Register</button>
            </form>
            <div class="help-block">
                Already have an account ?
                <RouterLink to="/login"><b>Login</b></RouterLink>
            </div>
        </div>
    </div>
</template>

<style scoped>

.help-block {
    margin-top: 24px;
}

.help-block a {
    color: var(--c-text);
}
</style>