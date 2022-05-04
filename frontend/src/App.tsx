import { useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import ChatRoom from "./ChatRoom";
import Login from "./Login";

const App = () => {
  const [user, setUser] = useState("");
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Login login={(name) => setUser(name)} />} />
          <Route path="/chat" element={<ChatRoom myUsername={user} />} />
        </Routes>
      </Router>
    </div>
  );
};

export default App;
