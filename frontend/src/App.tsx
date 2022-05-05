import { useState } from "react";
import "./App.css";
import ChatRoom from "./ChatRoom";
import Login from "./Login";

const App = () => {
  const [user, setUser] = useState<string | null>(null);
  return (
    <div className="App">
      {user ? (
        <ChatRoom myUsername={user} />
      ) : (
        <Login
          login={(name: string) => {
            setUser(name);
          }}
        />
      )}
    </div>
  );
};

export default App;
