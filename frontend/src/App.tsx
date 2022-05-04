import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import ChatRoom from "./ChatRoom";
import Login from "./Login";

const App = () => (
  <div className="App">
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/room" element={<ChatRoom />} />
      </Routes>
    </Router>
  </div>
);

export default App;
