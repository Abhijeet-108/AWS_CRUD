import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useState,
} from "react";

import { getProfile } from "../services/authApi";

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    const checkAuth = useCallback(async () => {
        try {
            const profile = await getProfile();

            setUser(profile);
        } catch (error) {
            setUser(null);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        checkAuth();
    }, [checkAuth]);

    const logout = () => {
        window.location.replace(
            "http://localhost:8080/api/auth/logout"
        );
    };

    return (
        <AuthContext.Provider
            value={{
                user,
                loading,
                logout,
                isAuthenticated: !!user,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);

    if (!context) {
        throw new Error(
            "useAuth must be used inside AuthProvider"
        );
    }

    return context;
};