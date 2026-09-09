import { Link } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import ThemeToggle from "./ThemeToggle";
import "./Header.css";

const Header = () => {
    const { isAuthenticated, user, logout } = useAuth();

    return (
        <header className="header">
            <div className="header-inner">
                <Link to="/" className="header-brand">EmpManager</Link>
                <div className="header-actions">
                    <ThemeToggle />
                    {isAuthenticated && (
                        <>
                            {user?.name && (
                                <span className="header-user">{user.name}</span>
                            )}
                            <button className="btn-logout" onClick={logout}>
                                Logout
                            </button>
                        </>
                    )}
                </div>
            </div>
        </header>
    );
};

export default Header;
