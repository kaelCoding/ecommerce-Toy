import { init_auth } from "./auth"

// chua dang nhap
export const init_store = async () => {
    await init_auth()
}

