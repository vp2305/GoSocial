import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import React from "react";
import "./App.css";
import ConfirmPage from "./ConfirmPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <div className="App-header">
              <h1>Welcome to GoSocial!</h1>
            </div>
          }
        />
        <Route path="/confirm/:token" element={<ConfirmPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
