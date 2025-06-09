import React from "react";

export default function LoginLayout({ children }: { children: React.ReactNode }) {
  return (
    <div style={{ minHeight: "100vh", display: "flex", alignItems: "center", justifyContent: "center", background: "#f5f6fa" }}>
      <div style={{ width: 360, padding: 32, background: "#fff", borderRadius: 8, boxShadow: "0 2px 16px rgba(0,0,0,0.08)" }}>
        <h2 style={{ textAlign: "center", marginBottom: 24 }}>登录后台</h2>
        {children}
      </div>
    </div>
  );
}