import { useState } from "react";

type LoginProps = {
  login: (name: string) => void;
};

const Login = (props: LoginProps) => {
  const { login } = props;
  const [user, setUser] = useState("");
  return (
    <form
      className="login"
      onSubmit={() => {
        login(user);
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
