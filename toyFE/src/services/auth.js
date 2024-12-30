
import { api } from "@/models/api";

export const auth_login_api = async (data) => {
    try {
        return await api("POST", "/login", data)
    } catch (error) {
        throw error
    }
}

export const auth_register_api = async (data) => {
    try {
        return await api("POST", "/register", data)
    } catch (error) {
        throw error
    }
}

export const auth_info_api = async () => {
    try {
        return await api("GET", "/auth/info")
    } catch (error) {
        throw error
    }
}
