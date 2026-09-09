import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
    withCredentials: true,
});

// api.interceptors.response.use(
//     // (response) => response,
//     // (error) => {
//     //     if (error.response?.status === 401) {
//     //         window.location.href = "/login";
//     //     }

//     //     return Promise.reject(error);
//     // }

//     (response) => response,
//     (error) => {s
//         console.error(
//             "API Error:",
//             error.response?.status,
//             error.response?.data
//         );

//         return Promise.reject(error);
//     }
// );

export const getProfile = async () => {
    const response = await api.get("/profile");

    return response.data;
};

// export const logout = async () => {
//     const response = await api.post("/auth/logout");
//     return response.data;
// }

export default api;