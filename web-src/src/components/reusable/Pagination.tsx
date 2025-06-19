"use client";

import React, { useMemo } from "react";

interface PaginationProps {
  currentPage: number;
  totalItems: number;
  pageSize: number;
  onPageChange: (page: number) => void;
  className?: string;
  showPageSize?: boolean;
}

export const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalItems,
  pageSize,
  onPageChange,
  className = "",
  showPageSize = false,
}) => {
  // 计算总页数
  const totalPages = useMemo(() => {
    return Math.ceil(totalItems / pageSize);
  }, [totalItems, pageSize]);

  // 生成页码数组
  const pageNumbers = useMemo(() => {
    const pages = [];
    const rangeStart = Math.max(1, currentPage - 2);
    const rangeEnd = Math.min(totalPages, currentPage + 2);

    for (let i = rangeStart; i <= rangeEnd; i++) {
      pages.push(i);
    }

    return pages;
  }, [currentPage, totalPages]);

  // 当总页数小于等于1时，不显示分页器
  if (totalPages <= 1) return null;

  return (
    <div className={`flex flex-col md:flex-row justify-between items-center gap-4 ${className}`}>
      {/* 页码信息 */}
      <div className="text-sm text-gray-500 dark:text-gray-400">
        共 {totalItems} 项，第 {currentPage} 页 / 共 {totalPages} 页
      </div>
      {/* 分页控制 */}
      <div className="flex items-center">
        {/* 上一页按钮 */}
        <button
          onClick={() => currentPage > 1 && onPageChange(currentPage - 1)}
          disabled={currentPage === 1}
          className={`px-3 py-1 mx-1 rounded-md 
                    ${
                      currentPage === 1
                        ? "bg-gray-100 text-gray-400 dark:bg-gray-800 dark:text-gray-600 cursor-not-allowed"
                        : "bg-gray-100 hover:bg-gray-200 text-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-gray-300"
                    }`}
          aria-label="上一页"
        >
          上一页
        </button>

        {/* 首页 */}
        {pageNumbers[0] > 1 && (
          <>
            <button
              onClick={() => onPageChange(1)}
              className="px-3 py-1 mx-1 rounded-md bg-gray-100 hover:bg-gray-200 text-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-gray-300"
            >
              1
            </button>
            {pageNumbers[0] > 2 && <span className="mx-1">...</span>}
          </>
        )}

        {/* 页码按钮 */}
        {pageNumbers.map((page) => (
          <button
            key={page}
            onClick={() => onPageChange(page)}
            className={`px-3 py-1 mx-1 rounded-md 
                      ${
                        currentPage === page
                          ? "bg-blue-500 text-white"
                          : "bg-gray-100 hover:bg-gray-200 text-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-gray-300"
                      }`}
          >
            {page}
          </button>
        ))}

        {/* 尾页 */}
        {pageNumbers[pageNumbers.length - 1] < totalPages && (
          <>
            {pageNumbers[pageNumbers.length - 1] < totalPages - 1 && (
              <span className="mx-1">...</span>
            )}
            <button
              onClick={() => onPageChange(totalPages)}
              className="px-3 py-1 mx-1 rounded-md bg-gray-100 hover:bg-gray-200 text-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-gray-300"
            >
              {totalPages}
            </button>
          </>
        )}
        {/* 下一页按钮 */}
        <button
          onClick={() => currentPage < totalPages && onPageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
          className={`px-3 py-1 mx-1 rounded-md 
                    ${
                      currentPage === totalPages
                        ? "bg-gray-100 text-gray-400 dark:bg-gray-800 dark:text-gray-600 cursor-not-allowed"
                        : "bg-gray-100 hover:bg-gray-200 text-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-gray-300"
                    }`}
          aria-label="下一页"
        >
          下一页
        </button>
      </div>

      {/* 可选的页面大小选择器 */}
      {showPageSize && (
        <select
          className="px-2 py-1 rounded-md bg-gray-100 text-gray-700 dark:bg-gray-800 dark:text-gray-300 border-none"
          onChange={() => {}}
          value={pageSize}
        >
          <option value={5}>5条/页</option>
          <option value={10}>10条/页</option>
          <option value={20}>20条/页</option>
          <option value={50}>50条/页</option>
        </select>
      )}
    </div>
  );
};
