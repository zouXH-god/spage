"use client";

import { useDevice } from "@/contexts/DeviceContext";
import { useState, ReactNode, useEffect } from "react";
import { usePathname } from "next/navigation";
import { Link } from "react-router-dom";
import Image from "next/image";
import { GetMetaInfoResponse } from "@/api/meta.models";
import { getMetaInfo } from "@/api/meta.api";
import { useSession } from "@/contexts/SessionContext";
import { logout } from "@/api/user.api";
import { LOGIN_PATH } from "@/consts";
import { getGravatarByUser } from "../reusable/Gravatar";
import { NavigationMenuDemo } from "./NavigationMenu";
import { AvatarDropdownMenu } from "./AvatarDropdownMenu";

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
    <nav className="sticky top-0 z-10 pt-4 px-4 w-full border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
      {/* 使用网格布局替换原来的flex布局 */}
      <div className="grid grid-cols-3 pb-4 items-center">
        {/* 左侧区域：品牌信息 */}
        <div className="flex items-center">
          <Link to="/" className="flex items-center">
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
            <NavigationMenuDemo />
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
            {/* 用户下拉菜单 */}
            <div id="user-dropdown">
              <AvatarDropdownMenu
                trigger={
                  <button
                    id="user-avatar"
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
                }
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