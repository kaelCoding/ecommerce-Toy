import { get_auth_user } from "@/stores/auth";
import { watch } from "vue";

export default class {
    #router;

    constructor(router) {
        this.#router = router;
        this.checkPath();
        this.watchAuth();
    }

    checkPath(){
        this.#router.beforeEach((to, from) => {
            const path = to.path;

            if (!get_auth_user.value){
                if (path == "/" ||
                    path == `/product/:nameProduct/:idProduct`
                ){
                    this.#router.push("/login");
                }
            }
            else {
                if (
                    path == "/login" ||
                    path == "/register"
                ){
                    this.#router.push("/");
                }
            }
        })
    }

    watchAuth() {
        watch(() => get_auth_user.value, () => {
            if (!get_auth_user.value){
                this.#router.push("/login")
            }
            else {
                this.#router.push("/");
            }
        })
    }
}