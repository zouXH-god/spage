"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const router = useRouter();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 简单示例：用户名和密码都为 admin 时通过
    if (username === "admin" && password === "admin") {
      localStorage.setItem("token", "mock-token");
      router.replace("/dashboard");
    } else {
      setError("用户名或密码错误");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div style={{ marginBottom: 16 }}>
        <input
          type="text"
          placeholder="用户名"
          value={username}
          onChange={e => setUsername(e.target.value)}
          style={{ width: "100%", padding: 8, borderRadius: 4, border: "1px solid #ddd" }}
        />
      </div>
      <div style={{ marginBottom: 16 }}>
        <input
          type="password"
          placeholder="密码"
          value={password}
          onChange={e => setPassword(e.target.value)}
          style={{ width: "100%", padding: 8, borderRadius: 4, border: "1px solid #ddd" }}
        />
      </div>
      {error && <div style={{ color: "red", marginBottom: 12 }}>{error}</div>}
      <button
        type="submit"
        style={{
          width: "100%",
          padding: 10,
          borderRadius: 4,
          border: "none",
          background: "#1677ff",
          color: "#fff",
          fontWeight: "bold",
          cursor: "pointer"
        }}
      >
        登录
      </button>
    </form>
  );
}