import { useState } from "react";
import { useNavigate } from "react-router-dom";

type LoginProps = {
  login: (name: string) => void;
};

const Login = (props: LoginProps) => {
  const { login } = props;
  const [user, setUser] = useState("");
  const navigate = useNavigate();
  return (
    <form
      onSubmit={() => {
        login(user);
        navigate("/chat");
      }}
    >
      <h1>Login</h1>
      <input
        type="text"
        value={user}
        placeholder="Username"
        onChange={(e) => {
          setUser(e.target.value);
        }}
      />
      <input type="submit" value="Login" />
    </form>
  );
};
export default Login;
