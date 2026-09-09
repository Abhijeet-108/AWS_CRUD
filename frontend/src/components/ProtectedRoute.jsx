import { useEffect } from "react";
import { useAuth } from "../context/AuthContext";

const ProtectedRoute = ({ children }) => {
    const { user, loading } = useAuth();

    useEffect(() => {
        if (!loading && !user) {
            window.location.replace("http://localhost:8080/api/auth/login");
        }
    }, [loading, user]);

    if (loading) {
        return (
            <div style={{
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                height: "60vh",
                gap: "0.75rem",
                color: "var(--text-muted)",
                fontSize: "0.875rem",
            }}>
                <span style={{
                    width: "18px",
                    height: "18px",
                    border: "2px solid var(--border)",
                    borderTopColor: "var(--text-sub)",
                    borderRadius: "50%",
                    display: "inline-block",
                    animation: "spin 0.6s linear infinite",
                }} />
                Checking authentication…
            </div>
        );
    }

    if (!user) return null;

    return children;
};

export default ProtectedRoute;
