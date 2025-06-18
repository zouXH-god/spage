"use client";

import React, { useEffect, useMemo, useState } from "react";

import { getUser, getUserProjects } from "@/api/user.api";
import { User } from "@/api/user.models";
import { getGravatarByUser } from "@/components/reusable/Gravatar";
import { Project } from "@/api/project.models";
import { Pagination } from "@/components/reusable/Pagination";
import { ProjectCard } from "./ProjectCard";

export function UserPage({ id }: { id: number }) {
  const [user, setUser] = useState<User | null>(null);
  // 项目分页查询
  const [page, setPage] = useState(1);
  const [pageSize, ] = useState(10);
  const [totalItems, setTotalItems] = useState(0);
  const [projects, setProjects] = useState<Project[]>([]);
  const [loading, setLoading] = useState(false);

  // 使用 useMemo 计算总页数，避免不必要的重复计算
  const totalPages = useMemo(() => {
    return Math.ceil(totalItems / pageSize);
  }, [totalItems, pageSize]);

  // 获取用户信息
  useEffect(() => {
    console.log(`AAAFetching user data for ID: ${id}`);
    getUser(id)
      .then((response) => {
        console.log("AAAUser data fetched successfully:", response.data);
        const data = response.data;
        setUser(data.user);
      })
      .catch((error) => {
        console.error("Error fetching user data:", error);
      });
  }, [id]);

  useEffect(() => {
    if (user) {
      setLoading(true);
      console.log(`Fetching projects for user ${user.id}, page: ${page}, pageSize: ${pageSize}`);
      getUserProjects(user.id, page, pageSize)
        .then((response) => {
          const data = response.data;
          setProjects(data.projects);
          setTotalItems(data.total);
        })
        .catch((error) => {
          console.error("Error fetching user projects:", error);
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [user, page, pageSize]); // 添加 page 和 pageSize 作为依赖

  // 页码变更处理函数
  const handlePageChange = (newPage: number) => {
    if (newPage < 1 || newPage > totalPages || newPage === page) return;
    setPage(newPage);
    // 滚动到页面顶部，提升用户体验
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-100 to-slate-300 dark:from-gray-900 dark:to-gray-800 flex justify-center items-start py-10">
      <div className="flex w-full max-w-6xl gap-8">
        {/* 侧边栏：用户信息卡片 */}
        <aside className="w-full max-w-xs">
          <div className="rounded-3xl shadow-xl bg-white/80 dark:bg-gray-800/80 p-6 flex flex-col items-center">
            {/* 头像 */}
            <div className="w-32 h-32 rounded-full overflow-hidden border-4 border-white dark:border-gray-800 shadow-lg mb-4">
              {getGravatarByUser(user ?? undefined, "w-full h-full object-cover")}
            </div>
            {/* 显示名称 */}
            <h2 className="text-2xl font-bold mb-1">{user?.displayName || user?.name}</h2>
            {/* 用户名 */}
            <div className="text-sm text-gray-500 dark:text-gray-400 mb-2">{user?.name}</div>
            {/* 个人信息 */}
            <div className="flex items-center text-sm text-gray-500 dark:text-gray-400 mb-2">
              <span>{user?.email}</span>
            </div>
            {/* 摘要 */}
            <div className="w-full mt-4">
              <div className="font-semibold text-gray-700 dark:text-gray-200 mb-1">摘要</div>
              <div className="text-sm text-gray-600 dark:text-gray-400">{user?.description}</div>
            </div>
          </div>
        </aside>
        {/* 主体内容 */}
        <main className="flex-1">
          <div className="rounded-3xl shadow-xl bg-white/80 dark:bg-gray-800/80 p-6">
            {/* 顶部tab和搜索 */}
            <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mb-6">
              <div className="flex gap-2 flex-wrap">
                <button className="px-4 py-1 rounded-full bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300">
                  项目
                </button>
                <button className="px-4 py-1 rounded-full bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300">
                  公开活动
                </button>
                <button className="px-4 py-1 rounded-full bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300">
                  已收藏
                </button>
              </div>
              <div className="flex gap-2">
                <input
                  type="text"
                  placeholder="搜索仓库..."
                  className="rounded-full px-4 py-1 bg-gray-100 dark:bg-gray-800 focus:outline-none"
                />
                <button className="px-3 py-1 rounded-full bg-gray-200 dark:bg-gray-700 text-gray-600 dark:text-gray-300">
                  过滤
                </button>
              </div>
            </div>
            {/* 仓库列表区域 - 使用 flex-1 让它自动填充可用空间 */}
            <div className="flex-1 overflow-y-auto pr-2">
              <div className="space-y-4">
                {loading || !projects ? (
                  <div className="text-center py-8">
                    <div className="inline-block animate-spin rounded-full h-8 w-8 border-4 border-gray-300 dark:border-gray-600 border-t-gray-600 dark:border-t-gray-300"></div>
                    <p className="mt-2 text-gray-600 dark:text-gray-400">加载中...</p>
                  </div>
                ) : projects.length > 0 ? (
                  projects.map((project) => (
                    <ProjectCard key={project.id} {...project} />
                  ))
                ) : (
                  <div className="text-center py-8 text-gray-500 dark:text-gray-400">
                    没有找到项目
                  </div>
                )}
              </div>
            </div>
            {/* 分页组件 - 保持在底部固定 */}
            {totalItems > pageSize && (
              <div className="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700">
                <Pagination
                  currentPage={page}
                  totalItems={totalItems}
                  pageSize={pageSize}
                  onPageChange={handlePageChange}
                />
              </div>
            )}
          </div>
        </main>
      </div>
    </div>
  );
}
