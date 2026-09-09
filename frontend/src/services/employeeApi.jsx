import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
});

// Attach Bearer token from localStorage to every request
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('emp_auth_token');
    if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
});

// Redirect to login on 401 Unauthorized
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('emp_auth_token');
            localStorage.removeItem('emp_auth_user');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export const createEmployee = async(employee) => {
    const response = await api.post("/employees", employee)
    return response.data;
}

export const getEmployees = async() => {
    const response = await api.get("/employees")
    return response.data;
}

export const getEmployeeById = async(id) => {
    const response = await api.get(`/employees/${id}`)
    return response.data;
}

export const searchEmployees = async(query) => {
    const response = await api.get(`/employees/search`, { params: { q : query, }, })
    return response.data;
}

export const updateEmployee = async(id, employee) => {
    const response = await api.put(`/employees/${id}`, employee)
    return response.data;
}

export const deleteEmployee = async(id) => {
    const response = await api.delete(`/employees/${id}`)
    return response.data;
}