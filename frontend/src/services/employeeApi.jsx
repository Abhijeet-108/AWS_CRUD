import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
    withCredentials: true, // send session cookie with every request
});

// MongoDB returns `_id` — normalize it to `id` so the rest of the app works uniformly
const normalizeEmployee = (emp) => ({ ...emp, id: emp._id ?? emp.id });

export const createEmployee = async(employee) => {
    const response = await api.post("/employees", employee)
    return response.data;
}

export const getEmployees = async() => {
    const response = await api.get("/employees");
    const raw = response.data;
    const list = Array.isArray(raw) ? raw : (raw.employees ?? raw.data ?? []);
    return list.map(normalizeEmployee);
}

export const getEmployeeById = async(id) => {
    const response = await api.get(`/employees/${id}`)
    return response.data;
}

export const searchEmployees = async(query) => {
    const response = await api.get(`/employees/search`, { params: { q: query } });
    const raw = response.data;
    const list = Array.isArray(raw) ? raw : (raw.employees ?? raw.data ?? []);
    return list.map(normalizeEmployee);
}

export const updateEmployee = async(id, employee) => {
    const response = await api.put(`/employees/${id}`, employee)
    return response.data;
}

export const deleteEmployee = async(id) => {
    const response = await api.delete(`/employees/${id}`)
    return response.data;
}