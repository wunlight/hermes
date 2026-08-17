import useAuthStore from "@/core/stores/authStore";
import { useState } from "react";
import { useNavigate } from "react-router";
import { Login } from "../service/service";
import type { LoginRequest } from "../types/types";

const LoginView = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [showPassword, setShowPassword] = useState(false);

  const setAuth = useAuthStore((state) => state.setAuth);

  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const req: LoginRequest = { email, password };

    const res = await Login(req);

    setAuth(res.userID, res.accessToken);

    navigate("/", { replace: true });
  };

  return (
    <div className="grid place-content-center h-dvh bg-slate-200">
      <form onSubmit={handleLogin} className="p-4 bg-white rounded-lg">
        <h6 className="font-semibold text-2xl">Login</h6>
        <div className="flex flex-col gap-3 pt-16 pb-12 min-w-72">
          <div className="flex flex-col gap-1">
            <label htmlFor="email" className="text-xs text-slate-500">
              Email
            </label>
            <input
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              type="text"
              className="px-3 py-1.5 text-sm border border-slate-300 rounded-md outline-none"
              placeholder="Enter your email"
              id="email"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label htmlFor="password" className="text-xs text-slate-500">
              Password
            </label>
            <div className="relative w-full">
              <input
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type={showPassword ? "text" : "password"}
                className="px-3 py-1.5 w-full text-sm border border-slate-300 rounded-md outline-none"
                placeholder="Enter your password"
                id="password"
              />
              <button
                type="button"
                onClick={() => setShowPassword((value) => !value)}
                className="absolute right-2 top-1/2 -translate-1/2"
              >
                <span
                  className={`${showPassword ? "icon-[mdi--eye-off]" : "icon-[mdi--eye]"} text-slate-500`}
                />
              </button>
            </div>
          </div>
        </div>

        <button
          type="submit"
          className="px-4 h-9 w-full text-sm text-white bg-blue-500 rounded-md"
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default LoginView;
