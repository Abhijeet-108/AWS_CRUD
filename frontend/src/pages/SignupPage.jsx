import { useEffect } from "react";
import { useAuth } from "../context/AuthContext";
import "./LoginPage.css";

const SignupPage = () => {
    const { login } = useAuth();

    useEffect(() => {
        login();
    }, [login]);

    return (
        <div className="auth-page">
            <div className="auth-card signup-card glass">
                <div className="auth-header">
                    <h1 className="auth-title">
                        Redirecting to signup...
                    </h1>

                    <p className="auth-subtitle">
                        Please wait while we redirect you to Cognito.
                    </p>
                </div>
            </div>
        </div>
    );
};

export default SignupPage;