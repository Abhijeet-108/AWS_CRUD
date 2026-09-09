import EmployeeCard from "./EmployeeCard";
import "./EmployeeList.css";

const EmployeeList = ({ employees, onEdit, onDelete }) => {
    if (!employees || employees.length === 0) {
        return (
            <div className="empty-state">
                No employees found.
            </div>
        );
    }

    return (
        <div className="employee-grid">
            {employees.map((employee) => (
                <EmployeeCard
                    key={employee.id}
                    employee={employee}
                    onEdit={onEdit}
                    onDelete={onDelete}
                />
            ))}
        </div>
    );
};

export default EmployeeList;
