import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import SignIn from "./components/SignIn";
import Home from "./components/Home";
import Navbar from "./components/Navbar";
import OutpatientScreeningCreate from "./components/OutpatientScreeningCreate";


function App() {
  const [token, setToken] = useState<string | null>();
  useEffect(() => {
    setToken(localStorage.getItem("token"));
  }, []);

  // if (!token) {
  //   return (
  //     <SignIn />
  //   );
  // }
// Create return function to get OutpatientScreeningCreate route
  return (
    <Router>
      <div>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/outpatientscreeningcreate" element={<OutpatientScreeningCreate />} />
        </Routes>
      </div>
    </Router>
  );
}
export default App;
