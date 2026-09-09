import { Routes, Route } from "react-router-dom";

import Header from "./components/Header";
import Footer from "./components/Footer";
import ProtectedRoute from "./components/ProtectedRoute";

import EmployeePage from "./pages/EmployeePage";

function App() {
    return (
        <>
            <Header />

            <main className="app-main">
                <Routes>
                    <Route
                        path="/"
                        element={
                            <ProtectedRoute>
                                <EmployeePage />
                            </ProtectedRoute>
                        }
                    />
                </Routes>
            </main>

            <Footer />
        </>
    );
}

export default App;

// import { useAuth } from "./context/AuthContext";

// function App() {
//     const {
//         user,
//         loading,
//         login,
//         logout,
//         isAuthenticated,
//     } = useAuth();

//     if (loading) {
//         return <p>Checking authentication...</p>;
//     }

//     if (!isAuthenticated) {
//         return (
//             <div>
//                 <h1>Employee Management</h1>

//                 <p>You are not logged in.</p>

//                 <button onClick={login}>
//                     Login
//                 </button>
//             </div>
//         );
//     }

//     return (
//         <div>
//             <h1>Employee Management</h1>

//             <p>User ID: {user.user_id}</p>

//             <p>Username: {user.username}</p>

//             <p>Name: {user.name}</p>

//             <button onClick={logout}>
//                 Logout
//             </button>
//         </div>
//     );
// }

// export default App;