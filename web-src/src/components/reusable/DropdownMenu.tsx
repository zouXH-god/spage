"use client";

import React, { useEffect, useRef, useState } from "react";

import { ChevronDown } from "lucide-react";

export interface DropdownMenuItem {
  key: string;
  label: React.ReactNode;
  icon?: React.ReactNode;
  children?: DropdownMenuItem[];
}

interface DropdownMenuProps {
  trigger: React.ReactNode;
  menu: DropdownMenuItem[];
  onSelect: (item: DropdownMenuItem) => void;
}

export default function DropdownMenu({ trigger, menu, onSelect }: DropdownMenuProps) {
  const [open, setOpen] = useState(false);
  const [submenuOpenKey, setSubmenuOpenKey] = useState<string | null>(null);
  const menuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
        setOpen(false);
        setSubmenuOpenKey(null);
      }
    }
    if (open) {
      document.addEventListener("mousedown", handleClickOutside);
    } else {
      document.removeEventListener("mousedown", handleClickOutside);
    }
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [open]);

  const renderMenu = (items: DropdownMenuItem[], isSub = false) => (
    <div
      className={`${
        isSub ? "absolute left-full top-0 mt-0 ml-1 min-w-[180px]" : "min-w-[200px]"
      } bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700 rounded shadow-lg py-1 z-50`}
    >
      {items.map((item) => (
        <div
          key={item.key}
          className="relative group"
          onMouseEnter={() => item.children && setSubmenuOpenKey(item.key)}
          onMouseLeave={() => item.children && setSubmenuOpenKey(null)}
        >
          <button
            className={`flex items-center w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-800 transition ${
              item.children ? "justify-between" : ""
            }`}
            onClick={() => {
              if (!item.children) {
                onSelect(item);
                setOpen(false);
                setSubmenuOpenKey(null);
              }
            }}
            type="button"
          >
            {item.icon && <span className="mr-2">{item.icon}</span>}
            <span>{item.label}</span>
            {item.children && <ChevronDown className="w-4 h-4 ml-auto transform -rotate-90" />}
          </button>
          {item.children && submenuOpenKey === item.key && renderMenu(item.children, true)}
        </div>
      ))}
    </div>
  );

  return (
    <div className="relative inline-block" ref={menuRef}>
      <div onClick={() => setOpen((v) => !v)}>{trigger}</div>
      {open && <div className="absolute left-0 top-full mt-2">{renderMenu(menu)}</div>}
    </div>
  );
}
