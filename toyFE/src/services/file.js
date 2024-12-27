
import { api } from "@/models/api";

export const files_upload_api = async (data) => {
    try {
        return await api("POST", "/upload", data)
    } catch (error) {
        throw error
    }
}
