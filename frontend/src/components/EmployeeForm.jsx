import { useEffect, useState } from "react";
import "./EmployeeForm.css";

const initialForm = {
    name: "",
    email: "",
    department: "",
    position: "",
    salary: "",
};

const EmployeeForm = ({ employee, onSubmit, onCancel }) => {
    const [form, setForm] = useState(initialForm);

    useEffect(() => {
        setForm(
            employee
                ? {
                      name: employee.name || "",
                      email: employee.email || "",
                      department: employee.department || "",
                      position: employee.position || "",
                      salary: employee.salary || "",
                  }
                : initialForm
        );
    }, [employee]);

    const handleChange = (e) =>
        setForm((prev) => ({ ...prev, [e.target.name]: e.target.value }));

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit({ ...form, salary: Number(form.salary) });
    };

    return (
        <form className="employee-form" onSubmit={handleSubmit}>
            <h2 className="form-title">
                {employee ? "Edit Employee" : "Add Employee"}
            </h2>

            <div className="form-grid">
                <div className="form-group">
                    <label className="form-label" htmlFor="name">Name</label>
                    <input
                        id="name"
                        name="name"
                        className="form-input"
                        placeholder="Full name"
                        value={form.name}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label className="form-label" htmlFor="email">Email</label>
                    <input
                        id="email"
                        name="email"
                        type="email"
                        className="form-input"
                        placeholder="name@company.com"
                        value={form.email}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label className="form-label" htmlFor="department">Department</label>
                    <input
                        id="department"
                        name="department"
                        className="form-input"
                        placeholder="e.g. Engineering"
                        value={form.department}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label className="form-label" htmlFor="position">Position</label>
                    <input
                        id="position"
                        name="position"
                        className="form-input"
                        placeholder="e.g. Software Engineer"
                        value={form.position}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group full-width">
                    <label className="form-label" htmlFor="salary">Salary (₹)</label>
                    <input
                        id="salary"
                        name="salary"
                        type="number"
                        min="0"
                        className="form-input"
                        placeholder="e.g. 800000"
                        value={form.salary}
                        onChange={handleChange}
                        required
                    />
                </div>
            </div>

            <div className="form-actions">
                <button type="submit" className="btn-primary">
                    {employee ? "Save Changes" : "Add Employee"}
                </button>
                {onCancel && (
                    <button type="button" className="btn-ghost" onClick={onCancel}>
                        Cancel
                    </button>
                )}
            </div>
        </form>
    );
};

export default EmployeeForm;
