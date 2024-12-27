import { api } from "@/models/api";

export const add_product_api = async (data) => {
    try {
        return await api("POST", "/products", data)
    } catch (error) {
        throw error
    }
}

export const get_products_api = async () => {
    try {
        return await api("GET", "/products")
    } catch (error) {
        throw error
    }
}

export const get_productID_api = async (id) => {
    try {
        return await api("GET", `/products/${id}`)
    } catch (error) {
        throw error
    }
}

export const update_product_api = async (id, data) => {
    try {
        return await api("PUT", `/products/${id}`, data)
    } catch (error) {
        throw error
    }
}

export const delete_product_api = async (id) => {
    try {
        return await api("DELETE", `/products/${id}`)
    } catch (error) {
        throw error
    }
}