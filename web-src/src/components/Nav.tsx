"use client";

import { useDevice } from "@/contexts/DeviceContext";
import { useState, ReactNode, useEffect } from "react";
import { usePathname } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import { GetMetaInfoResponse } from "@/api/meta.models";
import { getMetaInfo } from "@/api/meta.api";
import { useSession } from "@/contexts/SessionContext";
import { logout } from "@/api/user.api";
import { LOGIN_PATH } from "@/consts";
import { getGravatarByUser } from "./reusable/Gravatar";

type NavMenuProps = {
  children: ReactNode;
};

// 导航菜单类型定义
type MenuType = 'home' | 'project' | 'owner' | 'settings';

// 中间菜单项接口
interface MenuItem {
  label: string;
  href: string;
  icon?: ReactNode; // 可选的图标
}

// 不同页面类型的菜单配置
const MENU_CONFIGS: Record<MenuType, MenuItem[]> = {
  home: [
    { label: '浏览', href: '/explore' },
    { label: '文档', href: '/docs' },
    { label: '示例', href: '/examples' }
  ],
  project: [
    { label: '概览', href: '#overview' },
    { label: '代码', href: '#code' },
    { label: '问题', href: '#issues' },
    { label: '讨论', href: '#discussions' }
  ],
  owner: [
    { label: '仓库', href: '#repos' },
    { label: '活动', href: '#activity' },
    { label: '关注者', href: '#followers' }
  ],
  settings: [
    { label: '个人资料', href: '/settings/profile' },
    { label: '账号', href: '/settings/account' },
    { label: '安全', href: '/settings/security' }
  ]
};



