import { api } from "@/models/api";

export const add_category_api = async (data) => {
    try {
        return await api("POST", "/category", data)
    } catch (error) {
        throw error
    }
}

export const get_categoryID_api = async () => {
    try {
        return await api("GET", `/category/${id}`)
    } catch (error) {
        throw error
    }
}

export const update_category_api = async (id, data) => {
    try {
        return await api("PUT", `/category/${id}`, data)
    } catch (error) {
        throw error
    }
}

export const delete_product_api = async (id) => {
    try {
        return await api("DELETE", `/category/${id}`)
    } catch (error) {
        throw error
    }
}