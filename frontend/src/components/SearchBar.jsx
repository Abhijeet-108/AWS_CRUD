import "./SearchBar.css";

const SearchBar = ({ value, onChange }) => (
    <div className="search-bar">
        <svg
            className="search-icon"
            viewBox="0 0 16 16"
            fill="none"
            stroke="currentColor"
            strokeWidth="1.5"
            aria-hidden="true"
        >
            <circle cx="6.5" cy="6.5" r="4.5" />
            <line x1="10.5" y1="10.5" x2="14" y2="14" />
        </svg>

        <input
            type="text"
            className="search-input"
            placeholder="Search employees…"
            value={value}
            onChange={(e) => onChange(e.target.value)}
            aria-label="Search employees"
        />

        {value && (
            <button
                className="search-clear"
                onClick={() => onChange("")}
                aria-label="Clear search"
            >
                ✕
            </button>
        )}
    </div>
);

export default SearchBar;
