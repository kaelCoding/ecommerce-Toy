
import { apiFormData } from "@/models/api";

export const files_upload_api = async (data) => {
    try {
        return await apiFormData("POST", "/upload", data)
    } catch (error) {
        throw error
    }
}
