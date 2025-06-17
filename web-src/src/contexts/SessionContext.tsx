"use client";

import React, { createContext, useContext, useEffect, useState } from "react";

import { getUser } from "@/api/user.api";
// 假设 getUser 函数用于获取当前登录用户
import { User } from "@/api/user.models";

// 导入用户模型

interface SessionContextType {
  /** 当前登录用户对象，未登录时为 null */
  currentUser: User | null;
  /** 设置当前登录用户 */
  setUser: (user: User | null) => void;
  /** 用户数据是否正在加载中 */
  loading: boolean;
  /** 刷新当前用户数据 */
  refreshUser: () => Promise<void>;
}

const SessionContext = createContext<SessionContextType | undefined>(undefined);

export const SessionProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  // 刷新用户数据的函数
  const refreshUser = async () => {
    setLoading(true);
    try {
      // 调用 API 获取当前用户，假设 getUser() 无参数时获取当前登录用户
      const response = await getUser();
      setUser(response.data.user);
    } catch (error) {
      console.error("Failed to fetch user session:", error);
      setUser(null); // 获取失败则设置为未登录状态
    } finally {
      setLoading(false);
    }
  };

  // 组件挂载时获取一次用户数据
  useEffect(() => {
    refreshUser();
  }, []); // 空依赖数组确保只在组件初次渲染时执行

  return (
    <SessionContext.Provider value={{ currentUser: user, setUser, loading, refreshUser }}>
      {children}
    </SessionContext.Provider>
  );
};

export const useSession = () => {
  const context = useContext(SessionContext);
  if (context === undefined) {
    throw new Error("useSession must be used within a SessionProvider");
  }
  return context;
};
