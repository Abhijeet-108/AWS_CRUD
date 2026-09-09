import { useEffect, useState } from "react";
import "./ThemeToggle.css";

const ThemeToggle = () => {
    const [dark, setDark] = useState(() => {
        const saved = localStorage.getItem("theme");
        return saved ? saved === "dark" : window.matchMedia("(prefers-color-scheme: dark)").matches;
    });

    useEffect(() => {
        document.documentElement.className = dark ? "dark" : "light";
        localStorage.setItem("theme", dark ? "dark" : "light");
    }, [dark]);

    return (
        <button
            className="theme-toggle"
            onClick={() => setDark((d) => !d)}
            aria-label={dark ? "Switch to light mode" : "Switch to dark mode"}
            title={dark ? "Light mode" : "Dark mode"}
        >
            {dark ? "☀" : "☽"}
        </button>
    );
};

export default ThemeToggle;
