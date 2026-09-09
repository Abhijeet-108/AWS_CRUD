import { useEffect, useState } from "react";

import EmployeeForm from "../components/EmployeeForm";
import EmployeeList from "../components/EmployeeList";
import SearchBar from "../components/SearchBar";

import {
    createEmployee,
    deleteEmployee,
    getEmployees,
    searchEmployees,
    updateEmployee,
} from "../services/employeeApi";

import "./EmployeePage.css";

function EmployeePage() {
    const [employees, setEmployees] = useState([]);
    const [search, setSearch] = useState("");
    const [editingEmployee, setEditingEmployee] = useState(null);
    const [showForm, setShowForm] = useState(false);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const loadEmployees = async () => {
        try {
            setLoading(true);
            setError("");
            const data = await getEmployees();
            setEmployees(data);
        } catch (err) {
            console.error(err);
            setError("Failed to load employees.");
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => { loadEmployees(); }, []);

    const handleSearch = async (query) => {
        setSearch(query);
        if (!query.trim()) { loadEmployees(); return; }
        try {
            setLoading(true);
            setError("");
            const data = await searchEmployees(query);
            setEmployees(data);
        } catch (err) {
            console.error(err);
            setError("Search failed.");
        } finally {
            setLoading(false);
        }
    };

    const handleCreate = async (employee) => {
        try {
            setError("");
            await createEmployee(employee);
            setShowForm(false);
            await loadEmployees();
        } catch (err) {
            console.error(err);
            setError(err.response?.data?.error || "Failed to create employee.");
        }
    };

    const handleUpdate = async (employee) => {
        try {
            setError("");
            await updateEmployee(editingEmployee.id, employee);
            setEditingEmployee(null);
            await loadEmployees();
        } catch (err) {
            console.error(err);
            setError(err.response?.data?.error || "Failed to update employee.");
        }
    };

    const handleDelete = async (id) => {
        if (!window.confirm("Delete this employee?")) return;
        try {
            setError("");
            await deleteEmployee(id);
            await loadEmployees();
        } catch (err) {
            console.error(err);
            setError(err.response?.data?.error || "Failed to delete employee.");
        }
    };

    const handleEdit = (employee) => {
        setEditingEmployee(employee);
        setShowForm(false);
        window.scrollTo({ top: 0, behavior: "smooth" });
    };

    const isFormVisible = showForm || !!editingEmployee;

    return (
        <div className="employee-page">

            {/* ── Page header ── */}
            <div className="page-header">
                <div className="page-heading">
                    <h1 className="page-title">Employees</h1>
                    <p className="page-subtitle">
                        {employees.length} team member{employees.length !== 1 ? "s" : ""}
                    </p>
                </div>
            </div>

            {/* ── Error ── */}
            {error && (
                <div className="error-banner" role="alert">
                    {error}
                </div>
            )}

            {/* ── Toolbar: search left, add-button right ── */}
            <div className="page-toolbar">
                <SearchBar value={search} onChange={handleSearch} />
                <button
                    className={`btn-add ${isFormVisible ? "btn-add--active" : ""}`}
                    onClick={() => {
                        if (isFormVisible) {
                            setShowForm(false);
                            setEditingEmployee(null);
                        } else {
                            setShowForm(true);
                        }
                    }}
                >
                    {isFormVisible ? "✕ Cancel" : "+ Add Employee"}
                </button>
            </div>

            {/* ── Form — below toolbar ── */}
            {isFormVisible && (
                <div className="form-section">
                    <EmployeeForm
                        employee={editingEmployee}
                        onSubmit={editingEmployee ? handleUpdate : handleCreate}
                        onCancel={() => {
                            setShowForm(false);
                            setEditingEmployee(null);
                        }}
                    />
                </div>
            )}

            {/* ── List ── */}
            {loading ? (
                <div className="loading-container">
                    <div className="spinner" />
                    <span>Loading…</span>
                </div>
            ) : (
                <EmployeeList
                    employees={employees}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                />
            )}
        </div>
    );
}

export default EmployeePage;
