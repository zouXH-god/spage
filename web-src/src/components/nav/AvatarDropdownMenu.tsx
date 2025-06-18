import React, { ReactNode } from "react";
import { useSession } from "@/contexts/SessionContext";
import { logout } from "@/api/user.api";
import { LOGIN_PATH } from "@/consts";
import Link from "next/link";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

interface AvatarDropdownMenuProps {
  /** 自定义触发器 */
  trigger: ReactNode;
  /** 移动端菜单项 */
  menuItems?: Array<{ label: string; href: string }>;
  /** 是否为移动端 */
  isMobile?: boolean;
  /** 下拉菜单关闭回调 */
  onClose?: () => void;
}

export function AvatarDropdownMenu({ 
  trigger,
  menuItems = [],
  isMobile = false,
  onClose = () => {}
}: AvatarDropdownMenuProps) {
  const { currentUser } = useSession();

  const handleLogout = async () => {
    const currentUrl = window.location.pathname +
      (window.location.search || '') +
      (window.location.hash || '');
    const redirectUrl = encodeURIComponent(currentUrl);
    
    try {
      await logout();
      onClose();
      window.location.href = `${LOGIN_PATH}?redirect=${redirectUrl}`;
    } catch (error) {
      console.error("退出登录失败:", error);
    }
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        {trigger}
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="end">
        <DropdownMenuLabel>
          已登录用户 {currentUser?.name}
        </DropdownMenuLabel>
        
        <DropdownMenuGroup>
          <DropdownMenuItem asChild>
            <Link href={`/${currentUser?.name}`}>个人信息</Link>
          </DropdownMenuItem>
          <DropdownMenuItem asChild>
            <Link href="/starred">已点赞</Link>
          </DropdownMenuItem>
          <DropdownMenuItem asChild>
            <Link href="/notifications">订阅</Link>
          </DropdownMenuItem>
          <DropdownMenuItem asChild>
            <Link href="/settings">设置</Link>
          </DropdownMenuItem>
          <DropdownMenuItem asChild>
            <Link href="/help">帮助</Link>
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem asChild>
          <Link href="/admin">管理后台</Link>
        </DropdownMenuItem>
        
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={handleLogout}>
          退出
        </DropdownMenuItem>
        {isMobile && menuItems.length > 0 && (
          <>
            <DropdownMenuSeparator />
            <DropdownMenuLabel>当前页面菜单</DropdownMenuLabel>
            {menuItems.map((item) => (
              <DropdownMenuItem key={item.label} asChild>
                <Link href={item.href}>{item.label}</Link>
              </DropdownMenuItem>
            ))}
          </>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  )
}