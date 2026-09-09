import { useAuth } from "../context/AuthContext";
import "./LoginPage.css";

const LoginPage = () => {
    const { login } = useAuth();

    return (
        <div className="auth-page">
            <div className="auth-card glass">

                <div className="auth-header">
                    <h1 className="auth-title">
                        Welcome back
                    </h1>

                    <p className="auth-subtitle">
                        Sign in to your EmpManager account
                    </p>
                </div>

                <button
                    className="btn btn-primary auth-submit"
                    onClick={login}
                >
                    Sign In with Cognito
                </button>

            </div>
        </div>
    );
};

export default LoginPage;