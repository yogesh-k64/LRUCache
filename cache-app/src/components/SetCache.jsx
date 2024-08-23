import "../App.css";

import React, { useState } from "react";

function SetCache() {
    const [key, setKey] = useState("");
    const [value, setValue] = useState("");
    const [expiration, setExpiration] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        fetch("http://localhost:8080/set", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ key, value, expiration: parseInt(expiration, 10) }),
        }).then((response) => {
            if (response.ok) {
                alert("Cache set successfully!");
                setKey("");
                setValue("");
                setExpiration("");
            }
        });
    };

    return (
        <div className="container">
            <h1>Set Cache</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Key"
                    value={key}
                    onChange={(e) => setKey(e.target.value)}
                    required
                />
                <input
                    type="text"
                    placeholder="Value"
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                    required
                />
                <input
                    type="number"
                    placeholder="Expiration (seconds)"
                    value={expiration}
                    onChange={(e) => setExpiration(e.target.value)}
                    required
                />
                <button type="submit">Set Cache</button>
            </form>
        </div>
    );
}

export default SetCache;
