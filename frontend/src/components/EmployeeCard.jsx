import "./EmployeeCard.css";

const EmployeeCard = ({ employee, onEdit, onDelete }) => {
    const initials = employee.name
        ? employee.name.split(" ").map((w) => w[0]).join("").toUpperCase().slice(0, 2)
        : "??";

    return (
        <div className="employee-card">
            <div className="card-header">
                <div className="card-avatar">{initials}</div>
                <div>
                    <h3 className="card-name">{employee.name}</h3>
                    <span className="card-position">{employee.position}</span>
                </div>
            </div>

            <div className="card-divider" />

            <div className="card-details">
                <div className="detail-row">
                    <span className="detail-label">Email</span>
                    <span className="detail-value">{employee.email}</span>
                </div>
                <div className="detail-row">
                    <span className="detail-label">Department</span>
                    <span className="detail-value">{employee.department}</span>
                </div>
                <div className="detail-row">
                    <span className="detail-label">Salary</span>
                    <span className="detail-value">
                        ₹{Number(employee.salary).toLocaleString("en-IN")}
                    </span>
                </div>
            </div>

            <div className="card-actions">
                <button className="btn-card btn-edit" onClick={() => onEdit(employee)}>
                    Edit
                </button>
                <button className="btn-card btn-delete" onClick={() => onDelete(employee.id)}>
                    Delete
                </button>
            </div>
        </div>
    );
};

export default EmployeeCard;
