import "../App.css";

import React, { useState } from "react";

function GetCache() {
    const [key, setKey] = useState("");
    const [value, setValue] = useState(null);

    const handleSubmit = (e) => {
        e.preventDefault();
        fetch(`http://localhost:8080/get?key=${key}`)
            .then((response) => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error("Key not found or expired.");
                }
            })
            .then((data) => setValue(data.value))
            .catch((error) => alert(error.message));
    };

    return (
        <div className="container">
            <h1>Get Cache</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Key"
                    value={key}
                    onChange={(e) => setKey(e.target.value)}
                    required
                />
                <button type="submit">Get Cache</button>
            </form>
            {value && (
                <div className="result">
                    <h2>Value: {value}</h2>
                </div>
            )}
        </div>
    );
}

export default GetCache;
