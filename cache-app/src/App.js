import "./App.css";

import GetCache from "./components/GetCache";
import React from "react";
import SetCache from "./components/SetCache";

function App() {
    return (
        <div className="main-container" >
            <SetCache />
            <GetCache />
        </div>
    );
}

export default App;