export default function Nav() {
  const [metaInfo, setMetaInfo] = useState<GetMetaInfoResponse | null>(null);
  const [userMenuOpen, setUserMenuOpen] = useState(false);
  const pathname = usePathname();
  const { isMobile } = useDevice();
  const { currentUser, refreshUser } = useSession();

  // 首次渲染effect
  useEffect(() => {
    if (!currentUser) {
      refreshUser(); // 确保当前用户信息是最新的
    }
    getMetaInfo()
      .then(response => {
        setMetaInfo(response.data);
      })
      .catch(error => {
        console.error("获取元信息失败:", error);
      });
  }, []);

  // 点击外部关闭菜单
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      const target = event.target as Node;
      const dropdown = document.getElementById("user-dropdown");
      const avatar = document.getElementById("user-avatar");

      if (dropdown && !dropdown.contains(target) && avatar && !avatar.contains(target)) {
        setUserMenuOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // 根据路径判断当前页面类型
  const getMenuType = (): MenuType => {
    if (pathname.includes('/settings')) return 'settings';
    if (pathname.match(/\/[^/]+\/[^/]+$/)) return 'project'; // /:owner/:project
    if (pathname.match(/\/[^/]+$/)) return 'owner'; // /:owner
    return 'home';
  };

  const menuType = getMenuType();
  const menuItems = MENU_CONFIGS[menuType];

  return (
    <nav className="sticky top-0 z-10 pt-4 px-4 w-full border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-900">
      {/* 使用网格布局替换原来的flex布局 */}
      <div className="grid grid-cols-3 pb-4 items-center">
        {/* 左侧区域：品牌信息 */}
        <div className="flex items-center">
          <Link href="/" className="flex items-center">
            <div className="h-10 relative" style={{ aspectRatio: 'auto' }}>
              <Image
                src={metaInfo?.icon || "/apage.svg"}
                alt="Logo"
                width={40}
                height={40}
                className="h-full w-auto object-contain"
                priority
              />
            </div>
            {!isMobile && (
              <span className="ml-2 text-xl font-semibold">{metaInfo?.name || "Spage"}</span>
            )}
          </Link>
        </div>

        {/* 中间区域：页面菜单 (仅桌面版) */}
        <div className="flex justify-center items-center">
          {!isMobile && (
            <div className="flex space-x-6">
              {menuItems.map((item) => (
                <Link
                  key={item.label}
                  href={item.href}
                  className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white font-medium"
                >
                  {item.label}
                </Link>
              ))}
            </div>
          )}
        </div>

        {/* 右侧区域：个人菜单 */}
        <div className="flex items-center justify-end">
          {!isMobile && (
            <>
              <button className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white mr-4">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                  <path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd" />
                </svg>
              </button>
              <button className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white mr-4">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                </svg>
              </button>
            </>
          )}
          <div className="relative">
            <button
              id="user-avatar"
              onClick={() => setUserMenuOpen(!userMenuOpen)}
              className="flex items-center focus:outline-none"
            >
              <div className="h-8 w-8 rounded-full overflow-hidden">
                {getGravatarByUser(currentUser ?? undefined)}
              </div>
              {!isMobile && (
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4 ml-1 text-gray-500"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path fillRule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clipRule="evenodd" />
                </svg>
              )}
            </button>

            {/* 用户下拉菜单 */}
            <div id="user-dropdown">
              <UserDropdownMenu
                isOpen={userMenuOpen}
                onClose={() => setUserMenuOpen(false)}
                menuItems={menuItems}
                isMobile={isMobile}
              />
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}

// 用户下拉菜单组件
function UserDropdownMenu({
  isOpen,
  onClose,
  menuItems,
  isMobile
}: {
  isOpen: boolean;
  onClose: () => void;
  menuItems: MenuItem[];
  isMobile: boolean;
}) {
  if (!isOpen) return null;
  const { currentUser, refreshUser } = useSession();

  useEffect(() => {
    if (!currentUser) {
      refreshUser(); // 确保当前用户信息是最新的
    }
  }, [])

  const handleLogout = async () => {
    const currentUrl = window.location.pathname +
      (window.location.search || '') +
      (window.location.hash || '');
    const redirectUrl = encodeURIComponent(currentUrl);
    logout()
      .then(() => {
        onClose();
        window.location.href = `${LOGIN_PATH}?redirect=${redirectUrl}`;
      })
      .catch((error) => {
        console.error("退出登录失败:", error);
      });
  }

  return (
    <div className="absolute right-0 mt-2 w-48 bg-slate-100 dark:bg-gray-700 rounded-md shadow-lg py-1 z-10 text-gray-700 dark:text-gray-200">
      {/* 用户信息头部 */}
      <div className="px-4 py-3 border-b border-gray-200 dark:border-gray-600">
        <div className="font-medium">已登录用户 {currentUser?.name}</div>
      </div>

      {/* 用户菜单项 */}
      <div>
        <Link
          href={`/${currentUser?.name}`}
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
          </svg>
          个人信息
        </Link>

        <Link
          href="/starred"
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
          </svg>
          已点赞
        </Link>

        <Link
          href="/notifications"
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
          </svg>
          订阅
        </Link>

        <Link
          href="/settings"
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          设置
        </Link>

        <Link
          href="/help"
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          帮助
        </Link>
      </div>

      {/* 管理后台 (带分隔线) */}
      <div className="border-t border-b border-gray-200 dark:border-gray-600 py-1">
        <Link
          href="/admin"
          className="flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
          onClick={onClose}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01" />
          </svg>
          管理后台
        </Link>
      </div>

      {/* 退出按钮 */}
      <div>
        <button
          className="w-full flex items-center px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600 text-left"
          onClick={handleLogout}
        >
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          退出
        </button>
      </div>

      {/* 移动端显示页面菜单项 */}
      {isMobile && (
        <>
          <hr className="my-1 border-gray-200 dark:border-gray-600" />
          <div className="px-4 py-2 text-xs text-gray-500 dark:text-gray-400">当前页面菜单</div>
          {menuItems.map((item) => (
            <Link
              key={item.label}
              href={item.href}
              className="block px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-600"
              onClick={onClose}
            >
              {item.label}
            </Link>
          ))}
        </>
      )}
    </div>
  );
}