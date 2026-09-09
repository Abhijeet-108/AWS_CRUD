// import { useState } from "react";

// const initialForm = {
//     name: "",
//     email: "",
//     department: "",
//     position: "",
//     salary: "",
// };

// function EmployeeForm({
//     employee,
//     onSubmit,
//     onCancel,
// }) {
//     const [form, setForm] = useState(
//         employee
//         ? {
//             name: employee.name,
//             email: employee.email,
//             department: employee.department,
//             position: employee.position,
//             salary: employee.salary,
//             }
//         : initialForm
//     );

//     const handleChange = (event) => {
//         const { name, value } = event.target;

//         setForm((previous) => ({
//         ...previous,
//         [name]: value,
//         }));
//     };

//     const handleSubmit = (event) => {
//         event.preventDefault();

//         onSubmit({
//         ...form,
//         salary: Number(form.salary),
//         });
//     };

//     return (
//         <form onSubmit={handleSubmit}>
//         <div style={{ marginBottom: "10px" }}>
//             <label>Name</label>
//             <input
//             name="name"
//             value={form.name}
//             onChange={handleChange}
//             />
//         </div>

//         <div style={{ marginBottom: "10px" }}>
//             <label>Email</label>
//             <input
//             type="email"
//             name="email"
//             value={form.email}
//             onChange={handleChange}
//             />
//         </div>

//         <div style={{ marginBottom: "10px" }}>
//             <label>Department</label>
//             <input
//             name="department"
//             value={form.department}
//             onChange={handleChange}
//             />
//         </div>

//         <div style={{ marginBottom: "10px" }}>
//             <label>Position</label>
//             <input
//             name="position"
//             value={form.position}
//             onChange={handleChange}
//             />
//         </div>

//         <div style={{ marginBottom: "10px" }}>
//             <label>Salary</label>
//             <input
//             type="number"
//             name="salary"
//             value={form.salary}
//             onChange={handleChange}
//             />
//         </div>

//         <button type="submit" style={{ marginBottom: "10px" }}>
//             {employee ? "Update Employee" : "Add Employee"}
//         </button>

//         {onCancel && (
//             <button
//             type="button"
//             onClick={onCancel}
//             >
//             Cancel
//             </button>
//         )}
//         </form>
//     );
// }

// export default EmployeeForm;