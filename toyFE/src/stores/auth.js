import { computed, ref } from "vue";

export const token = ref("")

const auth_user = ref(null)

export const get_auth_user = computed(() => {
    return auth_user.value
})

export const save_token_local = (tk) => {
    token.value = "Bearer " + tk
    localStorage.setItem("TOKEN", token.value)
}

export const load_token_local = () => {
    const tokenLocal = localStorage.getItem("TOKEN")
    if (tokenLocal) {
        token.value = tokenLocal;
    }
}

export const init_auth = async () => {
    load_token_local();

    if(token.value){
        console.log("hello world!")
    }
}

export const logout = () => {
    token.value = null
    auth_user.value = null

    localStorage.clear()
}

// export const get_auth_info = async () => {
//     try {
//         const data = await au
//     } catch (error) {
//         console.log(error)
//     }
// }
